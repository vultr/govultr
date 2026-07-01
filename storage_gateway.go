package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const storageGatewaysPath = "/v2/storage-gateways"

// StorageGatewayService is the interface to interact with the storage gateway endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/storage-gateways
type StorageGatewayService interface {
	Create(ctx context.Context, options *StorageGatewayCreateReq) (*StorageGateway, *http.Response, error)
	Get(ctx context.Context, storageGatewayID string) (*StorageGateway, *http.Response, error)
	Update(ctx context.Context, storageGatewayID string, options *StorageGatewayUpdateReq) (*StorageGateway, *http.Response, error)
	List(ctx context.Context, options *ListOptions) ([]StorageGateway, *Meta, *http.Response, error)
	Delete(ctx context.Context, storageGatewayID string) error

	CreateExport(ctx context.Context, storageGatewayID string, options *StorageGatewayExportConfig) (*StorageGatewayExportConfig, *http.Response, error) //nolint:lll
	UpdateExport(ctx context.Context, storageGatewayID, exportID string, options *StorageGatewayExportUpdateReq) error
	DeleteExport(ctx context.Context, storageGatewayID, exportID string) error
}

// StorageGatewayServiceHandler handles interaction with the storage gateway methods for the Vultr API
type StorageGatewayServiceHandler struct {
	client *Client
}

// StorageGateway represents a Vultr storage gateway
type StorageGateway struct {
	ID             string                       `json:"id"`
	DateCreated    string                       `json:"date_created"`
	Status         string                       `json:"status"`
	Type           string                       `json:"type"`
	Region         string                       `json:"region"`
	Label          string                       `json:"label"`
	PendingCharges float32                      `json:"pending_charges"`
	Tags           []string                     `json:"tags"`
	Health         string                       `json:"health"`
	NetworkConfig  *StorageGatewayNetworkConfig `json:"network_config"`
	ExportConfig   *StorageGatewayExportConfig  `json:"export_config"`
}

// storageGatewayBase represents the base response for creating a storage gateway
type storageGatewayBase struct {
	StorageGateway *StorageGateway `json:"storage_gateway"`
}

// storageGatewaysBase represents the base response for listing storage gateways
type storageGatewaysBase struct {
	StorageGateways []StorageGateway `json:"storage_gateways"`
	Meta            *Meta            `json:"meta"`
}

// StorageGatewayNetworkConfig represents the network configuration for a Vultr storage gateway
type StorageGatewayNetworkConfig struct {
	Primary *StorageGatewayNetworkConfigPrimary `json:"primary,omitempty"`
}

// StorageGatewayNetworkConfigPrimary represents the primary network configuration for a Vultr storage gateway
type StorageGatewayNetworkConfigPrimary struct {
	IPv4PublicEnabled bool                            `json:"ipv4_public_enabled,omitempty"`
	IPv6PublicEnabled bool                            `json:"ipv6_public_enabled,omitempty"`
	VPC               *StorageGatewayNetworkConfigVPC `json:"vpc,omitempty"`
}

// StorageGatewayNetworkConfigVPC represents the VPC configuration for a Vultr storage gateway
type StorageGatewayNetworkConfigVPC struct {
	IPAddress   string `json:"vpc_ip_address,omitempty"`
	UUID        string `json:"vpc_uuid,omitempty"`
	Description string `json:"vpc_description,omitempty"`
}

// StorageGatewayExportConfig represents the export configuration for a Vultr storage gateway
type StorageGatewayExportConfig struct {
	Label          string   `json:"label,omitempty"`
	VfsUuid        string   `json:"vfs_uuid,omitempty"`
	PseudoRootPath string   `json:"pseudo_root_path,omitempty"`
	AllowedIPs     []string `json:"allowed_ips,omitempty"`
}

// StorageGatewayCreateReq is used for creating a new storage gateway
type StorageGatewayCreateReq struct {
	Label         string                      `json:"label"`
	Type          string                      `json:"type"`
	Region        string                      `json:"region"`
	ExportConfig  StorageGatewayExportConfig  `json:"export_config"`
	NetworkConfig StorageGatewayNetworkConfig `json:"network_config"`
}

