package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vpc2Path = "/v2/vpc2"

// VPC2Service is the interface to interact with the VPC 2.0 endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/vpc2
type VPC2Service interface { //nolint:dupl
	Create(ctx context.Context, createReq *VPC2Req) (*VPC2, *http.Response, error)
	Get(ctx context.Context, vpcID string) (*VPC2, *http.Response, error)
	Update(ctx context.Context, vpcID string, description string) error
	Delete(ctx context.Context, vpcID string) error
	List(ctx context.Context, options *ListOptions) ([]VPC2, *Meta, *http.Response, error)
}

// VPC2ServiceHandler handles interaction with the VPC 2.0 methods for the Vultr API
type VPC2ServiceHandler struct {
	client *Client
}

// VPC2 represents a Vultr VPC 2.0
type VPC2 struct {
	ID           string `json:"id"`
	Region       string `json:"region"`
	Description  string `json:"description"`
	IPBlock      string `json:"ip_block"`
	PrefixLength int    `json:"prefix_length"`
	DateCreated  string `json:"date_created"`
}

// VPC2Req represents parameters to create or update a VPC 2.0 resource
type VPC2Req struct {
	Region       string `json:"region"`
	Description  string `json:"description"`
	IPType       string `json:"ip_type"`
	IPBlock      string `json:"ip_block"`
	PrefixLength int    `json:"prefix_length"`
}

type vpcs2Base struct {
	VPCs []VPC2 `json:"vpcs"`
	Meta *Meta  `json:"meta"`
}

type vpc2Base struct {
	VPC *VPC2 `json:"vpc"`
}

// Create creates a new VPC 2.0. A VPC 2.0 can only be used at the location for which it was created.
func (n *VPC2ServiceHandler) Create(ctx context.Context, createReq *VPC2Req) (*VPC2, *http.Response, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, vpc2Path, createReq)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpc2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Get gets the VPC 2.0 of the requested ID
func (n *VPC2ServiceHandler) Get(ctx context.Context, vpcID string) (*VPC2, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpc2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Update updates a VPC 2.0
func (n *VPC2ServiceHandler) Update(ctx context.Context, vpcID, description string) error {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)

	vpcReq := RequestBody{"description": description}
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, vpcReq)
	if err != nil {
		return err
	}

	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete deletes a VPC 2.0. Before deleting, a VPC 2.0 must be disabled from all instances
func (n *VPC2ServiceHandler) Delete(ctx context.Context, vpcID string) error {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// List lists all VPCs 2.0 on the current account
func (n *VPC2ServiceHandler) List(ctx context.Context, options *ListOptions) ([]VPC2, *Meta, *http.Response, error) { //nolint:dupl
	req, err := n.client.NewRequest(ctx, http.MethodGet, vpc2Path, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vpcs := new(vpcs2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, nil, resp, err
	}

	return vpcs.VPCs, vpcs.Meta, resp, nil
}
