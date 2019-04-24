package govultr

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// BlockStorageService is the interface to interact with Block-Storage endpoint on the Vultr API
// Link: https://www.vultr.com/api/#block
type BlockStorageService interface {
	Attach(ctx context.Context, blockID, vpsID string) error
	Create(ctx context.Context, regionID, size int, label string) (*BlockStorage, error)
	Delete(ctx context.Context, blockID string) error
	Detach(ctx context.Context, blockID string) error
	SetLabel(ctx context.Context, blockID, label string) error
	//GetList(ctx context.Context, blockID string) ([]BlockStorage, error)
	//Resize(ctx context.Context, blockID string, size int) error
}

// BlockStorageServiceHandler handles interaction with the block-storage methods for the Vultr API
type BlockStorageServiceHandler struct {
	client *Client
}

// BlockStorage represents Vultr Block-Storage
type BlockStorage struct {
	BlockStorageID string `json:"SUBID"`
	DateCreated    string `json:"date_created"`
	Cost           string `json:"cost_per_month"`
	Status         string `json:"status"`
	Size           int    `json:"size_gb"`
	RegionID       int    `json:"region_id"`
	VpsID          string `json:"attached_to_SUBID"`
	Label          string `json:"label"`
}

// Attach will link a given block storage to a given Vultr vps
func (b *BlockStorageServiceHandler) Attach(ctx context.Context, blockID, vpsID string) error {

	uri := "/v1/block/attach"

	values := url.Values{
		"SUBID":           {blockID},
		"attach_to_SUBID": {vpsID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// Create builds out a block storage
func (b *BlockStorageServiceHandler) Create(ctx context.Context, regionID, size int, label string) (*BlockStorage, error) {

	uri := "/v1/block/create"

	values := url.Values{
		"DCID":    {strconv.Itoa(regionID)},
		"size_gb": {strconv.Itoa(size)},
		"label":   {label},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return nil, err
	}

	blockStorage := new(BlockStorage)

	err = b.client.DoWithContext(ctx, req, blockStorage)

	if err != nil {
		return nil, err
	}

	blockStorage.RegionID = regionID
	blockStorage.Label = label
	blockStorage.Size = size

	return blockStorage, nil
}

// Delete will remove block storage instance from your Vultr account
func (b *BlockStorageServiceHandler) Delete(ctx context.Context, blockID string) error {

	uri := "/v1/block/delete"

	values := url.Values{
		"SUBID": {blockID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// Detach will de-link a given block storage to the Vultr vps it is attached to
func (b *BlockStorageServiceHandler) Detach(ctx context.Context, blockID string) error {

	uri := "/v1/block/detach"

	values := url.Values{
		"SUBID": {blockID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// SetLabel allows you to set/update the label on your Vultr Block storage
func (b *BlockStorageServiceHandler) SetLabel(ctx context.Context, blockID, label string) error {
	uri := "/v1/block/label_set"

	values := url.Values{
		"SUBID": {blockID},
		"label": {label},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}
