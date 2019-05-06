package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServerServiceHandler_ChangeApp(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/app_change", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.ChangeApp(ctx, "1234", "24")

	if err != nil {
		t.Errorf("Server.ChangeApp returned %+v, ", err)
	}
}

func TestServerServiceHandler_ListApps(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/app_change_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1": {"APPID": "1","name": "LEMP","short_name": "lemp","deploy_name": "LEMP on CentOS 6 x64","surcharge": 0}}`
		fmt.Fprint(writer, response)
	})

	application, err := client.Server.ListApps(ctx, "1234")

	if err != nil {
		t.Errorf("Server.ListApps returned %+v, ", err)
	}

	expected := []Application{
		{
			AppID:      "1",
			Name:       "LEMP",
			ShortName:  "lemp",
			DeployName: "LEMP on CentOS 6 x64",
			Surcharge:  0,
		},
	}

	if !reflect.DeepEqual(application, expected) {
		t.Errorf("Server.ListApps returned %+v, expected %+v", application, expected)
	}
}

func TestServerServiceHandler_AppInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/get_app_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"app_info": "test"}`
		fmt.Fprint(writer, response)
	})

	appInfo, err := client.Server.AppInfo(ctx, "1234")

	if err != nil {
		t.Errorf("Server.AppInfo returned %+v, ", err)
	}

	expected := &AppInfo{AppInfo: "test"}

	if !reflect.DeepEqual(appInfo, expected) {
		t.Errorf("Server.AppInfo returned %+v, expected %+v", appInfo, expected)
	}
}

func TestServerServiceHandler_EnableBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.EnableBackup(ctx, "1234")

	if err != nil {
		t.Errorf("Server.EnableBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_DisableBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_disable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DisableBackup(ctx, "1234")

	if err != nil {
		t.Errorf("Server.DisableBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_GetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_get_schedule", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "enabled": true,"cron_type": "weekly","next_scheduled_time_utc": "2016-05-07 08:00:00","hour": 8,"dow": 6,"dom": 0}`
		fmt.Fprint(writer, response)
	})

	backup, err := client.Server.GetBackupSchedule(ctx, "1234")

	if err != nil {
		t.Errorf("Server.GetBackupSchedule returned %+v, ", err)
	}

	expected := &BackupSchedule{
		Enabled:  true,
		CronType: "weekly",
		NextRun:  "2016-05-07 08:00:00",
		Hour:     8,
		Dow:      6,
		Dom:      0,
	}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Server.GetBackupSchedule returned %+v, expected %+v", backup, expected)
	}
}

func TestServerServiceHandler_SetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_set_schedule", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	bs := &BackupSchedule{
		CronType: "",
		Hour:     23,
		Dow:      2,
		Dom:      3,
	}

	err := client.Server.SetBackupSchedule(ctx, "1234", bs)

	if err != nil {
		t.Errorf("Server.SetBackupSchedule returned %+v, ", err)
	}
}

func TestServerServiceHandler_RestoreBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/restore_backup", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.RestoreBackup(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.RestoreBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_RestoreSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/restore_snapshot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.RestoreSnapshot(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.RestoreSnapshot returned %+v, ", err)
	}
}

func TestServerServiceHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetLabel(ctx, "1234", "new-label")

	if err != nil {
		t.Errorf("Server.SetLabel returned %+v, ", err)
	}
}

func TestServerServiceHandler_SetTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/tag_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetTag(ctx, "1234", "new-tag")

	if err != nil {
		t.Errorf("Server.SetTag returned %+v, ", err)
	}
}

func TestServerServiceHandler_Neighbors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/neighbors", func(writer http.ResponseWriter, request *http.Request) {
		response := `[2345,1234]`
		fmt.Fprint(writer, response)
	})

	neighbors, err := client.Server.Neighbors(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Neighbors returned %+v, ", err)
	}

	expected := []int{2345, 1234}

	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Server.Neighbors returned %+v, expected %+v", neighbors, expected)
	}
}

