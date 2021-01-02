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

	net, err := client.Network.Create(ctx, options)

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
		fmt.Fprintf(writer, response)
	})

	networks, _, err := client.Network.List(ctx, nil)

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
