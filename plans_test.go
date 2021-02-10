package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPlanServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/plans", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "plans":[{ "id": "vc2-16c-64gb", "vcpu_count": 16, "ram": 65536, "disk": 1280, "disk_count": 1, "bandwidth": 10240, "monthly_cost": 320, "type": "vc2", "locations": [  "dfw"]}], "meta": { "total": 19, "links": { "next": "", "prev": "" } }}`
		fmt.Fprint(writer, response)
	})

	plans, meta, err := client.Plan.List(ctx, "vc2", nil)
	if err != nil {
		t.Errorf("Plan.List returned %+v", err)
	}

	expectedPlan := []Plan{
		{
			ID:          "vc2-16c-64gb",
			VCPUCount:   16,
			RAM:         65536,
			Disk:        1280,
			DiskCount:   1,
			Bandwidth:   10240,
			MonthlyCost: 320.00,
			Type:        "vc2",
			Locations: []string{
				"dfw",
			},
		},
	}

	expectedMeta := &Meta{
		Total: 19,
		Links: &Links{},
	}

	if !reflect.DeepEqual(plans, expectedPlan) {
		t.Errorf("Plan.List  plans returned %+v, expected %+v", plans, expectedPlan)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Plan.List  meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestPlanServiceHandler_GetBareMetalList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/plans-metal", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "plans_metal":[{"id": "vbm-4c-32gb","cpu_count": 4,"cpu_threads": 8,"cpu_model": "E3-1270v6","ram": 32768,"disk": 240, "disk_count": 1, "bandwidth": 5120,"monthly_cost": 300,"type": "SSD", "locations": [ "dwf"]}], "meta": { "total": 19, "links": { "next": "", "prev": "" } }}`
		fmt.Fprint(writer, response)
	})

	bareMetalPlans, meta, err := client.Plan.ListBareMetal(ctx, nil)
	if err != nil {
		t.Errorf("Plan.GetBareMetalList returned %+v", err)
	}

	expectedPlan := []BareMetalPlan{
		{
			ID:          "vbm-4c-32gb",
			CPUCount:    4,
			CPUModel:    "E3-1270v6",
			CPUThreads:  8,
			RAM:         32768,
			Disk:        240,
			DiskCount:   1,
			Bandwidth:   5120,
			MonthlyCost: 300,
			Type:        "SSD",
			Locations: []string{
				"dwf",
			},
		},
	}

	expectedMeta := &Meta{
		Total: 19,
		Links: &Links{},
	}

	if !reflect.DeepEqual(bareMetalPlans, expectedPlan) {
		t.Errorf("Plan.GetBareMetalList  returned %+v, expected %+v", bareMetalPlans, expectedPlan)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Plan.List  meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
