// Package govultr contains the functionality to interact with the Vultr public
// HTTP REST API.
package govultr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	version           = "3.33.2"
	defaultBase       = "https://api.vultr.com"
	userAgent         = "govultr/" + version
	rateLimit         = 500 * time.Millisecond
	rateLimitRetryMax = 3
	retryLimit        = 3
)

// RequestBody is used to create JSON bodies for one off calls
type RequestBody map[string]interface{}

// apiError is a generic JSON body found in some server error API responses
type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

// Client manages interaction with the Vultr API
type Client struct {
	// Http Client used to interact with the Vultr API
	client *retryablehttp.Client

	// BASE URL for APIs
	BaseURL *url.URL

	// User Agent for the client
	UserAgent string

	// Services used to interact with the API
	Account                  AccountService
	Application              ApplicationService
	Backup                   BackupService
	BareMetalServer          BareMetalServerService
	Billing                  BillingService
	BlockStorage             BlockStorageService
	CDN                      CDNService
	ContainerRegistry        ContainerRegistryService
	Database                 DatabaseService
	Domain                   DomainService
	DomainRecord             DomainRecordService
	FirewallGroup            FirewallGroupService
	FirewallRule             FireWallRuleService
	Instance                 InstanceService
	ISO                      ISOService
	Kubernetes               KubernetesService
	LoadBalancer             LoadBalancerService
	Logs                     LogsService
	Marketplace              MarketplaceService
	ObjectStorage            ObjectStorageService
	OIDC                     OIDCService
	Organization             OrganizationService
	OS                       OSService
	Plan                     PlanService
	Region                   RegionService
	ReservedIP               ReservedIPService
	Inference                InferenceService
	Snapshot                 SnapshotService
	SSHKey                   SSHKeyService
	StartupScript            StartupScriptService
	SubAccount               SubAccountService
	User                     UserService
	VirtualFileSystemStorage VirtualFileSystemStorageService
	VPC                      VPCService

	// Optional function called after every successful request made to the Vultr API
	onRequestCompleted RequestCompletionCallback

	// Tracks when the API rate limit bucket has refilled enough to send
	// another request. Guarded by rateLimitMu.
	rateLimitMu      sync.Mutex
	rateLimitedUntil time.Time
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// RateLimit holds the values reported by the API via the X-RateLimit-*
// response headers.
type RateLimit struct {
	// Limit is the maximum number of requests the bucket can hold (X-RateLimit-Limit).
	Limit int
	// Remaining is the number of whole tokens left in the bucket (X-RateLimit-Remaining).
	Remaining int
	// Reset is the number of seconds until the bucket is fully refilled (X-RateLimit-Reset).
	Reset int
	// OverLimit indicates the request was made over the limit (X-RateLimit-OverLimit).
	OverLimit bool
}

// parseRateLimit extracts the rate limit headers from an API response.
// Returns nil if no rate limit headers are present.
func parseRateLimit(res *http.Response) *RateLimit {
	if res == nil {
		return nil
	}

	if res.Header.Get("X-RateLimit-Limit") == "" &&
		res.Header.Get("X-RateLimit-Remaining") == "" {
		return nil
	}

	rl := &RateLimit{}
	rl.Limit, _ = strconv.Atoi(res.Header.Get("X-RateLimit-Limit"))
	rl.Remaining, _ = strconv.Atoi(res.Header.Get("X-RateLimit-Remaining"))
	rl.Reset, _ = strconv.Atoi(res.Header.Get("X-RateLimit-Reset"))
	rl.OverLimit = res.Header.Get("X-RateLimit-OverLimit") == "1"

	return rl
}

// NewClient returns a Vultr API Client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   90 * time.Second,
					KeepAlive: 90 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   30 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   -1,
				DisableKeepAlives:     true,
			},
			Timeout: 60 * time.Second,
		}
	}

	baseURL, _ := url.Parse(defaultBase)

	client := &Client{
		client:    retryablehttp.NewClient(),
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	client.client.HTTPClient = httpClient
	client.client.Logger = nil
	client.client.ErrorHandler = client.vultrErrorHandler
	client.SetRetryLimit(retryLimit)
	client.SetRateLimit(rateLimit)

	client.Account = &AccountServiceHandler{client}
	client.Application = &ApplicationServiceHandler{client}
	client.Backup = &BackupServiceHandler{client}
	client.BareMetalServer = &BareMetalServerServiceHandler{client}
	client.Billing = &BillingServiceHandler{client}
	client.BlockStorage = &BlockStorageServiceHandler{client}
	client.ContainerRegistry = &ContainerRegistryServiceHandler{client}
	client.CDN = &CDNServiceHandler{client}
	client.Database = &DatabaseServiceHandler{client}
	client.Domain = &DomainServiceHandler{client}
	client.DomainRecord = &DomainRecordsServiceHandler{client}
	client.FirewallGroup = &FireWallGroupServiceHandler{client}
	client.FirewallRule = &FireWallRuleServiceHandler{client}
	client.Instance = &InstanceServiceHandler{client}
	client.ISO = &ISOServiceHandler{client}
	client.Kubernetes = &KubernetesHandler{client}
	client.LoadBalancer = &LoadBalancerHandler{client}
	client.Logs = &LogsServiceHandler{client}
	client.Marketplace = &MarketplaceServiceHandler{client}
	client.ObjectStorage = &ObjectStorageServiceHandler{client}
	client.OIDC = &OIDCServiceHandler{client}
	client.Organization = &OrganizationServiceHandler{client}
	client.OS = &OSServiceHandler{client}
	client.Plan = &PlanServiceHandler{client}
	client.Region = &RegionServiceHandler{client}
	client.ReservedIP = &ReservedIPServiceHandler{client}
	client.Inference = &InferenceServiceHandler{client}
	client.Snapshot = &SnapshotServiceHandler{client}
	client.SSHKey = &SSHKeyServiceHandler{client}
	client.StartupScript = &StartupScriptServiceHandler{client}
	client.SubAccount = &SubAccountServiceHandler{client}
	client.User = &UserServiceHandler{client}
	client.VirtualFileSystemStorage = &VirtualFileSystemStorageServiceHandler{client}
	client.VPC = &VPCServiceHandler{client}

	return client
}

