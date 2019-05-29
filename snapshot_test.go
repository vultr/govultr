package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSnapshotServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/snapshot/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SNAPSHOTID": "1234567"}`

		fmt.Fprint(writer, response)
	})

	snapshot, err := client.Snapshot.Create(ctx, "987654321", "unit-test-desc")

	if err != nil {
		t.Errorf("Snapshot.Create returned error: %v", err)
	}

	expected := &Snapshot{SnapshotID: "1234567", Description: "unit-test-desc"}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("Snapshot.Create returned %+v, expected %+v", snapshot, expected)
	}
}

func TestSnapshotServiceHandler_CreateFromURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/snapshot/create_from_url", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SNAPSHOTID": "544e52f31c706"}`

		fmt.Fprint(writer, response)
	})

	snapshot, err := client.Snapshot.CreateFromURL(ctx, "http://localhost/some.iso")

	if err != nil {
		t.Errorf("Snapshot.CreateFromURL returned error: %v", err)
	}

	expected := &Snapshot{SnapshotID: "544e52f31c706"}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("Snapshot.CreateFromURL returned %+v, expected %+v", snapshot, expected)
	}
}

func TestSnapshotServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/snapshot/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Snapshot.Delete(ctx, "7a05cbf361d98")

	if err != nil {
		t.Errorf("Snapshot.Delete returned %+v, expected %+v", err, nil)
	}

}

func TestSnapshotServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/snapshot/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
		"5359435dc1df3": {
		"SNAPSHOTID": "5359435dc1df3",
		"date_created": "2014-04-22 16:11:46",
		"description": "",
		"size": "10000000",
		"status": "complete",
		"OSID": "127",
		"APPID": "0"
		}
		}
		`
		fmt.Fprint(writer, response)
	})

	snapshots, err := client.Snapshot.GetList(ctx)

	if err != nil {
		t.Errorf("Snapshot.GetList returned error: %v", err)
	}
	expected := []Snapshot{
		{SnapshotID: "5359435dc1df3", DateCreated: "2014-04-22 16:11:46", Description: "", Size: "10000000", Status: "complete", OsID: "127", AppID: "0"},
	}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("Snapshot.GetList returned %+v, expected %+v", snapshots, expected)

	}
}

func TestSnapshotServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/snapshot/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
			"5359435dc1df3": {
			"SNAPSHOTID": "5359435dc1df3",
			"date_created": "2014-04-22 16:11:46",
			"description": "",
			"size": "10000000",
			"status": "complete",
			"OSID": "127",
			"APPID": "0"
			}
			}
			`
		fmt.Fprint(writer, response)
	})

	snapshots, err := client.Snapshot.Get(ctx, "5359435dc1df3")

	if err != nil {
		t.Errorf("Snapshot.Get returned error: %v", err)
	}
	expected := &Snapshot{SnapshotID: "5359435dc1df3", DateCreated: "2014-04-22 16:11:46", Description: "", Size: "10000000", Status: "complete", OsID: "127", AppID: "0"}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("Snapshot.Get returned %+v, expected %+v", snapshots, expected)

	}
}
