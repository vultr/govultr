package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOSServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/os", func(w http.ResponseWriter, r *http.Request) {
		response := `{"os":[{"id":124,"name":"Windows 2012 R2 x64","arch":"x64","family":"windows"}],"meta":{"total":27,"links":{"next":"","prev":""}}}`
		fmt.Fprint(w, response)
	})

	os, meta, _, err := client.OS.List(ctx, nil)
	if err != nil {
		t.Errorf("OS.List returned error: %v", err)
	}

	expectedOS := []OS{
		{
			ID:     124,
			Name:   "Windows 2012 R2 x64",
			Arch:   "x64",
			Family: "windows",
		},
	}
	expectedMeta := &Meta{
		Total: 27,
		Links: &Links{},
	}

	if !reflect.DeepEqual(os, expectedOS) {
		t.Errorf("OS.List os returned %+v, expected %+v", os, expectedOS)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("OS.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
