package govultr

// import (
// 	"fmt"
// 	"net/http"
// 	"reflect"
// 	"testing"
// )

// func TestBareMetalServerServiceHandler_GetServer(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/abc123", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 		{
// 			"bare_metal": {
// 				"id": "abc123",
// 				"os": "CentOS 6 x64",
// 				"ram": "65536 MB",
// 				"disk": "2x 240 GB SSD",
// 				"main_ip": "203.0.113.10",
// 				"cpu_count": 1,
// 				"region": "ewr",
// 				"date_created": "2017-04-12 18:45:41",
// 				"status": "active",
// 				"netmask_v4": "255.255.255.0",
// 				"gateway_v4": "203.0.113.1",
// 				"plan": "vbm-4c-32gb",
// 				"v6_network": "2001:DB8:9000::",
// 				"v6_main_ip": "2001:DB8:9000::100",
// 				"v6_subnet": 64,
// 				"label": "my label",
// 				"tag": "my tag",
// 				"os_id": 127,
// 				"app_id": 0
// 			}
// 		}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	bm, err := client.BareMetalServer.Get(ctx, "abc123")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.GetServer returned error: %v", err)
// 	}

// 	expected := &BareMetalServer{
// 		ID:          "abc123",
// 		Os:          "CentOS 6 x64",
// 		RAM:         "65536 MB",
// 		Disk:        "2x 240 GB SSD",
// 		MainIP:      "203.0.113.10",
// 		CPUCount:    1,
// 		Region:      "ewr",
// 		DateCreated: "2017-04-12 18:45:41",
// 		Status:      "active",
// 		NetmaskV4:   "255.255.255.0",
// 		GatewayV4:   "203.0.113.1",
// 		Plan:        "vbm-4c-32gb",
// 		V6Network:   "2001:DB8:9000::",
// 		V6MainIP:    "2001:DB8:9000::100",
// 		V6Subnet:    64,
// 		Label:       "my label",
// 		Tag:         "my tag",
// 		OsID:        127,
// 		AppID:       0,
// 	}

// 	if !reflect.DeepEqual(bm, expected) {
// 		t.Errorf("BareMetalServer.GetServer returned %+v, expected %+v", bm, expected)
// 	}
// }

// func TestBareMetalServerServiceHandler_Create(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 			{
// 				"bare_metal": {
// 					"id": "900000",
// 					"os": "CentOS 6 x64",
// 					"ram": "65536 MB",
// 					"disk": "2x 240 GB SSD",
// 					"main_ip": "203.0.113.10",
// 					"cpu_count": 1,
// 					"region": "ewr",
// 					"default_password": "ab81u!ryranq",
// 					"date_created": "2017-04-12 18:45:41",
// 					"status": "active",
// 					"netmask_v4": "255.255.255.0",
// 					"gateway_v4": "203.0.113.1",
// 					"plan": "vbm-4c-32gb",
// 					"v6_network": "2001:DB8:9000::",
// 					"v6_main_ip": "2001:DB8:9000::100",
// 					"v6_subnet": 64,
// 					"label": "go-bm-test",
// 					"tag": "my tag",
// 					"os_id": 127,
// 					"app_id": 0
// 				}
// 			}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	options := &BareMetalReq{
// 		StartupScriptID: "1",
// 		Region:          "ewr",
// 		Plan:            "vbm-4c-32gb",
// 		SnapshotID:      "1",
// 		EnableIPV6:      "yes",
// 		Label:           "go-bm-test",
// 		SSHKeyIDs:       []string{"6b80207b1821f"},
// 		AppID:           1,
// 		UserData:        "echo Hello World",
// 		NotifyActivate:  "yes",
// 		Hostname:        "test",
// 		Tag:             "go-test",
// 		ReservedIPV4:    "111.111.111.111",
// 	}

// 	bm, err := client.BareMetalServer.Create(ctx, options)

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Create returned error: %v", err)
// 	}

