package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServerServiceHandler_GetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/backup-schedule", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"backup_schedule":{"enabled": true,"type": "weekly","next_scheduled_time_utc": "2016-05-07 08:00:00","hour": 8,"dow": 6,"dom": 0}}`
		fmt.Fprint(writer, response)
	})

	backup, err := client.Instance.GetBackupSchedule(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetBackupSchedule returned %+v, ", err)
	}

	expected := &BackupSchedule{
		Enabled:             BoolToBoolPtr(true),
		Type:                "weekly",
		NextScheduleTimeUTC: "2016-05-07 08:00:00",
		Hour:                8,
		Dow:                 6,
		Dom:                 0,
	}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Instance.GetBackupSchedule returned %+v, expected %+v", backup, expected)
	}
}

func TestServerServiceHandler_SetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/backup-schedule", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"backup_schedule":{"enabled": true,"type": "weekly","next_scheduled_time_utc": "2016-05-07 08:00:00","hour": 22,"dow": 2,"dom": 3}}`
		fmt.Fprint(writer, response)
	})

	bs := &BackupScheduleReq{
		Type: "weekly",
		Hour: IntToIntPtr(22),
		Dow:  IntToIntPtr(2),
		Dom:  3,
	}

	if err := client.Instance.SetBackupSchedule(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", bs); err != nil {
		t.Errorf("Instance.SetBackupSchedule returned %+v, ", err)
	}
}

func TestServerServiceHandler_RestoreBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/restore", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	restoreReq := &RestoreReq{
		BackupID: "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
	}

	if err := client.Instance.Restore(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", restoreReq); err != nil {
		t.Errorf("Instance.Restore returned %+v, ", err)
	}
}

func TestServerServiceHandler_Neighbors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/neighbors", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"neighbors":["14b3e7d6-ffb5-4994-8502-57fcd9db3b33","14b3e7d6-ffb5-4994-8502-57fcd9db3b33"]}`
		fmt.Fprint(writer, response)
	})

	neighbors, err := client.Instance.GetNeighbors(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.Neighbors returned %+v, ", err)
	}

	expected := &Neighbors{
		Neighbors: []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"},
	}

	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Instance.Neighbors returned %+v, expected %+v", neighbors, expected)
	}
}

