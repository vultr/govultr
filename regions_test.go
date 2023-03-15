package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRegionServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/regions", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"regions":[{"id":"ams","city": "test", "country":"NL","continent":"Europe","options":["ddos_protection"]}],"meta":{"total":1,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	region, meta, _, err := client.Region.List(ctx, nil)

	if err != nil {
		t.Errorf("Region.List returned error: %v", err)
	}

	expectedRegion := []Region{
		{
			ID:        "ams",
			City:      "test",
			Country:   "NL",
			Continent: "Europe",
			Options:   []string{"ddos_protection"},
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(region, expectedRegion) {
		t.Errorf("Region.List region returned %+v, expected %+v", region, expectedRegion)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Region.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestRegionServiceHandler_Availability(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/regions/ewr/availability", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"available_plans":["vc2-1c-1gb","vc2-1c-2gb","vc2-2c-4gb"]}`
		fmt.Fprint(writer, response)
	})

	region, _, err := client.Region.Availability(ctx, "ewr", "")

	if err != nil {
		t.Errorf("Region.Availability returned error: %v", err)
	}

	expected := &PlanAvailability{AvailablePlans: []string{"vc2-1c-1gb", "vc2-1c-2gb", "vc2-2c-4gb"}}
	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Region.Availability returned %+v, expected %+v", region, expected)
	}
}
