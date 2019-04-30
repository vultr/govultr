package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBareMetalServerServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SUBID": "900000"
		}
		`
		fmt.Fprint(writer, response)
	})

	options := &BareMetalServerOptions{
		StartupScriptID: "1",
		SnapshotID:      "1",
		EnableIPV6:      "yes",
		Label:           "go-bm-test",
		SSHKeyID:        "6b80207b1821f",
		AppID:           "1",
		UserData:        "ZWNobyBIZWxsbyBXb3JsZA==",
		NotifyActivate:  "yes",
		Hostname:        "test",
		Tag:             "go-test",
		ReservedIPV4:    "111.111.111.111",
	}

	bm, err := client.BareMetalServer.Create(ctx, "1", "1", "1", options)

	if err != nil {
		t.Errorf("BareMetalServer.Create returned error: %v", err)
	}

	expected := &BareMetalServer{BareMetalServerID: "900000"}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.Create returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"900000": {
					"SUBID": "900000",
					"os": "CentOS 6 x64",
					"ram": "65536 MB",
					"disk": "2x 240 GB SSD",
					"main_ip": "203.0.113.10",
					"cpu_count": 1,
					"location": "New Jersey",
					"DCID": "1",
					"default_password": "ab81u!ryranq",
					"date_created": "2017-04-12 18:45:41",
					"status": "active",
					"netmask_v4": "255.255.255.0",
					"gateway_v4": "203.0.113.1",
					"METALPLANID": 28,
					"v6_networks": [
						{
							"v6_network": "2001:DB8:9000::",
							"v6_main_ip": "2001:DB8:9000::100",
							"v6_network_size": 64
						}
					],
					"label": "my label",
					"tag": "my tag",
					"OSID": "127",
					"APPID": "0"
				}
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.GetList(ctx, "900000", "my tag", "my label", "203.0.113.10")

	if err != nil {
		t.Errorf("BareMetalServer.GetList returned error: %v", err)
	}

	expected := []BareMetalServer{
		{
			BareMetalServerID: "900000",
			Os:                "CentOS 6 x64",
			RAM:               "65536 MB",
			Disk:              "2x 240 GB SSD",
			MainIP:            "203.0.113.10",
			CPUCount:          1,
			Location:          "New Jersey",
			RegionID:          "1",
			DefaultPassword:   "ab81u!ryranq",
			DateCreated:       "2017-04-12 18:45:41",
			Status:            "active",
			NetmaskV4:         "255.255.255.0",
			GatewayV4:         "203.0.113.1",
			BareMetalPlanID:   28,
			V6Networks: []V6Network{
				{
					Network:     "2001:DB8:9000::",
					MainIP:      "2001:DB8:9000::100",
					NetworkSize: 64,
				},
			},
			Label: "my label",
			Tag:   "my tag",
			OsID:  "127",
			AppID: "0",
		},
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.Get returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_Halt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/halt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Halt(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Halt returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/reboot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Reboot(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Reboot returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_Reinstall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/reinstall", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Reinstall(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Reinstall returned %+v, expected %+v", err, nil)
	}
}
