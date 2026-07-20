package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBlockStorageServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"block":{"id":"123456","cost":10,"pending_charges":2.5,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","attached_to_instance_ip":"","attached_to_instance_label":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "ewr-123abc", "block_type": "test", "os_id": 0, "snapshot_id": "", "bootable": false}}`
		fmt.Fprint(writer, response)
	})
	blockReq := &BlockStorageCreate{
		Region:    "ewr",
		SizeGB:    100,
		Label:     "mylabel",
		BlockType: "test",
	}
	blockStorage, _, err := client.BlockStorage.Create(ctx, blockReq)
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{
		ID:                      "123456",
		Cost:                    10,
		PendingCharges:          2.5,
		Status:                  "active",
		SizeGB:                  100,
		Region:                  "ewr",
		DateCreated:             "01-01-1960",
		AttachedToInstance:      "",
		AttachedToInstanceIP:    "",
		AttachedToInstanceLabel: "",
		Label:                   "mylabel",
		MountID:                 "ewr-123abc",
		BlockType:               "test",
		OSID:                    0,
		SnapshotID:              "",
		Bootable:                false,
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"block":{"id":"123456","cost":10,"pending_charges":2.5,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","attached_to_instance_ip":"","attached_to_instance_label":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "ewr-123abc", "block_type": "test", "os_id": 0, "snapshot_id": "", "bootable": false}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, _, err := client.BlockStorage.Get(ctx, "123456")
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{
		ID:                      "123456",
		Cost:                    10,
		PendingCharges:          2.5,
		Status:                  "active",
		SizeGB:                  100,
		Region:                  "ewr",
		DateCreated:             "01-01-1960",
		AttachedToInstance:      "",
		AttachedToInstanceIP:    "",
		AttachedToInstanceLabel: "",
		Label:                   "mylabel",
		MountID:                 "ewr-123abc",
		BlockType:               "test",
		OSID:                    0,
		SnapshotID:              "",
		Bootable:                false,
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Get returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	blockUpdate := &BlockStorageUpdate{
		Label: "unit-test-label-setter",
	}
	err := client.BlockStorage.Update(ctx, "123456", blockUpdate)
	if err != nil {
		t.Errorf("BlockStorage.Update returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Delete(ctx, "123456")
	if err != nil {
		t.Errorf("BlockStorage.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"blocks":[{"id":"123456","cost":10,"pending_charges":2.5,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","attached_to_instance_ip":"","attached_to_instance_label":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "ewr-123abc", "block_type": "test", "os_id": 0, "snapshot_id": "", "bootable": false}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, meta, _, err := client.BlockStorage.List(ctx, nil)
	if err != nil {
		t.Errorf("BlockStorage.List returned error: %v", err)
	}

	expected := []BlockStorage{
		{
			ID:                      "123456",
			Cost:                    10,
			PendingCharges:          2.5,
			Status:                  "active",
			SizeGB:                  100,
			Region:                  "ewr",
			DateCreated:             "01-01-1960",
			AttachedToInstance:      "",
			AttachedToInstanceIP:    "",
			AttachedToInstanceLabel: "",
			Label:                   "mylabel",
			MountID:                 "ewr-123abc",
			BlockType:               "test",
			OSID:                    0,
			SnapshotID:              "",
			Bootable:                false,
		},
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.List returned %+v, expected %+v", blockStorage, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("BlockStorage.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBlockStorageServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/12345/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	attach := &BlockStorageAttach{
		InstanceID: "1234",
		Live:       BoolToBoolPtr(true),
	}
	err := client.BlockStorage.Attach(ctx, "12345", attach)
	if err != nil {
		t.Errorf("BlockStorage.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})
	detach := &BlockStorageDetach{Live: BoolToBoolPtr(true)}
	err := client.BlockStorage.Detach(ctx, "123456", detach)
	if err != nil {
		t.Errorf("BlockStorage.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_ListSnapshots(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/snapshots", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
	"snapshots": [
		{
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"block_id": "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
			"description": "block storage snapshot",
			"state": "COMPLETE",
			"added_at": "2026-07-17 16:46:05",
			"updated_at": "2026-07-17 16:46:09",
			"size": 10737418240,
			"next_invoice_date": "",
			"next_invoice_price": ""
		},
		{
			"id": "96aef714-ce33-46a7-88e1-2eb1e8e0646e",
			"description": "another snapshot test",
			"block_id": "f004aea7-aec2-47fe-9ad7-7ebd830f225e",
			"state": "COMPLETE",
			"added_at": "2026-07-17 16:46:05",
			"updated_at": "2026-07-17 16:46:09",
			"next_invoice_date": "",
			"next_invoice_price": "",
			"size": 107374182400
		}
	],
		"meta": {
			"total": 2,
			"links": {
				"next": "",
				"prev": ""
		}
	}
}`
		fmt.Fprint(writer, response)
	})

	snapshots, meta, _, err := client.BlockStorage.ListSnapshots(ctx, nil)
	if err != nil {
		t.Errorf("BlockStorage.ListSnapshots returned error: %v", err)
	}

	expected := []BlockStorageSnapshot{
		{
			ID:               "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			BlockID:          "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
			Description:      "block storage snapshot",
			State:            "COMPLETE",
			DateCreated:      "2026-07-17 16:46:05",
			DateUpdated:      "2026-07-17 16:46:09",
			InvoiceNextDate:  "",
			InvoiceNextPrice: "",
			Size:             10737418240,
		},
		{
			ID:               "96aef714-ce33-46a7-88e1-2eb1e8e0646e",
			Description:      "another snapshot test",
			BlockID:          "f004aea7-aec2-47fe-9ad7-7ebd830f225e",
			State:            "COMPLETE",
			DateCreated:      "2026-07-17 16:46:05",
			DateUpdated:      "2026-07-17 16:46:09",
			InvoiceNextDate:  "",
			InvoiceNextPrice: "",
			Size:             107374182400,
		},
	}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("BlockStorage.ListSnapshots returned %+v, expected %+v", snapshots, expected)
	}

	expectedMeta := &Meta{
		Total: 2,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("BlockStorage.ListSnapshots meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBlockStorageServiceHandler_GetSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/snapshots/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
	"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
	"block_id": "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
	"description": "block storage snapshot",
	"state": "COMPLETE",
	"added_at": "2026-07-17 16:46:05",
	"updated_at": "2026-07-17 16:46:09",
	"size": 10737418240,
	"next_invoice_date": "",
	"next_invoice_price": ""
}
`
		fmt.Fprint(writer, response)
	})

	snapshot, _, err := client.BlockStorage.GetSnapshot(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("BlockStorage.GetSnapshot returned error: %v", err)
	}

	expected := &BlockStorageSnapshot{
		ID:               "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		BlockID:          "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
		Description:      "block storage snapshot",
		State:            "COMPLETE",
		DateCreated:      "2026-07-17 16:46:05",
		DateUpdated:      "2026-07-17 16:46:09",
		InvoiceNextDate:  "",
		InvoiceNextPrice: "",
		Size:             10737418240,
	}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("BlockStorage.GetSnapshot returned %+v, expected %+v", snapshot, expected)
	}
}

func TestBlockStorageServiceHandler_CreateSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/snapshots", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
	"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
	"block_id": "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
	"description": "block storage snapshot",
	"state": "COMPLETE",
	"added_at": "2026-07-17 16:46:05",
	"updated_at": "2026-07-17 16:46:09",
	"size": 10737418240,
	"next_invoice_date": "",
	"next_invoice_price": ""
}
`
		fmt.Fprint(writer, response)
	})
	snapshotReq := &BlockStorageSnapshotReq{
		BlockID:     "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
		Description: "block storage snapshot",
	}
	snapshot, _, err := client.BlockStorage.CreateSnapshot(ctx, snapshotReq)
	if err != nil {
		t.Errorf("BlockStorage.CreateSnapshot returned error: %v", err)
	}

	expected := &BlockStorageSnapshot{
		ID:               "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		BlockID:          "ay2c5b1a-66fd-4dfb-b839-443f2e6c0b60",
		Description:      "block storage snapshot",
		State:            "COMPLETE",
		DateCreated:      "2026-07-17 16:46:05",
		DateUpdated:      "2026-07-17 16:46:09",
		InvoiceNextDate:  "",
		InvoiceNextPrice: "",
		Size:             10737418240,
	}

	if !reflect.DeepEqual(snapshot, expected) {
		t.Errorf("BlockStorage.CreateSnapshot returned %+v, expected %+v", snapshot, expected)
	}
}

func TestBlockStorageServiceHandler_UpdateSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/snapshots/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	snapshotReq := &BlockStorageSnapshotReq{
		Description: "block storage snapshot updated",
	}

	if err := client.BlockStorage.UpdateSnapshot(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60", snapshotReq); err != nil {
		t.Errorf("BlockStorage.UpdateSnapshot returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_DeleteSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/snapshots/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.BlockStorage.DeleteSnapshot(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60"); err != nil {
		t.Errorf("BlockStorage.DeleteSnapshot returned %+v, expected %+v", err, nil)
	}
}
