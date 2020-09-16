package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestIsoServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"iso":{"id":"9931","date_created":"2020-07-0917:15:27","filename":"CentOS-8.1.1911-x86_64-dvd1.iso","status":"pending"}}`
		fmt.Fprint(writer, response)
	})

	isoReq := &ISOReq{Url: "http://centos.com/CentOS-8.1.1911-x86_64-dvd1.iso"}
	iso, err := client.ISO.Create(ctx, isoReq)
	if err != nil {
		t.Errorf("Iso.Create returned %+v, expected %+v", err, nil)
	}

	expected := &ISO{
		ID:          "9931",
		DateCreated: "2020-07-0917:15:27",
		FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
		Size:        0,
		MD5Sum:      "",
		SHA512Sum:   "",
		Status:      "pending",
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Iso.Create returned %+v, expected %+v", iso, expected)
	}
}

func TestIsoServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso/9931", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"iso":{"id":"9931","date_created":"2020-07-0917:15:27","filename":"CentOS-8.1.1911-x86_64-dvd1.iso","status":"pending"}}`
		fmt.Fprint(writer, response)
	})

	iso, err := client.ISO.Get(ctx, "9931")
	if err != nil {
		t.Errorf("Iso.Get returned %+v, expected %+v", err, nil)
	}

	expected := &ISO{
		ID:          "9931",
		DateCreated: "2020-07-0917:15:27",
		FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
		Size:        0,
		MD5Sum:      "",
		SHA512Sum:   "",
		Status:      "pending",
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Iso.Get returned %+v, expected %+v", iso, expected)
	}
}

func TestIsoServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso/24", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ISO.Delete(ctx, "24")

	if err != nil {
		t.Errorf("Iso.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestIsoServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"isos":[{"id":"9931","date_created":"2020-07-0917:15:27","filename":"CentOS-8.1.1911-x86_64-dvd1.iso","status":"pending"}],"meta":{"total":8,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	iso, meta, err := client.ISO.List(ctx, nil)
	if err != nil {
		t.Errorf("Iso.List returned %+v, expected %+v", err, nil)
	}

	expectedIso := []ISO{
		{
			ID:          "9931",
			DateCreated: "2020-07-0917:15:27",
			FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
			Size:        0,
			MD5Sum:      "",
			SHA512Sum:   "",
			Status:      "pending",
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}
	if !reflect.DeepEqual(iso, expectedIso) {
		t.Errorf("Iso.List iso returned %+v, expected %+v", iso, expectedIso)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Iso.List returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestIsoServiceHandler_ListPublic(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso-public", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"public_isos": [{"id": "204515","name": "CentOS 7","description": "7 x86_64 Minimal"}],"meta":{"total":8,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	iso, meta, err := client.ISO.ListPublic(ctx, nil)
	if err != nil {
		t.Errorf("Iso.ListPublic returned %+v, expected %+v", err, nil)
	}

	expectedIso := []PublicISO{
		{ID: "204515", Name: "CentOS 7", Description: "7 x86_64 Minimal"},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}
	if !reflect.DeepEqual(iso, expectedIso) {
		t.Errorf("Iso.ListPublic  iso returned %+v, expected %+v", iso, expectedIso)
	}
	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Iso.ListPublic meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
