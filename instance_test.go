package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestInstanceServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"instances": [
		{
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"os": "CentOS SELinux 8 x64",
			"ram": 2048,
			"disk": 55,
			"main_ip": "192.0.2.123",
			"vcpu_count": 1,
			"region": "atl",
			"plan": "vc2-6c-16gb",
			"date_created": "2020-10-10T01:56:20+00:00",
			"status": "active",
			"allowed_bandwidth": 2000,
			"netmask_v4": "255.255.252.0",
			"gateway_v4": "192.0.2.1",
			"power_status": "running",
			"server_status": "ok",
			"v6_network": "2001:0db8:1112:18fb::",
			"v6_main_ip": "2001:0db8:1112:18fb:0200:00ff:fe00:0000",
			"v6_network_size": 64,
			"label": "Example Instance",
			"internal_ip": "",
			"vpc_only": false,
			"vpcs": [
			{
				"id": "775e26b3-f67d-46b7-87ed-1a0457fb3a5e",
				"version": 1,
				"subnet": "10.1.96.3"
			},
			{
				"id": "090a49c0-a1a2-4aab-a263-5d58f180c905",
				"version": 2,
				"subnet": "10.1.128.3"
			}
			],
			"kvm": "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
			"hostname": "my_hostname",
			"os_id": 215,
			"app_id": 0,
			"image_id": "",
			"snapshot_id": "",
			"firewall_group_id": "",
			"features": [
				"ddos_protection",
				"ipv6",
				"auto_backups"
			],
			"tags": [
				"a tag",
				"another"
			],
			"user_scheme": "root",
			"pending_charges": 5.42
		}
	],
	"meta": {
		"total": 3,
		"links": {
			"next": "WxYzExampleNext",
			"prev": ""
		}
	}
}`))

	server, meta, _, err := client.Instance.List(ctx, nil)
	if err != nil {
		t.Errorf("Instance.List returned %+v", err)
	}

	expected := []Instance{
		{
			ID:               "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			Os:               "CentOS SELinux 8 x64",
			RAM:              2048,
			Disk:             55,
			Plan:             "vc2-6c-16gb",
			MainIP:           "192.0.2.123",
			VPCOnly:          false,
			VCPUCount:        1,
			Region:           "atl",
			DateCreated:      "2020-10-10T01:56:20+00:00",
			Status:           "active",
			AllowedBandwidth: 2000,
			NetmaskV4:        "255.255.252.0",
			GatewayV4:        "192.0.2.1",
			PowerStatus:      "running",
			ServerStatus:     "ok",
			V6Network:        "2001:0db8:1112:18fb::",
			V6MainIP:         "2001:0db8:1112:18fb:0200:00ff:fe00:0000",
			V6NetworkSize:    64,
			Label:            "Example Instance",
			InternalIP:       "",
			KVM:              "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
			OsID:             215,
			AppID:            0,
			ImageID:          "",
			SnapshotID:       "",
			FirewallGroupID:  "",
			Features:         []string{"ddos_protection", "ipv6", "auto_backups"},
			Hostname:         "my_hostname",
			Tags:             []string{"a tag", "another"},
			UserScheme:       "root",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.List returned %+v, expected %+v", server, expected)
	}

	expectedMeta := &Meta{
		Total: 3,
		Links: &Links{
			Next: "WxYzExampleNext",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestInstanceServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/cb676a46-66fd-4dfb-b839-443f2e6c0b60", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"instance": {
		"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		"os": "CentOS SELinux 8 x64",
		"ram": 2048,
		"disk": 55,
		"main_ip": "192.0.2.123",
		"vcpu_count": 1,
		"region": "atl",
		"plan": "vc2-6c-16gb",
		"date_created": "2020-10-10T01:56:20+00:00",
		"status": "active",
		"allowed_bandwidth": 2000,
		"netmask_v4": "255.255.252.0",
		"gateway_v4": "192.0.2.1",
		"power_status": "running",
		"server_status": "ok",
		"v6_network": "2001:0db8:1112:18fb::",
		"v6_main_ip": "2001:0db8:1112:18fb:0200:00ff:fe00:0000",
		"v6_network_size": 64,
		"label": "Example Instance",
		"internal_ip": "",
		"vpc_only": false,
		"vpcs": [
		{
			"id": "775e26b3-f67d-46b7-87ed-1a0457fb3a5e",
			"version": 1,
			"subnet": "10.1.96.3"
		},
		{
			"id": "090a49c0-a1a2-4aab-a263-5d58f180c905",
			"version": 2,
			"subnet": "10.1.128.3"
		}
		],
		"kvm": "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		"hostname": "my_hostname",
		"os_id": 215,
		"app_id": 0,
		"image_id": "",
		"snapshot_id": "",
		"firewall_group_id": "",
		"features": [
			"ddos_protection",
			"ipv6",
			"auto_backups"
		],
		"tags": [
			"a tag",
			"another"
		],
		"user_scheme": "root",
		"pending_charges": 5.42
	}
}`))

	inst, _, err := client.Instance.Get(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("Instance.Get returned %+v", err)
	}

	expected := &Instance{
		ID:               "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Os:               "CentOS SELinux 8 x64",
		RAM:              2048,
		Disk:             55,
		Plan:             "vc2-6c-16gb",
		MainIP:           "192.0.2.123",
		VPCOnly:          false,
		VCPUCount:        1,
		Region:           "atl",
		DateCreated:      "2020-10-10T01:56:20+00:00",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "255.255.252.0",
		GatewayV4:        "192.0.2.1",
		PowerStatus:      "running",
		ServerStatus:     "ok",
		V6Network:        "2001:0db8:1112:18fb::",
		V6MainIP:         "2001:0db8:1112:18fb:0200:00ff:fe00:0000",
		V6NetworkSize:    64,
		Label:            "Example Instance",
		InternalIP:       "",
		KVM:              "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		OsID:             215,
		AppID:            0,
		ImageID:          "",
		SnapshotID:       "",
		FirewallGroupID:  "",
		Features:         []string{"ddos_protection", "ipv6", "auto_backups"},
		Hostname:         "my_hostname",
		Tags:             []string{"a tag", "another"},
		UserScheme:       "root",
	}

	if !reflect.DeepEqual(inst, expected) {
		t.Errorf("Instance.Get returned %+v, expected %+v", inst, expected)
	}
}

func TestInstanceServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"instance": {
		"id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		"os": "CentOS 8 Stream",
		"ram": 8192,
		"disk": 0,
		"main_ip": "0.0.0.0",
		"vcpu_count": 1,
		"region": "ewr",
		"plan": "vc2-4c-8gb",
		"date_created": "2021-09-14T13:22:20+00:00",
		"status": "pending",
		"allowed_bandwidth": 2000,
		"netmask_v4": "",
		"gateway_v4": "0.0.0.0",
		"power_status": "running",
		"server_status": "none",
		"v6_network": "",
		"v6_main_ip": "",
		"v6_network_size": 0,
		"label": "Example Instance",
		"internal_ip": "",
		"vpc_only": false,
		"kvm": "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		"hostname": "my_hostname",
		"os_id": 215,
		"app_id": 0,
		"image_id": "",
		"snapshot_id": "",
		"firewall_group_id": "",
		"features": [],
		"default_password": "v5{Fkvb#2ycPGwHs",
		"tags": [
			"a tag",
			"another"
		],
		"user_scheme": "root"
	}
}`))

	options := &InstanceCreateReq{
		Region:   "ewr",
		Plan:     "vc2-4c-8gb",
		Label:    "Example Instance",
		Backups:  "enabled",
		OsID:     215,
		UserData: "QmFzZTY0IEV4YW1wbGUgRGF0YQ==",
		Hostname: "my_hostname",
		Tags:     []string{"a tag", "another"},
	}

	inst, _, err := client.Instance.Create(ctx, options)
	if err != nil {
		t.Errorf("Instance.Create returned %+v", err)
	}

	expected := &Instance{
		ID:               "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		Os:               "CentOS 8 Stream",
		RAM:              8192,
		Disk:             0,
		Plan:             "vc2-4c-8gb",
		MainIP:           "0.0.0.0",
		VPCOnly:          false,
		VCPUCount:        1,
		Region:           "ewr",
		DateCreated:      "2021-09-14T13:22:20+00:00",
		Status:           "pending",
		AllowedBandwidth: 2000,
		NetmaskV4:        "",
		GatewayV4:        "0.0.0.0",
		PowerStatus:      "running",
		ServerStatus:     "none",
		V6Network:        "",
		V6MainIP:         "",
		V6NetworkSize:    0,
		Label:            "Example Instance",
		InternalIP:       "",
		KVM:              "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		OsID:             215,
		AppID:            0,
		ImageID:          "",
		SnapshotID:       "",
		FirewallGroupID:  "",
		Hostname:         "my_hostname",
		Tags:             []string{"a tag", "another"},
		UserScheme:       "root",
		DefaultPassword:  "v5{Fkvb#2ycPGwHs",
		Features:         []string{},
	}

	if !reflect.DeepEqual(inst, expected) {
		t.Errorf("Instance.Create returned %+v, expected %+v", inst, expected)
	}
}

func TestInstanceServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/4f0f12e5-1f84-404f-aa84-85f431ea5ec2", testJSONResponseHandlerFunc(http.StatusAccepted, `
{
	"instance": {
		"id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		"os": "CentOS 8 Stream",
		"ram": 8192,
		"disk": 0,
		"main_ip": "10.2.3.4",
		"vcpu_count": 1,
		"region": "ewr",
		"plan": "vc2-4c-8gb",
		"date_created": "2021-09-14T13:22:20+00:00",
		"status": "active",
		"allowed_bandwidth": 2000,
		"netmask_v4": "",
		"gateway_v4": "10.0.0.1",
		"power_status": "running",
		"server_status": "none",
		"v6_network": "",
		"v6_main_ip": "",
		"v6_network_size": 0,
		"label": "Example Instance",
		"internal_ip": "",
		"vpc_only": false,
		"kvm": "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		"hostname": "my_hostname",
		"os_id": 215,
		"app_id": 0,
		"image_id": "",
		"snapshot_id": "",
		"firewall_group_id": "a35eac93-9f56-4824-bb4e-bc3ac3814225",
		"features": [],
		"default_password": "",
		"tags": [
			"my tag"
		],
		"user_scheme": "root"
	}
}`))

	options := &InstanceUpdateReq{
		EnableIPv6:      BoolToBoolPtr(true),
		Tags:            []string{"my tag"},
		FirewallGroupID: "a35eac93-9f56-4824-bb4e-bc3ac3814225",
	}

	server, _, err := client.Instance.Update(ctx, "4f0f12e5-1f84-404f-aa84-85f431ea5ec2", options)
	if err != nil {
		t.Errorf("Instance.Update returned %+v", err)
	}

	expected := &Instance{
		ID:               "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		Os:               "CentOS 8 Stream",
		RAM:              8192,
		Disk:             0,
		Plan:             "vc2-4c-8gb",
		MainIP:           "10.2.3.4",
		VPCOnly:          false,
		VCPUCount:        1,
		Region:           "ewr",
		DateCreated:      "2021-09-14T13:22:20+00:00",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "",
		GatewayV4:        "10.0.0.1",
		PowerStatus:      "running",
		ServerStatus:     "none",
		V6Network:        "",
		V6MainIP:         "",
		V6NetworkSize:    0,
		Label:            "Example Instance",
		InternalIP:       "",
		KVM:              "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		OsID:             215,
		AppID:            0,
		ImageID:          "",
		SnapshotID:       "",
		FirewallGroupID:  "a35eac93-9f56-4824-bb4e-bc3ac3814225",
		Hostname:         "my_hostname",
		Tags:             []string{"my tag"},
		UserScheme:       "root",
		DefaultPassword:  "",
		Features:         []string{},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Instance.Update returned %+v, expected %+v", server, expected)
	}
}

func TestInstanceServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	err := client.Instance.Delete(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

	if err != nil {
		t.Errorf("Instance.Delete returned %+v", err)
	}
}

func TestInstanceServiceHandler_GetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/backup-schedule", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"backup_schedule": {
		"enabled": true,
		"type": "weekly",
		"next_scheduled_time_utc": "2016-05-07 08:00:00",
		"hour": 8,
		"dow": 6,
		"dom": 0
	}
}`))

	backup, _, err := client.Instance.GetBackupSchedule(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
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

func TestInstanceServiceHandler_SetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/backup-schedule", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"backup_schedule": {
		"enabled": true,
		"type": "weekly",
		"next_scheduled_time_utc": "2016-05-07 08:00:00",
		"hour": 22,
		"dow": 2,
		"dom": 3
	}
}`))

	bs := &BackupScheduleReq{
		Type: "weekly",
		Hour: IntToIntPtr(22),
		Dow:  IntToIntPtr(2),
		Dom:  3,
	}

	if _, err := client.Instance.SetBackupSchedule(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", bs); err != nil {
		t.Errorf("Instance.SetBackupSchedule returned %+v, ", err)
	}
}

func TestInstanceServiceHandler_RestoreBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/restore", testJSONResponseHandlerFunc(http.StatusAccepted, ""))

	restoreReq := &RestoreReq{
		BackupID: "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
	}

	if _, err := client.Instance.Restore(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", restoreReq); err != nil {
		t.Errorf("Instance.Restore returned %+v, ", err)
	}
}

func TestInstanceServiceHandler_RestoreSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/restore", testJSONResponseHandlerFunc(http.StatusAccepted, ""))

	restoreReq := &RestoreReq{
		SnapshotID: "14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
	}

	if _, err := client.Instance.Restore(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", restoreReq); err != nil {
		t.Errorf("Instance.Restore returned %+v, ", err)
	}
}

func TestInstanceServiceHandler_Neighbors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/neighbors", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"neighbors": [
		"14b3e7d6-ffb5-4994-8502-57fcd9db3b33",
		"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"
	]
}`))

	neighbors, _, err := client.Instance.GetNeighbors(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
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

func TestInstanceServiceHandler_ListVPCInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/vpcs", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"vpcs": [
		{
			"id": "v1-net539626f0798d7", 
			"mac_address": "5a:02:00:00:24:e9",
			"ip_address": "10.99.0.3"
		}
	],
	"meta": {
		"total":1,
		"links": {
			"next":"thisismycusror",
			"prev":""
		}
	}
}`))

	vpc, meta, _, err := client.Instance.ListVPCInfo(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)
	if err != nil {
		t.Errorf("Instance.ListVPCInfo returned %+v, ", err)
	}

	expected := []VPCInfo{
		{
			ID:         "v1-net539626f0798d7",
			MacAddress: "5a:02:00:00:24:e9",
			IPAddress:  "10.99.0.3",
		},
	}

	if !reflect.DeepEqual(vpc, expected) {
		t.Errorf("Instance.ListVPCInfo returned %+v, expected %+v", vpc, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Instance.ListVPCInfo meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestInstanceServiceHandler_GetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/user-data", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"user_data": {
		"data" : "ZWNobyBIZWxsbyBXb3JsZA=="
	}
}`))

	userData, _, err := client.Instance.GetUserData(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetUserData return %+v ", err)
	}

	expected := &UserData{Data: "ZWNobyBIZWxsbyBXb3JsZA=="}

	if !reflect.DeepEqual(userData, expected) {
		t.Errorf("Instance.GetUserData returned %+v, expected %+v", userData, expected)
	}
}

