package govultr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestAccountServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account", func(w http.ResponseWriter, r *http.Request) {
		response := `
		{
		"account" : {
			"balance": -5519.11,
			"pending_charges": 57.03,
			"last_payment_date": "2014-07-18 15:31:01",
			"last_payment_amount": -1.00,
			"name": "Test Tester",
			"email" : "example@vultr.com",
			"acls": [
				"subscriptions",
				"billing",
				"support",
				"provisioning"
			]
		}
		}
		`

		fmt.Fprint(w, response)
	})

	account, _, err := client.Account.Get(ctx)
	if err != nil {
		t.Errorf("Account.Get returned error: %v", err)
	}

	expected := &Account{Balance: -5519.11, PendingCharges: 57.03, LastPaymentDate: "2014-07-18 15:31:01", LastPaymentAmount: -1.00, Name: "Test Tester", Email: "example@vultr.com", ACL: []string{"subscriptions", "billing", "support", "provisioning"}}

	if !reflect.DeepEqual(account, expected) {
		t.Errorf("Account.Get returned %+v, expected %+v", account, expected)
	}
}

func TestAccountServiceHandler_GetBandwidth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/bandwidth", func(w http.ResponseWriter, r *http.Request) {
		response := `
		{
			"bandwidth": {
				"previous_month": {
					"timestamp_start": "1735689600",
					"timestamp_end": "1738367999",
					"gb_in": 0,
					"gb_out": 0,
					"total_instance_hours": 1,
					"total_instance_count": 1,
					"instance_bandwidth_credits": 1,
					"free_bandwidth_credits": 2048,
					"purchased_bandwidth_credits": 0,
					"overage": 0,
					"overage_unit_cost": 0.01,
					"overage_cost": 0
				},
				"current_month_to_date": {
					"timestamp_start": "1738368000",
					"timestamp_end": "1739577600",
					"gb_in": 0,
					"gb_out": 0,
					"total_instance_hours": 0,
					"total_instance_count": 0,
					"instance_bandwidth_credits": 0,
					"free_bandwidth_credits": 2048,
					"purchased_bandwidth_credits": 0,
					"overage": 0,
					"overage_unit_cost": 0.01,
					"overage_cost": 0
				},
				"current_month_projected": {
					"timestamp_start": "1738368000",
					"timestamp_end": "1740787199",
					"gb_in": 0,
					"gb_out": 0,
					"total_instance_hours": 0,
					"total_instance_count": 0,
					"instance_bandwidth_credits": 0,
					"free_bandwidth_credits": 2048,
					"purchased_bandwidth_credits": 0,
					"overage": 0,
					"overage_unit_cost": 0.01,
					"overage_cost": 0
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	accountBandwidth, _, err := client.Account.GetBandwidth(ctx)
	if err != nil {
		t.Errorf("Account.GetBandwidth returned error: %v", err)
	}

	expected := &AccountBandwidth{
		PreviousMonth: AccountBandwidthPeriod{
			TimestampStart:            "1735689600",
			TimestampEnd:              "1738367999",
			GBIn:                      0,
			GBOut:                     0,
			TotalInstanceHours:        1,
			TotalInstanceCount:        1,
			InstanceBandwidthCredits:  1,
			FreeBandwidthCredits:      2048,
			PurchasedBandwidthCredits: 0,
			Overage:                   0,
			OverageUnitCost:           0.01,
			OverageCost:               0,
		},
		CurrentMonthToDate: AccountBandwidthPeriod{
			TimestampStart:            "1738368000",
			TimestampEnd:              "1739577600",
			GBIn:                      0,
			GBOut:                     0,
			TotalInstanceHours:        0,
			TotalInstanceCount:        0,
			InstanceBandwidthCredits:  0,
			FreeBandwidthCredits:      2048,
			PurchasedBandwidthCredits: 0,
			Overage:                   0,
			OverageUnitCost:           0.01,
			OverageCost:               0,
		},
		CurrentMonthProjected: AccountBandwidthPeriod{
			TimestampStart:            "1738368000",
			TimestampEnd:              "1740787199",
			GBIn:                      0,
			GBOut:                     0,
			TotalInstanceHours:        0,
			TotalInstanceCount:        0,
			InstanceBandwidthCredits:  0,
			FreeBandwidthCredits:      2048,
			PurchasedBandwidthCredits: 0,
			Overage:                   0,
			OverageUnitCost:           0.01,
			OverageCost:               0,
		},
	}

	if !reflect.DeepEqual(accountBandwidth, expected) {
		t.Errorf("Account.GetBandwidth returned %+v, expected %+v", accountBandwidth, expected)
	}
}

func TestAccountServiceHandler_GetBGP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/bgp", func(w http.ResponseWriter, r *http.Request) {
		response := `
		{
		  "enabled": true,
		  "asn": 20473,
		  "password": "12345"
		}
		`

		fmt.Fprint(w, response)
	})

	accountBGP, _, err := client.Account.GetBGP(ctx)
	if err != nil {
		t.Errorf("Account.GetBGP returned error: %v", err)
	}

	expected := &AccountBGP{
		Enabled:  true,
		ASN:      20473,
		Password: "12345",
	}

	if !reflect.DeepEqual(accountBGP, expected) {
		t.Errorf("Account.GetBGP returned %+v, expected %+v", accountBGP, expected)
	}
}

func TestAccountServiceHandler_SetupBGP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/bgp/setup", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	r := &AccountBGPSetup{
		Prefixes: []string{
			"192.0.2.0/24",
			"2001:db8:0:1::/64",
		},
		ASN:                   1234,
		Password:              "topsecret",
		LetterOfAuthorization: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII=",
		RequestedRoutes:       "full",
		UseCase:               "Use my IP space on Vultr",
	}
	err := client.Account.SetupBGP(ctx, r)

	if err != nil {
		t.Errorf("Account.SetupBGP returned %+v, expected %+v", err, nil)
	}
}

func TestAccountServiceHandler_AddBGPPrefixes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/bgp/prefixes", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	prefixes := []string{
		"192.0.2.0/24",
		"2001:db8:0:1::/64",
	}
	err := client.Account.AddBGPPrefixes(ctx, prefixes)
	if err != nil {
		t.Errorf("Account.AddBGPPrefixes returned %+v, expected %+v", err, nil)
	}

}

func TestAccountServiceHandler_ListBGPPrefixes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/bgp/prefixes", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "prefixes": [
			{
			  "prefix": "192.0.2.0/24",
			  "date_added": "2020-01-01T13:10:10+00:00",
			  "rpki_status": "VALID"
			}
		  ],
		  "meta": {
			"total": 1,
			"links": {
			  "next": "",
			  "prev": ""
			}
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	prefixes, meta, _, err := client.Account.ListBGPPrefixes(ctx, options)
	if err != nil {
		t.Errorf("Account.ListBGPPrefixes returned %+v, expected %+v", err, nil)
	}

	expectedPrefixes := []AccountBGPPrefix{
		{
			Prefix:     "192.0.2.0/24",
			DateAdded:  "2020-01-01T13:10:10+00:00",
			RPKIStatus: "VALID",
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(prefixes, expectedPrefixes) {
		t.Errorf("Account.ListBGPPrefixes returned %+v, expected %+v", prefixes, expectedPrefixes)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Account.ListBGPPrefixes meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestAccountServiceHandler_ListCustomSubscriptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/account/custom-subscriptions", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "custom_subscriptions": [
			{
			  "id": "ab123456-7c89-0123-d456-78901e23f4g5",
			  "label": "test",
			  "description": "test",
			  "type": "Compute Test",
			  "region": "SEA",
			  "status": "active",
			  "date_created": "2026-01-26T14:27:35+00:00",
			  "cost": 5,
			  "pending_charges": 0
		    }
		  ],
		  "meta": {
			"total": 5,
			"links": {
			  "next": "bmV4dF9fYWIxMjM0NTYtN2M4OS0wMTIzLWQ0NTYtNzg5MDFlMjNmNGc1",
			  "prev": ""
			}
		  }
		}
		`

		subs := new(accountCustomSubscriptionBase)
		if err := json.NewDecoder(strings.NewReader(response)).Decode(subs); err != nil {
			t.Errorf("Account.ListCustomSubscriptions returned %+v, expected %+v", err, nil)
		}

		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	subscriptions, meta, _, err := client.Account.ListCustomSubscriptions(ctx, options)
	if err != nil {
		t.Errorf("Account.ListCustomSubscriptions returned %+v, expected %+v", err, nil)
	}

	expectedSubscriptions := []AccountCustomSubscription{
		{
			ID:             "ab123456-7c89-0123-d456-78901e23f4g5",
			Label:          "test",
			Description:    "test",
			Type:           "Compute Test",
			Region:         "SEA",
			Status:         "active",
			DateCreated:    "2026-01-26T14:27:35+00:00",
			Cost:           5,
			PendingCharges: 0,
		},
	}

	expectedMeta := &Meta{
		Total: 5,
		Links: &Links{
			Next: "bmV4dF9fYWIxMjM0NTYtN2M4OS0wMTIzLWQ0NTYtNzg5MDFlMjNmNGc1",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(subscriptions, expectedSubscriptions) {
		t.Errorf("Account.ListCustomSubscriptions returned %+v, expected %+v", subscriptions, expectedSubscriptions)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Account.ListCustomSubscriptions meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
