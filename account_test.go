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