func TestServerServiceHandler_EnablePrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_network_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.EnablePrivateNetwork(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.EnablePrivateNetwork returned %+v, ", err)
	}
}

func TestServerServiceHandler_DisablePrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_network_disable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DisablePrivateNetwork(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.DisablePrivateNetwork returned %+v, ", err)
	}
}

func TestServerServiceHandler_ListPrivateNetworks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_networks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"net539626f0798d7": {"NETWORKID": "net539626f0798d7","mac_address": "5a:02:00:00:24:e9","ip_address": "10.99.0.3"}}`
		fmt.Fprint(writer, response)
	})

	privateNetwork, err := client.Server.ListPrivateNetworks(ctx, "12345")

	if err != nil {
		t.Errorf("Server.ListPrivateNetworks return %+v, ", err)
	}

	expected := []PrivateNetwork{
		{
			NetworkID:  "net539626f0798d7",
			MacAddress: "5a:02:00:00:24:e9",
			IPAddress:  "10.99.0.3",
		},
	}

	if !reflect.DeepEqual(privateNetwork, expected) {
		t.Errorf("Server.ListPrivateNetworks returned %+v, expected %+v", privateNetwork, expected)
	}
}

func TestServerServiceHandler_ListUpgradePlan(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/upgrade_plan_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `[1, 2, 3, 4]`
		fmt.Fprint(writer, response)
	})

	plans, err := client.Server.ListUpgradePlan(ctx, "123")

	if err != nil {
		t.Errorf("Server.ListUpgradePlan return %+v ", err)
	}

	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(plans, expected) {
		t.Errorf("Server.ListUpgradePlan returned %+v, expected %+v", plans, expected)
	}
}

func TestServerServiceHandler_UpgradePlan(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/upgrade_plan", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.UpgradePlan(ctx, "12351", "123")

	if err != nil {
		t.Errorf("Server.UpgradePlan return %+v ", err)
	}
}

func TestServerServiceHandler_ListOS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/os_change_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"127": {"OSID": 127,"name": "CentOS 6 x64","arch": "x64","family": "centos","windows": false,"surcharge": "0.00"}}`
		fmt.Fprint(writer, response)
	})

	os, err := client.Server.ListOS(ctx, "1234")

	if err != nil {
		t.Errorf("Server.ListOS return %+v ", err)
	}

	expected := []OS{
		{
			OsID:    127,
			Name:    "CentOS 6 x64",
			Arch:    "x64",
			Family:  "centos",
			Windows: false,
		},
	}

	if !reflect.DeepEqual(os, expected) {
		t.Errorf("Server.ListOS returned %+v, expected %+v", os, expected)
	}
}

func TestServerServiceHandler_ChangeOS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/os_change", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.ChangeOS(ctx, "1234", "1")

	if err != nil {
		t.Errorf("Server.ChangeOS return %+v ", err)
	}
}

func TestServerServiceHandler_IsoAttach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/iso_attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.IsoAttach(ctx, "1234", "1")

	if err != nil {
		t.Errorf("Server.IsoAttach return %+v ", err)
	}
}

func TestServerServiceHandler_IsoDetach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/iso_detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.IsoDetach(ctx, "1234")

	if err != nil {
		t.Errorf("Server.IsoDetach return %+v ", err)
	}
}

func TestServerServiceHandler_IsoStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/iso_status", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"state": "ready","ISOID": "12345"}`
		fmt.Fprint(writer, response)
	})

	isoStatus, err := client.Server.IsoStatus(ctx, "1234")

	if err != nil {
		t.Errorf("Server.IsoStatus return %+v ", err)
	}

	expected := &ServerIso{State: "ready", IsoID: "12345"}

	if !reflect.DeepEqual(isoStatus, expected) {
		t.Errorf("Server.IsoStatus returned %+v, expected %+v", isoStatus, expected)
	}
}

