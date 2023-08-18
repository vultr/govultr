package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVPC2ServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpc": {
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"ip_block": "10.99.0.0",
				"prefix_length": 24
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &VPC2Req{
		Region:      "ewr",
		Description: "test1",
	}

	net, _, err := client.VPC2.Create(ctx, options)

	if err != nil {
		t.Errorf("VPC2.Create returned %+v, expected %+v", err, nil)
	}

	expected := &VPC2{
		ID:           "net539626f0798d7",
		Region:       "ewr",
		Description:  "test1",
		IPBlock:      "10.99.0.0",
		PrefixLength: 24,
		DateCreated:  "2017-08-25 12:23:45",
	}

	if !reflect.DeepEqual(net, expected) {
		t.Errorf("VPC2.Create returned %+v, expected %+v", net, expected)
	}
}

func TestVPC2ServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC2.Delete(ctx, "net539626f0798d7")

	if err != nil {
		t.Errorf("VPC2.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestVPC2ServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpcs": [{
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"ip_block": "10.99.0.0",
				"prefix_length": 24
			}]
		}
		`
		fmt.Fprint(writer, response)
	})

	vpcs, _, _, err := client.VPC2.List(ctx, nil)

	if err != nil {
		t.Errorf("VPC2.List returned error: %v", err)
	}

	expected := []VPC2{
		{
			ID:           "net539626f0798d7",
			Region:       "ewr",
			Description:  "test1",
			IPBlock:      "10.99.0.0",
			PrefixLength: 24,
			DateCreated:  "2017-08-25 12:23:45",
		},
	}

	if !reflect.DeepEqual(vpcs, expected) {
		t.Errorf("VPC2.List returned %+v, expected %+v", vpcs, expected)
	}
}

func TestVPC2ServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC2.Update(ctx, "net539626f0798d7", "update")

	if err != nil {
		t.Errorf("VPC2.Update returned %+v, expected %+v", err, nil)
	}
}

func TestVPC2ServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		req := `{"vpc": {"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60","date_created": "2020-10-10T01:56:20+00:00","region": "ewr","description": "sample desc","ip_block": "10.99.0.0","prefix_length": 24}}`
		fmt.Fprint(writer, req)
	})

	vpc, _, err := client.VPC2.Get(ctx, "net539626f0798d7")
	if err != nil {
		t.Errorf("VPC2.Get returned %+v, expected %+v", err, nil)
	}

	expected := &VPC2{
		ID:           "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Region:       "ewr",
		Description:  "sample desc",
		IPBlock:      "10.99.0.0",
		PrefixLength: 24,
		DateCreated:  "2020-10-10T01:56:20+00:00",
	}

	if !reflect.DeepEqual(vpc, expected) {
		t.Errorf("VPC2.Get returned %+v, expected %+v", vpc, expected)
	}
}

func TestVPC2ServiceHandler_ListNodes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/84fee086-6691-417a-b2db-e2a71061fa17/nodes", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"nodes": [
				{
					"id": "35dbcffe-58bf-46fe-bd68-964d95488dd8",
					"ip_address": "10.1.112.5",
					"mac_address": 98956121034033,
					"description": "bbbbbb-8ac448299844",
					"type": "vps",
					"node_status": "active"
				},
				{
					"id": "1f5d784a-1011-430c-a2e2-39ba045abe3c",
					"ip_address": "10.1.112.6",
					"mac_address": 98956121034034,
					"description": "bbbbbb-c76d8fc029d6",
					"type": "vps",
					"node_status": "active"
				}
			],
			"meta": {
				"total": 2,
				"links": {
					"next": "",
					"prev": ""
				}
			}
		}
		`
		fmt.Fprint(writer, response)
	})

	nodes, _, _, err := client.VPC2.ListNodes(ctx, "84fee086-6691-417a-b2db-e2a71061fa17", nil)

	if err != nil {
		t.Errorf("VPC2.ListNodes returned error: %v", err)
	}

	expected := []VPC2Node{
		{
			ID:          "35dbcffe-58bf-46fe-bd68-964d95488dd8",
			IPAddress:   "10.1.112.5",
			MACAddress:  98956121034033,
			Description: "bbbbbb-8ac448299844",
			Type:        "vps",
			NodeStatus:  "active",
		},
		{
			ID:          "1f5d784a-1011-430c-a2e2-39ba045abe3c",
			IPAddress:   "10.1.112.6",
			MACAddress:  98956121034034,
			Description: "bbbbbb-c76d8fc029d6",
			Type:        "vps",
			NodeStatus:  "active",
		},
	}

	if !reflect.DeepEqual(nodes, expected) {
		t.Errorf("VPC2.ListNode returned %+v, expected %+v", nodes, expected)
	}
}

func TestVPC2ServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/84fee086-6691-417a-b2db-e2a71061fa17/nodes/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	nodes := []string{"ce44b37a-bbe7-4e30-bfae-695c2e633bff", "45b794b7-4dd1-48b1-beb7-0b7bf3a16941"}
	options := &VPC2AttachDetachReq{
		Nodes: nodes,
	}

	err := client.VPC2.Attach(ctx, "84fee086-6691-417a-b2db-e2a71061fa17", options)

	if err != nil {
		t.Errorf("VPC2.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestVPC2ServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpc2/84fee086-6691-417a-b2db-e2a71061fa17/nodes/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	nodes := []string{"ce44b37a-bbe7-4e30-bfae-695c2e633bff", "45b794b7-4dd1-48b1-beb7-0b7bf3a16941"}
	options := &VPC2AttachDetachReq{
		Nodes: nodes,
	}

	err := client.VPC2.Detach(ctx, "84fee086-6691-417a-b2db-e2a71061fa17", options)

	if err != nil {
		t.Errorf("VPC2.Detach returned %+v, expected %+v", err, nil)
	}
}
