package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// get test
func TestReservedIPServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/12345/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.ReservedIP.Attach(ctx, "12345", "1234"); err != nil {
		t.Errorf("ReservedIP.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Convert(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/convert", func(writer http.ResponseWriter, request *http.Request) {
		response := `
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
		}
		`

		fmt.Fprint(writer, response)
	})

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

func TestReservedIPServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips", func(writer http.ResponseWriter, request *http.Request) {
		response := `
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
		}
		`
		fmt.Fprint(writer, response)
	})

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

	mux.HandleFunc("/v2/reserved-ips/12345", func(writer http.ResponseWriter, request *http.Request) {
		response := `
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
		}
		`
		fmt.Fprint(writer, response)
	})

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

	mux.HandleFunc("/v2/reserved-ips/12345", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.Delete(ctx, "12345")

	if err != nil {
		t.Errorf("ReservedIP.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/12345/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.Detach(ctx, "12345")

	if err != nil {
		t.Errorf("ReservedIP.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestReservedIPServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044", func(writer http.ResponseWriter, request *http.Request) {
		response := `
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
		}
		`
		fmt.Fprint(writer, response)
	})

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

func TestReservedIPServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"reserved_ips": [{
				"id": "1313044",
				"region": "ewr",
				"ip_type": "v4",
				"subnet": "10.234.22.53",
				"subnet_size": 32,
				"label": "my first reserved ip",
				"instance_id": "123456"
			}]
		}
		`
		fmt.Fprint(writer, response)
	})

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

func TestReservedIPServiceHandler_GetReverseDNS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044/reverse-dns", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "ipv6": [
			{
			  "ip": "2001:19f0:5401:98b::",
			  "domain": "0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.b.8.9.0.1.0.4.5.0.f.9.1.1.0.0.2.ip6.arpa"
			},
			{
			  "ip": "2001:19f0:5401:98b:ffff:ffff:ffff:ffff",
			  "domain": "f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.b.8.9.0.1.0.4.5.0.f.9.1.1.0.0.2.ip6.arpa"
			}
		  ]
		}`
		fmt.Fprint(writer, response)
	})

	rdns, _, err := client.ReservedIP.GetReverseDNS(ctx, "1313044")
	if err != nil {
		t.Errorf("ReservedIP.GetReverseDNS returned error: %v", err)
	}

	expected := &ReservedIPReverseDNS{
		IPv6: []ReservedIPReverseDNSIPv6{
			{
				IP:     "2001:19f0:5401:98b::",
				Domain: "0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.b.8.9.0.1.0.4.5.0.f.9.1.1.0.0.2.ip6.arpa",
			},
			{
				IP:     "2001:19f0:5401:98b:ffff:ffff:ffff:ffff",
				Domain: "f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.b.8.9.0.1.0.4.5.0.f.9.1.1.0.0.2.ip6.arpa",
			},
		},
	}

	if !reflect.DeepEqual(rdns, expected) {
		t.Errorf("ReservedIP.GetReverseDNS returned %+v, expected %+v", rdns, expected)
	}
}

func TestReservedIPServiceHandler_UpdateReverseDNS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044/reverse-dns/ipv4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &ReservedIPReverseDNSUpdateReq{
		V4: "test.com",
	}
	err := client.ReservedIP.UpdateReverseDNS(ctx, "1313044", req)
	if err != nil {
		t.Errorf("ReservedIP.UpdateReverseDNS returned error: %v", err)
	}
}

func TestReservedIPServiceHandler_CreateReverseDNS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044/reverse-dns/ipv6", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &ReservedIPReverseDNSCreateReq{
		V6: []ReservedIPReverseDNSIPv6{
			{
				IP:     "2001:19f0:5401:98b::",
				Domain: "test.com",
			},
		},
	}
	err := client.ReservedIP.CreateReverseDNS(ctx, "1313044", req)
	if err != nil {
		t.Errorf("ReservedIP.CreateReverseDNS returned error: %v", err)
	}
}

func TestReservedIPServiceHandler_DeleteReverseDNS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/reserved-ips/1313044/reverse-dns/ipv6/2001:19f0:5401:98b::", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ReservedIP.DeleteReverseDNS(ctx, "1313044", "2001:19f0:5401:98b::")
	if err != nil {
		t.Errorf("ReservedIP.DeleteReverseDNS returned error: %v", err)
	}
}