// NewRequest creates an API request
func (c *Client) NewRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	resolvedURL, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if !isNilInterface(body) {
		if err2 := json.NewEncoder(buf).Encode(body); err2 != nil {
			return nil, err2
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, resolvedURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// DoWithContext sends an API request and returns back the response. The API
// response is checked to see if it was a successful call. A successful call is
// then checked to see if we need to unmarshal since some resources have their
// own implementations of unmarshal.
func (c *Client) DoWithContext(ctx context.Context, r *http.Request, data interface{}) (*http.Response, error) {
	var err error

	var res *http.Response
	var errDo error

	for attempt := 0; ; attempt++ {
		// Honor the rate limit reported by the API : if a previous response
		// told us the bucket was exhausted, block here until it refills.
		if err := c.waitForRateLimit(ctx); err != nil {
			return nil, err
		}

		// Rewind the request body so it can be replayed on retries.
		if r.GetBody != nil {
			if body, gerr := r.GetBody(); gerr == nil {
				r.Body = body
			}
		}

		retriable, ferr := retryablehttp.FromRequest(r)
		if ferr != nil {
			return nil, ferr
		}
		retriable = retriable.WithContext(ctx)

		res, errDo = c.client.Do(retriable)
		if errDo != nil {
			break
		}

		// Track the bucket state so subsequent requests wait instead of fail.
		if rl := parseRateLimit(res); rl != nil && (rl.OverLimit || rl.Remaining == 0) && rl.Reset > 0 {
			c.setRateLimitedUntil(rl.Reset)
		}

		// The API rate limited us (429) : wait for the bucket to refill and retry.
		if res.StatusCode == http.StatusTooManyRequests && attempt < rateLimitRetryMax {
			if err := c.waitForRateLimit(ctx); err != nil {
				return nil, err
			}
			continue
		}
		break
	}

	defer func() {
		err = errors.Join(err, closeResponseBody(res))
	}()

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(r, res)
	}

	if errDo != nil {
		return nil, errDo
	}

	body, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}

	if errClose := res.Body.Close(); errClose != nil {
		return nil, errClose
	}

	// replace response body with in-memory bytes
	res.Body = io.NopCloser(bytes.NewReader(body))

	// the api can return a json error response so check that before continuing
	if len(body) > 0 {
		var responseError apiError
		if err := json.Unmarshal(body, &responseError); err != nil {
			return nil, fmt.Errorf("error unmarshaling response body : %v", err)
		}

		if responseError.Message != "" {
			return nil, fmt.Errorf("error returned from api : %v", responseError.Message)
		}
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusPartialContent:
		if res.ContentLength == 0 || isNilInterface(data) {
			return res, nil
		}

		switch res.Header.Get("Content-Type") { //nolint:gocritic
		case "application/json":
			if err := json.Unmarshal(body, data); err != nil {
				return nil, err
			}
		}

		return res, nil

	case http.StatusNoContent, http.StatusResetContent:
		return res, nil

	default:
		return nil, errors.New(string(body))
	}
}

