package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFireWallRuleServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123/rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"firewall_rule":{"id":1,"type":"v4","action":"accept","protocol":"tcp","port":"80","subnet":"127.0.0.1","subnet_size":32,"source":"","notes":"thisisanote"}}`
		fmt.Fprint(writer, response)
	})

	rule := &FirewallRuleReq{
		IPType:     "v4",
		Protocol:   "tcp",
		Subnet:     "127.0.0.1",
		SubnetSize: 30,
		Port:       "80",
		Notes:      "thisisanote",
	}

	firewallRule, err := client.FirewallRule.Create(ctx, "abc123", rule)
	if err != nil {
		t.Errorf("FirewallRule.Create returned error: %v", err)
	}

	expected := &FirewallRule{
		ID:         1,
		Action:     "accept",
		Type:       "v4",
		Protocol:   "tcp",
		Port:       "80",
		Subnet:     "127.0.0.1",
		SubnetSize: 32,
		Source:     "",
		Notes:      "thisisanote",
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.Create returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestFireWallRuleServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123/rules/1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.FirewallRule.Delete(ctx, "abc123", 1)

	if err != nil {
		t.Errorf("FirewallRule.Delete returned error: %v", err)
	}
}

func TestFireWallRuleServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123/rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"firewall_rules":[{"id":1,"type":"v4","action":"accept","protocol":"tcp","port":"22","subnet":"0.0.0.0","subnet_size":0,"source":"","notes":""}],"meta":{"total":5,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	firewallRule, meta, err := client.FirewallRule.List(ctx, "abc123", nil)
	if err != nil {
		t.Errorf("FirewallRule.List returned error: %v", err)
	}

	expectedRule := []FirewallRule{
		{
			ID:         1,
			Action:     "accept",
			Type:       "v4",
			Protocol:   "tcp",
			Port:       "22",
			Subnet:     "0.0.0.0",
			SubnetSize: 0,
			Source:     "",
			Notes:      "",
		},
	}

	expectedMeta := &Meta{
		Total: 5,
		Links: &Links{},
	}

	if !reflect.DeepEqual(firewallRule, expectedRule) {
		t.Errorf("FirewallRule.List rules returned %+v, expected %+v", firewallRule, expectedRule)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("FirewallRule.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestFireWallRuleServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/firewalls/abc123/rules/1", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"firewall_rule":{"id":1,"type":"v4","action":"accept","protocol":"tcp","port":"22","subnet":"0.0.0.0","subnet_size":0,"source":"","notes":""}}`
		fmt.Fprint(writer, response)
	})

	firewallRule, err := client.FirewallRule.Get(ctx, "abc123", 1)
	if err != nil {
		t.Errorf("FirewallRule.Get returned error: %v", err)
	}

	expectedRule := &FirewallRule{
		ID:         1,
		Action:     "accept",
		Type:       "v4",
		Protocol:   "tcp",
		Port:       "22",
		Subnet:     "0.0.0.0",
		SubnetSize: 0,
		Source:     "",
		Notes:      "",
	}

	if !reflect.DeepEqual(firewallRule, expectedRule) {
		t.Errorf("FirewallRule.Get returned %+v, expected %+v", firewallRule, expectedRule)
	}
}
