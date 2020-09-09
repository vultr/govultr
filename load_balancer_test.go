package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLoadBalancerHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(lbPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"load_balancers" : [
				{
					"id": "1317575",
					"date_created": "2020-01-07 17:24:23",
					"region": "ewr",
					"label": "my label",
					"status": "active",
					"ipv4": "123.123.123.123",
					"ipv6": "2001:DB8:1000::100",
					"generic_info": {
						"balancing_algorithm": "roundrobin",
						"ssl_redirect": false,
						"proxy_protocol": "off",
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
		}
		`
		fmt.Fprintf(writer, response)
	})

	list, meta, err := client.LoadBalancer.List(ctx, nil)

	if err != nil {
		t.Errorf("LoadBalancer.List returned %+v", err)
	}

	expected := []LoadBalancer{
		{
			ID:          "1317575",
			DateCreated: "2020-01-07 17:24:23",
			Label:       "my label",
			Status:      "active",
			Region:      "ewr",
			IPV6:        "2001:DB8:1000::100",
			IPV4:        "123.123.123.123",
			SSLInfo:     false,
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
				SSLRedirect:        false,
				ProxyProtocol:      "off",
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

func TestLoadBalancerHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	uri := fmt.Sprintf("%s/%d", lbPath, 1317575)
	mux.HandleFunc(uri, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.LoadBalancer.Delete(ctx, "1317575"); err != nil {
		t.Errorf("LoadBalancer.Delete returned %+v", err)
	}
}

func TestLoadBalancerHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	uri := fmt.Sprintf("%s/%d", lbPath, 1317575)
	mux.HandleFunc(uri, func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"load_balancer" : {
				"id": "1317575",
				"date_created": "2020-01-07 17:24:23",
				"region": "ewr",
				"label": "my label",
				"status": "active",
				"ipv4": "123.123.123.123",
				"ipv6": "2001:DB8:1000::100",
				"generic_info": {
					"balancing_algorithm": "roundrobin",
					"ssl_redirect": false,
					"proxy_protocol": "off",
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
				"instances": [
					"12345"
				]
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	info, err := client.LoadBalancer.Get(ctx, "1317575")

	if err != nil {
		t.Errorf("LoadBalancer.Get returned %+v", err)
	}

	expected := &LoadBalancer{
		ID:          "1317575",
		DateCreated: "2020-01-07 17:24:23",
		Label:       "my label",
		Status:      "active",
		Region:      "ewr",
		IPV6:        "2001:DB8:1000::100",
		IPV4:        "123.123.123.123",
		SSLInfo:     false,
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
			SSLRedirect:        false,
			ProxyProtocol:      "off",
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
	}

	if !reflect.DeepEqual(info, expected) {
		t.Errorf("LoadBalancer.Get returned %+v, expected %+v", info, expected)
	}
}

func TestLoadBalancerHandler_ListForwardingRules(t *testing.T) {
	setup()
	defer teardown()
	uri := fmt.Sprintf("%s/%d/forwarding-rules", lbPath, 12345)
	mux.HandleFunc(uri, func(writer http.ResponseWriter, request *http.Request) {
		response := `{
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
		}
		`
		fmt.Fprintf(writer, response)
	})

	list, meta, err := client.LoadBalancer.ListForwardingRules(ctx, "12345", nil)

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

func TestLoadBalancerHandler_DeleteForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	uri := fmt.Sprintf("%s/%d/forwarding-rules/%s", lbPath, 12345, "abcde1234")
	mux.HandleFunc(uri, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.LoadBalancer.DeleteForwardingRule(ctx, "12345", "abcde1234"); err != nil {
		t.Errorf("LoadBalancer.DeleteForwardingRule returned %+v", err)
	}
}

func TestLoadBalancerHandler_CreateForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	uri := fmt.Sprintf("%s/%d/forwarding-rules", lbPath, 1317575)
	mux.HandleFunc(uri, func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"forwarding_rule" : {
				"id":"0690a322c25890bc",
				"frontend_protocol":"http",
				"frontend_port":80,
				"backend_protocol":"http",
				"backend_port":80
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	rule := &ForwardingRule{
		RuleID:           "0690a322c25890bc",
		FrontendProtocol: "http",
		FrontendPort:     80,
		BackendProtocol:  "http",
		BackendPort:      80,
	}

	ruleID, err := client.LoadBalancer.CreateForwardingRule(ctx, "1317575", rule)
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

func TestLoadBalancerHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/load-balancers", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"load_balancer" :
				{
					"id": "1317575",
					"date_created": "2020-01-07 17:24:23",
					"region": "ewr",
					"label": "my label",
					"status": "active",
					"ipv4": "123.123.123.123",
					"ipv6": "2001:DB8:1000::100",
					"generic_info": {
						"balancing_algorithm": "roundrobin",
						"ssl_redirect": false,
						"proxy_protocol": "off",
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
					"instances": [
						"1234"
					]
				}
		}
		`
		fmt.Fprintf(writer, response)
	})

	lbCreate := &LoadBalancerReq{
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
		SSLRedirect:        false,
		ProxyProtocol:      false,
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

	lb, err := client.LoadBalancer.Create(ctx, lbCreate)
	if err != nil {
		t.Errorf("LoadBalancer.Create returned %+v", err)
	}

	expected := &LoadBalancer{
		ID:          "1317575",
		DateCreated: "2020-01-07 17:24:23",
		Label:       "my label",
		Status:      "active",
		Region:      "ewr",
		IPV6:        "2001:DB8:1000::100",
		IPV4:        "123.123.123.123",
		SSLInfo:     false,
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
			SSLRedirect:        false,
			ProxyProtocol:      "off",
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
	}

	if !reflect.DeepEqual(lb, expected) {
		t.Errorf("LoadBalancer.Create returned %+v, expected %+v", lb, expected)
	}
}
