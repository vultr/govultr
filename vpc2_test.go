package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVPCServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpc": {
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

	options := &VPCReq{
		Region:      "ewr",
		Description: "test1",
	}

	net, _, err := client.VPC.Create(ctx, options)

	if err != nil {
		t.Errorf("VPC.Create returned %+v, expected %+v", err, nil)
	}

	expected := &VPC{
		ID:           "net539626f0798d7",
		Region:       "ewr",
		Description:  "test1",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2017-08-25 12:23:45",
	}

	if !reflect.DeepEqual(net, expected) {
		t.Errorf("VPC.Create returned %+v, expected %+v", net, expected)
	}
}

func TestVPCServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.Delete(ctx, "net539626f0798d7")

	if err != nil {
		t.Errorf("VPC.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpcs": [{
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

	vpcs, _, _, err := client.VPC.List(ctx, nil)

	if err != nil {
		t.Errorf("VPC.List returned error: %v", err)
	}

	expected := []VPC{
		{
			ID:           "net539626f0798d7",
			Region:       "ewr",
			Description:  "test1",
			V4Subnet:     "10.99.0.0",
			V4SubnetMask: 24,
			DateCreated:  "2017-08-25 12:23:45",
		},
	}

	if !reflect.DeepEqual(vpcs, expected) {
		t.Errorf("VPC.List returned %+v, expected %+v", vpcs, expected)
	}
}

func TestVPCServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.Update(ctx, "net539626f0798d7", "update")

	if err != nil {
		t.Errorf("VPC.Update returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		req := `{"vpc": {"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60","date_created": "2020-10-10T01:56:20+00:00","region": "ewr","description": "sample desc","v4_subnet": "10.99.0.0","v4_subnet_mask": 24}}`
		fmt.Fprint(writer, req)
	})

	vpc, _, err := client.VPC.Get(ctx, "net539626f0798d7")
	if err != nil {
		t.Errorf("VPC.Get returned %+v, expected %+v", err, nil)
	}

	expected := &VPC{
		ID:           "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Region:       "ewr",
		Description:  "sample desc",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2020-10-10T01:56:20+00:00",
	}

	if !reflect.DeepEqual(vpc, expected) {
		t.Errorf("VPC.Get returned %+v, expected %+v", vpc, expected)
	}
}
