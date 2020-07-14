package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestBody map[string]interface{}

// BlockStorageService is the interface to interact with Block-Storage endpoint on the Vultr API
type BlockStorageService interface {
	Create(ctx context.Context, blockReq *BlockStorageReq) (*BlockStorage, error)
	Get(ctx context.Context, blockID int) (*BlockStorage, error)
	Update(ctx context.Context, blockID int, label string) error
	Delete(ctx context.Context, blockID int) error
	List(ctx context.Context, options *ListOptions) ([]BlockStorage, *Meta, error)

	Attach(ctx context.Context, blockID, instanceID int, liveAttach string) error
	Detach(ctx context.Context, blockID int, liveDetach string) error
	Resize(ctx context.Context, blockID int, sizeGB int) error
}

// BlockStorageServiceHandler handles interaction with the block-storage methods for the Vultr API
type BlockStorageServiceHandler struct {
	client *Client
}

// BlockStorage represents Vultr Block-Storage
type BlockStorage struct {
	ID                 int    `json:"id"`
	Cost               int    `json:"cost"`
	Status             string `json:"status"`
	SizeGB             int    `json:"size_gb"`
	Region             string `json:"region"`
	DateCreated        string `json:"date_created"`
	AttachedToInstance int    `json:"attached_to_instance"`
	Label              string `json:"label"`
}

// BlockStorageReq
type BlockStorageReq struct {
	Region string `json:"region"`
	SizeGB int    `json:"size_gb"`
	Label  string `json:"label,omitempty"`
}

type blockStoragesBase struct {
	Blocks []BlockStorage `json:"blocks"`
	Meta   *Meta          `json:"meta"`
}

type blockStorageBase struct {
	Block *BlockStorage `json:"block"`
}

// Create builds out a block storage
func (b *BlockStorageServiceHandler) Create(ctx context.Context, blockReq *BlockStorageReq) (*BlockStorage, error) {
	uri := "/v2/blocks"

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, blockReq)
	if err != nil {
		return nil, err
	}

	block := new(blockStorageBase)
	if err = b.client.DoWithContext(ctx, req, block); err != nil {
		return nil, err
	}

	return block.Block, nil
}

// Get returns a single block storage instance based ony our blockID you provide from your Vultr Account
func (b *BlockStorageServiceHandler) Get(ctx context.Context, blockID int) (*BlockStorage, error) {
	uri := fmt.Sprintf("/v2/blocks/%d", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	block := new(blockStorageBase)
	if err = b.client.DoWithContext(ctx, req, block); err != nil {
		return nil, err
	}

	return block.Block, nil
}

// SetLabel allows you to set/update the label on your Vultr Block storage
func (b *BlockStorageServiceHandler) Update(ctx context.Context, blockID int, label string) error {
	uri := fmt.Sprintf("/v2/blocks/%d", blockID)
	put := &RequestBody{"label": label}
	req, err := b.client.NewRequest(ctx, http.MethodPatch, uri, put)
	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Delete will remove block storage instance from your Vultr account
func (b *BlockStorageServiceHandler) Delete(ctx context.Context, blockID int) error {
	uri := fmt.Sprintf("/v2/blocks/%d", blockID)

	req, err := b.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err !=nil {
		return err
	}

	return nil
}

// List returns a list of all block storage instances on your Vultr Account
func (b *BlockStorageServiceHandler) List(ctx context.Context, options *ListOptions) ([]BlockStorage, *Meta, error) {
	uri := "/v2/blocks"

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	blocks := new(blockStoragesBase)
	if err = b.client.DoWithContext(ctx, req, blocks); err != nil {
		return nil, nil, err
	}

	return blocks.Blocks, blocks.Meta, nil
}

// Attach will link a given block storage to a given Vultr vps
// If liveAttach is set to "yes" the block storage will be attached without reloading the instance
func (b *BlockStorageServiceHandler) Attach(ctx context.Context, blockID, instanceID int, liveAttach string) error {
	uri := fmt.Sprintf("/v2/blocks/%d/attach", blockID)

	t := make(map[string]interface{})
	t["instance_id"] = instanceID
	if liveAttach == "yes" {
		t["live"] = liveAttach
	}

	updates := RequestBody{}
	updates = t

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, updates)
	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

//
// Detach will de-link a given block storage to the Vultr instance it is attached to
// If liveDetach is set to "yes" the block storage will be detached without reloading the instance
func (b *BlockStorageServiceHandler) Detach(ctx context.Context, blockID int, liveDetach string) error {
	uri := fmt.Sprintf("/v2/blocks/%d/detach", blockID)

	t := make(map[string]interface{})
	if liveDetach == "yes" {
		t["live"] = liveDetach
	}

	updates := RequestBody{}
	updates = t

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, updates)
	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Resize allows you to resize your Vultr block storage
func (b *BlockStorageServiceHandler) Resize(ctx context.Context, blockID int, sizeGB int) error {
	uri := fmt.Sprintf("/v2/blocks/%d/resize", blockID)
	body := &RequestBody{"size_gb": sizeGB}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
