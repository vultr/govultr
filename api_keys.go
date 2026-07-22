package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// APIKeyService is the interface to interact with the API Keys endpoint on the Vultr API
// Link: https://www.vultr.com/api/#tag/api-keys
type APIKeyService interface {
	Create(ctx context.Context, apiReq *APIKeyCreate) (*APIKey, *http.Response, error)
	Get(ctx context.Context, apiKeyID string) (*APIKey, *http.Response, error)
	Delete(ctx context.Context, apiKeyID string) error
	List(ctx context.Context, options *ListOptions) ([]APIKey, *http.Response, error)
}

// APIKeyServiceHandler handles interaction with the API key methods for the Vultr API
type APIKeyServiceHandler struct {
	client *Client
}

// APIKey represents a Vultr API key
type APIKey struct {
	ID         string `json:"id"`
	APIKey     string `json:"api_key"`
	Name       string `json:"name"`
	Expire     bool   `json:"expire"`
	DateExpire string `json:"date_expire"`
}

// APIKeyCreate struct is used for creating API keys.
type APIKeyCreate struct {
	Name       string `json:"name,omitempty"`
	Expire     bool   `json:"expire,omitempty"`
	DateExpire string `json:"date_expire,omitempty"`
}

type apiKeyBase struct {
	APIKey *APIKey `json:"api_key"`
}

type apiKeysBase struct {
	APIKeys []APIKey `json:"api_keys"`
}

// Create adds an API key to the currently authenticated user's API key list
func (b *APIKeyServiceHandler) Create(ctx context.Context, apiReq *APIKeyCreate) (*APIKey, *http.Response, error) {
	uri := "/v2/apikeys"

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, apiReq)
	if err != nil {
		return nil, nil, err
	}

	apiKey := new(apiKeyBase)
	resp, err := b.client.DoWithContext(ctx, req, apiKey)
	if err != nil {
		return nil, resp, err
	}

	return apiKey.APIKey, resp, nil
}

// Get returns a single API key instance for the currently authenticated user based on the apiKeyID you provide
func (b *APIKeyServiceHandler) Get(ctx context.Context, apiKeyID string) (*APIKey, *http.Response, error) {
	uri := fmt.Sprintf("/v2/apikeys/%s", apiKeyID)

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	apiKey := new(apiKeyBase)
	resp, err := b.client.DoWithContext(ctx, req, apiKey)
	if err != nil {
		return nil, resp, err
	}

	return apiKey.APIKey, resp, nil
}

// Delete an API key from your Vultr account
func (b *APIKeyServiceHandler) Delete(ctx context.Context, apiKeyID string) error {
	uri := fmt.Sprintf("/v2/apikeys/%s", apiKeyID)

	req, err := b.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// List returns a list of all API key instances for the currently authenticated user
func (b *APIKeyServiceHandler) List(ctx context.Context, options *ListOptions) ([]APIKey, *http.Response, error) { //nolint:dupl
	uri := "/v2/apikeys"

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	apiKeys := new(apiKeysBase)
	resp, err := b.client.DoWithContext(ctx, req, apiKeys)
	if err != nil {
		return nil, resp, err
	}

	return apiKeys.APIKeys, resp, nil
}
