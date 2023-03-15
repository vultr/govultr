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

	mux.HandleFunc("/v2/backups", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"backups": [
				{
				"id": "543d34149403a",
				"date_created": "2014-10-14 12:40:40",
				"description": "Automatic server backup",
				"size": 42949672960,
				"status": "complete"
				}
			],
			"meta": {
				"total":8,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backups, meta, _, err := client.Backup.List(ctx, nil)
	if err != nil {
		t.Errorf("Backup.List returned error: %v", err)
	}

	expected := []Backup{
		{
			ID:          "543d34149403a",
			DateCreated: "2014-10-14 12:40:40",
			Description: "Automatic server backup",
			Size:        42949672960,
			Status:      "complete",
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.List returned %+v, expected %+v", backups, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Backup.List returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBackupServiceHandler_ListEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/backups", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"backups": [],
			"meta": {
				"total":0,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backups, meta, _, err := client.Backup.List(ctx, nil)
	if err != nil {
		t.Errorf("Backup.List returned error: %v", err)
	}

	expected := []Backup{}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("Backup.List returned %+v, expected %+v", backups, expected)
	}

	expectedMeta := &Meta{
		Total: 0,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Backup.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBackupServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/backups/543d34149403a", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"backup": {
				"id": "543d34149403a",
				"date_created": "2014-10-14 12:40:40",
				"description": "Automatic server backup",
				"size": 42949672960,
				"status": "complete"
			}
		}
		`

		fmt.Fprint(w, response)
	})

	backup, _, err := client.Backup.Get(ctx, "543d34149403a")
	if err != nil {
		t.Errorf("Backup.Get returned error: %v", err)
	}

	expected := &Backup{
		ID:          "543d34149403a",
		DateCreated: "2014-10-14 12:40:40",
		Description: "Automatic server backup",
		Size:        42949672960,
		Status:      "complete",
	}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Backup.Get returned %+v, expected %+v", backup, expected)
	}
}

func TestBackupServiceHandler_GetEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/backups/543d34149403a", func(w http.ResponseWriter, r *http.Request) {
		response := `			
			{
				"backup": {}
			}
		`

		fmt.Fprint(w, response)
	})

	backup, _, err := client.Backup.Get(ctx, "543d34149403a")
	if err != nil {
		t.Errorf("Backup.Get returned error: %v", err)
	}

	expected := &Backup{}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Backup.Get returned %+v, expected %+v", backup, expected)
	}
}
