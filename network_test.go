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

	mux.HandleFunc("/v1/network/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"NETWORKID": "net59a0526477dd3"
		}
		`

		fmt.Fprint(writer, response)
	})

	net, err := client.Network.Create(ctx, "1", "go-test", "111.111.111.111/24")

	if err != nil {
		t.Errorf("Network.Create returned %+v, expected %+v", err, nil)
	}

	expected := &Network{
		NetworkID: "net59a0526477dd3",
	}

	if !reflect.DeepEqual(net, expected) {
		t.Errorf("Network.Create returned %+v, expected %+v", net, expected)
	}
}

func TestNetworkServiceHandler_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/network/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Network.Destroy(ctx, "foo")

	if err != nil {
		t.Errorf("Network.Destroy returned %+v, expected %+v", err, nil)
	}
}

func TestNetworkServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/network/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"net539626f0798d7": {
				"DCID": "1",
				"NETWORKID": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"description": "test1",
				"v4_subnet": "10.99.0.0",
				"v4_subnet_mask": 24
			},
			"net53962b0f2341f": {
				"DCID": "1",
				"NETWORKID": "net53962b0f2341f",
				"date_created": "2014-06-09 17:45:51",
				"description": "vultr",
				"v4_subnet": "0.0.0.0",
				"v4_subnet_mask": 0
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	networks, err := client.Network.GetList(ctx)

	if err != nil {
		t.Errorf("Network.GetList returned error: %v", err)
	}

	expected := []Network{
		{
			NetworkID:    "net539626f0798d7",
			RegionID:     "1",
			Description:  "test1",
			V4Subnet:     "10.99.0.0",
			V4SubnetMask: 24,
			DateCreated:  "2017-08-25 12:23:45",
		},
		{
			NetworkID:    "net53962b0f2341f",
			RegionID:     "1",
			Description:  "vultr",
			V4Subnet:     "0.0.0.0",
			V4SubnetMask: 0,
			DateCreated:  "2014-06-09 17:45:51",
		},
	}

	if !reflect.DeepEqual(networks, expected) {
		t.Errorf("Network.GetList returned %+v, expected %+v", networks, expected)
	}
}