func TestServerServiceHandler_SetFirewallGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/firewall_group_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetFirewallGroup(ctx, "1234", "123")

	if err != nil {
		t.Errorf("Server.SetFirewallGroup return %+v ", err)
	}
}

func TestServerServiceHandler_SetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/set_user_data", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetUserData(ctx, "1234", "user-test-data")

	if err != nil {
		t.Errorf("Server.SetUserData return %+v ", err)
	}
}

func TestServerServiceHandler_GetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/get_user_data", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"userdata": "ZWNobyBIZWxsbyBXb3JsZA=="}`
		fmt.Fprint(writer, response)
	})

	userData, err := client.Server.GetUserData(ctx, "1234")

	if err != nil {
		t.Errorf("Server.GetUserData return %+v ", err)
	}

	expected := &UserData{UserData: "ZWNobyBIZWxsbyBXb3JsZA=="}

	if !reflect.DeepEqual(userData, expected) {
		t.Errorf("Server.GetUserData returned %+v, expected %+v", userData, expected)
	}
}

func TestServerServiceHandler_IPV4Info(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "576965": [{"ip": "123.123.123.123","netmask": "255.255.255.248","gateway": "123.123.123.1","type": "main_ip","reverse": "host1.example.com"}]}`
		fmt.Fprint(writer, response)
	})

	ipv4, err := client.Server.IPV4Info(ctx, "1234", true)

	if err != nil {
		t.Errorf("Server.IPV4Info returned %+v", err)
	}

	expected := []IPV4{
		{
			IP:      "123.123.123.123",
			Netmask: "255.255.255.248",
			Gateway: "123.123.123.1",
			Type:    "main_ip",
			Reverse: "host1.example.com",
		},
	}

	if !reflect.DeepEqual(ipv4, expected) {
		t.Errorf("Server.IPV4Info returned %+v, expected %+v", ipv4, expected)
	}
}

func TestServerServiceHandler_IPV6Info(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list_ipv6", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": [{"ip": "2001:DB8:1000::100","network": "2001:DB8:1000::","network_size": "64","type": "main_ip"}]}`
		fmt.Fprint(writer, response)
	})

	ipv6, err := client.Server.IPV6Info(ctx, "1234")

	if err != nil {
		t.Errorf("Server.IPV6Info returned %+v", err)
	}

	expected := []IPV6{
		{
			IP:          "2001:DB8:1000::100",
			Network:     "2001:DB8:1000::",
			NetworkSize: "64",
			Type:        "main_ip",
		},
	}

	if !reflect.DeepEqual(ipv6, expected) {
		t.Errorf("Server.IPV6Info returned %+v, expected %+v", ipv6, expected)
	}
}

func TestServerServiceHandler_AddIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/create_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.AddIPV4(ctx, "1234")

	if err != nil {
		t.Errorf("Server.AddIPV4 returned %+v", err)
	}
}

func TestServerServiceHandler_DestroyIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/destroy_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DestroyIPV4(ctx, "1234", "192.168.0.1")

	if err != nil {
		t.Errorf("Server.DestroyIPV4 returned %+v", err)
	}
}

func TestServerServiceHandler_EnableIPV6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/ipv6_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.EnableIPV6(ctx, "1234")

	if err != nil {
		t.Errorf("Server.EnableIPV6 returned %+v", err)
	}
}

func TestServerServiceHandler_Bandwidth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/bandwidth", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"incoming_bytes": [["2014-06-10","81072581"]],"outgoing_bytes": [["2014-06-10","4059610"]]}`
		fmt.Fprint(writer, response)
	})

	bandwidth, err := client.Server.Bandwidth(ctx, "123")

	if err != nil {
		t.Errorf("Server.Bandwidth returned %+v", err)
	}

	expected := []map[string]string{
		{"date": "2014-06-10", "incoming": "81072581", "outgoing": "4059610"},
	}

	if !reflect.DeepEqual(bandwidth, expected) {
		t.Errorf("Server.Bandwidth returned %+v, expected %+v", bandwidth, expected)
	}
}

