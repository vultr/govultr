package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFireWallGroupServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"firewall_group": {
		"id": "44d0f934",
		"description": "govultr test",
		"date_created": "2020-10-10T01:56:20+00:00",
		"date_modified": "2020-10-10T01:56:20+00:00",
		"instance_count": 15,
		"rule_count": 6,
		"max_rule_count": 999
	}
}`))

	group := &FirewallGroupReq{Description: "govultr test"}
	firewallGroup, _, err := client.FirewallGroup.Create(ctx, group)
	if err != nil {
		t.Errorf("FirewallGroup.Create returned error: %v", err)
	}

	expected := &FirewallGroup{
		ID:            "44d0f934",
		Description:   "govultr test",
		DateCreated:   "2020-10-10T01:56:20+00:00",
		DateModified:  "2020-10-10T01:56:20+00:00",
		InstanceCount: 15,
		RuleCount:     6,
		MaxRuleCount:  999,
	}

	if !reflect.DeepEqual(firewallGroup, expected) {
		t.Errorf("FirewallGroup.Create returned %+v, expected %+v", firewallGroup, expected)
	}
}

func TestFireWallGroupServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/44d0f934", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"firewall_group": {
		"id": "44d0f934",
		"description": "govultr test",
		"date_created": "2020-10-10T01:56:20+00:00",
		"date_modified": "2020-10-10T01:56:20+00:00",
		"instance_count": 15,
		"rule_count": 6,
		"max_rule_count": 999
	}
}`))

	firewallGroup, _, err := client.FirewallGroup.Get(ctx, "44d0f934")
	if err != nil {
		t.Errorf("FirewallGroup.Create returned error: %v", err)
	}

	expected := &FirewallGroup{
		ID:            "44d0f934",
		Description:   "govultr test",
		DateCreated:   "2020-10-10T01:56:20+00:00",
		DateModified:  "2020-10-10T01:56:20+00:00",
		InstanceCount: 15,
		RuleCount:     6,
		MaxRuleCount:  999,
	}

	if !reflect.DeepEqual(firewallGroup, expected) {
		t.Errorf("FirewallGroup.Create returned %+v, expected %+v", firewallGroup, expected)
	}
}

func TestFireWallGroupServiceHandler_ChangeDescription(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	put := &FirewallGroupReq{Description: "test"}
	err := client.FirewallGroup.Update(ctx, "abc123", put)
	if err != nil {
		t.Errorf("FirewallGroup.ChangeDescription returned error: %v", err)
	}
}

func TestFireWallGroupServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	err := client.FirewallGroup.Delete(ctx, "abc123")

	if err != nil {
		t.Errorf("FirewallGroup.Delete returned error: %v", err)
	}
}

func TestFireWallGroupServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"firewall_groups": [
		{
			"id": "44d0f934",
			"description": "govultr test",
			"date_created": "2020-10-10T01:56:20+00:00",
			"date_modified": "2020-10-10T01:56:20+00:00",
			"instance_count": 15,
			"rule_count": 6,
			"max_rule_count": 999
		}
	],
	"meta": {
		"total": 5,
		"links": {
			"next": "",
			"prev": ""
		}
	}
}`))

	firewallGroup, meta, _, err := client.FirewallGroup.List(ctx, nil)
	if err != nil {
		t.Errorf("FirewallGroup.List returned error: %v", err)
	}

	expectedGroup := []FirewallGroup{
		{
			ID:            "44d0f934",
			Description:   "govultr test",
			DateCreated:   "2020-10-10T01:56:20+00:00",
			DateModified:  "2020-10-10T01:56:20+00:00",
			InstanceCount: 15,
			RuleCount:     6,
			MaxRuleCount:  999,
		},
	}

	expectedMeta := &Meta{
		Total: 5,
		Links: &Links{},
	}

	if !reflect.DeepEqual(firewallGroup, expectedGroup) {
		t.Errorf("FirewallGroup.List groups returned %+v, expected %+v", firewallGroup, expectedGroup)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("FirewallGroup.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
