package govultr

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// LoadBalancerService is the interface to interact with the server endpoints on the Vultr API
// Link: https://www.vultr.com/api/#loadbalancer
type LoadBalancerService interface {
	List(ctx context.Context) ([]LoadBalancers, error)
	Delete(ctx context.Context, ID int) error
	SetLabel(ctx context.Context, ID int, label string) error
	AttachedInstances(ctx context.Context, ID int) (*InstanceList, error)
	AttachInstance(ctx context.Context, ID, backendNode int) error
	DetachInstance(ctx context.Context, ID, backendNode int) error
	GetHealthCheck(ctx context.Context, ID int) (*HealthCheck, error)
	SetHealthCheck(ctx context.Context, ID int, healthConfig *HealthCheck) error
	GetGenericInfo(ctx context.Context, ID int) (*GenericInfo, error)
	ListForwardingRules(ctx context.Context, ID int) (*ForwardingRules, error)
	DeleteForwardingRule(ctx context.Context, ID int, RuleID string) error
	CreateForwardingRule(ctx context.Context, ID int, rule *ForwardingRule) (*ForwardingRule, error)
}

// LoadBalancerHandler handles interaction with the server methods for the Vultr API
type LoadBalancerHandler struct {
	client *Client
}

// LoadBalancers represent a basic structure of a load balancer
type LoadBalancers struct {
	ID          int    `json:"SUBID"`
	DateCreated string `json:"date_created"`
	RegionID    int    `json:"DCID"`
	Location    string
	Label       string
	Status      string
	IPV4        string `json:"main_ipv4"`
	IPV6        string `json:"main_ipv6"`
}

// InstanceList represents instances that attached to your load balancer
type InstanceList struct {
	InstanceList []string `json:"instance_list"`
}

// HealthCheck represents your health check configuration for your load balancer.
type HealthCheck struct {
	Protocol           string
	Port               int
	Path               string
	CheckInterval      int `json:"check_interval"`
	ResponseTimeout    int `json:"response_timeout"`
	UnhealthyThreshold int `json:"unhealthy_threshold"`
	HealthyThreshold   int `json:"healthy_threshold"`
}

// GenericInfo represents generic configuration of your load balancer
type GenericInfo struct {
	BalancingAlgorithm string     `json:"balancing_algorithm"`
	SSLRedirect        bool       `json:"ssl_redirect"`
	StickySessions     CookieName `json:"sticky_sessions"`
}

// CookieName represents cookie for your load balancer
type CookieName struct {
	CookieName string `json:"cookie_name"`
}

// ForwardingRules represent a list of forwarding rules
type ForwardingRules struct {
	ForwardRuleList []ForwardingRule `json:"forward_rule_list"`
}

// ForwardingRule represent a single forwarding rule
type ForwardingRule struct {
	RuleID           string `json:"RULEID,omitempty"`
	FrontendProtocol string `json:"frontend_protocol,omitempty"`
	FrontendPort     int    `json:"frontend_port,omitempty"`
	BackendProtocol  string `json:"backend_protocol,omitempty"`
	BackendPort      int    `json:"backend_port,omitempty"`
}

// List all load balancer subscriptions on the current account.
func (l *LoadBalancerHandler) List(ctx context.Context) ([]LoadBalancers, error) {
	uri := "/v1/loadbalancer/list"

	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var lbList []LoadBalancers

	err = l.client.DoWithContext(ctx, req, &lbList)
	if err != nil {
		return nil, err
	}

	return lbList, nil
}

