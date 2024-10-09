package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSubAccountServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/subaccounts", func(writer http.ResponseWriter, request *http.Request) {
		response := `
{
	"subaccount": { 
		"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60", 
		"email": "subaccount@vultr.com", 
		"subaccount_name": "Acme Widgets LLC", 
		"subaccount_id": "472924", 
		"activated": false, 
		"balance": 0, 
		"pending_charges": 0 
	}
}`

		fmt.Fprint(writer, response)
	})

	saReq := &SubAccountReq{
		Email:   "subaccount@vultr.com",
		Name:    "Acme Widgets LLC",
		OtherID: "472924",
	}

	sa, _, err := client.SubAccount.Create(ctx, saReq)
	if err != nil {
		t.Errorf("SubAccount.Create returned error: %v", err)
	}

	expected := &SubAccount{
		ID:             "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Email:          "subaccount@vultr.com",
		Name:           "Acme Widgets LLC",
		OtherID:        "472924",
		Activated:      false,
		Balance:        0,
		PendingCharges: 0,
	}

	if !reflect.DeepEqual(sa, expected) {
		t.Errorf("SubAccount.Create returned %+v, expected %+v", sa, expected)
	}
}

func TestSubAccountServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/subaccounts", func(writer http.ResponseWriter, request *http.Request) {
		response := `
{ 
	"subaccounts": [
		{ 
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60", 
			"email": "subaccount@vultr.com", 
			"subaccount_name": "Acme Widgets LLC", 
			"subaccount_id": "472924", 
			"activated": false, 
			"balance": 0, 
			"pending_charges": 0 
		}
	],
	"meta": {
		"total": 1,
		"links": {
				"next": "",
				"prev": ""
		}
	}
}`
		fmt.Fprint(writer, response)
	})

	sas, meta, _, err := client.SubAccount.List(ctx, nil)
	if err != nil {
		t.Errorf("SubAccount.List returned error: %v", err)
	}

	expected := []SubAccount{
		{
			ID:             "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			Email:          "subaccount@vultr.com",
			Name:           "Acme Widgets LLC",
			OtherID:        "472924",
			Activated:      false,
			Balance:        0,
			PendingCharges: 0,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(sas, expected) {
		t.Errorf("SubAccount.List returned %+v, expected %+v", sas, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("SubAccount.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
