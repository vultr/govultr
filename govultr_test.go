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
	"testing/iotest"
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
	thisURL, _ := url.Parse(server.URL)
	client.BaseURL = thisURL
}

func teardown() {
	server.Close()
}

// testJSONResponseHandlerFunc is used to build a http handler func for
// mocking endpoint responses with JSON bodies
func testJSONResponseHandlerFunc(statusCode int, body string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, body)
	}
}

// testYAMLResponseHandlerFunc is used to build a http handler func for
// mocking endpoint responses with YAML bodies
func testYAMLResponseHandlerFunc(statusCode int, body string) func(http.ResponseWriter, *http.Request) { //nolint:unused
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/yaml")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, body)
	}
}

func TestNewClient(t *testing.T) {
	setup()
	defer teardown()

	if client.BaseURL == nil || client.BaseURL.String() != server.URL {
		t.Errorf("NewClient BaseURL = %v, expected %v", client.BaseURL, server.URL)
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
		Bird string `json:"bird"`
	}

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if method := http.MethodGet; method != request.Method {
			t.Errorf("Request method = %v, expecting %v", request.Method, method)
		}
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, `{"bird": "vultr"}`)
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
		client.DoWithContext(context.Background(), req, nil)
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

	tTime := 600 * time.Millisecond
	client.SetRateLimit(tTime)

	if client.client.RetryWaitMax != tTime {
		t.Errorf("NewClient max RateLimit = %v, expected %v", client.client.RetryWaitMax, tTime)
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

	mux.HandleFunc("/wrong", testJSONResponseHandlerFunc(http.StatusOK, "{"))

	req, _ := client.NewRequest(ctx, http.MethodGet, "/wrong", nil)
	if _, err := client.DoWithContext(ctx, req, struct{}{}); err == nil {
		t.Error("Expected response body to be invalid")
	}
}

// trackedResponseBody records closure of the underlying transport body.
type trackedResponseBody struct {
	io.Reader
	closes     int
	closeError error
}

func (b *trackedResponseBody) Close() error {
	b.closes++
	b.Reader = iotest.ErrReader(errors.New("read on closed body"))
	return b.closeError
}

// responseBodyTransport supplies a response with an observable body lifetime.
type responseBodyTransport struct {
	status int
	body   *trackedResponseBody
	err    error
}

func (tr responseBodyTransport) RoundTrip(*http.Request) (*http.Response, error) {
	if tr.err != nil {
		return nil, tr.err
	}
	return &http.Response{StatusCode: tr.status, Header: http.Header{"Content-Type": {"application/json"}}, Body: tr.body, ContentLength: -1}, nil
}

func TestDoWithContextClosesOriginalBody(t *testing.T) {
	for _, tc := range []struct {
		name                                       string
		status                                     int
		body                                       string
		noDecode, readError, closeError, wantError bool
	}{
		{name: "decoded_success", status: 200, body: `{"bird":"vultr"}`},
		{name: "without_decode", status: 200, body: `{"bird":"vultr"}`, noDecode: true},
		{name: "no_content", status: 204},
		{name: "api_error", status: 400, body: `{"error":"invalid request"}`, wantError: true},
		{name: "invalid_json", status: 200, body: `{`, wantError: true},
		{name: "read_error", status: 200, readError: true, wantError: true},
		{name: "close_error", status: 200, body: `{"bird":"vultr"}`, closeError: true, wantError: true},
		{name: "read_and_close_error", status: 200, readError: true, closeError: true, wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := &trackedResponseBody{Reader: strings.NewReader(tc.body)}
			readError := errors.New("read failed")
			closeError := errors.New("close failed")
			if tc.closeError {
				body.closeError = closeError
			}
			if tc.readError {
				body.Reader = iotest.ErrReader(readError)
			}
			testClient := NewClient(&http.Client{Transport: responseBodyTransport{status: tc.status, body: body}})
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.invalid/", http.NoBody)
			if err != nil {
				t.Fatal(err)
			}
			var decoded map[string]string
			var target interface{} = &decoded
			if tc.noDecode {
				target = nil
			}
			res, err := testClient.DoWithContext(context.Background(), req, target)
			if (err != nil) != tc.wantError {
				t.Fatalf("error = %v, wantError = %t", err, tc.wantError)
			}
			if tc.readError && !errors.Is(err, readError) {
				t.Errorf("error = %v, want read error", err)
			}
			if tc.closeError && !tc.readError && !errors.Is(err, closeError) {
				t.Errorf("error = %v, want close error", err)
			}
			if body.closes != 1 {
				t.Errorf("original body closed %d times, want 1", body.closes)
			}
			if res != nil {
				defer res.Body.Close()
				retained, err := io.ReadAll(res.Body)
				if err != nil {
					t.Fatal(err)
				}
				if string(retained) != tc.body {
					t.Errorf("retained body = %q, want %q", retained, tc.body)
				}
			}
			if tc.name == "decoded_success" && decoded["bird"] != "vultr" {
				t.Errorf("decoded = %v", decoded)
			}
		})
	}
}

func TestDoWithContextTransportError(t *testing.T) {
	transportError := errors.New("transport failed")
	testClient := NewClient(&http.Client{Transport: responseBodyTransport{err: transportError}})
	testClient.SetRetryLimit(0)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.invalid/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.DoWithContext(context.Background(), req, nil)
	if res != nil || err == nil || !strings.Contains(err.Error(), transportError.Error()) {
		t.Fatalf("response = %v, error = %v, want transport error", res, err)
	}
}
