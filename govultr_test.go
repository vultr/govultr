package govultr

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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
	client.BaseUrl = url
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	setup()
	defer teardown()

	if client.BaseUrl == nil || client.BaseUrl.String() != server.URL {
		t.Errorf("NewClient BaseURL = %v, expected %v", client.BaseUrl, server.URL)
	}

	if client.RateLimit == 0 || client.RateLimit.String() != "200ms" {
		t.Errorf("NewClient RateLimit = %v, expected %v", client.RateLimit, rateLimit.String())
	}

	if client.ApiKey.key != "dummy-key" {
		t.Errorf("NewClient ApiKey = %v, expected %v", client.ApiKey.key, "dummy-key")
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

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(nil, "dum-dum")

	in := "/unit"
	out := defaultBase + "/unit"

	inRequest := url.Values{
		"balance": {"500"},
	}
	outRequest := `balance=500`

	req, _ := c.NewRequest(ctx, http.MethodGet, in, inRequest)

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

	if c.ApiKey.key != "dum-dum" {
		t.Errorf("NewRequest() API-Key = %v, expected %v", c.ApiKey.key, "dum-dum")
	}

}

func TestClient_SetBaseUrl(t *testing.T) {
	setup()
	defer teardown()

	base := "http://localhost/vultr"
	err := client.SetBaseUrl(base)

	if err != nil {
		t.Fatalf("SetBaseUrl unexpected error: %v", err)
	}

	if client.BaseUrl.String() != base {
		t.Errorf("NewClient BaseUrl = %v, expected %v", client.BaseUrl, base)
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

	time := 500 * time.Millisecond
	client.SetRateLimit(time)

	if client.RateLimit != time {
		t.Errorf("NewClient RateLimit = %v, expected %v", client.RateLimit, time)
	}
}

func TestClient_OnRequestCompleted(t *testing.T) {
	//todo
}