// 	expected := &BareMetalServer{
// 		ID:              "900000",
// 		Os:              "CentOS 6 x64",
// 		RAM:             "65536 MB",
// 		Disk:            "2x 240 GB SSD",
// 		MainIP:          "203.0.113.10",
// 		CPUCount:        1,
// 		DefaultPassword: "ab81u!ryranq",
// 		DateCreated:     "2017-04-12 18:45:41",
// 		Status:          "active",
// 		NetmaskV4:       "255.255.255.0",
// 		GatewayV4:       "203.0.113.1",
// 		Plan:            "vbm-4c-32gb",
// 		V6Network:       "2001:DB8:9000::",
// 		V6MainIP:        "2001:DB8:9000::100",
// 		V6Subnet:        64,
// 		Label:           "go-bm-test",
// 		Tag:             "my tag",
// 		OsID:            127,
// 		Region:          "ewr",
// 		AppID:           0,
// 	}

// 	if !reflect.DeepEqual(bm, expected) {
// 		t.Errorf("BareMetalServer.Create returned %+v, expected %+v", bm, expected)
// 	}
// }

// func TestBareMetalServerServiceHandler_Update(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/123456", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprint(writer)
// 	})

// 	options := &BareMetalReq{
// 		Label: "my new label",
// 	}

// 	err := client.BareMetalServer.Update(ctx, "123456", options)
// 	if err != nil {
// 		t.Errorf("BareMetal.Update returned %+v, expected %+v", err, nil)
// 	}
// }
// func TestBareMetalServerServiceHandler_Delete(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprint(writer)
// 	})

// 	err := client.BareMetalServer.Delete(ctx, "900000")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Delete returned %+v, expected %+v", err, nil)
// 	}
// }

// func TestBareMetalServerServiceHandler_List(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 			{
// 				"bare_metals": [{
// 					"id": "90000",
// 					"os": "CentOS 6 x64",
// 					"ram": "65536 MB",
// 					"disk": "2x 240 GB SSD",
// 					"main_ip": "203.0.113.10",
// 					"cpu_count": 1,
// 					"region": "ewr",
// 					"date_created": "2017-04-12 18:45:41",
// 					"status": "active",
// 					"netmask_v4": "255.255.255.0",
// 					"gateway_v4": "203.0.113.1",
// 					"plan": "vbm-4c-32gb",
// 					"v6_network": "2001:DB8:9000::",
// 					"v6_main_ip": "2001:DB8:9000::100",
// 					"v6_subnet": 64,
// 					"label": "my label",
// 					"tag": "my tag",
// 					"os_id": 127,
// 					"app_id": 0
// 				}]
// 			}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	bm, _, err := client.BareMetalServer.List(ctx, nil)

// 	if err != nil {
// 		t.Errorf("BareMetalServer.List returned error: %v", err)
// 	}

// 	expected := []BareMetalServer{
// 		{
// 			ID:              "90000",
// 			Os:              "CentOS 6 x64",
// 			RAM:             "65536 MB",
// 			Disk:            "2x 240 GB SSD",
// 			MainIP:          "203.0.113.10",
// 			CPUCount:        1,
// 			Region:          "ewr",
// 			DateCreated:     "2017-04-12 18:45:41",
// 			Status:          "active",
// 			NetmaskV4:       "255.255.255.0",
// 			GatewayV4:       "203.0.113.1",
// 			Plan:            "vbm-4c-32gb",
// 			V6Network:       "2001:DB8:9000::",
// 			V6MainIP:        "2001:DB8:9000::100",
// 			V6Subnet:        64,
// 			Label:           "my label",
// 			Tag:             "my tag",
// 			OsID:            127,
// 			AppID:           0,
// 		},
// 	}

// 	if !reflect.DeepEqual(bm, expected) {
// 		t.Errorf("BareMetalServer.List returned %+v, expected %+v", bm, expected)
// 	}
// }

// func TestBareMetalServerServiceHandler_Bandwidth(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/123456/bandwidth", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 		{
// 			"bandwidth": {
// 				"2017-04-01": {
// 					"incoming_bytes": 91571055,
// 					"outgoing_bytes": 3084731
// 				}
// 			}
// 		}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	bandwidth, err := client.BareMetalServer.Bandwidth(ctx, "123456")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Bandwidth returned %+v", err)
// 	}

