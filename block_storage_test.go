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
		response := `{"block":{"id":"123456","cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "ewr-123abc", "block_type": "test"}}`
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
		ID:                 "123456",
		Cost:               10,
		Status:             "active",
		SizeGB:             100,
		Region:             "ewr",
		DateCreated:        "01-01-1960",
		AttachedToInstance: "",
		Label:              "mylabel",
		MountID:            "ewr-123abc",
		BlockType:          "test",
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"block":{"id":"123456","cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "123abc", "block_type": "test"}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, _, err := client.BlockStorage.Get(ctx, "123456")
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{
		ID:                 "123456",
		Cost:               10,
		Status:             "active",
		SizeGB:             100,
		Region:             "ewr",
		DateCreated:        "01-01-1960",
		AttachedToInstance: "",
		Label:              "mylabel",
		MountID:            "123abc",
		BlockType:          "test",
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
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
		t.Errorf("BlockStorage.SetLabel returned %+v, expected %+v", err, nil)
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
		response := `{"blocks":[{"id":"123456","cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":"","date_created":"01-01-1960","label":"mylabel", "mount_id": "123abc", "block_type": "test"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, meta, _, err := client.BlockStorage.List(ctx, nil)
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := []BlockStorage{
		{
			ID:                 "123456",
			Cost:               10,
			Status:             "active",
			SizeGB:             100,
			Region:             "ewr",
			DateCreated:        "01-01-1960",
			AttachedToInstance: "",
			Label:              "mylabel",
			MountID:            "123abc",
			BlockType:          "test",
		},
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("User.List meta returned %+v, expected %+v", meta, expectedMeta)
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
