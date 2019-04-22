package govultr

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	version     = "0.0.1"
	DefaultBase = "https://api.vultr.com"
	UserAgent   = "govultr/" + version
)

type ApiKey struct {
	key string
}

type Client struct {
	client *http.Client

	BaseUrl *url.URL

	UserAgent string

	Account AccountService

	ApiKey ApiKey
}

func NewClient(httpClient *http.Client, key string) *Client {

	baseUrl, _ := url.Parse(DefaultBase)

	client := &Client{
		client:    httpClient,
		BaseUrl:   baseUrl,
		UserAgent: UserAgent,
	}

	client.Account = &AccountServiceHandler{client}

	apiKey := ApiKey{key: key}
	client.ApiKey = apiKey

	return client
}

func (c *Client) NewRequest(ctx context.Context, method, uri string, body url.Values) (*http.Request, error) {

	path, err := url.Parse(uri)
	resolvedUrl := c.BaseUrl.ResolveReference(path)

	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	if body != nil {
		strings.NewReader(body.Encode())
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(method, resolvedUrl.String(), reqBody)

	if err != nil {
		return nil, err
	}

	req.Header.Add("API-key", c.ApiKey.key)
	// todo review the Accept and content types
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")

	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

//todo DoWithContext
func (c *Client) DoWithContext(ctx context.Context, r *http.Request, data interface{}) error {

	// todo slow this down with a sleep for the rate limit
	req := r.WithContext(ctx)
	res, err := c.client.Do(req)

	//todo handle the error this might throw
	defer res.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		if data != nil {
			if string(body) == "[]" {
				data = nil
			} else {
				if err := json.Unmarshal(body, data); err != nil {
					return err
				}
			}
			return nil
		}
	}

	return errors.New(string(body))
}