// StorageGatewayUpdateReq is used for updating an existing storage gateway
type StorageGatewayUpdateReq struct {
	Label string `url:"label"`
}

// StorageGatewayExportUpdateReq represents the request body for updating a storage gateway export
type StorageGatewayExportUpdateReq struct {
	AllowedIPs []string `json:"allowed_ips"`
}

// Create a new storage gateway
func (s StorageGatewayServiceHandler) Create(ctx context.Context, options *StorageGatewayCreateReq) (*StorageGateway, *http.Response, error) { //nolint:lll
	req, err := s.client.NewRequest(ctx, http.MethodPost, storageGatewaysPath, options)
	if err != nil {
		return nil, nil, err
	}

	storageGateway := new(storageGatewayBase)
	resp, err := s.client.DoWithContext(ctx, req, storageGateway)
	if err != nil {
		return nil, resp, err
	}

	return storageGateway.StorageGateway, resp, nil
}

// Get information from an existing storage gateway
func (s StorageGatewayServiceHandler) Get(ctx context.Context, storageGatewayID string) (*StorageGateway, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", storageGatewaysPath, storageGatewayID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	storageGateway := new(storageGatewayBase)
	resp, err := s.client.DoWithContext(ctx, req, storageGateway)
	if err != nil {
		return nil, resp, err
	}

	return storageGateway.StorageGateway, resp, nil
}

// Update an existing storage gateway
func (s StorageGatewayServiceHandler) Update(ctx context.Context, storageGatewayID string, options *StorageGatewayUpdateReq) (*StorageGateway, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", storageGatewaysPath, storageGatewayID)
	req, err := s.client.NewRequest(ctx, http.MethodPut, uri, options)
	if err != nil {
		return nil, nil, err
	}

	storageGateway := new(storageGatewayBase)
	resp, err := s.client.DoWithContext(ctx, req, storageGateway)
	if err != nil {
		return nil, resp, err
	}

	return storageGateway.StorageGateway, resp, nil
}

// List all storage gateways in your account
func (s StorageGatewayServiceHandler) List(ctx context.Context, options *ListOptions) ([]StorageGateway, *Meta, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, storageGatewaysPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	storageGateways := new(storageGatewaysBase)
	resp, err := s.client.DoWithContext(ctx, req, storageGateways)
	if err != nil {
		return nil, nil, resp, err
	}

	return storageGateways.StorageGateways, storageGateways.Meta, resp, nil
}

// Delete an existing storage gateway
func (s StorageGatewayServiceHandler) Delete(ctx context.Context, storageGatewayID string) error {
	uri := fmt.Sprintf("%s/%s", storageGatewaysPath, storageGatewayID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}

// CreateExport add a new export to an existing storage gateway
func (s StorageGatewayServiceHandler) CreateExport(ctx context.Context, storageGatewayID string, options *StorageGatewayExportConfig) (*StorageGatewayExportConfig, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/exports", storageGatewaysPath, storageGatewayID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, options)
	if err != nil {
		return nil, nil, err
	}

	exportConfig := new(StorageGatewayExportConfig)
	resp, err := s.client.DoWithContext(ctx, req, exportConfig)
	return exportConfig, resp, err
}

// UpdateExport updates an existing storage gateway export
func (s StorageGatewayServiceHandler) UpdateExport(ctx context.Context, storageGatewayID, exportID string, options *StorageGatewayExportUpdateReq) error { //nolint:lll
	uri := fmt.Sprintf("%s/%s/exports/%s", storageGatewaysPath, storageGatewayID, exportID)
	req, err := s.client.NewRequest(ctx, http.MethodPatch, uri, options)
	if err != nil {
		return err
	}

	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}

// DeleteExport deletes an export from an existing storage gateway
func (s StorageGatewayServiceHandler) DeleteExport(ctx context.Context, storageGatewayID, exportID string) error {
	uri := fmt.Sprintf("%s/%s/exports/%s", storageGatewaysPath, storageGatewayID, exportID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}
