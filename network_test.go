package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNetworkServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/private-networks", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"network": {
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"v4_subnet": "10.99.0.0",
				"v4_subnet_mask": 24
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NetworkReq{
		Region:      "ewr",
		Description: "test1",
	}

	net,_, err := client.Network.Create(ctx, options)

	if err != nil {
		t.Errorf("Network.Create returned %+v, expected %+v", err, nil)
	}

	expected := &Network{
		NetworkID:    "net539626f0798d7",
		Region:       "ewr",
		Description:  "test1",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2017-08-25 12:23:45",
	}

	if !reflect.DeepEqual(net, expected) {
		t.Errorf("Network.Create returned %+v, expected %+v", net, expected)
	}
}

func TestNetworkServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/private-networks/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Network.Delete(ctx, "net539626f0798d7")

	if err != nil {
		t.Errorf("Network.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestNetworkServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/private-networks", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"networks": [{
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"v4_subnet": "10.99.0.0",
				"v4_subnet_mask": 24
			}]
		}
		`
		fmt.Fprint(writer, response)
	})

	networks, _, _, err := client.Network.List(ctx, nil)

	if err != nil {
		t.Errorf("Network.List returned error: %v", err)
	}

	expected := []Network{
		{
			NetworkID:    "net539626f0798d7",
			Region:       "ewr",
			Description:  "test1",
			V4Subnet:     "10.99.0.0",
			V4SubnetMask: 24,
			DateCreated:  "2017-08-25 12:23:45",
		},
	}

	if !reflect.DeepEqual(networks, expected) {
		t.Errorf("Network.List returned %+v, expected %+v", networks, expected)
	}
}

func TestNetworkServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/private-networks/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Network.Update(ctx, "net539626f0798d7", "update")

	if err != nil {
		t.Errorf("Network.Update returned %+v, expected %+v", err, nil)
	}
}

func TestNetworkServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/private-networks/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		req := `{"network": {"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60","date_created": "2020-10-10T01:56:20+00:00","region": "ewr","description": "sample desc","v4_subnet": "10.99.0.0","v4_subnet_mask": 24}}`
		fmt.Fprint(writer, req)
	})

	network,_,err := client.Network.Get(ctx, "net539626f0798d7")
	if err != nil {
		t.Errorf("Network.Get returned %+v, expected %+v", err, nil)
	}

	expected := &Network{
		NetworkID:    "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Region:       "ewr",
		Description:  "sample desc",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2020-10-10T01:56:20+00:00",
	}

	if !reflect.DeepEqual(network, expected) {
		t.Errorf("Instance.Get returned %+v, expected %+v", network, expected)
	}
}
