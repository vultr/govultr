package govultr

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestFireWallRuleServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/rule_create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"rulenumber": 2}`

		fmt.Fprint(writer, response)
	})

	firewallRule, err := client.FirewallRule.Create(ctx, "123456", "tcp", "8080", "10.0.0.0/32", "note")

	if err != nil {
		t.Errorf("FirewallRule.Create returned error: %v", err)
	}

	expected := &FirewallRule{RuleNumber: 2}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.Create returned %+v, expected %+v", firewallRule, expected)
	}

	firewallRule, err = client.FirewallRule.Create(ctx, "123456", "tcp", "8080", "::/0", "note")

	if err != nil {
		t.Errorf("FirewallRule.Create returned error: %v", err)
	}

	expected = &FirewallRule{RuleNumber: 2}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.Create returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestFireWallRuleServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/rule_delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.FirewallRule.Delete(ctx, "123456", "123")

	if err != nil {
		t.Errorf("FirewallRule.Delete returned error: %v", err)
	}

}

func TestFireWallRuleServiceHandler_GetAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/rule_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "1": {"rulenumber": 1,"action": "accept","protocol": "icmp","port": "","subnet": "","subnet_size": 0,"notes": ""}}`
		fmt.Fprint(writer, response)
	})

	firewallRule, err := client.FirewallRule.ListByIPType(ctx, "12345", "v4")
	if err != nil {
		t.Errorf("FirewallRule.ListByIPType returned error: %v", err)
	}

	expected := []FirewallRule{
		{
			RuleNumber: 1,
			Action:     "accept",
			Protocol:   "icmp",
			Network:    nil,
		},
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.ListByIPType returned %+v, expected %+v", firewallRule, expected)
	}

	firewallRule, err = client.FirewallRule.ListByIPType(ctx, "12345", "v6")
	if err != nil {
		t.Errorf("FirewallRule.ListByIPType returned error: %v", err)
	}

	expected = []FirewallRule{
		{
			RuleNumber: 1,
			Action:     "accept",
			Protocol:   "icmp",
			Network:    nil,
		},
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.ListByIPType returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestFireWallRuleServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/rule_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "1": {"rulenumber": 1,"action": "accept","protocol": "icmp","port": "8080","subnet": "10.0.0.0","subnet_size": 32,"notes": ""}}`
		fmt.Fprint(writer, response)
	})

	firewallRule, err := client.FirewallRule.List(ctx, "12345")
	if err != nil {
		t.Errorf("FirewallRule.List returned error: %v", err)
	}

	_, ip, _ := net.ParseCIDR("10.0.0.0/32")
	expected := []FirewallRule{
		{
			RuleNumber: 1,
			Action:     "accept",
			Protocol:   "icmp",
			Port:       "8080",
			Network:    ip,
		},
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("FirewallRule.List returned %+v, expected %+v", firewallRule, expected)
	}
}
