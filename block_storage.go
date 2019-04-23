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
	//Attach(ctx context.Context, blockID, vpsID, live string) error
	Create(ctx context.Context, regionID, size int, label string) (*BlockStorage, error)
	//Delete(ctx context.Context, subID string) error
	//Detach(ctx context.Context, blockID, live string) error
	//SetLabel(ctx context.Context, blockID, label string) error
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
	CostPerMonth   int    `json:"cost_per_month"`
	Status         string `json:"status"`
	Size           int    `json:"size_gb"`
	RegionID       int    `json:"region_id"`
	VpsID          int    `json:"attached_to_SUBID"`
	Label          string `json:"label"`
}

//todo
//attach

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

	return nil, nil
}