func TestServerServiceHandler_ListReverseIPV6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reverse_list_ipv6", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": [{"ip": "2001:DB8:1000::101","reverse": "host1.example.com"}]}`
		fmt.Fprint(writer, response)
	})

	reverseIPV6, err := client.Server.ListReverseIPV6(ctx, "123890")

	if err != nil {
		t.Errorf("Server.ListReverseIPV6 returned error: %v", err)
	}

	expected := []ReverseIPV6{
		{IP: "2001:DB8:1000::101", Reverse: "host1.example.com"},
	}

	if !reflect.DeepEqual(reverseIPV6, expected) {
		t.Errorf("Server.ListReverseIPV6returned %+v, expected %+v", reverseIPV6, expected)
	}
}

func TestServerServiceHandler_SetDefaultReverseIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reverse_default_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetDefaultReverseIPV4(ctx, "1234", "129.123.123.1")

	if err != nil {
		t.Errorf("Server.SetDefaultReverseIPV4 returned %+v", err)
	}
}

func TestServerServiceHandler_DeleteReverseIPV6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reverse_delete_ipv6", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DeleteReverseIPV6(ctx, "1234", "2001:19f0:8001:1480:5400:2ff:fe00:8228")

	if err != nil {
		t.Errorf("Server.DeleteReverseIPV6 returned %+v", err)
	}
}

func TestServerServiceHandler_SetReverseIPV4(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reverse_set_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetReverseIPV4(ctx, "1234", "192.168.0.1", "test.com")

	if err != nil {
		t.Errorf("Server.SetReverseIPV4 returned %+v", err)
	}
}

func TestServerServiceHandler_SetReverseIPV6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reverse_set_ipv6", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetReverseIPV6(ctx, "1234", "2001:19f0:8001:1480:5400:2ff:fe00:8228", "test.com")

	if err != nil {
		t.Errorf("Server.SetReverseIPV6 returned %+v", err)
	}
}

func TestServerServiceHandler_Halt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/halt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.Halt(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Halt returned %+v", err)
	}
}

func TestServerServiceHandler_Start(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/start", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.Start(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Start returned %+v", err)
	}
}

func TestServerServiceHandler_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reboot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.Reboot(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Reboot returned %+v", err)
	}
}

func TestServerServiceHandler_Reinstall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/reinstall", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.Reinstall(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Reinstall returned %+v", err)
	}
}

func TestServerServiceHandler_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.Destroy(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Destroy returned %+v", err)
	}
}

func TestServerServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SUBID": "1234151"}`
		fmt.Fprint(writer, response)
	})

	optionsWithPrivateNetwork := &ServerOptions{
		IPXEChain:            "test.org",
		IsoID:                1,
		ScriptID:             "213",
		EnableIPV6:           true,
		EnablePrivateNetwork: true,
		AutoBackups:          true,
		UserData:             "uno-dos-tres",
		NotifyActivate:       true,
		DDOSProtection:       true,
		Hostname:             "hostname-3000",
		Tag:                  "tagger",
		Label:                "label-extreme",
		SSHKeyID:             "1234",
		ReservedIPV4:         "63.209.35.79",
		FirewallGroupID:      "1234",
		AppID:                "1234",
	}

	server, err := client.Server.Create(ctx, 1, 2, 3, optionsWithPrivateNetwork)

	if err != nil {
		t.Errorf("Server.Create returned %+v", err)
	}

	expected := &Server{VpsID: "1234151"}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.Create returned %+v, expected %+v", server, expected)
	}

	options := &ServerOptions{
		NetworkID: []string{"1", "2", "3"},
	}

	serverWithNetwork, err := client.Server.Create(ctx, 1, 2, 3, options)

	if err != nil {
		t.Errorf("Server.Create returned %+v", err)
	}

	if !reflect.DeepEqual(serverWithNetwork, expected) {
		t.Errorf("Server.Create returned %+v, expected %+v", serverWithNetwork, expected)
	}
}

func TestServerServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": {"SUBID": "576965","os": "CentOS 6 x64","ram": "4096 MB","disk": "Virtual 60 GB","main_ip": "123.123.123.123","vcpu_count": "2","location": "New Jersey","DCID": "1","default_password": "nreqnusibni","date_created": "2013-12-19 14:45:41","pending_charges": "46.67","status": "active","cost_per_month": "10.05","current_bandwidth_gb": 131.512,"allowed_bandwidth_gb": "1000","netmask_v4": "255.255.255.248","gateway_v4": "123.123.123.1","power_status": "running","server_state": "ok","VPSPLANID": "28","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64","v6_network": "2001:DB8:1000::","v6_networks": [{"v6_network": "2001:DB8:1000::","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64"}],"label": "my new server","internal_ip": "10.99.0.10","kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV","auto_backups": "yes","tag": "mytag","OSID": "127","APPID": "0","FIREWALLGROUPID": "0"}}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Server.GetList(ctx)

	if err != nil {
		t.Errorf("Server.GetList returned %+v", err)
	}

	expected := []Server{
		{
			VpsID:            "576965",
			OS:               "CentOS 6 x64",
			RAM:              "4096 MB",
			Disk:             "Virtual 60 GB",
			MainIP:           "123.123.123.123",
			VPSCpus:          "2",
			Location:         "New Jersey",
			RegionID:         "1",
			DefaultPassword:  "nreqnusibni",
			Created:          "2013-12-19 14:45:41",
			PendingCharges:   "46.67",
			Status:           "active",
			Cost:             "10.05",
			CurrentBandwidth: 131.512,
			AllowedBandwidth: "1000",
			NetmaskV4:        "255.255.255.248",
			GatewayV4:        "123.123.123.1",
			PowerStatus:      "running",
			ServerState:      "ok",
			PlanID:           "28",
			V6Networks:       []V6Network{{Network: "2001:DB8:1000::", MainIP: "2001:DB8:1000::100", NetworkSize: "64"}},
			Label:            "my new server",
			InternalIP:       "10.99.0.10",
			KVMUrl:           "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
			AutoBackups:      "yes",
			Tag:              "mytag",
			OsID:             "127",
			AppID:            "0",
			FirewallGroupID:  "0",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.GetList returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_GetListByLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": {"SUBID": "576965","os": "CentOS 6 x64","ram": "4096 MB","disk": "Virtual 60 GB","main_ip": "123.123.123.123","vcpu_count": "2","location": "New Jersey","DCID": "1","default_password": "nreqnusibni","date_created": "2013-12-19 14:45:41","pending_charges": "46.67","status": "active","cost_per_month": "10.05","current_bandwidth_gb": 131.512,"allowed_bandwidth_gb": "1000","netmask_v4": "255.255.255.248","gateway_v4": "123.123.123.1","power_status": "running","server_state": "ok","VPSPLANID": "28","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64","v6_network": "2001:DB8:1000::","v6_networks": [{"v6_network": "2001:DB8:1000::","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64"}],"label": "my new server","internal_ip": "10.99.0.10","kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV","auto_backups": "yes","tag": "mytag","OSID": "127","APPID": "0","FIREWALLGROUPID": "0"}}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Server.GetListByLabel(ctx, "label")

	if err != nil {
		t.Errorf("Server.GetListByLabel returned %+v", err)
	}

	expected := []Server{
		{
			VpsID:            "576965",
			OS:               "CentOS 6 x64",
			RAM:              "4096 MB",
			Disk:             "Virtual 60 GB",
			MainIP:           "123.123.123.123",
			VPSCpus:          "2",
			Location:         "New Jersey",
			RegionID:         "1",
			DefaultPassword:  "nreqnusibni",
			Created:          "2013-12-19 14:45:41",
			PendingCharges:   "46.67",
			Status:           "active",
			Cost:             "10.05",
			CurrentBandwidth: 131.512,
			AllowedBandwidth: "1000",
			NetmaskV4:        "255.255.255.248",
			GatewayV4:        "123.123.123.1",
			PowerStatus:      "running",
			ServerState:      "ok",
			PlanID:           "28",
			V6Networks:       []V6Network{{Network: "2001:DB8:1000::", MainIP: "2001:DB8:1000::100", NetworkSize: "64"}},
			Label:            "my new server",
			InternalIP:       "10.99.0.10",
			KVMUrl:           "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
			AutoBackups:      "yes",
			Tag:              "mytag",
			OsID:             "127",
			AppID:            "0",
			FirewallGroupID:  "0",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.GetListByLabel returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_GetListByMainIP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": {"SUBID": "576965","os": "CentOS 6 x64","ram": "4096 MB","disk": "Virtual 60 GB","main_ip": "123.123.123.123","vcpu_count": "2","location": "New Jersey","DCID": "1","default_password": "nreqnusibni","date_created": "2013-12-19 14:45:41","pending_charges": "46.67","status": "active","cost_per_month": "10.05","current_bandwidth_gb": 131.512,"allowed_bandwidth_gb": "1000","netmask_v4": "255.255.255.248","gateway_v4": "123.123.123.1","power_status": "running","server_state": "ok","VPSPLANID": "28","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64","v6_network": "2001:DB8:1000::","v6_networks": [{"v6_network": "2001:DB8:1000::","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64"}],"label": "my new server","internal_ip": "10.99.0.10","kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV","auto_backups": "yes","tag": "mytag","OSID": "127","APPID": "0","FIREWALLGROUPID": "0"}}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Server.GetListByMainIP(ctx, "label")

	if err != nil {
		t.Errorf("Server.GetListByMainIP returned %+v", err)
	}

	expected := []Server{
		{
			VpsID:            "576965",
			OS:               "CentOS 6 x64",
			RAM:              "4096 MB",
			Disk:             "Virtual 60 GB",
			MainIP:           "123.123.123.123",
			VPSCpus:          "2",
			Location:         "New Jersey",
			RegionID:         "1",
			DefaultPassword:  "nreqnusibni",
			Created:          "2013-12-19 14:45:41",
			PendingCharges:   "46.67",
			Status:           "active",
			Cost:             "10.05",
			CurrentBandwidth: 131.512,
			AllowedBandwidth: "1000",
			NetmaskV4:        "255.255.255.248",
			GatewayV4:        "123.123.123.1",
			PowerStatus:      "running",
			ServerState:      "ok",
			PlanID:           "28",
			V6Networks:       []V6Network{{Network: "2001:DB8:1000::", MainIP: "2001:DB8:1000::100", NetworkSize: "64"}},
			Label:            "my new server",
			InternalIP:       "10.99.0.10",
			KVMUrl:           "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
			AutoBackups:      "yes",
			Tag:              "mytag",
			OsID:             "127",
			AppID:            "0",
			FirewallGroupID:  "0",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.GetListByMainIP returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_GetListByTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"576965": {"SUBID": "576965","os": "CentOS 6 x64","ram": "4096 MB","disk": "Virtual 60 GB","main_ip": "123.123.123.123","vcpu_count": "2","location": "New Jersey","DCID": "1","default_password": "nreqnusibni","date_created": "2013-12-19 14:45:41","pending_charges": "46.67","status": "active","cost_per_month": "10.05","current_bandwidth_gb": 131.512,"allowed_bandwidth_gb": "1000","netmask_v4": "255.255.255.248","gateway_v4": "123.123.123.1","power_status": "running","server_state": "ok","VPSPLANID": "28","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64","v6_network": "2001:DB8:1000::","v6_networks": [{"v6_network": "2001:DB8:1000::","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64"}],"label": "my new server","internal_ip": "10.99.0.10","kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV","auto_backups": "yes","tag": "mytag","OSID": "127","APPID": "0","FIREWALLGROUPID": "0"}}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Server.GetListByTag(ctx, "label")

	if err != nil {
		t.Errorf("Server.GetListByTag returned %+v", err)
	}

	expected := []Server{
		{
			VpsID:            "576965",
			OS:               "CentOS 6 x64",
			RAM:              "4096 MB",
			Disk:             "Virtual 60 GB",
			MainIP:           "123.123.123.123",
			VPSCpus:          "2",
			Location:         "New Jersey",
			RegionID:         "1",
			DefaultPassword:  "nreqnusibni",
			Created:          "2013-12-19 14:45:41",
			PendingCharges:   "46.67",
			Status:           "active",
			Cost:             "10.05",
			CurrentBandwidth: 131.512,
			AllowedBandwidth: "1000",
			NetmaskV4:        "255.255.255.248",
			GatewayV4:        "123.123.123.1",
			PowerStatus:      "running",
			ServerState:      "ok",
			PlanID:           "28",
			V6Networks:       []V6Network{{Network: "2001:DB8:1000::", MainIP: "2001:DB8:1000::100", NetworkSize: "64"}},
			Label:            "my new server",
			InternalIP:       "10.99.0.10",
			KVMUrl:           "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
			AutoBackups:      "yes",
			Tag:              "mytag",
			OsID:             "127",
			AppID:            "0",
			FirewallGroupID:  "0",
		},
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.GetListByTag returned %+v, expected %+v", server, expected)
	}
}

func TestServerServiceHandler_GetServer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SUBID": "576965","os": "CentOS 6 x64","ram": "4096 MB","disk": "Virtual 60 GB","main_ip": "123.123.123.123","vcpu_count": "2","location": "New Jersey","DCID": "1","default_password": "nreqnusibni","date_created": "2013-12-19 14:45:41","pending_charges": "46.67","status": "active","cost_per_month": "10.05","current_bandwidth_gb": 131.512,"allowed_bandwidth_gb": "1000","netmask_v4": "255.255.255.248","gateway_v4": "123.123.123.1","power_status": "running","server_state": "ok","VPSPLANID": "28","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64","v6_network": "2001:DB8:1000::","v6_networks": [{"v6_network": "2001:DB8:1000::","v6_main_ip": "2001:DB8:1000::100","v6_network_size": "64"}],"label": "my new server","internal_ip": "10.99.0.10","kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV","auto_backups": "yes","tag": "mytag","OSID": "127","APPID": "0","FIREWALLGROUPID": "0"}`
		fmt.Fprint(writer, response)
	})

	server, err := client.Server.GetServer(ctx, "1234")

	if err != nil {
		t.Errorf("Server.GetServer returned %+v", err)
	}

	expected := &Server{
		VpsID:            "576965",
		OS:               "CentOS 6 x64",
		RAM:              "4096 MB",
		Disk:             "Virtual 60 GB",
		MainIP:           "123.123.123.123",
		VPSCpus:          "2",
		Location:         "New Jersey",
		RegionID:         "1",
		DefaultPassword:  "nreqnusibni",
		Created:          "2013-12-19 14:45:41",
		PendingCharges:   "46.67",
		Status:           "active",
		Cost:             "10.05",
		CurrentBandwidth: 131.512,
		AllowedBandwidth: "1000",
		NetmaskV4:        "255.255.255.248",
		GatewayV4:        "123.123.123.1",
		PowerStatus:      "running",
		ServerState:      "ok",
		PlanID:           "28",
		V6Networks:       []V6Network{{Network: "2001:DB8:1000::", MainIP: "2001:DB8:1000::100", NetworkSize: "64"}},
		Label:            "my new server",
		InternalIP:       "10.99.0.10",
		KVMUrl:           "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV",
		AutoBackups:      "yes",
		Tag:              "mytag",
		OsID:             "127",
		AppID:            "0",
		FirewallGroupID:  "0",
	}

	if !reflect.DeepEqual(server, expected) {
		t.Errorf("Server.GetServer returned %+v, expected %+v", server, expected)
	}
}
