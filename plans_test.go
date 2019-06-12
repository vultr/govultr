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

	mux.HandleFunc("/v1/plans/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"201": {"VPSPLANID": "201","name": "1024 MB RAM,25 GB SSD,1.00 TB BW","vcpu_count": "1","ram": "1024","disk": "25","bandwidth": "1.00","bandwidth_gb": "1024","price_per_month": "5.00","plan_type": "SSD","windows": false,"available_locations": [1,2,3,4,5,6]}}`
		fmt.Fprint(writer, response)
	})

	plans, err := client.Plan.List(ctx, "vc2")

	if err != nil {
		t.Errorf("Plan.List returned %+v", err)
	}

	expected := []Plan{{
		VpsID:       201,
		Name:        "1024 MB RAM,25 GB SSD,1.00 TB BW",
		VCPUs:       1,
		RAM:         "1024",
		Disk:        "25",
		Price:       "5.00",
		Bandwidth:   "1.00",
		BandwidthGB: "1024",
		Windows:     false,
		PlanType:    "SSD",
		Regions:     []int{1, 2, 3, 4, 5, 6},
	},
	}

	if !reflect.DeepEqual(plans, expected) {
		t.Errorf("Plan.List  returned %+v, expected %+v", plans, expected)
	}
}

func TestPlanServiceHandler_GetBareMetalList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/plans/list_baremetal", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"99": {"METALPLANID": "99","name": "32768 MB RAM,4x 240 GB SSD,1.00 TB BW","cpu_count": 12,"cpu_model": "E-2186G","ram": 32768,"disk": "4x 240 GB SSD","bandwidth_tb": 1,"price_per_month": 600,"plan_type": "SSD","deprecated": false,"available_locations": [1]}}`
		fmt.Fprint(writer, response)
	})

	bareMetalPlans, err := client.Plan.GetBareMetalList(ctx)

	if err != nil {
		t.Errorf("Plan.GetBareMetalList returned %+v", err)
	}

	expected := []BareMetalPlan{
		{
			BareMetalID: "99",
			Name:        "32768 MB RAM,4x 240 GB SSD,1.00 TB BW",
			CPUs:        12,
			CPUModel:    "E-2186G",
			RAM:         32768,
			Disk:        "4x 240 GB SSD",
			BandwidthTB: 1,
			Price:       600,
			PlanType:    "SSD",
			Deprecated:  false,
			Regions:     []int{1},
		},
	}

	if !reflect.DeepEqual(bareMetalPlans, expected) {
		t.Errorf("Plan.GetBareMetalList  returned %+v, expected %+v", bareMetalPlans, expected)
	}
}

func TestPlanServiceHandler_GetVc2List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list_vc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"201": {"VPSPLANID": "201","name": "1024 MB RAM,25 GB SSD,1.00 TB BW","vcpu_count": "1","ram": "1024","disk": "25","bandwidth": "1.00","bandwidth_gb": "1024","price_per_month": "5.00","plan_type": "SSD"}}`
		fmt.Fprint(writer, response)
	})

	vc2, err := client.Plan.GetVc2List(ctx)

	if err != nil {
		t.Errorf("Plan.GetVc2List returned %+v", err)
	}

	expected := []VCPlan{
		{
			VpsID:       "201",
			Name:        "1024 MB RAM,25 GB SSD,1.00 TB BW",
			VCPUs:       "1",
			RAM:         "1024",
			Disk:        "25",
			Bandwidth:   "1.00",
			BandwidthGB: "1024",
			Price:       "5.00",
			PlanType:    "SSD",
		},
	}
	if !reflect.DeepEqual(vc2, expected) {
		t.Errorf("Plan.GetVc2List  returned %+v, expected %+v", vc2, expected)
	}
}

func TestPlanServiceHandler_GetVdc2List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list_vdc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"115": {"VPSPLANID": "115","name": "8192 MB RAM,110 GB SSD,10.00 TB BW","vcpu_count": "2","ram": "8192","disk": "110","bandwidth": "10.00","bandwidth_gb": "10240","price_per_month": "60.00","plan_type": "DEDICATED"}}`
		fmt.Fprint(writer, response)
	})

	vdc2, err := client.Plan.GetVdc2List(ctx)

	if err != nil {
		t.Errorf("Plan.GetVdc2List returned %+v", err)
	}

	expected := []VCPlan{
		{
			VpsID:       "115",
			Name:        "8192 MB RAM,110 GB SSD,10.00 TB BW",
			VCPUs:       "2",
			RAM:         "8192",
			Disk:        "110",
			Bandwidth:   "10.00",
			BandwidthGB: "10240",
			Price:       "60.00",
			PlanType:    "DEDICATED",
		},
	}

	if !reflect.DeepEqual(vdc2, expected) {
		t.Errorf("Plan.GetVdc2List  returned %+v, expected %+v", vdc2, expected)
	}
}

func TestPlanServiceHandler_GetVc2zList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list_vc2z", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "401": {"VPSPLANID": "401","name": "2048 MB RAM,64 GB SSD,2.00 TB BW","vcpu_count": "1","ram": "2048","disk": "64","bandwidth": "2.00","price_per_month": "12.00","plan_type": "HIGHFREQUENCY"}}`
		fmt.Fprint(writer, response)
	})

	vc2z, err := client.Plan.GetVc2zList(ctx)

	if err != nil {
		t.Errorf("Plan.GetVc2zList returned %+v", err)
	}

	expected := []VCPlan{
		{
			VpsID:     "401",
			Name:      "2048 MB RAM,64 GB SSD,2.00 TB BW",
			VCPUs:     "1",
			RAM:       "2048",
			Disk:      "64",
			Bandwidth: "2.00",
			Price:     "12.00",
			PlanType:  "HIGHFREQUENCY",
		},
	}

	if !reflect.DeepEqual(vc2z, expected) {
		t.Errorf("Plan.GetVc2zList  returned %+v, expected %+v", vc2z, expected)
	}
}
