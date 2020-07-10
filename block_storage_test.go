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
		response := `{"block":{"id":123456,"cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":0,"date_created":"01-01-1960","label":"mylabel"}}`
		fmt.Fprint(writer, response)
	})
	blockReq := &BlockStorageReq{
		Region: "ewr",
		SizeGB: 100,
		Label:  "mylabel",
	}
	blockStorage, err := client.BlockStorage.Create(ctx, blockReq)
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{
		ID:                 123456,
		Cost:               10,
		Status:             "active",
		SizeGB:             100,
		Region:             "ewr",
		DateCreated:        "01-01-1960",
		AttachedToInstance: 0,
		Label:              "mylabel",
	}

	if !reflect.DeepEqual(blockStorage, expected) {
		t.Errorf("BlockStorage.Create returned %+v, expected %+v", blockStorage, expected)
	}
}

func TestBlockStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"block":{"id":123456,"cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":0,"date_created":"01-01-1960","label":"mylabel"}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, err := client.BlockStorage.Get(ctx, 123456)
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := &BlockStorage{
		ID:                 123456,
		Cost:               10,
		Status:             "active",
		SizeGB:             100,
		Region:             "ewr",
		DateCreated:        "01-01-1960",
		AttachedToInstance: 0,
		Label:              "mylabel",
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

	err := client.BlockStorage.Update(ctx, 123456, "unit-test-label-setter")
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

	err := client.BlockStorage.Delete(ctx, 123456)
	if err != nil {
		t.Errorf("BlockStorage.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"blocks":[{"id":123456,"cost":10,"status":"active","size_gb":100,"region":"ewr","attached_to_instance":0,"date_created":"01-01-1960","label":"mylabel"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	blockStorage, meta, err := client.BlockStorage.List(ctx, nil)
	if err != nil {
		t.Errorf("BlockStorage.Create returned error: %v", err)
	}

	expected := []BlockStorage{
		{
			ID:                 123456,
			Cost:               10,
			Status:             "active",
			SizeGB:             100,
			Region:             "ewr",
			DateCreated:        "01-01-1960",
			AttachedToInstance: 0,
			Label:              "mylabel",
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

	err := client.BlockStorage.Attach(ctx, 12345, 1234, "yes")
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

	err := client.BlockStorage.Detach(ctx, 123456, "yes")
	if err != nil {
		t.Errorf("BlockStorage.Detach returned %+v, expected %+v", err, nil)
	}
}

func TestBlockStorageServiceHandler_Resize(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blocks/123456/resize", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BlockStorage.Resize(ctx, 123456, 50)
	if err != nil {
		t.Errorf("BlockStorage.Resize returned %+v, expected %+v", err, nil)
	}
}
