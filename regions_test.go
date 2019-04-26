package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRegionsServiceHandler_Availability(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/regions/availability", func(writer http.ResponseWriter, request *http.Request) {
		response := `[201,202,203,204,205,206,115,29,93,94,95,96,97,98,100]`
		fmt.Fprint(writer, response)
	})

	region, err := client.Regions.Availability(ctx, 1, "vc2")

	if err != nil {
		t.Errorf("Regions.Availability returned error: %v", err)
	}

	expected := []int{201, 202, 203, 204, 205, 206, 115, 29, 93, 94, 95, 96, 97, 98, 100}

	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Regions.Availability returned %+v, expected %+v", region, expected)
	}
}

func TestRegionsServiceHandler_BareMetalAvailability(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/regions/availability_baremetal", func(writer http.ResponseWriter, request *http.Request) {
		response := `[1,2,3,4]`
		fmt.Fprint(writer, response)
	})

	region, err := client.Regions.BareMetalAvailability(ctx, 1)

	if err != nil {
		t.Errorf("Regions.BareMetalAvailability returned error: %v", err)
	}

	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Regions.BareMetalAvailability returned %+v, expected %+v", region, expected)
	}
}

func TestRegionsServiceHandler_Vc2Availability(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/regions/availability_vc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `[115,29,93,94,95,96,97,98,100]`
		fmt.Fprint(writer, response)
	})

	region, err := client.Regions.Vc2Availability(ctx, 1)

	if err != nil {
		t.Errorf("Regions.Vc2Availability returned error: %v", err)
	}

	expected := []int{115, 29, 93, 94, 95, 96, 97, 98, 100}

	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Regions.Vc2Availability returned %+v, expected %+v", region, expected)
	}
}

func TestRegionsServiceHandler_Vdc2Availability(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/regions/availability_vdc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `[115,29,93,94]`
		fmt.Fprint(writer, response)
	})

	region, err := client.Regions.Vdc2Availability(ctx, 1)

	if err != nil {
		t.Errorf("Regions.Vdc2Availability returned error: %v", err)
	}

	expected := []int{115, 29, 93, 94}

	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Regions.Vdc2Availability returned %+v, expected %+v", region, expected)
	}
}

func TestRegionsServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/regions/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1": {"DCID": "1","name": "New Jersey","country": "US","continent": "North America","state": "NJ","ddos_protection": true,"block_storage": true,"regioncode": "EWR"},"2": {"DCID": "1","name": "New Jersey","country": "US","continent": "North America","state": "NJ","ddos_protection": true,"block_storage": true,"regioncode": "EWR"}}`
		fmt.Fprint(writer, response)
	})

	region, err := client.Regions.GetList(ctx)

	if err != nil {
		t.Errorf("Regions.GetList returned error: %v", err)
	}

	expected := []Region{
		{
			RegionID:     "1",
			Name:         "New Jersey",
			Country:      "US",
			Continent:    "North America",
			State:        "NJ",
			Ddos:         true,
			BlockStorage: true,
			RegionCode:   "EWR",
		},
		{
			RegionID:     "1",
			Name:         "New Jersey",
			Country:      "US",
			Continent:    "North America",
			State:        "NJ",
			Ddos:         true,
			BlockStorage: true,
			RegionCode:   "EWR",
		},
	}

	if !reflect.DeepEqual(region, expected) {
		t.Errorf("Regions.GetList returned %+v, expected %+v", region, expected)
	}
}
