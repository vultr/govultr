package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOSServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/os/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"127": {
				"OSID": 127,
				"name": "CentOS 6 x64",
				"arch": "x64",
				"family": "centos",
				"windows": false
			}
		}
		`

		fmt.Fprint(w, response)
	})

	apps, err := client.OS.GetList(ctx)
	if err != nil {
		t.Errorf("OS.GetList returned error: %v", err)
	}

	expected := []OS{
		{
			OsID:    127,
			Name:    "CentOS 6 x64",
			Arch:    "x64",
			Family:  "centos",
			Windows: false,
		},
	}

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("OS.GetList returned %+v, expected %+v", apps, expected)
	}
}

func TestOSServiceHandler_GetList_StringIDs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/os/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"127": {
				"OSID": 127,
				"name": "CentOS 6 x64",
				"arch": "x64",
				"family": "centos",
				"windows": false
			}
		}
		`

		fmt.Fprint(w, response)
	})

	apps, err := client.OS.GetList(ctx)
	if err != nil {
		t.Errorf("OS.GetList returned error: %v", err)
	}

	expected := []OS{
		{
			OsID:    127,
			Name:    "CentOS 6 x64",
			Arch:    "x64",
			Family:  "centos",
			Windows: false,
		},
	}

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("OS.GetList returned %+v, expected %+v", apps, expected)
	}
}

func TestOSServiceHandler_GetListEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/os/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	apps, err := client.OS.GetList(ctx)
	if err != nil {
		t.Errorf("OS.GetList returned error: %v", err)
	}

	var expected []OS

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("OS.GetList returned %+v, expected %+v", apps, expected)
	}
}
