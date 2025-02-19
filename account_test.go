package govultr

import (
	"fmt"
	"net/http"
	"reflect"
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
