package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBackupServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"543d34149403a": {
				"BACKUPID": "543d34149403a",
				"date_created": "2014-10-14 12:40:40",
				"description": "Automatic server backup",
				"size": "42949672960",
				"status": "complete"
			},
			"543d340f6dbce": {
				"BACKUPID": "543d340f6dbce",
				"date_created": "2014-10-13 16:11:46",
				"description": "",
				"size": "10000000",
				"status": "complete"
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.GetList(ctx)
	if err != nil {
		t.Errorf("Backup.GetList returned error: %v", err)
	}

	expected := []Backup{
		{
			BackupID:    "543d34149403a",
			DateCreated: "2014-10-14 12:40:40",
			Description: "Automatic server backup",
			Size:        "42949672960",
			Status:      "complete",
		},
		{
			BackupID:    "543d340f6dbce",
			DateCreated: "2014-10-13 16:11:46",
			Description: "",
			Size:        "10000000",
			Status:      "complete",
		},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.GetList returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_GetListEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.GetList(ctx)
	if err != nil {
		t.Errorf("Backup.GetList returned error: %v", err)
	}

	var expected []Backup

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.GetList returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_GetListBySub(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"543d34149403a": {
				"BACKUPID": "543d34149403a",
				"date_created": "2014-10-14 12:40:40",
				"description": "Automatic server backup",
				"size": "42949672960",
				"status": "complete"
			},
			"543d340f6dbce": {
				"BACKUPID": "543d340f6dbce",
				"date_created": "2014-10-13 16:11:46",
				"description": "",
				"size": "10000000",
				"status": "complete"
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.GetListBySub(ctx, "test-backupID")
	if err != nil {
		t.Errorf("Backup.GetListBySub returned error: %v", err)
	}

	expected := []Backup{
		{
			BackupID:    "543d34149403a",
			DateCreated: "2014-10-14 12:40:40",
			Description: "Automatic server backup",
			Size:        "42949672960",
			Status:      "complete",
		},
		{
			BackupID:    "543d340f6dbce",
			DateCreated: "2014-10-13 16:11:46",
			Description: "",
			Size:        "10000000",
			Status:      "complete",
		},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.GetListBySub returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_GetListBySubEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.GetListBySub(ctx, "test-backupID")
	if err != nil {
		t.Errorf("Backup.GetList returned error: %v", err)
	}

	var expected []Backup

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.GetList returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"543d34149403a": {
				"BACKUPID": "543d34149403a",
				"date_created": "2014-10-14 12:40:40",
				"description": "Automatic server backup",
				"size": "42949672960",
				"status": "complete"
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backup, err := client.Backup.Get(ctx, "543d34149403a")
	if err != nil {
		t.Errorf("Backup.Get returned error: %v", err)
	}

	expected := &Backup{
		BackupID:    "543d34149403a",
		DateCreated: "2014-10-14 12:40:40",
		Description: "Automatic server backup",
		Size:        "42949672960",
		Status:      "complete",
	}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Backup.Get returned %+v, expected %+v", backup, expected)
	}
}

func TestBackupServiceHandler_GetEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	backup, err := client.Backup.Get(ctx, "test-backupID")
	if err != nil {
		t.Errorf("Backup.Get returned error: %v", err)
	}

	expected := &Backup{}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Backup.Get returned %+v, expected %+v", backup, expected)
	}
}
