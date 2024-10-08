package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// Pull Zones =================================================================

func TestCDNPullZoneServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(cdnPullPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"pull_zone": {
		"id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		"date_created": "2024-01-25 09:41:05",
		"status": "active",
		"label": "my-pullzone",
		"origin_scheme": "https",
		"origin_domain": "constant.com",
		"cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
		"vanity_domain": "my.domain.com",
		"cache_size": 50000000,
		"requests": 0,
		"in_bytes": 0,
		"out_bytes": 0,
		"packets_per_sec": 50,
		"last_purge": "2024-01-25 11:39:04",
		"cors": false,
		"gzip": true,
		"block_ai": false,
		"block_bad_bots": true,
		"regions": [
		  "ord"
		]
	}
}`
		fmt.Fprint(writer, response)
	})

	pzReq := &CDNZoneReq{
		Label:        "my-pullzone",
		OriginScheme: "https",
		OriginDomain: "constant.com",
		VanityDomain: "my.domain.com",
		CORS:         BoolToBoolPtr(false),
		GZIP:         BoolToBoolPtr(true),
		BlockAI:      BoolToBoolPtr(false),
		BlockBadBots: BoolToBoolPtr(true),
	}

	pz, _, err := client.CDN.CreatePullZone(ctx, pzReq)
	if err != nil {
		t.Errorf("CDN.CreatePullZone returned error: %v", err)
	}

	expected := &CDNZone{
		ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		Label:         "my-pullzone",
		DateCreated:   "2024-01-25 09:41:05",
		Status:        "active",
		OriginScheme:  "https",
		OriginDomain:  "constant.com",
		CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
		VanityDomain:  "my.domain.com",
		CacheSize:     50000000,
		Requests:      0,
		BytesIn:       0,
		BytesOut:      0,
		PacketsPerSec: 50,
		DatePurged:    "2024-01-25 11:39:04",
		CORS:          false,
		GZIP:          true,
		BlockAI:       false,
		BlockBadBots:  true,
		Regions:       []string{"ord"},
	}

	if !reflect.DeepEqual(pz, expected) {
		t.Errorf("CDN.CreatePullZone returned %+v, expected %+v", pz, expected)
	}
}

func TestCDNPullZoneServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(cdnPullPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"pull_zones": [
		{
			"id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
			"date_created": "2024-01-25 09:41:05",
			"status": "active",
			"label": "my-pullzone",
			"origin_scheme": "https",
			"origin_domain": "constant.com",
			"cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
			"vanity_domain": "my.domain.com",
			"cache_size": 50000000,
			"requests": 0,
			"in_bytes": 0,
			"out_bytes": 0,
			"packets_per_sec": 50,
			"last_purge": "2024-01-25 11:39:04",
			"cors": false,
			"gzip": true,
			"block_ai": false,
			"block_bad_bots": true,
			"regions": [
			  "ord"
			]
		}
	],
	"meta": {
		"total": 1,
		"links": {
			"next": "",
			"prev": ""
		}
	}
}`
		fmt.Fprint(writer, response)
	})

	pzs, meta, _, err := client.CDN.ListPullZones(ctx)
	if err != nil {
		t.Errorf("CDN.ListPullZone returned error: %v", err)
	}

	expected := []CDNZone{
		{
			ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
			Label:         "my-pullzone",
			DateCreated:   "2024-01-25 09:41:05",
			Status:        "active",
			OriginScheme:  "https",
			OriginDomain:  "constant.com",
			CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
			VanityDomain:  "my.domain.com",
			CacheSize:     50000000,
			Requests:      0,
			BytesIn:       0,
			BytesOut:      0,
			PacketsPerSec: 50,
			DatePurged:    "2024-01-25 11:39:04",
			CORS:          false,
			GZIP:          true,
			BlockAI:       false,
			BlockBadBots:  true,
			Regions:       []string{"ord"},
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(pzs, expected) {
		t.Errorf("CDN.ListPullZone returned %+v, expected %+v", pzs, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("CDN.ListPullZone meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestCDNPullZoneServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	path := fmt.Sprintf("%s/%s", cdnPullPath, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"pull_zone": {
		"id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		"date_created": "2024-01-25 09:41:05",
		"status": "active",
		"label": "new-label",
		"origin_scheme": "https",
		"origin_domain": "constant.com",
		"cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
		"vanity_domain": "my.domain.com",
		"cache_size": 50000000,
		"requests": 0,
		"in_bytes": 0,
		"out_bytes": 0,
		"packets_per_sec": 50,
		"last_purge": "2024-01-25 11:39:04",
		"cors": true,
		"gzip": true,
		"block_ai": true,
		"block_bad_bots": true,
		"regions": [
		  "ord"
		]
	}
}`
		fmt.Fprint(writer, response)
	})

	pzUpdateReq := &CDNZoneReq{
		Label:   "new-label",
		CORS:    BoolToBoolPtr(true),
		BlockAI: BoolToBoolPtr(true),
	}

	pz, _, err := client.CDN.UpdatePullZone(ctx, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5", pzUpdateReq)
	if err != nil {
		t.Errorf("CDN.UpdatePullZone returned error: %v", err)
	}

	expected := &CDNZone{
		ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		Label:         "new-label",
		DateCreated:   "2024-01-25 09:41:05",
		Status:        "active",
		OriginScheme:  "https",
		OriginDomain:  "constant.com",
		CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
		VanityDomain:  "my.domain.com",
		CacheSize:     50000000,
		Requests:      0,
		BytesIn:       0,
		BytesOut:      0,
		PacketsPerSec: 50,
		DatePurged:    "2024-01-25 11:39:04",
		CORS:          true,
		GZIP:          true,
		BlockAI:       true,
		BlockBadBots:  true,
		Regions:       []string{"ord"},
	}

	if !reflect.DeepEqual(pz, expected) {
		t.Errorf("CDN.UpdatePullZone returned %+v, expected %+v", pz, expected)
	}
}

// Push Zones =================================================================

func TestCDNPushZoneServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(cdnPushPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
  "push_zone": {
    "id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
    "date_created": "2024-01-25 09:41:05",
    "status": "active",
    "label": "my-pushzone",
    "cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
    "vanity_domain": "my.domain.com",
    "cache_size": 50000000,
    "requests": null,
    "in_bytes": null,
    "out_bytes": null,
    "packets_per_sec": 50,
    "cors": false,
    "gzip": true,
    "block_ai": false,
    "block_bad_bots": true,
    "regions": [
      "yto"
    ]
  }
}`
		fmt.Fprint(writer, response)
	})

	pzReq := &CDNZoneReq{
		Label:        "my-pushzone",
		VanityDomain: "my.domain.com",
		CORS:         BoolToBoolPtr(false),
		GZIP:         BoolToBoolPtr(true),
		BlockAI:      BoolToBoolPtr(false),
		BlockBadBots: BoolToBoolPtr(true),
	}

	pz, _, err := client.CDN.CreatePushZone(ctx, pzReq)
	if err != nil {
		t.Errorf("CDN.CreatePushZone returned error: %v", err)
	}

	expected := &CDNZone{
		ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		Label:         "my-pushzone",
		DateCreated:   "2024-01-25 09:41:05",
		Status:        "active",
		CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
		VanityDomain:  "my.domain.com",
		CacheSize:     50000000,
		Requests:      0,
		BytesIn:       0,
		BytesOut:      0,
		PacketsPerSec: 50,
		CORS:          false,
		GZIP:          true,
		BlockAI:       false,
		BlockBadBots:  true,
		Regions:       []string{"yto"},
	}

	if !reflect.DeepEqual(pz, expected) {
		t.Errorf("CDN.CreatePushZone returned %+v, expected %+v", pz, expected)
	}
}

func TestCDNPushZoneServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(cdnPushPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"push_zones": [
		{
			"id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
			"date_created": "2024-01-25 09:41:05",
			"status": "active",
			"label": "my-pushzone",
			"cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
			"vanity_domain": "my.domain.com",
			"cache_size": 50000000,
			"requests": 0,
			"in_bytes": 0,
			"out_bytes": 0,
			"packets_per_sec": 50,
			"cors": false,
			"gzip": true,
			"block_ai": false,
			"block_bad_bots": true,
			"regions": [
			  "yto"
			]
		}
	],
	"meta": {
		"total": 1,
		"links": {
			"next": "",
			"prev": ""
		}
	}
}`
		fmt.Fprint(writer, response)
	})

	pzs, meta, _, err := client.CDN.ListPushZones(ctx)
	if err != nil {
		t.Errorf("CDN.ListPushZone returned error: %v", err)
	}

	expected := []CDNZone{
		{
			ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
			Label:         "my-pushzone",
			DateCreated:   "2024-01-25 09:41:05",
			Status:        "active",
			CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
			VanityDomain:  "my.domain.com",
			CacheSize:     50000000,
			Requests:      0,
			BytesIn:       0,
			BytesOut:      0,
			PacketsPerSec: 50,
			CORS:          false,
			GZIP:          true,
			BlockAI:       false,
			BlockBadBots:  true,
			Regions:       []string{"yto"},
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(pzs, expected) {
		t.Errorf("CDN.ListPushZone returned %+v, expected %+v", pzs, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("CDN.ListPushZone meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestCDNPushZoneServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	path := fmt.Sprintf("%s/%s", cdnPushPath, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"push_zone": {
		"id": "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		"date_created": "2024-01-25 09:41:05",
		"status": "active",
		"label": "new-label",
		"cdn_url": "https://cdn-wdghak5h67sm.vultrcdn.com",
		"vanity_domain": "my.domain.com",
		"cache_size": 50000000,
		"requests": 0,
		"in_bytes": 0,
		"out_bytes": 0,
		"packets_per_sec": 50,
		"cors": true,
		"gzip": true,
		"block_ai": true,
		"block_bad_bots": true,
		"regions": [
		  "yto"
		]
	}
}`
		fmt.Fprint(writer, response)
	})

	pzUpdateReq := &CDNZoneReq{
		Label:   "new-label",
		CORS:    BoolToBoolPtr(true),
		BlockAI: BoolToBoolPtr(true),
	}

	pz, _, err := client.CDN.UpdatePushZone(ctx, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5", pzUpdateReq)
	if err != nil {
		t.Errorf("CDN.UpdatePushZone returned error: %v", err)
	}

	expected := &CDNZone{
		ID:            "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5",
		Label:         "new-label",
		DateCreated:   "2024-01-25 09:41:05",
		Status:        "active",
		CDNURL:        "https://cdn-wdghak5h67sm.vultrcdn.com",
		VanityDomain:  "my.domain.com",
		CacheSize:     50000000,
		Requests:      0,
		BytesIn:       0,
		BytesOut:      0,
		PacketsPerSec: 50,
		CORS:          true,
		GZIP:          true,
		BlockAI:       true,
		BlockBadBots:  true,
		Regions:       []string{"yto"},
	}

	if !reflect.DeepEqual(pz, expected) {
		t.Errorf("CDN.UpdatePushZone returned %+v, expected %+v", pz, expected)
	}
}

