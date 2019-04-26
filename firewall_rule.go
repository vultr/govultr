package govultr

import (
	"context"
	"net"
	"net/http"
	"net/url"
)

type FireWallRuleService interface {
	Create(ctx context.Context, groupID string, rule *FirewallRule) (*FirewallRule, error)
	Delete(ctx context.Context, groupID, ruleID string) error
	GetList(ctx context.Context, groupID, direction, ipType string) ([]FirewallRule, error)
}

type FireWallRuleServiceHandler struct {
	client *Client
}

type FirewallRule struct {
	RuleNumber int        `json:"rulenumber"`
	Action     string     `json:"action"`
	Protocol   string     `json:"protocol"`
	Port       string     `json:"port"`
	Network    *net.IPNet `json:"network"`
	Notes      string     `json:"notes"`
}

func (f *FireWallRuleServiceHandler) Create(ctx context.Context, groupID string, rule *FirewallRule) (*FirewallRule, error) {
	return nil, nil
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
func (f *FireWallRuleServiceHandler) GetList(ctx context.Context, groupID, direction, ipType string) ([]FirewallRule, error) {
	return nil, nil
}
