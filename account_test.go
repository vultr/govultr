package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountServiceHandler_GetInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/account/info", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
		"balance": "-5519.11",
		"pending_charges": "57.03",
		"last_payment_date": "2014-07-18 15:31:01",
		"last_payment_amount": "-1.00"
		}
		`

		fmt.Fprint(w, response)
	})

	account, err := client.Account.GetInfo(ctx)
	if err != nil {
		t.Errorf("Account.GetInfo returned error: %v", err)
	}

	expected := &Account{Balance: "-5519.11", PendingCharges: "57.03", LastPaymentDate: "2014-07-18 15:31:01", LastPaymentAmount: "-1.00"}

	if !reflect.DeepEqual(account, expected) {
		t.Errorf("Account.GetInfo returned %+v, expected %+v", account, expected)
	}
}
