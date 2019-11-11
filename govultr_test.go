package govultr

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
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

	client = NewClient(nil, "dummy-key")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func handleString(method, pattern, resp string) {
	var failed bool // fail every other request to test retries but still work consistently
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`Invalid HTTP method. Check that the method (POST|GET) matches what the documentation indicates.`))
			return
		}

		failed = !failed
		if failed {
			w.WriteHeader(http.StatusServiceUnavailable) // this is the Vultr status for rate limits
			w.Write([]byte(`Rate limit hit. API requests are limited to an average of 3/s. Try your request again later.`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})
}

func TestNewClient(t *testing.T) {
	setup()
	defer teardown()

	if client.BaseURL == nil || client.BaseURL.String() != server.URL {
		t.Errorf("NewClient BaseURL = %v, expected %v", client.BaseURL, server.URL)
	}

	if client.APIKey.key != "dummy-key" {
		t.Errorf("NewClient ApiKey = %v, expected %v", client.APIKey.key, "dummy-key")
	}

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

	err := client.DoWithContext(context.Background(), req, data)

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

	err := client.DoWithContext(context.Background(), req, nil)

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
	}, "dummy-key")

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	var panicked string
	func() {
		defer func() {
			if err := recover(); err != nil {
				panicked = fmt.Sprint(err)
			}
		}()
		client.DoWithContext(context.Background(), req, nil)
	}()
	if panicked != "" {
		t.Errorf("unexpected panic: %s", panicked)
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(nil, "dum-dum")

	in := "/unit"
	out := defaultBase + "/unit"

	inRequest := url.Values{
		"balance": {"500"},
	}
	outRequest := `balance=500`

	req, _ := c.NewRequest(ctx, http.MethodPost, in, inRequest)

	if req.URL.String() != out {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", in, req.URL, out)
	}

	body, _ := ioutil.ReadAll(req.Body)

	if string(body) != outRequest {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inRequest, string(body), outRequest)
	}

	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}

	if c.APIKey.key != "dum-dum" {
		t.Errorf("NewRequest() API-Key = %v, expected %v", c.APIKey.key, "dum-dum")
	}

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
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

	err := client.DoWithContext(context.Background(), req, data)
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
