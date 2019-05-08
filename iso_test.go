package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestIsoServiceHandler_CreateFromURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/iso/create_from_url", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"ISOID": 24}`

		fmt.Fprint(writer, response)
	})

	iso, err := client.Iso.CreateFromURL(ctx, "domain.com/coolest-iso-ever.iso")

	if err != nil {
		t.Errorf("Iso.CreateFromURL returned %+v, expected %+v", err, nil)
	}

	expected := &Iso{IsoID: 24}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Iso.CreateFromURL returned %+v, expected %+v", iso, expected)
	}
}

func TestIsoServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/iso/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Iso.Delete(ctx, 24)

	if err != nil {
		t.Errorf("Iso.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestIsoServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/iso/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "24": { "ISOID": 24,"date_created": "2014-04-01 14:10:09","filename": "CentOS-6.5-x86_64-minimal.iso","size": 9342976,"md5sum": "ec066","sha512sum": "1741f890bce04613f60b0","status": "complete"}}`
		fmt.Fprint(writer, response)
	})

	iso, err := client.Iso.GetList(ctx)

	if err != nil {
		t.Errorf("Iso.GetList returned %+v, expected %+v", err, nil)
	}

	expected := []Iso{
		{IsoID: 24, DateCreated: "2014-04-01 14:10:09", FileName: "CentOS-6.5-x86_64-minimal.iso", Size: 9342976, MD5Sum: "ec066", SHA512Sum: "1741f890bce04613f60b0", Status: "complete"},
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Iso.GetList returned %+v, expected %+v", iso, expected)
	}
}

func TestIsoServiceHandler_GetPublicList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/iso/list_public", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"204515": {"ISOID": 204515,"name": "CentOS 7","description": "7 x86_64 Minimal"}}`
		fmt.Fprint(writer, response)
	})

	iso, err := client.Iso.GetPublicList(ctx)

	if err != nil {
		t.Errorf("Iso.GetPublicList returned %+v, expected %+v", err, nil)
	}

	expected := []PublicIso{
		{IsoID: 204515, Name: "CentOS 7", Description: "7 x86_64 Minimal"},
	}

	if !reflect.DeepEqual(iso, expected) {
		t.Errorf("Iso.GetPublicList returned %+v, expected %+v", iso, expected)
	}
}
