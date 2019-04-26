package govultr

import (
	"context"
)

// RegionsService is the interface to interact with Region endpoints on the Vultr API
// Link: https://www.vultr.com/api/#regions
type RegionsService interface {
	Availability(ctx context.Context, regionID int, planType string) ([]int, error)
	BareMetalAvailability(ctx context.Context, regionID int) ([]int, error) // returns plan Ids
	Vc2Availability(ctx context.Context, regionID int) ([]int, error)       // returns plan IDs
	Vdc2Availability(ctx context.Context, regionID int) ([]int, error)      // returns plan IDs
	GetList(ctx context.Context, available string) ([]Region, error)
}

// RegionsServiceHandler handles interaction with the region methods for the Vultr API
type RegionsServiceHandler struct {
	Client *Client
}

// Region represents a Vultr region
type Region struct {
	RegionID     string `json:"DCID"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	Continent    string `json:"continent"`
	State        string `json:"state"`
	Ddos         bool   `json:"ddos_protection"`
	BlockStorage bool   `json:"block_storage"`
	RegionCode   string `json:"regioncode"`
}

// Availability retrieves a list of the VPSPLANIDs currently available for a given location.
func (r *RegionsServiceHandler) Availability(ctx context.Context, regionID int, planType string) ([]int, error) {

	return nil, nil
}

// BareMetalAvailability retrieve a list of the METALPLANIDs currently available for a given location.
func (r *RegionsServiceHandler) BareMetalAvailability(ctx context.Context, regionID int) ([]int, error) {
	return nil, nil
}

// Vc2Availability retrieve a list of the vc2 VPSPLANIDs currently available for a given location.
func (r *RegionsServiceHandler) Vc2Availability(ctx context.Context, regionID int) ([]int, error) {
	return nil, nil
}

// Vdc2Availability retrieves a list of the vdc2 VPSPLANIDs currently available for a given location.
func (r *RegionsServiceHandler) Vdc2Availability(ctx context.Context, regionID int) ([]int, error) {
	return nil, nil
}

// GetList retrieves a list of all active regions
func (r *RegionsServiceHandler) GetList(ctx context.Context, available string) ([]Region, error) {
	return nil, nil
}
