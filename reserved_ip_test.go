package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestReservedIPServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"reserved_ips": [
		{
			"id": "1313044",
			"region": "ewr",
			"ip_type": "v4",
			"subnet": "10.234.22.53",
			"subnet_size": 32,
			"label": "my first reserved ip",
			"instance_id": "123456"
		}
	],
	"meta": {
		"total": 1,
		"links": {
			"next": "",
			"prev": ""
		}
	}
}`))

	ips, _, _, err := client.ReservedIP.List(ctx, nil)
	if err != nil {
		t.Errorf("ReservedIP.List returned error: %v", err)
	}

	expected := []ReservedIP{
		{
			ID:         "1313044",
			Region:     "ewr",
			IPType:     "v4",
			Subnet:     "10.234.22.53",
			SubnetSize: 32,
			Label:      "my first reserved ip",
			InstanceID: "123456",
		},
	}

	if !reflect.DeepEqual(ips, expected) {
		t.Errorf("ReservedIP.List returned %+v, expected %+v", ips, expected)
	}
}

func TestReservedIPServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"reserved_ip": {
		"id": "1313044",
		"region": "ewr",
		"ip_type": "v4",
		"subnet": "10.234.22.53",
		"subnet_size": 32,
		"label": "my first reserved ip",
		"instance_id": "123456"
	}
}`))

	ip, _, err := client.ReservedIP.Get(ctx, "1313044")
	if err != nil {
		t.Errorf("ReservedIP.Get returned error: %v", err)
	}

	expected := &ReservedIP{
		ID:         "1313044",
		Region:     "ewr",
		IPType:     "v4",
		Subnet:     "10.234.22.53",
		SubnetSize: 32,
		Label:      "my first reserved ip",
		InstanceID: "123456",
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Get returned %+v, expected %+v", ip, expected)
	}
}

func TestReservedIPServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"reserved_ip": {
		"id": "1313044",
		"region": "ewr",
		"ip_type": "v4",
		"subnet": "10.234.22.53",
		"subnet_size": 32,
		"label": "my first reserved ip",
		"instance_id": ""
	}
}`))

	options := &ReservedIPReq{
		IPType: "v4",
		Label:  "my first reserved ip",
		Region: "ewr",
	}

	ip, _, err := client.ReservedIP.Create(ctx, options)
	if err != nil {
		t.Errorf("ReservedIP.Create returned %+v, expected %+v", err, nil)
	}

	expected := &ReservedIP{
		ID:         "1313044",
		Region:     "ewr",
		IPType:     "v4",
		Subnet:     "10.234.22.53",
		SubnetSize: 32,
		Label:      "my first reserved ip",
		InstanceID: "",
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Create returned %+v, expected %+v", ip, expected)
	}
}

func TestReservedIPServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	options := &ReservedIPUpdateReq{
		Label: StringToStringPtr("my first reserved ip updated"),
	}

	mux.HandleFunc("/v2/reserved-ips/12345", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"reserved_ip": {
		"id": "12345",
		"region": "yto",
		"ip_type": "v4",
		"subnet": "10.234.22.53",
		"subnet_size": 32,
		"label": "my first reserved ip updated",
		"instance_id": "123456"
	}
}`))

	ip, _, err := client.ReservedIP.Update(ctx, "12345", options)
	expected := &ReservedIP{
		ID:         "12345",
		Region:     "yto",
		IPType:     "v4",
		Subnet:     "10.234.22.53",
		SubnetSize: 32,
		Label:      "my first reserved ip updated",
		InstanceID: "123456",
	}

	if err != nil {
		t.Errorf("ReservedIP.Update returned %+v, expected %+v", err, nil)
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Update returned %+v, expected %+v", ip, expected)
	}
}

func TestReservedIPServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/12345", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.ReservedIP.Delete(ctx, "12345"); err != nil {
		t.Errorf("ReservedIP.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/12345/attach", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.ReservedIP.Attach(ctx, "12345", "1234"); err != nil {
		t.Errorf("ReservedIP.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/12345/detach", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.ReservedIP.Detach(ctx, "12345"); err != nil {
		t.Errorf("ReservedIP.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Convert(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/convert", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"reserved_ip": {
		"id": "1312965",
		"region": "ewr",
		"ip_type": "v4",
		"subnet": "111.111.111.111",
		"subnet_size": 32,
		"label": "my first reserved ip",
		"instance_id": "1234"
	}
}`))

	options := &ReservedIPConvertReq{
		IPAddress: "111.111.111.111",
		Label:     "my first reserved ip",
	}
	ip, _, err := client.ReservedIP.Convert(ctx, options)
	if err != nil {
		t.Errorf("ReservedIP.Convert returned %+v, expected %+v", err, nil)
	}

	expected := &ReservedIP{
		ID:         "1312965",
		Region:     "ewr",
		IPType:     "v4",
		Subnet:     "111.111.111.111",
		SubnetSize: 32,
		Label:      "my first reserved ip",
		InstanceID: "1234",
	}

	if !reflect.DeepEqual(ip, expected) {
		t.Errorf("ReservedIP.Convert returned %+v, expected %+v", ip, expected)
	}
}
