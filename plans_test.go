package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPlansServiceHandler_GetAllList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"201": {"VPSPLANID": "201","name": "1024 MB RAM,25 GB SSD,1.00 TB BW","vcpu_count": "1","ram": "1024","disk": "25","bandwidth": "1.00","bandwidth_gb": "1024","price_per_month": "5.00","plan_type": "SSD","windows": false,"available_locations": [1,2,3,4,5,6]}}`
		fmt.Fprint(writer, response)
	})

	plans, err := client.Plans.GetAllList(ctx, "vc2")

	if err != nil {
		t.Errorf("Plans.GetAllList returned %+v", err)
	}

	expected := []Plans{{
		VpsID:       201,
		Name:        "1024 MB RAM,25 GB SSD,1.00 TB BW",
		VCpus:       1,
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
		t.Errorf("Plans.GetAllList  returned %+v, expected %+v", plans, expected)
	}
}

func TestPlansServiceHandler_GetBareMetalList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/plans/list_baremetal", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"99": {"METALPLANID": "99","name": "32768 MB RAM,4x 240 GB SSD,1.00 TB BW","cpu_count": 12,"ram": 32768,"disk": "4x 240 GB SSD","bandwidth_tb": 1,"price_per_month": 600,"plan_type": "SSD","deprecated": false,"available_locations": [1]}}`
		fmt.Fprint(writer, response)
	})

	bareMetalPlans, err := client.Plans.GetBareMetalList(ctx)

	if err != nil {
		t.Errorf("Plans.GetBareMetalList returned %+v", err)
	}

	expected := []BareMetalPlan{
		{
			BareMetalID: "99",
			Name:        "32768 MB RAM,4x 240 GB SSD,1.00 TB BW",
			Cpus:        12,
			RAM:         32768,
			Disk:        "4x 240 GB SSD",
			Bandwidth:   1,
			Price:       600,
			PlanType:    "SSD",
			Deprecated:  false,
			Regions:     []int{1},
		},
	}

	if !reflect.DeepEqual(bareMetalPlans, expected) {
		t.Errorf("Plans.GetBareMetalList  returned %+v, expected %+v", bareMetalPlans, expected)
	}
}

func TestPlansServiceHandler_GetVc2List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list_vc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"201": {"VPSPLANID": "201","name": "1024 MB RAM,25 GB SSD,1.00 TB BW","vcpu_count": "1","ram": "1024","disk": "25","bandwidth": "1.00","bandwidth_gb": "1024","price_per_month": "5.00","plan_type": "SSD"}}`
		fmt.Fprint(writer, response)
	})

	vc2, err := client.Plans.GetVc2List(ctx)

	if err != nil {
		t.Errorf("Plans.GetBareMetalList returned %+v", err)
	}

	expected := []VCPlan{
		{
			VpsID:       "201",
			Name:        "1024 MB RAM,25 GB SSD,1.00 TB BW",
			Cpus:        "1",
			RAM:         "1024",
			Disk:        "25",
			Bandwidth:   "1.00",
			BandwidthGB: "1024",
			Cost:        "5.00",
			PlanType:    "SSD",
		},
	}
	if !reflect.DeepEqual(vc2, expected) {
		t.Errorf("Plans.GetVc2List  returned %+v, expected %+v", vc2, expected)
	}
}

func TestPlansServiceHandler_GetVdc2List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/plans/list_vdc2", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"115": {"VPSPLANID": "115","name": "8192 MB RAM,110 GB SSD,10.00 TB BW","vcpu_count": "2","ram": "8192","disk": "110","bandwidth": "10.00","bandwidth_gb": "10240","price_per_month": "60.00","plan_type": "DEDICATED"}}`
		fmt.Fprint(writer, response)
	})

	vdc2, err := client.Plans.GetVdc2List(ctx)

	if err != nil {
		t.Errorf("Plans.GetVdc2List returned %+v", err)
	}

	expected := []VCPlan{
		{
			VpsID:       "115",
			Name:        "8192 MB RAM,110 GB SSD,10.00 TB BW",
			Cpus:        "2",
			RAM:         "8192",
			Disk:        "110",
			Bandwidth:   "10.00",
			BandwidthGB: "10240",
			Cost:        "60.00",
			PlanType:    "DEDICATED",
		},
	}

	if !reflect.DeepEqual(vdc2, expected) {
		t.Errorf("Plans.GetVdc2List  returned %+v, expected %+v", vdc2, expected)
	}
}