func TestServerServiceHandler_ListPrivateNetworks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/private-networks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"private_networks": [{"network_id": "v1-net539626f0798d7","mac_address": "5a:02:00:00:24:e9","ip_address": "10.99.0.3"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
		Cursor:  "",
	}
	privateNetwork, meta, err := client.Instance.ListPrivateNetworks(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", options)
	if err != nil {
		t.Errorf("Instance.ListPrivateNetworks return %+v, ", err)
	}

	expected := []PrivateNetwork{
		{
			NetworkID:  "v1-net539626f0798d7",
			MacAddress: "5a:02:00:00:24:e9",
			IPAddress:  "10.99.0.3",
		},
	}

	if !reflect.DeepEqual(privateNetwork, expected) {
		t.Errorf("Instance.ListPrivateNetworks returned %+v, expected %+v", privateNetwork, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.ListPrivateNetworks meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestServerServiceHandler_GetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/user-data", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"user_data": {"data" : "ZWNobyBIZWxsbyBXb3JsZA=="}}`
		fmt.Fprint(writer, response)
	})

	userData, err := client.Instance.GetUserData(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetUserData return %+v ", err)
	}

	expected := &UserData{Data: "ZWNobyBIZWxsbyBXb3JsZA=="}

	if !reflect.DeepEqual(userData, expected) {
		t.Errorf("Instance.GetUserData returned %+v, expected %+v", userData, expected)
	}
}

func TestServerServiceHandler_ListIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "ipv4s": [{"ip": "123.123.123.123","netmask": "255.255.255.248","gateway": "123.123.123.1","type": "main_ip","reverse": "host1.example.com"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	ipv4, meta, err := client.Instance.ListIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)

	if err != nil {
		t.Errorf("Instance.ListIPv4 returned %+v", err)
	}

	expected := []IPv4{
		{
			IP:      "123.123.123.123",
			Netmask: "255.255.255.248",
			Gateway: "123.123.123.1",
			Type:    "main_ip",
			Reverse: "host1.example.com",
		},
	}

	if !reflect.DeepEqual(ipv4, expected) {
		t.Errorf("Instance.ListIPv4 returned %+v, expected %+v", ipv4, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.ListIPv4 meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestServerServiceHandler_ListIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"ipv6s": [{"ip": "2001:DB8:1000::100","network": "2001:DB8:1000::","network_size": 64,"type": "main_ip"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	ipv6, meta, err := client.Instance.ListIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)
	if err != nil {
		t.Errorf("Instance.ListIPv6 returned %+v", err)
	}

	expected := []IPv6{
		{
			IP:          "2001:DB8:1000::100",
			Network:     "2001:DB8:1000::",
			NetworkSize: 64,
			Type:        "main_ip",
		},
	}

	if !reflect.DeepEqual(ipv6, expected) {
		t.Errorf("Instance.ListIPv6 returned %+v, expected %+v", ipv6, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.ListIPV6 meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestServerServiceHandler_CreateIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "ipv4": {"ip": "123.123.123.123","netmask": "255.255.255.248","gateway": "123.123.123.1","type": "main_ip","reverse": "host1.example.com"}}`
		fmt.Fprint(writer, response)
	})

	ipv4, err := client.Instance.CreateIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", BoolToBoolPtr(false))
	if err != nil {
		t.Errorf("Instance.CreateIPv4 returned %+v", err)
	}

	expected := &IPv4{
		IP:      "123.123.123.123",
		Netmask: "255.255.255.248",
		Gateway: "123.123.123.1",
		Type:    "main_ip",
		Reverse: "host1.example.com",
	}

	if !reflect.DeepEqual(ipv4, expected) {
		t.Errorf("Instance.CreateIPv4 returned %+v, expected %+v", ipv4, expected)
	}
}

func TestServerServiceHandler_DestroyIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/192.168.0.1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Instance.DeleteIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "192.168.0.1")

	if err != nil {
		t.Errorf("Instance.DestroyIPV4 returned %+v", err)
	}
}

func TestServerServiceHandler_GetBandwidth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/bandwidth", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"bandwidth": {
				"2017-04-01": {
					"incoming_bytes": 91571055,
					"outgoing_bytes": 3084731
				}
			}
		}		
		`
		fmt.Fprint(writer, response)
	})

	bandwidth, err := client.Instance.GetBandwidth(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetBandwidth returned %+v", err)
	}

	expected := &Bandwidth{
		Bandwidth: map[string]struct {
			IncomingBytes int `json:"incoming_bytes"`
			OutgoingBytes int `json:"outgoing_bytes"`
		}{
			"2017-04-01": {
				IncomingBytes: 91571055,
				OutgoingBytes: 3084731,
			},
		},
	}

	if !reflect.DeepEqual(bandwidth, expected) {
		t.Errorf("Instance.GetBandwidth returned %+v, expected %+v", bandwidth, expected)
	}
}

func TestServerServiceHandler_ListReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"reverse_ipv6s": [{"ip": "2001:DB8:1000::101","reverse": "host1.example.com"}]}`
		fmt.Fprint(writer, response)
	})

	reverseIPV6, err := client.Instance.ListReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

	if err != nil {
		t.Errorf("Instance.ListReverseIPv6 returned error: %v", err)
	}

	expected := []ReverseIP{
		{IP: "2001:DB8:1000::101", Reverse: "host1.example.com"},
	}

	if !reflect.DeepEqual(reverseIPV6, expected) {
		t.Errorf("Instance.ListReverseIPv6 returned %+v, expected %+v", reverseIPV6, expected)
	}
}

func TestServerServiceHandler_DefaultReverseIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/reverse/default", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.DefaultReverseIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "129.123.123.1"); err != nil {
		t.Errorf("Instance.DefaultReverseIPv4 returned %+v", err)
	}
}

func TestServerServiceHandler_DeleteReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse/2001:19f0:8001:1480:5400:2ff:fe00:8228", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.DeleteReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "2001:19f0:8001:1480:5400:2ff:fe00:8228"); err != nil {
		t.Errorf("Instance.DeleteReverseIPv6 returned %+v", err)
	}
}

func TestServerServiceHandler_CreateReverseIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/reverse", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	reverseReq := &ReverseIP{
		IP:      "192.168.0.1",
		Reverse: "test.com",
	}

	if err := client.Instance.CreateReverseIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", reverseReq); err != nil {
		t.Errorf("Instance.CreateReverseIPv4 returned %+v", err)
	}
}

func TestServerServiceHandler_CreateReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	reverseReq := &ReverseIP{
		IP:      "192.168.0.1",
		Reverse: "test.com",
	}

	if err := client.Instance.CreateReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", reverseReq); err != nil {
		t.Errorf("Instance.CreateReverseIPv6 returned %+v", err)
	}
}

func TestServerServiceHandler_Halt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/halt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.Halt(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.Halt returned %+v", err)
	}
}

func TestServerServiceHandler_Start(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/start", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.Start(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.Start returned %+v", err)
	}
}

func TestServerServiceHandler_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/reboot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Instance.Reboot(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

	if err != nil {
		t.Errorf("Instance.Reboot returned %+v", err)
	}
}

func TestServerServiceHandler_Reinstall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/reinstall", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instance": {
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "tagger",
				"os_id": 362,
				"app_id": 0,
				"firewall_group_id": "1234",
				"features": [
					"auto_backups", "ipv6"
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	instanceRes, err := client.Instance.Reinstall(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)
	if err != nil {
		t.Errorf("Instance.Reinstall returned %+v", err)
	}

	expected := &Instance{
		ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		Os:               "CentOS SELinux 8 x64",
		OsID:             362,
		RAM:              2048,
		Disk:             60,
		MainIP:           "123.123.123.123",
		VCPUCount:        2,
		Region:           "ewr",
		DefaultPassword:  "nreqnusibni",
		DateCreated:      "2013-12-19 14:45:41",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		Plan:             "vc2-1c-2gb",
		V6Network:        "2001:DB8:1000::",
		V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		Tag:              "tagger",
		AppID:            0,
		FirewallGroupID:  "1234",
		Features:         []string{"auto_backups", "ipv6"},
	}

	if !reflect.DeepEqual(instanceRes, expected) {
		t.Errorf("Instance.Create returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Instance.Delete(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

	if err != nil {
		t.Errorf("Instance.Delete returned %+v", err)
	}
}

func TestServerServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instance": {
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "tagger",
				"os_id": 362,
				"app_id": 0,
				"firewall_group_id": "1234",
				"features": [
					"auto_backups", "ipv6"
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &InstanceCreateReq{
		IPXEChainURL:    "test.org",
		ISOID:           "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		ScriptID:        "213",
		EnableIPv6:      BoolToBoolPtr(true),
		Backups:         "enabled",
		UserData:        "dW5vLWRvcy10cmVz",
		ActivationEmail: BoolToBoolPtr(true),
		DDOSProtection:  BoolToBoolPtr(true),
		SnapshotID:      "12ab",
		Hostname:        "hostname-3000",
		Tag:             "tagger",
		Label:           "label-extreme",
		SSHKeys:         []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "dev-preview-abc124"},
		ReservedIPv4:    "63.209.35.79",
		FirewallGroupID: "1234",
		AppID:           1,
	}

	server, err := client.Instance.Create(ctx, options)
	if err != nil {
		t.Errorf("Instance.Create returned %+v", err)
	}

	features := []string{"auto_backups", "ipv6"}

	expected := &Instance{
		ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		Os:               "CentOS SELinux 8 x64",
		OsID:             362,
		RAM:              2048,
		Disk:             60,
		MainIP:           "123.123.123.123",
		VCPUCount:        2,
		Region:           "ewr",
		DefaultPassword:  "nreqnusibni",
		DateCreated:      "2013-12-19 14:45:41",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		Plan:             "vc2-1c-2gb",
		V6Network:        "2001:DB8:1000::",
		V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		Tag:              "tagger",
		AppID:            0,
		FirewallGroupID:  "1234",
		Features:         features,
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.Create returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instances": [{
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "mytag",
				"os_id": 362,
				"app_id": 0,
				"firewall_group_id": "",
				"features": [
					"auto_backups"
				]
			}],
			"meta":{
				"total":1,
				"links":{
					"next":"thisismycusror",
					"prev":""
				}
			}			
		}`
		fmt.Fprint(writer, response)
	})

	server, meta, err := client.Instance.List(ctx, nil)
	if err != nil {
		t.Errorf("Instance.List returned %+v", err)
	}

	features := []string{"auto_backups"}

	expected := []Instance{
		{
			ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
			Os:               "CentOS SELinux 8 x64",
			OsID:             362,
			RAM:              2048,
			Disk:             60,
			MainIP:           "123.123.123.123",
			VCPUCount:        2,
			Region:           "ewr",
			DefaultPassword:  "nreqnusibni",
			DateCreated:      "2013-12-19 14:45:41",
			Status:           "active",
			AllowedBandwidth: 2000,
			NetmaskV4:        "255.255.255.248",
			GatewayV4:        "123.123.123.1",
			PowerStatus:      "running",
			ServerStatus:     "ok",
			Plan:             "vc2-1c-2gb",
			V6Network:        "2001:DB8:1000::",
			V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
			V6NetworkSize:    64,
			Label:            "my new server",
			InternalIP:       "10.99.0.10",
			KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
			Tag:              "mytag",
			AppID:            0,
			FirewallGroupID:  "",
			Features:         features,
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.List returned %+v, expected %+v", server, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestServerServiceHandler_GetServer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instance": {
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "mytag",
				"os_id": 362,
				"app_id": 0,
				"firewall_group_id": "",
				"features": [
					"auto_backups"
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Instance.Get(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetServer returned %+v", err)
	}

	features := []string{"auto_backups"}

	expected := &Instance{
		ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		Os:               "CentOS SELinux 8 x64",
		OsID:             362,
		RAM:              2048,
		Disk:             60,
		MainIP:           "123.123.123.123",
		VCPUCount:        2,
		Region:           "ewr",
		DefaultPassword:  "nreqnusibni",
		DateCreated:      "2013-12-19 14:45:41",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		Plan:             "vc2-1c-2gb",
		V6Network:        "2001:DB8:1000::",
		V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		Tag:              "mytag",
		AppID:            0,
		FirewallGroupID:  "",
		Features:         features,
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.GetServer returned %+v, expected %+v", server, expected)
	}
}

func TestInstanceServiceHandler_GetUpgrades(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/upgrades", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
   "upgrades":{
      "os":[
         {
            "id":127,
            "name":"CentOS 6 x64",
            "arch":"x64",
            "family":"centos"
         }
      ],
      "applications":[
         {
            "id":1,
            "name":"LEMP",
            "short_name":"lemp",
            "deploy_name":"LEMP on CentOS 6"
         }
      ],
      "plans":[
         "vc2-2c-4gb"
      ]
   }
}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Instance.GetUpgrades(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetUpgrades returned %+v", err)
	}

	expected := &Upgrades{
		Applications: []Application{
			{
				ID:         1,
				Name:       "LEMP",
				ShortName:  "lemp",
				DeployName: "LEMP on CentOS 6",
			},
		},
		OS: []OS{
			{
				ID:     127,
				Name:   "CentOS 6 x64",
				Arch:   "x64",
				Family: "centos",
			},
		},
		Plans: []string{
			"vc2-2c-4gb",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.GetUpgrades returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_MassStart(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/start", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.MassStart(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassStart returned %+v", err)
	}
}

func TestServerServiceHandler_MassReboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/reboot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.MassReboot(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassReboot returned %+v", err)
	}
}

func TestServerServiceHandler_MassHalt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/halt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.MassHalt(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassHalt returned %+v", err)
	}
}

func TestServerServiceHandler_AttachPrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/private-networks/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.AttachPrivateNetwork(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.AttachPrivateNetwork returned %+v", err)
	}
}

