package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBackupServiceHandler_List(t *testing.T) {
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

	backups, err := client.Backup.List(ctx)
	if err != nil {
		t.Errorf("Backup.List returned error: %v", err)
	}

	expected := []Backup{
		{
			BackupID:    "543d34149403a",
			DateCreated: "2014-10-14 12:40:40",
			Description: "Automatic server backup",
			Size:        "42949672960",
			Status:      "complete",
		},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.List returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_ListEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.List(ctx)
	if err != nil {
		t.Errorf("Backup.List returned error: %v", err)
	}

	var expected []Backup

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.List returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_ListBySub(t *testing.T) {
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

	backups, err := client.Backup.ListBySub(ctx, "test-backupID")
	if err != nil {
		t.Errorf("Backup.ListBySub returned error: %v", err)
	}

	expected := []Backup{
		{
			BackupID:    "543d34149403a",
			DateCreated: "2014-10-14 12:40:40",
			Description: "Automatic server backup",
			Size:        "42949672960",
			Status:      "complete",
		},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.ListBySub returned %+v, expected %+v", backups, expected)
	}
}

func TestBackupServiceHandler_ListBySubEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/backup/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	backups, err := client.Backup.ListBySub(ctx, "test-backupID")
	if err != nil {
		t.Errorf("Backup.ListBySub returned error: %v", err)
	}

	var expected []Backup

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.ListBySub returned %+v, expected %+v", backups, expected)
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
