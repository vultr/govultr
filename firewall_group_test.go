package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFireWallGroupServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/group_create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"FIREWALLGROUPID":"1234abcd"}`
		fmt.Fprint(writer, response)
	})

	firewallGroup, err := client.FirewallGroup.Create(ctx, "firewall-group-name")

	if err != nil {
		t.Errorf("FirewallGroup.Create returned error: %v", err)
	}

	expected := &FirewallGroup{FirewallGroupID: "1234abcd"}

	if !reflect.DeepEqual(firewallGroup, expected) {
		t.Errorf("FirewallGroup.Create returned %+v, expected %+v", firewallGroup, expected)
	}

}

func TestFireWallGroupServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/group_delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.FirewallGroup.Delete(ctx, "12345abcd")

	if err != nil {
		t.Errorf("FirewallGroup.Delete returned error: %v", err)
	}
}

func TestFireWallGroupServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/group_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1234abcd": { "FIREWALLGROUPID": "1234abcd", "description": "my http firewall","date_created": "2017-02-14 17:48:40","date_modified": "2017-02-14 17:48:40","instance_count": 2,"rule_count": 2, "max_rule_count": 50}}`
		fmt.Fprint(writer, response)
	})

	firewallGroup, err := client.FirewallGroup.GetList(ctx)

	if err != nil {
		t.Errorf("FirewallGroup.GetList returned error: %v", err)
	}

	expected := []FirewallGroup{
		{
			FirewallGroupID: "1234abcd",
			Description:     "my http firewall",
			DateCreated:     "2017-02-14 17:48:40",
			DateModified:    "2017-02-14 17:48:40",
			InstanceCount:   2,
			RuleCount:       2,
			MaxRuleCount:    50,
		},
	}

	if !reflect.DeepEqual(firewallGroup, expected) {
		t.Errorf("FirewallGroup.GetList returned %+v, expected %+v", firewallGroup, expected)
	}

}

func TestFireWallGroupServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/group_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1234abcd": { "FIREWALLGROUPID": "1234abcd", "description": "my http firewall","date_created": "2017-02-14 17:48:40","date_modified": "2017-02-14 17:48:40","instance_count": 2,"rule_count": 2, "max_rule_count": 50}}`
		fmt.Fprint(writer, response)
	})

	firewallGroup, err := client.FirewallGroup.Get(ctx, "1234abcd")

	if err != nil {
		t.Errorf("FirewallGroup.Get returned error: %v", err)
	}

	expected := &FirewallGroup{
		FirewallGroupID: "1234abcd",
		Description:     "my http firewall",
		DateCreated:     "2017-02-14 17:48:40",
		DateModified:    "2017-02-14 17:48:40",
		InstanceCount:   2,
		RuleCount:       2,
		MaxRuleCount:    50,
	}

	if !reflect.DeepEqual(firewallGroup, expected) {
		t.Errorf("FirewallGroup.Get returned %+v, expected %+v", firewallGroup, expected)
	}
}

func TestFireWallGroupServiceHandler_ChangeDescription(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/firewall/group_set_description", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.FirewallGroup.ChangeDescription(ctx, "12345abcd", "new description")

	if err != nil {
		t.Errorf("FirewallGroup.ChangeDescription returned error: %v", err)
	}
}