func TestServerServiceHandler_DetachPrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/private-networks/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.DetachPrivateNetwork(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.DetachPrivateNetwork returned %+v", err)
	}
}

func TestServerServiceHandler_ISOAttach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.AttachISO(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.AttachISO returned %+v", err)
	}
}

func TestServerServiceHandler_ISODetach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Instance.DetachISO(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.DetachISO returned %+v", err)
	}
}

func TestServerServiceHandler_ISO(t *testing.T) {
	setup()
	defer teardown()
	ret := `{"iso_status": {"state": "ready","iso_id": "0532a75b-14e8-48b8-b27e-1ebcf382a804"}}`
	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, ret)
	})

	iso, err := client.Instance.ISOStatus(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.ISOStatus returned %+v", err)
	}

	expected := &Iso{
		State: "ready",
		IsoID: "0532a75b-14e8-48b8-b27e-1ebcf382a804",
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Instance.ISOStatus returned %+v, expected %+v", iso, expected)
	}
}

func TestServerServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instance": {
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "tagger",
				"os_id": 362,
				"app_id": 0,
				"firewall_group_id": "1234",
				"features": [
					"auto_backups", "ipv6"
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &InstanceUpdateReq{
		EnableIPv6:      BoolToBoolPtr(true),
		Backups:         "enabled",
		UserData:        "dW5vLWRvcy10cmVz",
		DDOSProtection:  BoolToBoolPtr(true),
		Tag:             "tagger",
		Label:           "label-extreme",
		FirewallGroupID: "1234",
		AppID:           1,
	}

	server, err := client.Instance.Update(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", options)
	if err != nil {
		t.Errorf("Instance.Update returned %+v", err)
	}

	expected := &Instance{
		ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		Os:               "CentOS SELinux 8 x64",
		OsID:             362,
		RAM:              2048,
		Disk:             60,
		MainIP:           "123.123.123.123",
		VCPUCount:        2,
		Region:           "ewr",
		DefaultPassword:  "nreqnusibni",
		DateCreated:      "2013-12-19 14:45:41",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		Plan:             "vc2-1c-2gb",
		V6Network:        "2001:DB8:1000::",
		V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		Tag:              "tagger",
		AppID:            0,
		FirewallGroupID:  "1234",
		Features:         []string{"auto_backups", "ipv6"},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.Create returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_CreateMarketplace(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"instance": {
				"id": "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
				"os": "CentOS SELinux 8 x64",
				"ram": 2048,
				"disk": 60,
				"main_ip": "123.123.123.123",
				"vcpu_count": 2,
				"region": "ewr",
				"plan": "vc2-1c-2gb",
				"date_created": "2013-12-19 14:45:41",
				"status": "active",
				"allowed_bandwidth": 2000,
				"netmask_v4": "255.255.255.248",
				"gateway_v4": "123.123.123.1",
				"power_status": "running",
				"server_status": "ok",
				"v6_network": "2001:DB8:1000::",
				"v6_main_ip": "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
				"v6_network_size": 64,
				"label": "my new server",
				"internal_ip": "10.99.0.10",
				"kvm": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
				"default_password" : "nreqnusibni",
				"tag": "tagger",
				"os_id": 362,
				"image_id": "test",
				"app_id": 0,
				"firewall_group_id": "1234",
				"features": [
					"auto_backups", "ipv6"
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &InstanceCreateReq{
		IPXEChainURL:    "test.org",
		ISOID:           "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		ScriptID:        "213",
		EnableIPv6:      BoolToBoolPtr(true),
		Backups:         "enabled",
		UserData:        "dW5vLWRvcy10cmVz",
		ActivationEmail: BoolToBoolPtr(true),
		DDOSProtection:  BoolToBoolPtr(true),
		SnapshotID:      "12ab",
		Hostname:        "hostname-3000",
		Tag:             "tagger",
		Label:           "label-extreme",
		SSHKeys:         []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "dev-preview-abc124"},
		ReservedIPv4:    "63.209.35.79",
		FirewallGroupID: "1234",
		ImageID:         "test",
	}

	server, err := client.Instance.Create(ctx, options)
	if err != nil {
		t.Errorf("Instance.Create returned %+v", err)
	}

	features := []string{"auto_backups", "ipv6"}

	expected := &Instance{
		ID:               "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		Os:               "CentOS SELinux 8 x64",
		OsID:             362,
		RAM:              2048,
		Disk:             60,
		MainIP:           "123.123.123.123",
		VCPUCount:        2,
		Region:           "ewr",
		DefaultPassword:  "nreqnusibni",
		DateCreated:      "2013-12-19 14:45:41",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		Plan:             "vc2-1c-2gb",
		V6Network:        "2001:DB8:1000::",
		V6MainIP:         "fd11:1111:1112:1c02:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVM:              "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		Tag:              "tagger",
		AppID:            0,
		ImageID:          "test",
		FirewallGroupID:  "1234",
		Features:         features,
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.Create returned %+v, expected %+v", server, expected)
	}
}
