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

	mux.HandleFunc("/v2/snapshots", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"snapshot":{"id": "5359435d28b9a","date_created": "2014-04-18 12:40:40","description": "Test snapshot","size": "42949672960","status": "complete","os_id": 127,"app_id": 0}}`
		fmt.Fprint(writer, response)
	})

	snap := &SnapshotReq{
		InstanceID:  12345,
		Description: "Test snapshot",
	}

	snapshot, err := client.Snapshot.Create(ctx, snap)
	if err != nil {
		t.Errorf("Snapshot.Create returned error: %v", err)
	}

	expected := &Snapshot{
		ID:          "5359435d28b9a",
		DateCreated: "2014-04-18 12:40:40",
		Description: "Test snapshot",
		Size:        "42949672960",
		Status:      "complete",
		OsID:        127,
		AppID:       0,
	}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("Snapshot.Create returned %+v, expected %+v", snapshot, expected)
	}
}

func TestSnapshotServiceHandler_CreateFromURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/snapshots/create-from-url", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"snapshot":{"id": "5359435d28b9a","date_created": "2014-04-18 12:40:40","description": "Test snapshot","size": "42949672960","status": "complete","os_id": 127,"app_id": 0}}`
		fmt.Fprint(writer, response)
	})
	snap := SnapshotURLReq{URL: "http://vultr.com"}
	snapshot, err := client.Snapshot.CreateFromURL(ctx, snap)
	if err != nil {
		t.Errorf("Snapshot.CreateFromURL returned error: %v", err)
	}

	expected := &Snapshot{
		ID:          "5359435d28b9a",
		DateCreated: "2014-04-18 12:40:40",
		Description: "Test snapshot",
		Size:        "42949672960",
		Status:      "complete",
		OsID:        127,
		AppID:       0,
	}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("Snapshot.CreateFromURL returned %+v, expected %+v", snapshot, expected)
	}
}

func TestSnapshotServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/snapshots/5359435d28b9a", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"snapshot":{"id": "5359435d28b9a","date_created": "2014-04-18 12:40:40","description": "Test snapshot","size": "42949672960","status": "complete","os_id": 127,"app_id": 0}}`
		fmt.Fprint(writer, response)
	})

	snapshot, err := client.Snapshot.Get(ctx, "5359435d28b9a")
	if err != nil {
		t.Errorf("Snapshot.Get returned error: %v", err)
	}

	expected := &Snapshot{
		ID:          "5359435d28b9a",
		DateCreated: "2014-04-18 12:40:40",
		Description: "Test snapshot",
		Size:        "42949672960",
		Status:      "complete",
		OsID:        127,
		AppID:       0,
	}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("Snapshot.Get returned %+v, expected %+v", snapshot, expected)
	}
}

func TestSnapshotServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/snapshots/7a05cbf361d98", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Snapshot.Delete(ctx, "7a05cbf361d98")

	if err != nil {
		t.Errorf("Snapshot.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestSnapshotServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/snapshots", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"snapshots": [{"id": "885ee0f4f263c","date_created": "2014-04-18 12:40:40","description": "Test snapshot","size": "42949672960","status": "complete","os_id": 127,"app_id": 0}],"meta": {"total": 4,"links": {"next": "","prev": ""}}}`
		fmt.Fprint(writer, response)
	})

	snapshots, meta, err := client.Snapshot.List(ctx, nil)
	if err != nil {
		t.Errorf("Snapshot.List returned error: %v", err)
	}

	expectedSnap := []Snapshot{
		{
			ID:          "885ee0f4f263c",
			DateCreated: "2014-04-18 12:40:40",
			Description: "Test snapshot",
			Size:        "42949672960",
			Status:      "complete",
			OsID:        127,
			AppID:       0,
		},
	}

	expectedMeta := &Meta{
		Total: 4,
		Links: &Links{},
	}

	if !reflect.DeepEqual(snapshots, expectedSnap) {
		t.Errorf("Snapshot.list snapshots returned %+v, expected %+v", snapshots, expectedSnap)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Snapshot.list meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