func TestCDNPushZoneServiceHandler_CreateFileEndpoint(t *testing.T) {
	setup()
	defer teardown()

	path := fmt.Sprintf("%s/%s/files", cdnPushPath, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
  "upload_endpoint": {
    "URL": "https://cdn.s3-ewr-000.vultr.dev/v-cdn-agent-assets",
    "inputs": {
      "acl": "public-read",
      "key": "cdn-ady5dwsa6mdh.vultrcdn.com/my-file.jpg",
      "X-Amz-Credential": "kNCaYoUJZ6szuajKsgN/20240418/us-east-1/s3/aws4_request",
      "X-Amz-Algorithm": "AWS4-HMAC-SHA256",
      "Policy": "eyJleHBpcmF0aW9uIjoiMjAyNC0wNC0xOFQxMzowNzo0MloiLCJjb25kaXRpb25zIjpbeyJhY2wiOiJwdWJsaWMtcmVhZCJ9LHsiYnVja2V0Ijoidi1jZG4tYWdlbnQtYXNzZXRzIn0seyJrZXkiOiJjZG4tYWR5NWR3c2E2bWRoLnZ1bHRyY2RuLmNvbVwvcGF0Y2guanBnIn0sWyJjb250ZW50LWxlbmd0aC1yYW5nZSIsODU3NzczLDg1Nzc3M10seyJYLUFtei1EYXRlIjoiMjAyNDA0MThUMTMwMjQyWiJ9LHsiWC1BbXotQ3JlZGVudGlhbCI6InlrTkNhWW9VSlo2c3p1YWpLc2dOXC8yMDI0MDQxOFwvdXMtZWFzdC0xXC9zM1wvYXdzNF9yZXF1ZXN0In0seyJYLUFtei1BbGdvcml0aG0iOiJBV1M0LUhNQUMtU0hBMjU2In1dfQ==",
      "X-Amz-Signature": "8cc2328bf9bd9531ccae5f8b156e7f578f3ee4414bb60f5eac97bbb62a0f2536"
    }
  }
}`
		fmt.Fprint(writer, response)
	})

	endReq := &CDNZoneEndpointReq{
		Name: "my-image.jpg",
		Size: 857773,
	}

	fileEndpoint, _, err := client.CDN.CreatePushZoneFileEndpoint(ctx, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5", endReq)
	if err != nil {
		t.Errorf("CDN.CreatePushZoneFileEndpoint returned error: %v", err)
	}

	expected := &CDNZoneEndpoint{
		URL: "https://cdn.s3-ewr-000.vultr.dev/v-cdn-agent-assets",
		Inputs: CDNZoneEndpointInputs{
			ACL:        "public-read",
			Key:        "cdn-ady5dwsa6mdh.vultrcdn.com/my-file.jpg",
			Policy:     "eyJleHBpcmF0aW9uIjoiMjAyNC0wNC0xOFQxMzowNzo0MloiLCJjb25kaXRpb25zIjpbeyJhY2wiOiJwdWJsaWMtcmVhZCJ9LHsiYnVja2V0Ijoidi1jZG4tYWdlbnQtYXNzZXRzIn0seyJrZXkiOiJjZG4tYWR5NWR3c2E2bWRoLnZ1bHRyY2RuLmNvbVwvcGF0Y2guanBnIn0sWyJjb250ZW50LWxlbmd0aC1yYW5nZSIsODU3NzczLDg1Nzc3M10seyJYLUFtei1EYXRlIjoiMjAyNDA0MThUMTMwMjQyWiJ9LHsiWC1BbXotQ3JlZGVudGlhbCI6InlrTkNhWW9VSlo2c3p1YWpLc2dOXC8yMDI0MDQxOFwvdXMtZWFzdC0xXC9zM1wvYXdzNF9yZXF1ZXN0In0seyJYLUFtei1BbGdvcml0aG0iOiJBV1M0LUhNQUMtU0hBMjU2In1dfQ==",
			Credential: "kNCaYoUJZ6szuajKsgN/20240418/us-east-1/s3/aws4_request",
			Algorithm:  "AWS4-HMAC-SHA256",
			Signature:  "8cc2328bf9bd9531ccae5f8b156e7f578f3ee4414bb60f5eac97bbb62a0f2536",
		},
	}

	if !reflect.DeepEqual(fileEndpoint, expected) {
		t.Errorf("CDN.ListPullZoneFiles returned %+v, expected %+v", fileEndpoint, expected)
	}
}

func TestCDNPushZoneServiceHandler_ListFiles(t *testing.T) {
	setup()
	defer teardown()

	path := fmt.Sprintf("%s/%s/files", cdnPushPath, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
  "files": [
    {
      "name": "my-file.jpg",
      "size": 857773,
      "last_modified": "2024-04-02 14:11:12"
    }
  ],
  "count": 1,
  "total_size": 857773
}`
		fmt.Fprint(writer, response)
	})

	files, _, err := client.CDN.ListPushZoneFiles(ctx, "ef4d95d5-98dc-4710-94f5-0ee97e70a9a5")
	if err != nil {
		t.Errorf("CDN.ListPushZoneFiles returned error: %v", err)
	}

	expected := &CDNZoneFileData{
		Count: 1,
		Size:  857773,
		Files: []CDNZoneFile{
			{
				Name:         "my-file.jpg",
				Size:         857773,
				DateModified: "2024-04-02 14:11:12",
			},
		},
	}

	if !reflect.DeepEqual(files, expected) {
		t.Errorf("CDN.ListPullZoneFiles returned %+v, expected %+v", files, expected)
	}
}
