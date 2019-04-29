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

	expected := &ServerAppInfo{AppInfo: "test"}

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
		Enabled: true,
		CronType: "weekly",
		NextRun: "2016-05-07 08:00:00",
		Hour: 8,
		Dow: 6,
		Dom: 0,
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
		Hour: 23,
		Dow: 2,
		Dom: 3,
	}

	err := client.Server.SetBackupSchedule(ctx, "1234", bs)

	if err != nil {
		t.Errorf("Server.SetBackupSchedule returned %+v, ", err)
	}
}