package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBlockStorageServiceHandler_Attach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Attach(ctx, "123456", "876521")

	if err != nil {
		t.Errorf("BlockStorage.Attach returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SUBID": "1234566"}`
		fmt.Fprint(writer, response)
	})

	blockStorage, err := client.BlockStorage.Create(ctx, 1, 10, "unit-test")

	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{BlockStorageID: "1234566", RegionID: 1, Size: 10, Label: "unit-test"}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Delete(ctx, "123456")

	if err != nil {
		t.Errorf("BlockStorage.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_Detach(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Detach(ctx, "123456")

	if err != nil {
		t.Errorf("BlockStorage.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.SetLabel(ctx, "123456", "unit-test-label-setter")

	if err != nil {
		t.Errorf("BlockStorage.SetLabel returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `[
	 	{
        	"SUBID": 1313216,
        	"date_created": "2016-03-29 10:10:04",
        	"cost_per_month": 10,
        	"status": "pending",
        	"size_gb": 100,
        	"DCID": 1,
        	"attached_to_SUBID": null,
        	"label": "files1"
    	},
		{
        	"SUBID": 1313216,
        	"date_created": "2016-03-29 10:10:04",
        	"cost_per_month": 10,
        	"status": "pending",
        	"size_gb": 100,
        	"DCID": 1,
        	"attached_to_SUBID": null,
        	"label": "files1"
    	}
		]
		`
		fmt.Fprint(writer, response)
	})

	blockStorage, err := client.BlockStorage.GetList(ctx)

	if err != nil {
		t.Errorf("BlockStorage.Get returned error: %v", err)
	}

	expected := []BlockStorageGet{
		{BlockStorageID: 1313216, DateCreated: "2016-03-29 10:10:04", Cost: 10, Status: "pending", Size: 100, RegionID: 1, Label: "files1"},
		{BlockStorageID: 1313216, DateCreated: "2016-03-29 10:10:04", Cost: 10, Status: "pending", Size: 100, RegionID: 1, Label: "files1"},
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Get returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
	 	{
        	"SUBID": 1313216,
        	"date_created": "2016-03-29 10:10:04",
        	"cost_per_month": 10,
        	"status": "pending",
        	"size_gb": 100,
        	"DCID": 1,
        	"attached_to_SUBID": null,
        	"label": "files1"
    	}
		`
		fmt.Fprint(writer, response)
	})

	blockStorage, err := client.BlockStorage.Get(ctx, "1313216")

	if err != nil {
		t.Errorf("BlockStorage.Get returned error: %v", err)
	}

	expected := &BlockStorageGet{BlockStorageID: 1313216, DateCreated: "2016-03-29 10:10:04", Cost: 10, Status: "pending", Size: 100, RegionID: 1, Label: "files1"}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Get returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Resize(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/block/resize", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Resize(ctx, "123456", 50)

	if err != nil {
		t.Errorf("BlockStorage.Resize returned %+v, expected %+v", err, nil)
	}
}
