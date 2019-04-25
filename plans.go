package govultr

import (
	"context"
	"fmt"
	"net/http"
)

// PlansService is the interface to interact with the Plans endpoints on the Vultr API
// Link: https://www.vultr.com/api/#plans
type PlansService interface {
	GetAllList(ctx context.Context, planType string) ([]Plans, error)
	GetBareMetalList(ctx context.Context) ([]BareMetalPlan, error)
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

// BareMetalPlan represents bare metal plans
type BareMetalPlan struct {
	BareMetalID string `json:"METALPLANID"`
	Name        string `json:"name"`
	Cpus        int    `json:"cpu_count"`
	RAM         int    `json:"ram"`
	Disk        string `json:"disk"`
	Bandwidth   int    `json:"bandwidth_tb"`
	Price       int    `json:"price_per_month"`
	PlanType    string `json:"plan_type"`
	Deprecated  bool   `json:"deprecated"`
	Regions     []int  `json:"available_locations"`
}

// GetAllList retrieve a list of all active plans.
func (p *PlansServiceHandler) GetAllList(ctx context.Context, planType string) ([]Plans, error) {

	uri := "/v1/plans/list"

	req, err := p.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	if planType != "" {
		q := req.URL.Query()
		q.Add("type", planType)
		req.URL.RawQuery = q.Encode()
	}

	var planMap map[string]Plans
	err = p.Client.DoWithContext(ctx, req, &planMap)

	if err != nil {
		fmt.Println(err)
	}

	var plans []Plans
	for _, p := range planMap {
		plans = append(plans, p)
	}

	return plans, nil
}

// GetBareMetalList retrieves a list of all active bare metal plans.
func (p *PlansServiceHandler) GetBareMetalList(ctx context.Context) ([]BareMetalPlan, error) {

	uri := "/v1/plans/list_baremetal"

	req, err := p.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	var bareMetalMap map[string]BareMetalPlan
	err = p.Client.DoWithContext(ctx, req, &bareMetalMap)

	if err != nil {
		return nil, err
	}

	var bareMetalPlan []BareMetalPlan
	for _, b := range bareMetalMap {
		bareMetalPlan = append(bareMetalPlan, b)
	}

	return bareMetalPlan, nil
}

// GetVc2List retrieve a list of all active vc2 plans.
func (p *PlansServiceHandler) GetVc2List(ctx context.Context) ([]Plans, error) {
	uri := "/v1/plans/list_vc2"

	req, err := p.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	var planMap map[string]Plans
	err = p.Client.DoWithContext(ctx, req, &planMap)

	if err != nil {
		fmt.Println(err)
	}

	var plans []Plans
	for _, p := range planMap {
		plans = append(plans, p)
	}

	return plans, nil
}

// GetVdc2List Retrieve a list of all active vdc2 plans
func (p *PlansServiceHandler) GetVdc2List(ctx context.Context) ([]Plans, error) {
	return nil, nil
}
