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
	"time"
)

const (
	version     = "0.0.1"
	DefaultBase = "https://api.vultr.com"
	UserAgent   = "govultr/" + version
	RateLimit   = 200 * time.Millisecond
)

// ApiKey contains a users API Key for interacting with the API
type ApiKey struct {
	// API Key
	key string
}

// Client manages interaction with the Vultr V1 API
type Client struct {
	// Http Client used to interact with the Vultr V1 API
	client *http.Client

	// BASE URL for APIs
	BaseUrl *url.URL

	// User Agent for the client
	UserAgent string

	// API Key
	ApiKey ApiKey

	// API Rate Limit - Vultr rate limits based on time
	RateLimit time.Duration

	// Services used to interact with the API
	Account AccountService

	// Optional function called after every successful request made to the Vultr API
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// New Client returns a Vultr API Client
func NewClient(httpClient *http.Client, key string) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(DefaultBase)

	client := &Client{
		client:    httpClient,
		BaseUrl:   baseUrl,
		UserAgent: UserAgent,
		RateLimit: RateLimit,
	}

	client.Account = &AccountServiceHandler{client}

	apiKey := ApiKey{key: key}
	client.ApiKey = apiKey

	return client
}

// NewRequest creates an API Request
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

// DoWithContext sends an API Request and returns back the response. The API response is checked  to see if it was
// a successful call. A successful call is then checked to see if we need to unmarshal since some resources
// have their own implements of unmarshal.
func (c *Client) DoWithContext(ctx context.Context, r *http.Request, data interface{}) error {

	// Sleep this call
	time.Sleep(c.RateLimit)

	req := r.WithContext(ctx)
	res, err := c.client.Do(req)

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, res)
	}

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

// Overrides the default BaseUrl
func (c *Client) SetBaseUrl(baseUrl string) error {
	updatedUrl, err := url.Parse(baseUrl)

	if err != nil {
		return err
	}

	c.BaseUrl = updatedUrl
	return nil
}

// Overrides the default rateLimit
func (c *Client) SetRateLimit(time time.Duration) {
	c.RateLimit = time
}

// OnRequestCompleted sets the API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
