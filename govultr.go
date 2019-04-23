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
	defaultBase = "https://api.vultr.com"
	userAgent   = "govultr/" + version
	rateLimit   = 200 * time.Millisecond
)

// APIKey contains a users API Key for interacting with the API
type APIKey struct {
	// API Key
	key string
}

// Client manages interaction with the Vultr V1 API
type Client struct {
	// Http Client used to interact with the Vultr V1 API
	client *http.Client

	// BASE URL for APIs
	baseURL *url.URL

	// User Agent for the client
	UserAgent string

	// API Key
	APIKey APIKey

	// API Rate Limit - Vultr rate limits based on time
	RateLimit time.Duration

	// Services used to interact with the API
	Account     AccountService
	API         APIService
	Application ApplicationService
	Backup      BackupService
	OS          OSService
	Snapshot    SnapshotService

	// Optional function called after every successful request made to the Vultr API
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a Vultr API Client
func NewClient(httpClient *http.Client, key string) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBase)

	client := &Client{
		client:    httpClient,
		baseURL:   baseURL,
		UserAgent: userAgent,
		RateLimit: rateLimit,
	}

	client.Account = &AccountServiceHandler{client}
	client.API = &APIServiceHandler{client}
	client.Application = &ApplicationServiceHandler{client}
	client.Backup = &BackupServiceHandler{client}
	client.OS = &OSServiceHandler{client}
	client.Snapshot = &SnapshotServiceHandler{client}

	apiKey := APIKey{key: key}
	client.APIKey = apiKey

	return client
}

// NewRequest creates an API Request
func (c *Client) NewRequest(ctx context.Context, method, uri string, body url.Values) (*http.Request, error) {

	path, err := url.Parse(uri)
	resolvedURL := c.baseURL.ResolveReference(path)

	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	if body != nil {
		reqBody = strings.NewReader(body.Encode())
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(method, resolvedURL.String(), reqBody)

	if err != nil {
		return nil, err
	}

	req.Header.Add("API-key", c.APIKey.key)
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
		}
		return nil
	}

	return errors.New(string(body))
}

// SetBaseURL Overrides the default BaseUrl
func (c *Client) SetBaseURL(baseURL string) error {
	updatedURL, err := url.Parse(baseURL)

	if err != nil {
		return err
	}

	c.baseURL = updatedURL
	return nil
}

// SetRateLimit Overrides the default rateLimit
func (c *Client) SetRateLimit(time time.Duration) {
	c.RateLimit = time
}

// SetUserAgent Overrides the default UserAgent
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua
}

// OnRequestCompleted sets the API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
