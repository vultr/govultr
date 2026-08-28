package govultr

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestStartupScriptServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"startup_scripts": [
		{
			"id": "14350",
			"date_created": "2020-06-08 17:58:10",
			"date_modified": "2020-06-08 17:59:54",
			"name": "govultr",
			"type": "pxe",
			"script": "dGVzdA=="
		}
	],
	"meta": {
		"total": 1,
		"links": {
			"next": "",
			"prev": ""
		}
	}
}`))

	scripts, meta, _, err := client.StartupScript.List(ctx, nil)
	if err != nil {
		t.Errorf("StartupScript.List returned error: %v", err)
	}

	expectedScript := []StartupScript{
		{
			ID:           "14350",
			DateCreated:  "2020-06-08 17:58:10",
			DateModified: "2020-06-08 17:59:54",
			Name:         "govultr",
			Type:         "pxe",
			Script:       "dGVzdA==",
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(scripts, expectedScript) {
		t.Errorf("StartupScript.List scripts returned %+v, expected %+v", scripts, expectedScript)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("StartupScript.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestStartupScriptServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/14350", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"startup_script": {
		"id": "14350",
		"date_created": "2020-06-08 17:58:10",
		"date_modified": "2020-06-08 17:59:54",
		"name": "govultr",
		"type": "pxe",
		"script": "dGVzdA=="
	}
}`))

	scripts, _, err := client.StartupScript.Get(ctx, "14350")
	if err != nil {
		t.Errorf("StartupScript.Get returned error: %v", err)
	}

	expectedScript := &StartupScript{
		ID:           "14350",
		DateCreated:  "2020-06-08 17:58:10",
		DateModified: "2020-06-08 17:59:54",
		Name:         "govultr",
		Type:         "pxe",
		Script:       "dGVzdA==",
	}

	if !reflect.DeepEqual(scripts, expectedScript) {
		t.Errorf("StartupScript.Get scripts returned %+v, expected %+v", scripts, expectedScript)
	}
}

func TestStartupScriptServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"startup_script": {
		"id": "14356",
		"date_created": "2020-07-07 18:52:56",
		"date_modified": "2020-07-07 18:59:54",
		"name": "govultr",
		"type": "boot",
		"script": "dGVzdGFwaXVwZGF0ZQ=="
	}
}`))

	req := &StartupScriptReq{
		Name:   "govultr",
		Type:   "boot",
		Script: "dGVzdGFwaXVwZGF0ZQ==",
	}

	script, _, err := client.StartupScript.Create(context.Background(), req)
	if err != nil {
		t.Errorf("StartupScript.Create returned %+v, expected %+v", err, nil)
	}

	expected := &StartupScript{
		ID:           "14356",
		DateCreated:  "2020-07-07 18:52:56",
		DateModified: "2020-07-07 18:59:54",
		Name:         "govultr",
		Type:         "boot",
		Script:       "dGVzdGFwaXVwZGF0ZQ==",
	}

	if !reflect.DeepEqual(script, expected) {
		t.Errorf("StartupScript.Create returned %+v, expected %+v", script, expected)
	}
}

func TestStartupScriptServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/1234", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	script := &StartupScriptReq{
		Name:   "foo",
		Type:   "boot",
		Script: "dGVzdA==",
	}

	if err := client.StartupScript.Update(ctx, "1234", script); err != nil {
		t.Errorf("StartupScript.Update returned error: %+v", err)
	}
}

func TestStartupScriptServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/1234", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.StartupScript.Delete(ctx, "1234"); err != nil {
		t.Errorf("StartupScript.Delete returned %+v, expected %+v", err, nil)
	}
}
