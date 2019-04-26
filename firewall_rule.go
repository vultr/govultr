package govultr

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

// FireWallRuleService is the interface to interact with the firewall rule endpoints on the Vultr API
// Link: https://www.vultr.com/api/#firewall
type FireWallRuleService interface {
	Create(ctx context.Context, groupID, protocol, port, network, notes string) (*FirewallRule, error)
	Delete(ctx context.Context, groupID, ruleID string) error
	GetList(ctx context.Context, groupID, direction, ipType string) ([]FirewallRule, error)
}

// FireWallRuleServiceHandler handles interaction with the firewall rule methods for the Vultr API
type FireWallRuleServiceHandler struct {
	client *Client
}

// FirewallRule represents a Vultr firewall rule
type FirewallRule struct {
	RuleNumber int        `json:"rulenumber"`
	Action     string     `json:"action"`
	Protocol   string     `json:"protocol"`
	Port       string     `json:"port"`
	Network    *net.IPNet `json:"network"`
	Notes      string     `json:"notes"`
}

// Create will create a rule in a firewall rule.
func (f *FireWallRuleServiceHandler) Create(ctx context.Context, groupID, protocol, port, cdirBlock, notes string) (*FirewallRule, error) {

	uri := "/v1/firewall/rule_create"

	ip, ipNet, err := net.ParseCIDR(cdirBlock)

	if err != nil {
		return nil, err
	}

	values := url.Values{
		"FIREWALLGROUPID": {groupID},
		"direction":       {"in"},
		"protocol":        {protocol},
		"subnet":          {ip.String()},
	}

	// mask
	mask, _ := ipNet.Mask.Size()
	values.Add("subnet_size", strconv.Itoa(mask))

	// ip Type
	if ipNet.IP.To4() != nil {
		values.Add("ip_type", "v4")
	} else {
		values.Add("ip_type", "v6")
	}

	// Optional params
	if port != "" {
		values.Add("port", port)
	}

	if notes != "" {
		values.Add("notes", notes)
	}

	req, err := f.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return nil, err
	}

	firewallRule := new(FirewallRule)
	err = f.client.DoWithContext(ctx, req, firewallRule)

	if err != nil {
		return nil, err
	}

	return firewallRule, nil
}

// Delete will delete a firewall rule on your Vultr account
func (f *FireWallRuleServiceHandler) Delete(ctx context.Context, groupID, ruleID string) error {

	uri := "/v1/firewall/rule_delete"

	values := url.Values{
		"FIREWALLGROUPID": {groupID},
		"rulenumber":      {ruleID},
	}

	req, err := f.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = f.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetList will list the rules in a firewall rule
func (f *FireWallRuleServiceHandler) GetList(ctx context.Context, groupID, direction, ipType string) ([]FirewallRule, error) {
	return nil, nil
}
