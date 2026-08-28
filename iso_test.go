package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestISOServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"iso": {
		"id": "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
		"date_created": "2020-10-10T01:56:20+00:00",
		"filename": "CentOS-8.1.1911-x86_64-dvd1.iso",
		"status": "pending"
	}
}`))

	isoReq := &ISOReq{URL: "http://centos.com/CentOS-8.1.1911-x86_64-dvd1.iso"}
	iso, _, err := client.ISO.Create(ctx, isoReq)
	if err != nil {
		t.Errorf("ISO.Create returned %+v, expected %+v", err, nil)
	}

	expected := &ISO{
		ID:          "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
		DateCreated: "2020-10-10T01:56:20+00:00",
		FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
		Size:        0,
		MD5Sum:      "",
		SHA512Sum:   "",
		Status:      "pending",
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("ISO.Create returned %+v, expected %+v", iso, expected)
	}
}

func TestISOServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso/7368d5f6-1281-438a-9ec6-17d7c7930d1a", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"iso": {
		"id": "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
		"date_created": "2020-10-10T01:56:20+00:00",
		"filename": "CentOS-8.1.1911-x86_64-dvd1.iso",
		"status": "complete",
		"size": 120582323,
		"md5sum": "77ba289bdc966ec996278a5a740d96d8",
		"sha512sum": "2b31b6fcab34d6ea9a6b293601c39b90cb044e5679fcc5"
	}
}`))

	iso, _, err := client.ISO.Get(ctx, "7368d5f6-1281-438a-9ec6-17d7c7930d1a")
	if err != nil {
		t.Errorf("ISO.Get returned %+v, expected %+v", err, nil)
	}

	expected := &ISO{
		ID:          "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
		DateCreated: "2020-10-10T01:56:20+00:00",
		FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
		Size:        120582323,
		MD5Sum:      "77ba289bdc966ec996278a5a740d96d8",
		SHA512Sum:   "2b31b6fcab34d6ea9a6b293601c39b90cb044e5679fcc5",
		Status:      "complete",
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("ISO.Get returned %+v, expected %+v", iso, expected)
	}
}

func TestISOServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso/5c2e9c16-0cde-4f70-858a-59f20ef03118", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	err := client.ISO.Delete(ctx, "5c2e9c16-0cde-4f70-858a-59f20ef03118")

	if err != nil {
		t.Errorf("ISO.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestISOServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"isos": [
		{
			"id": "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
			"date_created": "2020-10-10T01:56:20+00:00",
			"filename": "CentOS-8.1.1911-x86_64-dvd1.iso",
			"status": "complete",
			"size": 120582323,
			"md5sum": "77ba289bdc966ec996278a5a740d96d8",
			"sha512sum": "2b31b6fcab34d6ea9a6b293601c39b90cb044e5679fcc5"
		}
	],
	"meta": {
		"total": 8,
		"links": {
			"next": "adflzxcvljadflzcv",
			"prev": ""
		}
	}
}`))

	iso, meta, _, err := client.ISO.List(ctx, nil)
	if err != nil {
		t.Errorf("ISO.List returned %+v, expected %+v", err, nil)
	}

	expectedISO := []ISO{
		{
			ID:          "7368d5f6-1281-438a-9ec6-17d7c7930d1a",
			DateCreated: "2020-10-10T01:56:20+00:00",
			FileName:    "CentOS-8.1.1911-x86_64-dvd1.iso",
			Size:        120582323,
			MD5Sum:      "77ba289bdc966ec996278a5a740d96d8",
			SHA512Sum:   "2b31b6fcab34d6ea9a6b293601c39b90cb044e5679fcc5",
			Status:      "complete",
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{
			Next: "adflzxcvljadflzcv",
			Prev: "",
		},
	}
	if !reflect.DeepEqual(iso, expectedISO) {
		t.Errorf("ISO.List iso returned %+v, expected %+v", iso, expectedISO)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ISO.List returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestISOServiceHandler_ListPublic(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/iso-public", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"public_isos": [
		{
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b604",
			"name": "CentOS 7",
			"description": "7 x86_64 Minimal",
			"md5sum": "7f4df50f42ee1b52b193e79855a3aa19"
		}
	],
	"meta": {
		"total":8,
		"links": {
			"next":"asdfcxvasdfz",
			"prev":""
		}
	}
}`))

	iso, meta, _, err := client.ISO.ListPublic(ctx, nil)
	if err != nil {
		t.Errorf("ISO.ListPublic returned %+v, expected %+v", err, nil)
	}

	expectedISO := []PublicISO{
		{
			ID:          "cb676a46-66fd-4dfb-b839-443f2e6c0b604",
			Name:        "CentOS 7",
			Description: "7 x86_64 Minimal",
			MD5Sum:      "7f4df50f42ee1b52b193e79855a3aa19",
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{
			Next: "asdfcxvasdfz",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(iso, expectedISO) {
		t.Errorf("ISO.ListPublic  iso returned %+v, expected %+v", iso, expectedISO)
	}
	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ISO.ListPublic meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
