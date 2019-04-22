package govultr

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
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
	//todo
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

func TestClient_SetRateLimit(t *testing.T) {
	setup()
	defer teardown()

	ua := "vultr/testing"
	client.SetUserAgent(ua)

	if client.UserAgent != ua {
		t.Errorf("NewClient UserAgent = %v, expected %v", client.UserAgent, userAgent)
	}
}

func TestClient_OnRequestCompleted(t *testing.T) {
	//todo
}
