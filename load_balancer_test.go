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

	mux.HandleFunc("/v1/loadbalancer/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `[{"SUBID":1317575,"date_created":"2020-01-07 17:24:23","location":"New Jersey","label":"test","status":"active"}]`
		fmt.Fprintf(writer, response)
	})

	list, err := client.LoadBalancer.List(ctx)

	if err != nil {
		t.Errorf("LoadBalancer.List returned %+v", err)
	}

	expected := []LoadBalancers{
		{
			ID:          1317575,
			DateCreated: "2020-01-07 17:24:23",
			Location:    "New Jersey",
			Label:       "test",
			Status:      "active",
			RegionID:    0,
			IPV6:        "",
			IPV4:        "",
		},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("LoadBalancer.List returned %+v, expected %+v", list, expected)
	}
}

func TestLoadBalancerHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.Delete(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.Delete returned %+v", err)
	}
}

func TestLoadBalancerHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.SetLabel(ctx, 12345, "label")

	if err != nil {
		t.Errorf("LoadBalancer.SetLabel returned %+v", err)
	}
}

func TestLoadBalancerHandler_AttachedInstances(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"instance_list": ["1234", "2341"]}`
		fmt.Fprintf(writer, response)
	})

	instanceList, err := client.LoadBalancer.AttachedInstances(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.AttachedInstances returned %+v ", err)
	}

	expected := &InstanceList{InstanceList: []string{"1234", "2341"}}

	if !reflect.DeepEqual(instanceList, expected) {
		t.Errorf("LoadBalancer.AttachedInstances returned %+v, expected %+v", instanceList, expected)
	}
}

func TestLoadBalancerHandler_AttachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.AttachInstance(ctx, 12345, 45678)

	if err != nil {
		t.Errorf("LoadBalancer.AttachInstance returned %+v", err)
	}
}

func TestLoadBalancerHandler_DetachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.DetachInstance(ctx, 12345, 45678)

	if err != nil {
		t.Errorf("LoadBalancer.DetachInstance returned %+v", err)
	}
}

func TestLoadBalancerHandler_GetHealthCheck(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/health_check_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "protocol": "http","port": 81,"path": "/test","check_interval": 10,"response_timeout": 45,"unhealthy_threshold": 1,"healthy_threshold": 2}`
		fmt.Fprintf(writer, response)
	})

	health, err := client.LoadBalancer.GetHealthCheck(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.GetHealthCheck returned %+v ", err)
	}

	expected := &HealthCheck{
		Protocol:           "http",
		Port:               81,
		Path:               "/test",
		CheckInterval:      10,
		ResponseTimeout:    45,
		UnhealthyThreshold: 1,
		HealthyThreshold:   2,
	}

	if !reflect.DeepEqual(health, expected) {
		t.Errorf("LoadBalancer.GetHealthCheck returned %+v, expected %+v", health, expected)
	}
}

func TestLoadBalancerHandler_SetHealthCheck(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/health_check_update", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	health := &HealthCheck{
		Protocol:           "HTTPS",
		Port:               8080,
		Path:               "/health",
		CheckInterval:      4,
		ResponseTimeout:    5,
		UnhealthyThreshold: 3,
		HealthyThreshold:   4,
	}
	err := client.LoadBalancer.SetHealthCheck(ctx, 12345, health)

	if err != nil {
		t.Errorf("LoadBalancer.SetHealthCheck returned %+v", err)
	}
}

