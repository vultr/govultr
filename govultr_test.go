package govultr

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	mux    *http.ServeMux
	ctx    = context.TODO()
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	setup()
	defer teardown()

	if client.BaseURL == nil || client.BaseURL.String() != server.URL {
		t.Errorf("NewClient BaseURL = %v, expected %v", client.BaseURL, server.URL)
	}

	//if client != "dummy-key" {
	//	t.Errorf("NewClient ApiKey = %v, expected %v", client.APIKey.key, "dummy-key")
	//}

	if client.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", client.UserAgent, userAgent)
	}

	services := []string{
		"Account",
	}

	copied := reflect.ValueOf(client)
	copyValue := reflect.Indirect(copied)

	for _, s := range services {
		if copyValue.FieldByName(s).IsNil() {
			t.Errorf("c.%s shouldn't be nil", s)
		}
	}
}

func TestClient_DoWithContext(t *testing.T) {
	setup()
	defer teardown()

	type vultr struct {
		Bird string
	}

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if method := http.MethodGet; method != request.Method {
			t.Errorf("Request method = %v, expecting %v", request.Method, method)
		}
		fmt.Fprint(writer, `{"Bird":"vultr"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	data := new(vultr)

	_, err := client.DoWithContext(context.Background(), req, data)

	if err != nil {
		t.Fatalf("DoWithContext(): %v", err)
	}

	expected := &vultr{"vultr"}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Response body = %v, expected %v", data, expected)
	}
}

func TestClient_DoWithContextFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if method := http.MethodGet; method != request.Method {
			t.Errorf("Request method = %v, expecting %v", request.Method, method)
		}
		writer.WriteHeader(500)
		fmt.Fprint(writer, `{Error}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	_, err := client.DoWithContext(context.Background(), req, nil)

	if !strings.Contains(err.Error(), "gave up after") || !strings.Contains(err.Error(), "last error") {
		t.Fatalf("DoWithContext(): %v: expected 'gave up after ..., last error ...'", err)
	}
}

type errRoundTripper struct{}

func (errRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("fake error")
}

func TestClient_DoWithContextError(t *testing.T) {
	setup()
	defer teardown()

	client = NewClient(&http.Client{
		Transport: errRoundTripper{},
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	var panicked string
	func() {
		defer func() {
			if err := recover(); err != nil {
				panicked = fmt.Sprint(err)
			}
		}()
		client.DoWithContext(context.Background(), req, nil) //nolint:all
	}()
	if panicked != "" {
		t.Errorf("unexpected panic: %s", panicked)
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(nil)

	in := "/unit"
	out := defaultBase + "/unit"

	inRequest := RequestBody{"balance": 500}
	outRequest := `{"balance":500}` + "\n"

	req, _ := c.NewRequest(ctx, http.MethodPost, in, inRequest)
	if req.URL.String() != out {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", in, req.URL, out)
	}

	body, _ := io.ReadAll(req.Body)

	if string(body) != outRequest {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inRequest, string(body), outRequest)
	}

	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("NewRequest() Header Content Type = %v, expected %v", contentType, "application/x-www-form-urlencoded")
	}
}

func TestClient_SetBaseUrl(t *testing.T) {
	setup()
	defer teardown()

	base := "http://localhost/vultr"
	err := client.SetBaseURL(base)

	if err != nil {
		t.Fatalf("SetBaseUrl unexpected error: %v", err)
	}

	if client.BaseURL.String() != base {
		t.Errorf("NewClient BaseUrl = %v, expected %v", client.BaseURL, base)
	}

	if err := client.SetBaseURL(":"); err == nil {
		t.Error("Expected invalid BaseURL to fail")
	}
}

func TestClient_SetUserAgent(t *testing.T) {
	setup()
	defer teardown()

	ua := "vultr/testing"
	client.SetUserAgent(ua)

	if client.UserAgent != ua {
		t.Errorf("NewClient UserAgent = %v, expected %v", client.UserAgent, userAgent)
	}
}

func TestClient_SetRateLimit(t *testing.T) {
	setup()
	defer teardown()

	time := 600 * time.Millisecond
	client.SetRateLimit(time)

	if client.client.RetryWaitMax != time {
		t.Errorf("NewClient max RateLimit = %v, expected %v", client.client.RetryWaitMax, time)
	}
}

func TestClient_OnRequestCompleted(t *testing.T) {
	setup()
	defer teardown()

	var completedReq *http.Request
	var completedRes string

	type Vultr struct {
		Bird string
	}

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, `{"Vultr":"bird"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	data := new(Vultr)

	client.OnRequestCompleted(func(request *http.Request, response *http.Response) {
		completedReq = req
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			t.Errorf("Failed to dump response: %s", err)
		}
		completedRes = string(dump)
	})

	_, err := client.DoWithContext(context.Background(), req, data)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	if !reflect.DeepEqual(req, completedReq) {
		t.Errorf("Completed request = %v, expected %v", completedReq, req)
	}

	expected := `{"Vultr":"bird"}`
	if !strings.Contains(completedRes, expected) {
		t.Errorf("expected response to contain %v, Response = %v", expected, completedRes)
	}
}

func TestClient_SetRetryLimit(t *testing.T) {
	setup()
	defer teardown()

	client.SetRetryLimit(4)

	if client.client.RetryMax != 4 {
		t.Errorf("NewClient RateLimit = %v, expected %v", client.client.RetryMax, 4)
	}
}

func TestNewRequest_badURI(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest(ctx, http.MethodGet, ":/1.", nil)
	if err == nil {
		t.Error("expected invalid URI to fail")
	}
}

func TestNewRequest_badBody(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest(ctx, http.MethodGet, "/", make(chan int))
	if err == nil {
		t.Error("expected invalid Body to fail")
	}
}

func TestRequest_InvalidCall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/wrong", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, nil)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/wrong", nil)
	if _, err := client.DoWithContext(ctx, req, nil); err == nil {
		t.Error("Expected invalid status code to bad request")
	}
}

func TestRequest_InvalidResponseBody(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/wrong", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, `{`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/wrong", nil)
	if _, err := client.DoWithContext(ctx, req, struct{}{}); err == nil {
		t.Error("Expected response body to be invalid")
	}
}