// waitForRateLimit blocks until the API rate limit bucket has refilled, or
// until the context is canceled. It is a no-op if we are not rate limited.
func (c *Client) waitForRateLimit(ctx context.Context) error {
	c.rateLimitMu.Lock()
	until := c.rateLimitedUntil
	c.rateLimitMu.Unlock()

	wait := time.Until(until)
	if wait <= 0 {
		return nil
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(wait):
		return nil
	}
}

// setRateLimitedUntil records when the rate limit bucket will be refilled,
// based on the X-RateLimit-Reset header, so subsequent requests wait instead
// of failing.
func (c *Client) setRateLimitedUntil(resetSeconds int) {
	if resetSeconds <= 0 {
		return
	}

	until := time.Now().Add(time.Duration(resetSeconds) * time.Second)

	c.rateLimitMu.Lock()
	defer c.rateLimitMu.Unlock()

	if until.After(c.rateLimitedUntil) {
		c.rateLimitedUntil = until
	}
}

func closeResponseBody(res *http.Response) error {
	if res == nil {
		return fmt.Errorf("response is invalid, can not close")
	}

	if _, err := res.Body.Read(nil); err != http.ErrBodyReadAfterClose {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return fmt.Errorf("error closing response body : %w", err)
	}

	return nil
}

// SetBaseURL Overrides the default BaseUrl
func (c *Client) SetBaseURL(baseURL string) error {
	updatedURL, err := url.Parse(baseURL)

	if err != nil {
		return err
	}

	c.BaseURL = updatedURL
	return nil
}

// SetRateLimit Overrides the default rateLimit. For performance, exponential
// backoff is used with the minimum wait being 2/3rds the time provided.
func (c *Client) SetRateLimit(t time.Duration) {
	c.client.RetryWaitMin = t / 3 * 2
	c.client.RetryWaitMax = t
}

// SetUserAgent Overrides the default UserAgent
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua
}

// OnRequestCompleted sets the API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// SetRetryLimit overrides the default RetryLimit
func (c *Client) SetRetryLimit(n int) {
	c.client.RetryMax = n
}

func (c *Client) vultrErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if resp == nil {
		if err != nil {
			return nil, fmt.Errorf("gave up after %d attempts, last error : %s", numTries, err.Error())
		}
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (resp == nil)", numTries)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (error reading response body: %v)", numTries, err)
	}
	return nil, fmt.Errorf("gave up after %d attempts, last error: %#v", numTries, strings.TrimSpace(string(buf)))
}

func isNilInterface(i interface{}) bool {
	v := reflect.ValueOf(i)

	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// BoolToBoolPtr helper function that returns a pointer from your bool value
func BoolToBoolPtr(value bool) *bool {
	return &value
}

// StringToStringPtr helper function that returns a pointer from your string value
func StringToStringPtr(value string) *string {
	return &value
}

// IntToIntPtr helper function that returns a pointer from your string value
func IntToIntPtr(value int) *int {
	return &value
}
