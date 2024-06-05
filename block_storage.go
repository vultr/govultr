package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// BlockStorageService is the interface to interact with Block-Storage endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/block
type BlockStorageService interface {
	Create(ctx context.Context, blockReq *BlockStorageCreate) (*BlockStorage, *http.Response, error)
	Get(ctx context.Context, blockID string) (*BlockStorage, *http.Response, error)
	Update(ctx context.Context, blockID string, blockReq *BlockStorageUpdate) error
	Delete(ctx context.Context, blockID string) error
	List(ctx context.Context, options *ListOptions) ([]BlockStorage, *Meta, *http.Response, error)

	Attach(ctx context.Context, blockID string, attach *BlockStorageAttach) error
	Detach(ctx context.Context, blockID string, detach *BlockStorageDetach) error
}

// BlockStorageServiceHandler handles interaction with the block-storage methods for the Vultr API
type BlockStorageServiceHandler struct {
	client *Client
}

// BlockStorage represents Vultr Block-Storage
type BlockStorage struct {
	ID                 string  `json:"id"`
	Status             string  `json:"status"`
	Region             string  `json:"region"`
	DateCreated        string  `json:"date_created"`
	AttachedToInstance string  `json:"attached_to_instance"`
	Label              string  `json:"label"`
	MountID            string  `json:"mount_id"`
	BlockType          string  `json:"block_type"`
	SizeGB             int     `json:"size_gb"`
	Cost               float32 `json:"cost"`
}

// BlockStorageCreate struct is used for creating Block Storage.
type BlockStorageCreate struct {
	Region    string `json:"region"`
	Label     string `json:"label,omitempty"`
	BlockType string `json:"block_type,omitempty"`
	SizeGB    int    `json:"size_gb"`
}

// BlockStorageUpdate struct is used to update Block Storage.
type BlockStorageUpdate struct {
	Label  string `json:"label,omitempty"`
	SizeGB int    `json:"size_gb,omitempty"`
}

// BlockStorageAttach struct used to define if a attach should be restart the instance.
type BlockStorageAttach struct {
	Live       *bool  `json:"live,omitempty"`
	InstanceID string `json:"instance_id"`
}

// BlockStorageDetach struct used to define if a detach should be restart the instance.
type BlockStorageDetach struct {
	Live *bool `json:"live,omitempty"`
}

type blockStoragesBase struct {
	Meta   *Meta          `json:"meta"`
	Blocks []BlockStorage `json:"blocks"`
}

type blockStorageBase struct {
	Block *BlockStorage `json:"block"`
}

// Create builds out a block storage
func (b *BlockStorageServiceHandler) Create(ctx context.Context, blockReq *BlockStorageCreate) (*BlockStorage, *http.Response, error) {
	uri := "/v2/blocks"

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, blockReq)
	if err != nil {
		return nil, nil, err
	}

	block := new(blockStorageBase)
	resp, err := b.client.DoWithContext(ctx, req, block)
	if err != nil {
		return nil, resp, err
	}

	return block.Block, resp, nil
}

// Get returns a single block storage instance based ony our blockID you provide from your Vultr Account
func (b *BlockStorageServiceHandler) Get(ctx context.Context, blockID string) (*BlockStorage, *http.Response, error) {
	uri := fmt.Sprintf("/v2/blocks/%s", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	block := new(blockStorageBase)
	resp, err := b.client.DoWithContext(ctx, req, block)
	if err != nil {
		return nil, resp, err
	}

	return block.Block, resp, nil
}

// Update a block storage subscription.
func (b *BlockStorageServiceHandler) Update(ctx context.Context, blockID string, blockReq *BlockStorageUpdate) error {
	uri := fmt.Sprintf("/v2/blocks/%s", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodPatch, uri, blockReq)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete a block storage subscription from your Vultr account
func (b *BlockStorageServiceHandler) Delete(ctx context.Context, blockID string) error {
	uri := fmt.Sprintf("/v2/blocks/%s", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// List returns a list of all block storage instances on your Vultr Account
func (b *BlockStorageServiceHandler) List(ctx context.Context, options *ListOptions) ([]BlockStorage, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/blocks"

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	blocks := new(blockStoragesBase)
	resp, err := b.client.DoWithContext(ctx, req, blocks)
	if err != nil {
		return nil, nil, resp, err
	}

	return blocks.Blocks, blocks.Meta, resp, nil
}

// Attach will link a given block storage to a given Vultr instance
// If Live is set to true the block storage will be attached without reloading the instance
func (b *BlockStorageServiceHandler) Attach(ctx context.Context, blockID string, attach *BlockStorageAttach) error {
	uri := fmt.Sprintf("/v2/blocks/%s/attach", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, attach)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// Detach will de-link a given block storage to the Vultr instance it is attached to
// If Live is set to true the block storage will be detached without reloading the instance
func (b *BlockStorageServiceHandler) Detach(ctx context.Context, blockID string, detach *BlockStorageDetach) error {
	uri := fmt.Sprintf("/v2/blocks/%s/detach", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, detach)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}
