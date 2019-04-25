package govultr

import (
	"Context"
)

// PlansService is the interface to interact with the Plans endpoints on the Vultr API
// Link: https://www.vultr.com/api/#plans
type PlansService interface {
	GetAllList(ctx context.Context, planType string) ([]Plans, error)
	GetBareMetalList(ctx context.Context) ([]Plans, error)
	GetVc2List(ctx context.Context) ([]Plans, error)
	GetVdc2List(ctx context.Context) ([]Plans, error)
}

// PlansServiceHandler handles interaction with the Plan methods for the Vultr API
type PlansServiceHandler struct {
	Client *Client
}

// Plans represents available plans that Vultr offers
type Plans struct {
	VpsID      int    `json:"VPSPLANID,string"`
	Name       string `json:"name"`
	VCpus      int    `json:"vcpu_count,string"`
	RAM        string `json:"ram"`
	Disk       string `json:"disk"`
	Bandwidth  string `json:"bandwidth"`
	Price      string `json:"price_per_month"`
	Windows    bool   `json:"windows"`
	PlanType   string `json:"plan_type"`
	Regions    []int  `json:"available_locations"`
	Deprecated bool   `json:"deprecated"`
}

// GetAllList retrieve a list of all active plans.
func (p *PlansServiceHandler) GetAllList(ctx context.Context, planType string) ([]Plans, error) {

	return nil, nil
}

// GetBareMetalList retrieves a list of all active bare metal plans.
func (p *PlansServiceHandler) GetBareMetalList(ctx context.Context) ([]Plans, error) {
	return nil, nil
}

// GetVc2List retrieve a list of all active vc2 plans.
func (p *PlansServiceHandler) GetVc2List(ctx context.Context) ([]Plans, error) {
	return nil, nil
}

// GetVdc2List Retrieve a list of all active vdc2 plans
func (p *PlansServiceHandler) GetVdc2List(ctx context.Context) ([]Plans, error) {
	return nil, nil
}
