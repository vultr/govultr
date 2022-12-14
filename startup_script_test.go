package govultr

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStartupScriptServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"startup_script": {"id": "14356","date_created": "2020-07-07 18:52:56","date_modified": "2020-07-07 18:59:54","name": "govultr","type": "boot","script": "dGVzdGFwaXVwZGF0ZQ=="}}`
		fmt.Fprint(writer, response)
	})

	script := &StartupScriptReq{
		Name:   "govultr",
		Type:   "boot",
		Script: "dGVzdGFwaXVwZGF0ZQ==",
	}
	s, err := client.StartupScript.Create(context.Background(), script)

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

	if !reflect.DeepEqual(s, expected) {
		t.Errorf("StartupScript.Create returned %+v, expected %+v", s, expected)
	}
}

func TestStartupScriptServiceHandler_GET(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/14350", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"startup_script": {"id": "14350","date_created": "2020-06-08 17:58:10","date_modified": "2020-06-08 17:59:54","name": "govultr","type": "pxe","script": "dGVzdA=="}}`
		fmt.Fprint(writer, response)
	})

	scripts, err := client.StartupScript.Get(ctx, "14350")

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

func TestStartupScriptServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/1234", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	script := &StartupScriptReq{
		Name:   "foo",
		Type:   "boot",
		Script: "dGVzdA==",
	}

	err := client.StartupScript.Update(ctx, "1234", script)

	if err != nil {
		t.Errorf("StartupScript.Update returned error: %+v", err)
	}
}

func TestStartupScriptServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts/1234", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.StartupScript.Delete(ctx, "1234")

	if err != nil {
		t.Errorf("StartupScript.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestStartupScriptServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/startup-scripts", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"startup_scripts": [{"id": "14350","date_created": "2020-06-08 17:58:10","date_modified": "2020-06-08 17:59:54","name": "govultr","type": "pxe","script": "dGVzdA=="}],"meta": {"total": 1,"links": {"next": "","prev": ""}}}`
		fmt.Fprint(writer, response)
	})

	scripts, meta, err := client.StartupScript.List(ctx, nil)

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
