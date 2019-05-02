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

	mux.HandleFunc("/v1/startupscript/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SCRIPTID": 5
		}
		`

		fmt.Fprint(writer, response)
	})

	s, err := client.StartupScript.Create(context.Background(), "foo", "#!/bin/bash\necho hello world > /root/hello", "pxe")

	if err != nil {
		t.Errorf("StartupScript.Create returned %+v, expected %+v", err, nil)
	}

	expected := &StartupScript{
		ScriptID:     "5",
		DateCreated:  "",
		DateModified: "",
		Name:         "foo",
		Type:         "pxe",
		Script:       "#!/bin/bash\necho hello world > /root/hello",
	}

	if !reflect.DeepEqual(s, expected) {
		t.Errorf("StartupScript.Create returned %+v, expected %+v", s, expected)
	}
}

func TestStartupScriptServiceHandler_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/startupscript/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.StartupScript.Destroy(ctx, "foo")

	if err != nil {
		t.Errorf("StartupScript.Destroy returned %+v, expected %+v", err, nil)
	}
}

func TestStartupScriptServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/startupscript/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"3": {
				"SCRIPTID": "3",
				"date_created": "2014-05-21 15:27:18",
				"date_modified": "2014-05-21 15:27:18",
				"name": "foo",
				"type": "boot",
				"script": "#!/bin/bash echo Hello World > /root/hello"
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	scripts, err := client.StartupScript.GetList(ctx)

	if err != nil {
		t.Errorf("StartupScript.GetList returned error: %v", err)
	}

	expected := []StartupScript{
		{
			ScriptID:     "3",
			Name:         "foo",
			Type:         "boot",
			Script:       "#!/bin/bash echo Hello World > /root/hello",
			DateCreated:  "2014-05-21 15:27:18",
			DateModified: "2014-05-21 15:27:18",
		},
	}

	if !reflect.DeepEqual(scripts, expected) {
		t.Errorf("StartupScript.GetList returned %+v, expected %+v", scripts, expected)
	}
}

func TestStartupScriptServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/startupscript/update", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writer)
	})

	script := &StartupScript{
		ScriptID: "1",
		Name:     "foo",
		Type:     "boot",
		Script:   "#!/bin/bash echo Hello World > /root/hello",
	}

	err := client.StartupScript.Update(ctx, script)

	if err != nil {
		t.Errorf("StartupScript.Update returned error: %+v", err)
	}
}
