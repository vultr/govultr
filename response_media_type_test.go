package govultr

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

// mediaTypeResponseTransport serves a response without opening a network connection.
type mediaTypeResponseTransport struct {
	contentType   string
	body          string
	contentLength int64
	calls         int
}

// RoundTrip returns the configured response to the public client.
func (tr *mediaTypeResponseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.calls++
	header := make(http.Header)
	header.Set("Content-Type", tr.contentType)
	return &http.Response{
		StatusCode:    http.StatusOK,
		Header:        header,
		Body:          io.NopCloser(strings.NewReader(tr.body)),
		ContentLength: tr.contentLength,
		Request:       req,
	}, nil
}

// TestAccountGetJSONMediaTypes covers valid JSON media types through the public service.
func TestAccountGetJSONMediaTypes(t *testing.T) {
	const body = `{"account":{"name":"demo","balance":12.5}}`
	const balance = 12.5
	for _, test := range []struct {
		name   string
		header string
	}{
		{"bare", "application/json"},
		{"charset", "application/json; charset=utf-8"},
		{"quoted_parameter", `application/json; charset="utf-8"`},
		{"mixed_case", "Application/JSON"},
		{"whitespace", " application/json ; charset=utf-8 "},
		{"multiple_parameters", `application/json; charset=utf-8; profile="example"`},
	} {
		for _, length := range []struct {
			name  string
			value int64
		}{
			{"known_length", int64(len(body))},
			{"unknown_length", -1},
		} {
			t.Run(test.name+"/"+length.name, func(t *testing.T) {
				transport := &mediaTypeResponseTransport{
					contentType: test.header, body: body, contentLength: length.value,
				}
				c := NewClient(&http.Client{Transport: transport})
				account, resp, err := c.Account.Get(context.Background())
				if err != nil {
					t.Fatal(err)
				}
				if transport.calls != 1 {
					t.Fatalf("requests = %d, want 1", transport.calls)
				}
				if resp == nil || account == nil || account.Name != "demo" || account.Balance != balance {
					t.Fatalf("account = %#v, response = %#v; want the decoded account", account, resp)
				}
			})
		}
	}
}

// TestResponseMediaTypePreservesOtherFormats keeps unsupported and invalid headers unchanged.
func TestResponseMediaTypePreservesOtherFormats(t *testing.T) {
	const body = `{"bird":"vultr"}`
	for _, test := range []struct {
		name   string
		header string
	}{
		{"absent", ""},
		{"yaml", "application/yaml"},
		{"json_suffix", "application/problem+json"},
		{"json_prefix", "application/jsonp"},
		{"invalid_parameter", "application/json; charset"},
		{"conflicting_parameters", "application/json; charset=utf-8; charset=ascii"},
		{"invalid_media_type", "application/json extra"},
	} {
		t.Run(test.name, func(t *testing.T) {
			transport := &mediaTypeResponseTransport{
				contentType: test.header, body: body, contentLength: int64(len(body)),
			}
			c := NewClient(&http.Client{Transport: transport})
			req, err := c.NewRequest(context.Background(), http.MethodGet, "/unit", nil)
			if err != nil {
				t.Fatal(err)
			}
			data := struct {
				Bird string `json:"bird"`
			}{Bird: "unchanged"}
			resp, err := c.DoWithContext(context.Background(), req, &data)
			if err != nil {
				t.Fatal(err)
			}
			if data.Bird != "unchanged" || transport.calls != 1 {
				t.Fatalf("data = %#v, requests = %d; want unchanged data and one request", data, transport.calls)
			}
			raw, err := io.ReadAll(resp.Body)
			if err != nil || string(raw) != body {
				t.Fatalf("response body = %q, error = %v; want the original body", raw, err)
			}
		})
	}
}

// TestResponseJSONMediaTypeDecodeErrors preserves JSON decoder errors with parameters.
func TestResponseJSONMediaTypeDecodeErrors(t *testing.T) {
	const body = `{"bird":123}`
	for _, header := range []string{"application/json", "application/json; charset=utf-8"} {
		t.Run(header, func(t *testing.T) {
			transport := &mediaTypeResponseTransport{
				contentType: header, body: body, contentLength: int64(len(body)),
			}
			c := NewClient(&http.Client{Transport: transport})
			req, err := c.NewRequest(context.Background(), http.MethodGet, "/unit", nil)
			if err != nil {
				t.Fatal(err)
			}
			var data struct {
				Bird string `json:"bird"`
			}
			if _, err := c.DoWithContext(context.Background(), req, &data); err == nil {
				t.Fatal("expected a JSON type error")
			}
		})
	}
}

// TestResponseJSONMediaTypeGuards keeps the nil-target and zero-length guards.
func TestResponseJSONMediaTypeGuards(t *testing.T) {
	const body = `{"bird":"vultr"}`
	for _, test := range []struct {
		name    string
		body    string
		length  int64
		nilData bool
	}{
		{"nil_destination", body, int64(len(body)), true},
		{"zero_length_guard", body, 0, false},
		{"empty_body", "", 0, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			transport := &mediaTypeResponseTransport{
				contentType: "application/json; charset=utf-8", body: test.body, contentLength: test.length,
			}
			c := NewClient(&http.Client{Transport: transport})
			req, err := c.NewRequest(context.Background(), http.MethodGet, "/unit", nil)
			if err != nil {
				t.Fatal(err)
			}
			data := struct {
				Bird string `json:"bird"`
			}{Bird: "unchanged"}
			var target interface{} = &data
			if test.nilData {
				target = nil
			}
			if _, err := c.DoWithContext(context.Background(), req, target); err != nil {
				t.Fatal(err)
			}
			if data.Bird != "unchanged" || transport.calls != 1 {
				t.Fatalf("data = %#v, requests = %d; want unchanged data and one request", data, transport.calls)
			}
		})
	}
}