// Delete a load balancer subscription.
func (l *LoadBalancerHandler) Delete(ctx context.Context, ID int) error {
	uri := "/v1/loadbalancer/destroy"

	values := url.Values{
		"SUBID": {strconv.Itoa(ID)},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetLabel sets the label for your load balancer subscription.
func (l *LoadBalancerHandler) SetLabel(ctx context.Context, ID int, label string) error {
	uri := "/v1/loadbalancer/label_set"

	values := url.Values{
		"SUBID": {strconv.Itoa(ID)},
		"label": {label},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return nil
	}

	return nil
}

// AttachedInstances lists the instances that are currently attached to a load balancer subscription.
func (l *LoadBalancerHandler) AttachedInstances(ctx context.Context, ID int) (*InstanceList, error) {
	uri := "/v1/loadbalancer/instance_list"

	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", strconv.Itoa(ID))
	req.URL.RawQuery = q.Encode()

	var instances InstanceList

	err = l.client.DoWithContext(ctx, req, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

// AttachInstance attaches a backend node to your load balancer subscription
func (l *LoadBalancerHandler) AttachInstance(ctx context.Context, ID, backendNode int) error {
	uri := "/v1/loadbalancer/instance_attach"

	values := url.Values{
		"SUBID":       {strconv.Itoa(ID)},
		"backendNode": {strconv.Itoa(backendNode)},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DetachInstance detaches a backend node to your load balancer subscription
func (l *LoadBalancerHandler) DetachInstance(ctx context.Context, ID, backendNode int) error {
	uri := "/v1/loadbalancer/instance_detach"

	values := url.Values{
		"SUBID":       {strconv.Itoa(ID)},
		"backendNode": {strconv.Itoa(backendNode)},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetHealthCheck retrieves the health check configuration for your load balancer subscription.
func (l *LoadBalancerHandler) GetHealthCheck(ctx context.Context, ID int) (*HealthCheck, error) {
	uri := "/v1/loadbalancer/health_check_info"

	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", strconv.Itoa(ID))
	req.URL.RawQuery = q.Encode()

	var healthCheck HealthCheck
	err = l.client.DoWithContext(ctx, req, &healthCheck)
	if err != nil {
		return nil, err
	}

	return &healthCheck, nil
}

// SetHealthCheck sets your health check configuration for your load balancer
func (l *LoadBalancerHandler) SetHealthCheck(ctx context.Context, ID int, healthConfig *HealthCheck) error {
	uri := "/v1/loadbalancer/health_check_update"

	values := url.Values{
		"SUBID": {strconv.Itoa(ID)},
	}

	if healthConfig != nil {
		if healthConfig.Protocol != "" {
			values.Add("protocol", healthConfig.Protocol)
		}

		if healthConfig.Port != 0 {
			values.Add("port", strconv.Itoa(healthConfig.Port))
		}

		if healthConfig.CheckInterval != 0 {
			values.Add("check_interval", strconv.Itoa(healthConfig.CheckInterval))
		}

		if healthConfig.ResponseTimeout != 0 {
			values.Add("response_timeout", strconv.Itoa(healthConfig.ResponseTimeout))
		}

		if healthConfig.UnhealthyThreshold != 0 {
			values.Add("unhealthy_threshold", strconv.Itoa(healthConfig.UnhealthyThreshold))
		}

		if healthConfig.HealthyThreshold != 0 {
			values.Add("healthy_threshold", strconv.Itoa(healthConfig.HealthyThreshold))
		}

		if healthConfig.Path != "" {
			values.Add("path", healthConfig.Path)
		}
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetGenericInfo is the generic configuration of a load balancer subscription
func (l *LoadBalancerHandler) GetGenericInfo(ctx context.Context, ID int) (*GenericInfo, error) {
	uri := "/v1/loadbalancer/generic_info"

	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", strconv.Itoa(ID))
	req.URL.RawQuery = q.Encode()

	var info GenericInfo

	err = l.client.DoWithContext(ctx, req, &info)
	if err != nil {
		return nil, err
	}

	return &info, err
}

// ListForwardingRules lists all forwarding rules for a load balancer subscription
func (l *LoadBalancerHandler) ListForwardingRules(ctx context.Context, ID int) (*ForwardingRules, error) {
	uri := "v1/loadbalancer/forward_rule_list"

	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", strconv.Itoa(ID))
	req.URL.RawQuery = q.Encode()

	var frList ForwardingRules

	err = l.client.DoWithContext(ctx, req, &frList)
	if err != nil {
		return nil, err
	}

	return &frList, nil
}

// DeleteForwardingRule removes a forwarding rule from a load balancer subscription
func (l *LoadBalancerHandler) DeleteForwardingRule(ctx context.Context, ID int, RuleID string) error {
	uri := "/v1/loadbalancer/forward_rule_delete"

	values := url.Values{
		"SUBID":  {strconv.Itoa(ID)},
		"RULEID": {RuleID},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return err
	}

	err = l.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateForwardingRule will create a new forwarding rule for your load balancer subscription.
// Note the RuleID will be returned in the ForwardingRule struct
func (l *LoadBalancerHandler) CreateForwardingRule(ctx context.Context, ID int, rule *ForwardingRule) (*ForwardingRule, error) {
	uri := "/v1/loadbalancer/forward_rule_create"

	values := url.Values{
		"SUBID":             {strconv.Itoa(ID)},
		"frontend_protocol": {rule.FrontendProtocol},
		"backend_protocol":  {rule.BackendProtocol},
		"frontend_port":     {strconv.Itoa(rule.FrontendPort)},
		"backend_port":      {strconv.Itoa(rule.BackendPort)},
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return nil, err
	}

	var fr ForwardingRule
	err = l.client.DoWithContext(ctx, req, &fr)
	if err != nil {
		return nil, err
	}

	return &fr, nil
}