func TestInstanceServiceHandler_ListIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"ipv4s": [
		{
			"ip": "123.123.123.123",
			"netmask": "255.255.255.248",
			"gateway": "123.123.123.1",
			"type": "main_ip",
			"reverse": "host1.example.com"
		}
	],
	"meta": {
		"total":1,
		"links": {
			"next":"thisismycusror",
			"prev":""
		}
	}
}`))

	ipv4, meta, _, err := client.Instance.ListIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)

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

func TestInstanceServiceHandler_ListIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"ipv6s":  [
		{
			"ip":  "2001:DB8:1000::100",
			"network":  "2001:DB8:1000::",
			"network_size":  64,
			"type":  "main_ip"
		}
	],
	"meta": {
		"total": 1,
		"links": {
			"next": "thisismycusror",
			"prev": ""
		}
	}
}`))

	ipv6, meta, _, err := client.Instance.ListIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", nil)
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

func TestInstanceServiceHandler_CreateIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"ipv4": {
		"ip": "123.123.123.123",
		"netmask": "255.255.255.248",
		"gateway": "123.123.123.1",
		"type": "main_ip",
		"reverse": "host1.example.com"
	}
}`))

	ipv4, _, err := client.Instance.CreateIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", BoolToBoolPtr(false))
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

func TestInstanceServiceHandler_DestroyIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/192.168.0.1",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	err := client.Instance.DeleteIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "192.168.0.1")

	if err != nil {
		t.Errorf("Instance.DestroyIPV4 returned %+v", err)
	}
}

func TestInstanceServiceHandler_GetBandwidth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/bandwidth", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"bandwidth": {
		"2017-04-01": {
			"incoming_bytes": 91571055,
			"outgoing_bytes": 3084731
		}
	}
}`))

	bandwidth, _, err := client.Instance.GetBandwidth(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Instance.GetBandwidth returned %+v", err)
	}

	expected := &Bandwidth{
		Bandwidth: map[string]struct {
			IncomingBytes int64 `json:"incoming_bytes"`
			OutgoingBytes int64 `json:"outgoing_bytes"`
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

func TestInstanceServiceHandler_ListReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"reverse_ipv6s": [
		{
			"ip": "2001:DB8:1000::101",
			"reverse": "host1.example.com"
		}
	]
}`))

	reverseIPV6, _, err := client.Instance.ListReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

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

func TestInstanceServiceHandler_DefaultReverseIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/reverse/default",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	if err := client.Instance.DefaultReverseIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "172.123.123.1"); err != nil {
		t.Errorf("Instance.DefaultReverseIPv4 returned %+v", err)
	}
}

func TestInstanceServiceHandler_DeleteReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse/2001:19f0:8001:1480:5400:2ff:fe00:8228",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	if err := client.Instance.DeleteReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "2001:19f0:8001:1480:5400:2ff:fe00:8228"); err != nil {
		t.Errorf("Instance.DeleteReverseIPv6 returned %+v", err)
	}
}

func TestInstanceServiceHandler_CreateReverseIPv4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv4/reverse",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	reverseReq := &ReverseIP{
		IP:      "192.168.0.1",
		Reverse: "test.com",
	}

	if err := client.Instance.CreateReverseIPv4(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", reverseReq); err != nil {
		t.Errorf("Instance.CreateReverseIPv4 returned %+v", err)
	}
}

func TestInstanceServiceHandler_CreateReverseIPv6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/ipv6/reverse",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	reverseReq := &ReverseIP{
		IP:      "192.168.0.1",
		Reverse: "test.com",
	}

	if err := client.Instance.CreateReverseIPv6(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", reverseReq); err != nil {
		t.Errorf("Instance.CreateReverseIPv6 returned %+v", err)
	}
}

func TestInstanceServiceHandler_Halt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/halt",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	if err := client.Instance.Halt(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.Halt returned %+v", err)
	}
}

func TestInstanceServiceHandler_Start(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/start", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.Start(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.Start returned %+v", err)
	}
}

func TestInstanceServiceHandler_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/reboot", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	err := client.Instance.Reboot(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")

	if err != nil {
		t.Errorf("Instance.Reboot returned %+v", err)
	}
}

func TestInstanceServiceHandler_Reinstall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/reinstall", testJSONResponseHandlerFunc(http.StatusAccepted, `
{
	"instance": {
		"id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		"os": "CentOS 8 Stream",
		"ram": 8192,
		"disk": 0,
		"main_ip": "10.2.3.4",
		"vcpu_count": 1,
		"region": "ewr",
		"plan": "vc2-4c-8gb",
		"date_created": "2021-09-14T13:22:20+00:00",
		"status": "active",
		"allowed_bandwidth": 2000,
		"netmask_v4": "",
		"gateway_v4": "10.0.0.1",
		"power_status": "running",
		"server_status": "none",
		"v6_network": "",
		"v6_main_ip": "",
		"v6_network_size": 0,
		"label": "Example Instance",
		"internal_ip": "",
		"vpc_only": false,
		"kvm": "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		"hostname": "my_hostname_reinstalled",
		"os_id": 215,
		"app_id": 0,
		"image_id": "",
		"snapshot_id": "",
		"firewall_group_id": "a35eac93-9f56-4824-bb4e-bc3ac3814225",
		"features": [],
		"default_password": "",
		"tags": [
			"my tag"
		],
		"user_scheme": "root"
	}
}`))

	req := &ReinstallReq{
		Hostname: "my_hostname_reinstalled",
	}

	inst, _, err := client.Instance.Reinstall(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", req)
	if err != nil {
		t.Errorf("Instance.Reinstall returned %+v", err)
	}

	expected := &Instance{
		ID:               "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
		Os:               "CentOS 8 Stream",
		RAM:              8192,
		Disk:             0,
		Plan:             "vc2-4c-8gb",
		MainIP:           "10.2.3.4",
		VPCOnly:          false,
		VCPUCount:        1,
		Region:           "ewr",
		DateCreated:      "2021-09-14T13:22:20+00:00",
		Status:           "active",
		AllowedBandwidth: 2000,
		NetmaskV4:        "",
		GatewayV4:        "10.0.0.1",
		PowerStatus:      "running",
		ServerStatus:     "none",
		V6Network:        "",
		V6MainIP:         "",
		V6NetworkSize:    0,
		Label:            "Example Instance",
		InternalIP:       "",
		KVM:              "https://console.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
		OsID:             215,
		AppID:            0,
		ImageID:          "",
		SnapshotID:       "",
		FirewallGroupID:  "a35eac93-9f56-4824-bb4e-bc3ac3814225",
		Hostname:         "my_hostname_reinstalled",
		Tags:             []string{"my tag"},
		UserScheme:       "root",
		DefaultPassword:  "",
		Features:         []string{},
	}

	if !reflect.DeepEqual(inst, expected) {
		t.Errorf("Instance.Reinstall returned %+v, expected %+v", inst, expected)
	}
}

func TestInstanceServiceHandler_GetUpgrades(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/upgrades", testJSONResponseHandlerFunc(http.StatusOK, `
{
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
}`))

	ups, _, err := client.Instance.GetUpgrades(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
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

	if !reflect.DeepEqual(ups, expected) {
		t.Errorf("Instance.GetUpgrades returned %+v, expected %+v", ups, expected)
	}
}

func TestInstanceServiceHandler_MassStart(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/start", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.MassStart(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassStart returned %+v", err)
	}
}

func TestInstanceServiceHandler_MassReboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/reboot", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.MassReboot(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassReboot returned %+v", err)
	}
}

func TestInstanceServiceHandler_MassHalt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/halt", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.MassHalt(ctx, []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Instance.MassHalt returned %+v", err)
	}
}

func TestInstanceServiceHandler_AttachVPC(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/vpcs/attach", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.AttachVPC(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.AttachVPC returned %+v", err)
	}
}

func TestInstanceServiceHandler_DetachVPC(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/vpcs/detach", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.Instance.DetachVPC(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.DetachVPC returned %+v", err)
	}
}

func TestInstanceServiceHandler_ISOAttach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso/attach", testJSONResponseHandlerFunc(http.StatusAccepted, ""))

	if _, err := client.Instance.AttachISO(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.AttachISO returned %+v", err)
	}
}

func TestInstanceServiceHandler_ISODetach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso/detach", testJSONResponseHandlerFunc(http.StatusAccepted, ""))

	if _, err := client.Instance.DetachISO(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Instance.DetachISO returned %+v", err)
	}
}

func TestInstanceServiceHandler_ISOStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/iso", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"iso_status": {
		"state": "ready",
		"iso_id": "0532a75b-14e8-48b8-b27e-1ebcf382a804"
	}
}`))

	iso, _, err := client.Instance.ISOStatus(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
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