// 	expected := &BandwidthBase{
// 		map[string]BareMetalServerBandwidth{
// 			"2017-04-01": {
// 				IncomingBytes: 91571055,
// 				OutgoingBytes: 3084731,
// 			},
// 		}}

// 	if !reflect.DeepEqual(bandwidth.BareMetalBandwidth, expected.BareMetalBandwidth) {
// 		t.Errorf("BareMetalServer.Bandwidth returned %+v, expected %+v", bandwidth, expected.BareMetalBandwidth)
// 	}
// }

// func TestBareMetalServerServiceHandler_Halt(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000/halt", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprint(writer)
// 	})

// 	err := client.BareMetalServer.Halt(ctx, "900000")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Halt returned %+v, expected %+v", err, nil)
// 	}
// }

// func TestBareMetalServerServiceHandler_IPV4Info(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000/ipv4", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 		{
// 			"baremetal_ipv4s": [
// 				{
// 					"ip": "203.0.113.10",
// 					"netmask": "255.255.255.0",
// 					"gateway": "203.0.113.1",
// 					"type": "main_ip",
// 					"reverse": "203.0.113.10.vultr.com"
// 				}
// 			],
//   "meta": {
//     "total": 1,
//     "links": {
//       "next": "",
//       "prev": ""
//     }
//   }
// 		}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	ipv4, _, err := client.BareMetalServer.IPV4Info(ctx, "900000", nil)

// 	fmt.Print(ipv4)
// 	if err != nil {
// 		t.Errorf("BareMetalServer.IPV4Info returned %+v", err)
// 	}

// 	expected := []BareMetalServerIPV4{
// 		{
// 			IP:      "203.0.113.10",
// 			Netmask: "255.255.255.0",
// 			Gateway: "203.0.113.1",
// 			Type:    "main_ip",
// 			Reverse: "203.0.113.10.vultr.com",
// 		},
// 	}

// 	if !reflect.DeepEqual(ipv4, expected) {
// 		t.Errorf("BareMetalServer.IPV4Info returned %+v, expected %+v", ipv4, expected)
// 	}
// }

// func TestBareMetalServerServiceHandler_IPV6Info(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000/ipv6", func(writer http.ResponseWriter, request *http.Request) {
// 		response := `
// 		{
// 			"baremetal_ipv6s": [
// 				{
// 					"ip": "2001:DB8:9000::100",
// 					"network": "2001:DB8:9000::",
// 					"network_size": 64,
// 					"type": "main_ip"
// 				}
// 			]
// 		}
// 		`
// 		fmt.Fprint(writer, response)
// 	})

// 	ipv4, _, err := client.BareMetalServer.IPV6Info(ctx, "900000", nil)

// 	if err != nil {
// 		t.Errorf("BareMetalServer.IPV6Info returned %+v", err)
// 	}

// 	expected := []BareMetalServerIPV6{
// 		{
// 			IP:          "2001:DB8:9000::100",
// 			Network:     "2001:DB8:9000::",
// 			NetworkSize: 64,
// 			Type:        "main_ip",
// 		},
// 	}

// 	if !reflect.DeepEqual(ipv4, expected) {
// 		t.Errorf("BareMetalServer.IPV6Info returned %+v, expected %+v", ipv4, expected)
// 	}
// }

// func TestBareMetalServerServiceHandler_Reboot(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000/reboot", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprint(writer)
// 	})

// 	err := client.BareMetalServer.Reboot(ctx, "900000")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Reboot returned %+v, expected %+v", err, nil)
// 	}
// }

// func TestBareMetalServerServiceHandler_Reinstall(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/v2/bare-metals/900000/reinstall", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprint(writer)
// 	})

// 	err := client.BareMetalServer.Reinstall(ctx, "900000")

// 	if err != nil {
// 		t.Errorf("BareMetalServer.Reinstall returned %+v, expected %+v", err, nil)
// 	}
// }
