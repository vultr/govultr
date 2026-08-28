package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestLoadBalancerHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(lbPath, testJSONResponseHandlerFunc(http.StatusOK, `
{
	"load_balancers" : [
		{
			"id": "ebfc2020-8be8-485f-8e01-74f75b24e390",
			"date_created": "2020-01-07 17:24:23",
			"region": "ewr",
			"label": "my label",
			"status": "active",
			"ipv4": "123.123.123.123",
			"ipv6": "2001:DB8:1000::100",
			"generic_info": {
				"balancing_algorithm": "roundrobin",
				"ssl_redirect": false,
				"proxy_protocol": false,
				"vpc": "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
				"sticky_sessions": {
					"cookie_name": "my-cookie"
				}
			},
			"health_check": {
				"protocol": "http",
				"port": 80,
				"path": "/",
				"check_interval": 15,
				"response_timeout": 5,
				"unhealthy_threshold": 5,
				"healthy_threshold": 5
			},
			"has_ssl": false,
			"forwarding_rules": [
				{
					"id": "abcd12345",
					"frontend_protocol": "http",
					"frontend_port": 80,
					"backend_protocol": "http",
					"backend_port": 80
				}
			],
			"firewall_rules": [
				{
					"id": "abcd12345",
					"port": 80,
					"source": "0.0.0.0/0",
					"ip_type": "v4"
				}
			],
			"nodes": 3,
			"instances": [
				"12345"
			]
		}
	],
	"meta": {
		"total":8,
		"links": {
			"next":"",
			"prev":""
		}
	}
}`))

	list, meta, _, err := client.LoadBalancer.List(ctx, nil)
	if err != nil {
		t.Errorf("LoadBalancer.List returned %+v", err)
	}

	expected := []LoadBalancer{
		{
			ID:          "ebfc2020-8be8-485f-8e01-74f75b24e390",
			DateCreated: "2020-01-07 17:24:23",
			Label:       "my label",
			Status:      "active",
			Region:      "ewr",
			IPV6:        "2001:DB8:1000::100",
			IPV4:        "123.123.123.123",
			SSLInfo:     BoolToBoolPtr(false),
			ForwardingRules: []ForwardingRule{
				{
					RuleID:           "abcd12345",
					FrontendProtocol: "http",
					FrontendPort:     80,
					BackendProtocol:  "http",
					BackendPort:      80,
				},
			},
			GenericInfo: &GenericInfo{
				BalancingAlgorithm: "roundrobin",
				SSLRedirect:        BoolToBoolPtr(false),
				ProxyProtocol:      BoolToBoolPtr(false),
				VPC:                "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
				StickySessions: &StickySessions{
					CookieName: "my-cookie",
				},
			},
			HealthCheck: &HealthCheck{
				Protocol:           "http",
				Port:               80,
				Path:               "/",
				CheckInterval:      15,
				ResponseTimeout:    5,
				UnhealthyThreshold: 5,
				HealthyThreshold:   5,
			},
			Instances: []string{"12345"},
			Nodes:     3,
			FirewallRules: []LBFirewallRule{
				{
					RuleID: "abcd12345",
					Port:   80,
					Source: "0.0.0.0/0",
					IPType: "v4",
				},
			},
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("LoadBalancer.List returned %+v, expected %+v", list, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("LoadBalancer.List returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestLoadBalancerHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"load_balancer" : {
		"id": "ebfc2020-8be8-485f-8e01-74f75b24e390",
		"date_created": "2020-01-07 17:24:23",
		"region": "ewr",
		"label": "my label",
		"status": "active",
		"ipv4": "123.123.123.123",
		"ipv6": "2001:DB8:1000::100",
		"generic_info": {
			"balancing_algorithm": "roundrobin",
			"ssl_redirect": false,
			"proxy_protocol": false,
			"vpc": "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
			"sticky_sessions": {
				"cookie_name": "my-cookie"
			}
		},
		"health_check": {
			"protocol": "http",
			"port": 80,
			"path": "/",
			"check_interval": 15,
			"response_timeout": 5,
			"unhealthy_threshold": 5,
			"healthy_threshold": 5
		},
		"has_ssl": false,
		"forwarding_rules": [
			{
				"id": "abcd12345",
				"frontend_protocol": "http",
				"frontend_port": 80,
				"backend_protocol": "http",
				"backend_port": 80
			}
		],
		"firewall_rules": [
			{
				"id": "abcd12345",
				"port": 80,
				"source": "0.0.0.0/0",
				"ip_type": "v4"
			}
		],
		"nodes": 3,
		"instances": [
			"12345"
		]
	}
}`))

	info, _, err := client.LoadBalancer.Get(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390")
	if err != nil {
		t.Errorf("LoadBalancer.Get returned %+v", err)
	}

	expected := &LoadBalancer{
		ID:          "ebfc2020-8be8-485f-8e01-74f75b24e390",
		DateCreated: "2020-01-07 17:24:23",
		Label:       "my label",
		Status:      "active",
		Region:      "ewr",
		IPV6:        "2001:DB8:1000::100",
		IPV4:        "123.123.123.123",
		SSLInfo:     BoolToBoolPtr(false),
		ForwardingRules: []ForwardingRule{
			{
				RuleID:           "abcd12345",
				FrontendProtocol: "http",
				FrontendPort:     80,
				BackendProtocol:  "http",
				BackendPort:      80,
			},
		},
		GenericInfo: &GenericInfo{
			BalancingAlgorithm: "roundrobin",
			SSLRedirect:        BoolToBoolPtr(false),
			ProxyProtocol:      BoolToBoolPtr(false),
			VPC:                "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
			StickySessions: &StickySessions{
				CookieName: "my-cookie",
			},
		},
		HealthCheck: &HealthCheck{
			Protocol:           "http",
			Port:               80,
			Path:               "/",
			CheckInterval:      15,
			ResponseTimeout:    5,
			UnhealthyThreshold: 5,
			HealthyThreshold:   5,
		},
		Instances: []string{"12345"},
		Nodes:     3,
		FirewallRules: []LBFirewallRule{
			{
				RuleID: "abcd12345",
				Port:   80,
				Source: "0.0.0.0/0",
				IPType: "v4",
			},
		},
	}

	if !reflect.DeepEqual(info, expected) {
		t.Errorf("LoadBalancer.Get returned %+v, expected %+v", info, expected)
	}
}

func TestLoadBalancerHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers", testJSONResponseHandlerFunc(http.StatusAccepted, `
{
	"load_balancer" :
		{
			"id": "ebfc2020-8be8-485f-8e01-74f75b24e390",
			"date_created": "2020-01-07 17:24:23",
			"region": "ewr",
			"label": "my label",
			"status": "active",
			"ipv4": "123.123.123.123",
			"ipv6": "2001:DB8:1000::100",
			"generic_info": {
				"balancing_algorithm": "roundrobin",
				"ssl_redirect": false,
				"proxy_protocol": false,
				"vpc": "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
				"sticky_sessions": {
					"cookie_name": "my-cookie"
				}
			},
			"health_check": {
				"protocol": "http",
				"port": 80,
				"path": "/",
				"check_interval": 15,
				"response_timeout": 5,
				"unhealthy_threshold": 5,
				"healthy_threshold": 5
			},
			"has_ssl": false,
			"forwarding_rules": [
				{
					"id": "abcd123",
					"frontend_protocol": "http",
					"frontend_port": 80,
					"backend_protocol": "http",
					"backend_port": 80
				}
			],
			"firewall_rules": [
				{
					"id": "abcd123",
					"port": 80,
					"source": "0.0.0.0/0",
					"ip_type": "v4"
				}
			],
			"nodes": 3,
			"instances": [
				"1234"
			]
		}
}`))

	lbCreate := &LoadBalancerReq{
		Label:  "my label",
		Region: "ewr",
		ForwardingRules: []ForwardingRule{
			{
				FrontendProtocol: "http",
				FrontendPort:     80,
				BackendProtocol:  "http",
				BackendPort:      80,
			},
		},
		BalancingAlgorithm: "roundrobin",
		SSLRedirect:        BoolToBoolPtr(false),
		ProxyProtocol:      BoolToBoolPtr(false),
		Nodes:              3,
		VPC:                StringToStringPtr("8d5bdbdb-3324-4d0c-b303-03e1315e1c02"),
		HealthCheck: &HealthCheck{
			Protocol:           "http",
			Port:               80,
			Path:               "/",
			CheckInterval:      15,
			ResponseTimeout:    5,
			UnhealthyThreshold: 5,
			HealthyThreshold:   5,
		},
	}

	lb, _, err := client.LoadBalancer.Create(ctx, lbCreate)
	if err != nil {
		t.Errorf("LoadBalancer.Create returned %+v", err)
	}

	expected := &LoadBalancer{
		ID:          "ebfc2020-8be8-485f-8e01-74f75b24e390",
		DateCreated: "2020-01-07 17:24:23",
		Label:       "my label",
		Status:      "active",
		Region:      "ewr",
		IPV6:        "2001:DB8:1000::100",
		IPV4:        "123.123.123.123",
		SSLInfo:     BoolToBoolPtr(false),
		ForwardingRules: []ForwardingRule{
			{
				RuleID:           "abcd123",
				FrontendProtocol: "http",
				FrontendPort:     80,
				BackendProtocol:  "http",
				BackendPort:      80,
			},
		},
		GenericInfo: &GenericInfo{
			BalancingAlgorithm: "roundrobin",
			SSLRedirect:        BoolToBoolPtr(false),
			ProxyProtocol:      BoolToBoolPtr(false),
			VPC:                "8d5bdbdb-3324-4d0c-b303-03e1315e1c02",
			StickySessions: &StickySessions{
				CookieName: "my-cookie",
			},
		},
		HealthCheck: &HealthCheck{
			Protocol:           "http",
			Port:               80,
			Path:               "/",
			CheckInterval:      15,
			ResponseTimeout:    5,
			UnhealthyThreshold: 5,
			HealthyThreshold:   5,
		},
		Instances: []string{"1234"},
		Nodes:     3,
		FirewallRules: []LBFirewallRule{
			{
				RuleID: "abcd123",
				Port:   80,
				Source: "0.0.0.0/0",
				IPType: "v4",
			},
		},
	}

	if !reflect.DeepEqual(lb, expected) {
		t.Errorf("LoadBalancer.Create returned %+v, expected %+v", lb, expected)
	}
}

func TestLoadBalancerHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	lbUpdate := &LoadBalancerReq{
		Label:  "my label",
		Region: "ewr",
		ForwardingRules: []ForwardingRule{
			{
				RuleID:           "abcd12345",
				FrontendProtocol: "http",
				FrontendPort:     80,
				BackendProtocol:  "http",
				BackendPort:      80,
			},
		},
		BalancingAlgorithm: "roundrobin",
		SSLRedirect:        BoolToBoolPtr(false),
		ProxyProtocol:      BoolToBoolPtr(false),
		VPC:                StringToStringPtr("8d5bdbdb-3324-4d0c-b303-03e1315e1c02"),
		Nodes:              5,
		HealthCheck: &HealthCheck{
			Protocol:           "http",
			Port:               80,
			Path:               "/",
			CheckInterval:      15,
			ResponseTimeout:    5,
			UnhealthyThreshold: 5,
			HealthyThreshold:   5,
		},
	}

	err := client.LoadBalancer.Update(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", lbUpdate)
	if err != nil {
		t.Errorf("LoadBalancer.Update returned %+v", err)
	}
}

func TestLoadBalancerHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.LoadBalancer.Delete(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390"); err != nil {
		t.Errorf("LoadBalancer.Delete returned %+v", err)
	}
}

func TestLoadBalancerHandler_ListForwardingRules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390/forwarding-rules", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"forwarding_rules":[
		{
			"id":"0690a322c25890bc",
			"frontend_protocol":"http",
			"frontend_port":80,
			"backend_protocol":"http",
			"backend_port":80
		}
	],
	"meta": {
		"total":8,
		"links": {
			"next":"",
			"prev":""
		}
	}
}`))

	list, meta, _, err := client.LoadBalancer.ListForwardingRules(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", nil)
	if err != nil {
		t.Errorf("LoadBalancer.ListForwardingRules returned %+v", err)
	}

	expected := []ForwardingRule{
		{
			RuleID:           "0690a322c25890bc",
			FrontendProtocol: "http",
			FrontendPort:     80,
			BackendProtocol:  "http",
			BackendPort:      80,
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("LoadBalancer.ListForwardingRules returned %+v, expected %+v", list, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("LoadBalancer.ListForwardingRules returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestLoadBalancerHandler_GetFowardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390/forwarding-rules/abc123",
		testJSONResponseHandlerFunc(http.StatusOK, `
{
	"forwarding_rule": {
		"id": "abc123",
		"frontend_protocol": "http",
		"frontend_port": 8080,
		"backend_protocol": "http",
		"backend_port": 80
	}
}`))

	rule, _, err := client.LoadBalancer.GetForwardingRule(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", "abc123")
	if err != nil {
		t.Errorf("LoadBalancer.GetForwardingRule returned %+v", err)
	}

	expected := &ForwardingRule{
		RuleID:           "abc123",
		FrontendProtocol: "http",
		FrontendPort:     8080,
		BackendProtocol:  "http",
		BackendPort:      80,
	}

	if !reflect.DeepEqual(rule, expected) {
		t.Errorf("LoadBalancer.GetForwardingRule returned %+v, expected %+v", rule, expected)
	}
}

func TestLoadBalancerHandler_CreateForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390/forwarding-rules", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"forwarding_rule" : {
		"id":"0690a322c25890bc",
		"frontend_protocol":"http",
		"frontend_port":80,
		"backend_protocol":"http",
		"backend_port":80
	}
}`))

	rule := &ForwardingRule{
		RuleID:           "0690a322c25890bc",
		FrontendProtocol: "http",
		FrontendPort:     80,
		BackendProtocol:  "http",
		BackendPort:      80,
	}

	ruleID, _, err := client.LoadBalancer.CreateForwardingRule(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", rule)
	if err != nil {
		t.Errorf("LoadBalancer.CreateForwardingRule returned %+v", err)
	}

	expected := &ForwardingRule{
		RuleID:           "0690a322c25890bc",
		FrontendProtocol: "http",
		FrontendPort:     80,
		BackendProtocol:  "http",
		BackendPort:      80,
	}

	if !reflect.DeepEqual(ruleID, expected) {
		t.Errorf("LoadBalancer.CreateForwardingRule returned %+v, expected %+v", ruleID, expected)
	}
}

func TestLoadBalancerHandler_DeleteForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390/forwarding-rules/abcde123",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	if err := client.LoadBalancer.DeleteForwardingRule(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", "abcde123"); err != nil {
		t.Errorf("LoadBalancer.DeleteForwardingRule returned %+v", err)
	}
}

func TestLoadBalancerHandler_GetFirewallRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/load-balancers/ebfc2020-8be8-485f-8e01-74f75b24e390/firewall-rules/abc123",
		testJSONResponseHandlerFunc(http.StatusOK, `
{
		"firewall_rule": {
		"id": "abc123",
		"port": 80,
		"source": "0.0.0.0/0",
		"ip_type": "v4"
	}
}`))

	rule, _, err := client.LoadBalancer.GetFirewallRule(ctx, "ebfc2020-8be8-485f-8e01-74f75b24e390", "abc123")
	if err != nil {
		t.Errorf("LoadBalancer.GetFirewallRule returned %+v", err)
	}

	expected := &LBFirewallRule{
		RuleID: "abc123",
		Port:   80,
		Source: "0.0.0.0/0",
		IPType: "v4",
	}

	if !reflect.DeepEqual(rule, expected) {
		t.Errorf("LoadBalancer.GetFirewallRule returned %+v, expected %+v", rule, expected)
	}
}
