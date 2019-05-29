package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestReservedIPServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.Attach(ctx, "111.111.111.111", "1")

	if err != nil {
		t.Errorf("ReservedIP.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Convert(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/convert", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SUBID": 1312965
		}
		`

		fmt.Fprint(writer, response)
	})

	ip, err := client.ReservedIP.Convert(ctx, "111.111.111.111", "1", "go-test")

	if err != nil {
		t.Errorf("ReservedIP.Convert returned %+v, expected %+v", err, nil)
	}

	expected := &ReservedIP{
		ReservedIPID: "1312965",
		RegionID:     0,
		IPType:       "",
		Subnet:       "",
		SubnetSize:   0,
		Label:        "go-test",
		AttachedID:   "",
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Convert returned %+v, expected %+v", ip, expected)
	}
}

func TestReservedIPServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SUBID": 1312965
		}
		`

		fmt.Fprint(writer, response)
	})

	ip, err := client.ReservedIP.Create(ctx, 1, "v4", "go-test")

	if err != nil {
		t.Errorf("ReservedIP.Create returned %+v, expected %+v", err, nil)
	}

	expected := &ReservedIP{
		ReservedIPID: "1312965",
		RegionID:     1,
		IPType:       "v4",
		Subnet:       "",
		SubnetSize:   0,
		Label:        "go-test",
		AttachedID:   "",
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Create returned %+v, expected %+v", ip, expected)
	}
}

func TestReservedIPServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.Delete(ctx, "111.111.111.111")

	if err != nil {
		t.Errorf("ReservedIP.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.Detach(ctx, "111.111.111.111", "1")

	if err != nil {
		t.Errorf("ReservedIP.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/reservedip/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"1313044": {
				"SUBID": 1313044,
				"DCID": 1,
				"ip_type": "v4",
				"subnet": "10.234.22.53",
				"subnet_size": 32,
				"label": "my first reserved ip",
				"attached_SUBID": 123456
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	ips, err := client.ReservedIP.GetList(ctx)

	if err != nil {
		t.Errorf("ReservedIP.GetList returned error: %v", err)
	}

	expected := []ReservedIP{
		{
			ReservedIPID: "1313044",
			RegionID:     1,
			IPType:       "v4",
			Subnet:       "10.234.22.53",
			SubnetSize:   32,
			Label:        "my first reserved ip",
			AttachedID:   "123456",
		},
	}

	if !reflect.DeepEqual(ips, expected) {
		t.Errorf("ReservedIP.GetList returned %+v, expected %+v", ips, expected)
	}
}
