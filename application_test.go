package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestApplicationServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/applications", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
		"applications": [
			{
            	"id": 1,
            	"name": "LEMP",
            	"short_name": "lemp",
            	"deploy_name": "LEMP on CentOS 6 x64"
        	}
    	],
    	"meta": {
        	"total": 29,
        	"links": {
            	"next": "bmV4dF9fNDM=",
            	"prev": ""
        		}
    		}
		}
		`

		fmt.Fprint(w, response)
	})

	options := &ListOptions{
		PerPage: 1,
		Cursor:  "",
	}
	apps, meta, err := client.Application.List(ctx, options)
	if err != nil {
		t.Errorf("Application.List returned error: %v", err)
	}

	expected := []Application{
		{
			ID:         1,
			Name:       "LEMP",
			ShortName:  "lemp",
			DeployName: "LEMP on CentOS 6 x64",
		},
	}

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("Application.List apps returned %+v, expected %+v", apps, expected)
	}

	expectedMeta := &Meta{
		Total: 29,
		Links: &Links{
			Next: "bmV4dF9fNDM=",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Application.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
