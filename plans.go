package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// PlanService is the interface to interact with the Plans endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/plans
type PlanService interface {
	List(ctx context.Context, planType string, options *ListOptions) ([]Plan, *Meta, *http.Response, error)
	ListBareMetal(ctx context.Context, options *ListOptions) ([]BareMetalPlan, *Meta, *http.Response, error)
}

// PlanServiceHandler handles interaction with the Plans methods for the Vultr API
type PlanServiceHandler struct {
	client *Client
}

// BareMetalPlan represents bare metal plans
type BareMetalPlan struct {
	ID          string   `json:"id"`
	CPUModel    string   `json:"cpu_model"`
	Type        string   `json:"type"`
	Locations   []string `json:"locations"`
	CPUCount    int      `json:"cpu_count"`
	CPUThreads  int      `json:"cpu_threads"`
	RAM         int      `json:"ram"`
	Disk        int      `json:"disk"`
	DiskCount   int      `json:"disk_count"`
	Bandwidth   int      `json:"bandwidth"`
	MonthlyCost float32  `json:"monthly_cost"`
}

// Plan represents vc2, vdc, or vhf
type Plan struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	GPUType     string   `json:"gpu_type,omitempty"`
	Locations   []string `json:"locations"`
	VCPUCount   int      `json:"vcpu_count"`
	RAM         int      `json:"ram"`
	Disk        int      `json:"disk"`
	DiskCount   int      `json:"disk_count"`
	Bandwidth   int      `json:"bandwidth"`
	GPUVRAM     int      `json:"gpu_vram_gb,omitempty"`
	MonthlyCost float32  `json:"monthly_cost"`
}

type plansBase struct {
	Meta  *Meta  `json:"meta"`
	Plans []Plan `json:"plans"`
}

type bareMetalPlansBase struct {
	Meta  *Meta           `json:"meta"`
	Plans []BareMetalPlan `json:"plans_metal"`
}

// List retrieves a list of all active plans.
// planType is optional - pass an empty string to get all plans
func (p *PlanServiceHandler) List(ctx context.Context, planType string, options *ListOptions) ([]Plan, *Meta, *http.Response, error) {
	uri := "/v2/plans"

	req, err := p.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	if planType != "" {
		newValues.Add("type", planType)
	}

	req.URL.RawQuery = newValues.Encode()

	plans := new(plansBase)
	resp, err := p.client.DoWithContext(ctx, req, plans)
	if err != nil {
		return nil, nil, resp, err
	}

	return plans.Plans, plans.Meta, resp, nil
}

// ListBareMetal all active bare metal plans.
func (p *PlanServiceHandler) ListBareMetal(ctx context.Context, options *ListOptions) ([]BareMetalPlan, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/plans-metal"

	req, err := p.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	bmPlans := new(bareMetalPlansBase)
	resp, err := p.client.DoWithContext(ctx, req, bmPlans)
	if err != nil {
		return nil, nil, nil, err
	}

	return bmPlans.Plans, bmPlans.Meta, resp, nil
}