func TestLoadBalancerHandler_GetGenericInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/generic_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"balancing_algorithm":"roundrobin","ssl_redirect":false,"sticky_sessions":{"cookie_name":"test"}}`
		fmt.Fprintf(writer, response)
	})

	info, err := client.LoadBalancer.GetGenericInfo(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.GetGenericInfo returned %+v", err)
	}

	expected := &GenericInfo{
		BalancingAlgorithm: "roundrobin",
		SSLRedirect:        false,
		StickySessions:     CookieName{CookieName: "test"},
	}

	if !reflect.DeepEqual(info, expected) {
		t.Errorf("LoadBalancer.GetGenericInfo returned %+v, expected %+v", info, expected)
	}
}

func TestLoadBalancerHandler_ListForwardingRules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/forward_rule_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"forward_rule_list":[{"RULEID":"0690a322c25890bc","frontend_protocol":"http","frontend_port":80,"backend_protocol":"http","backend_port":80}]}`
		fmt.Fprintf(writer, response)
	})

	list, err := client.LoadBalancer.ListForwardingRules(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.ListForwardingRules returned %+v", err)
	}

	expected := &ForwardingRules{ForwardRuleList: []ForwardingRule{{
		RuleID:           "0690a322c25890bc",
		FrontendProtocol: "http",
		FrontendPort:     80,
		BackendProtocol:  "http",
		BackendPort:      80,
	}}}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("LoadBalancer.ListForwardingRules returned %+v, expected %+v", list, expected)
	}
}

func TestLoadBalancerHandler_DeleteForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/forward_rule_delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.DeleteForwardingRule(ctx, 12345, "abcde1234")

	if err != nil {
		t.Errorf("LoadBalancer.DeleteForwardingRule returned %+v", err)
	}
}

func TestLoadBalancerHandler_CreateForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/forward_rule_create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"RULEID": "abc123"}`
		fmt.Fprintf(writer, response)
	})

	rule := &ForwardingRule{
		FrontendProtocol: "http",
		FrontendPort:     8080,
		BackendProtocol:  "http",
		BackendPort:      8080,
	}
	ruleID, err := client.LoadBalancer.CreateForwardingRule(ctx, 123, rule)
	if err != nil {
		t.Errorf("LoadBalancer.CreateForwardingRule returned %+v", err)
	}

	expected := &ForwardingRule{
		RuleID: "abc123",
	}

	if !reflect.DeepEqual(ruleID, expected) {
		t.Errorf("LoadBalancer.CreateForwardingRule returned %+v, expected %+v", ruleID, expected)
	}
}

func TestLoadBalancerHandler_GetFullConfig(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/conf_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"generic_info":{"balancing_algorithm":"roundrobin","ssl_redirect":true,"sticky_sessions":{"cookie_name":"cookiename"}},"health_checks_info":{"protocol":"http","port":80,"path":"\/","check_interval":15,"response_timeout":5,"unhealthy_threshold":5,"healthy_threshold":5},"has_ssl":true,"forward_rule_list":[{"RULEID":"b06ce4cd520eea15","frontend_protocol":"http","frontend_port":80,"backend_protocol":"http","backend_port":80}],"instance_list":["1317615"]}`
		fmt.Fprintf(writer, response)
	})

	config, err := client.LoadBalancer.GetFullConfig(ctx, 123)
	if err != nil {
		t.Errorf("LoadBalancer.GetFullConfig returned %+v", err)
	}

	expected := &LBConfig{
		GenericInfo: GenericInfo{
			BalancingAlgorithm: "roundrobin",
			SSLRedirect:        true,
			StickySessions:     CookieName{CookieName: "cookiename"},
		},
		HealthCheck: HealthCheck{
			Protocol:           "http",
			Port:               80,
			Path:               "/",
			CheckInterval:      15,
			ResponseTimeout:    5,
			UnhealthyThreshold: 5,
			HealthyThreshold:   5,
		},
		SSLInfo: false,
		ForwardingRules: ForwardingRules{ForwardRuleList: []ForwardingRule{{
			RuleID:           "b06ce4cd520eea15",
			FrontendProtocol: "http",
			FrontendPort:     80,
			BackendProtocol:  "http",
			BackendPort:      80,
		}}},
		InstanceList: InstanceList{InstanceList: []string{"1317615"}},
	}

	if !reflect.DeepEqual(config, expected) {
		t.Errorf("LoadBalancer.GetFullConfigreturned %+v, expected %+v", config, expected)
	}
}
