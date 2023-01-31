package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBillingServiceHandler_ListHistory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/history", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"billing_history": [
				{
					"id": 5317720,
					"date": "2018-04-01T00:30:05+00:00",
					"type": "invoice",
					"description": "Invoice #5317720",
					"amount": 2.35,
					"balance": -497.65
				  }
			],
			"meta": {
				"total":1,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	history, meta,_, err := client.Billing.ListHistory(ctx, nil)
	if err != nil {
		t.Errorf("Billing.ListHistory returned error: %v", err)
	}

	expected := []History{
		{
			ID:          5317720,
			Date:        "2018-04-01T00:30:05+00:00",
			Type:        "invoice",
			Description: "Invoice #5317720",
			Amount:      2.35,
			Balance:     -497.65,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(history, expected) {
		t.Errorf("Billing.ListHistory returned %+v, expected %+v", history, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Billing.ListHistory returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBillingServiceHandler_ListInvoices(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/invoices", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"billing_invoices": [
				{
					"id": 5317720,
					"date": "2018-04-01T00:30:05+00:00",
					"description": "Invoice #5317720",
					"amount": 2.35,
					"balance": -497.65
				  }
			],
			"meta": {
				"total":1,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	invoices, meta,_,err := client.Billing.ListInvoices(ctx, nil)
	if err != nil {
		t.Errorf("Billing.ListInvoices returned error: %v", err)
	}

	expected := []Invoice{
		{
			ID:          5317720,
			Date:        "2018-04-01T00:30:05+00:00",
			Description: "Invoice #5317720",
			Amount:      2.35,
			Balance:     -497.65,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(invoices, expected) {
		t.Errorf("Billing.ListInvoices returned %+v, expected %+v", invoices, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Billing.ListInvoices returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBillingServiceHandler_ListHistoryEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/history", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"billing_history": [],
			"meta": {
				"total":0,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	history, meta,_,err := client.Billing.ListHistory(ctx, nil)
	if err != nil {
		t.Errorf("Billing.ListHistory returned error: %v", err)
	}

	expected := []History{}

	if !reflect.DeepEqual(history, expected) {
		t.Errorf("Billing.ListHistory returned %+v, expected %+v", history, expected)
	}

	expectedMeta := &Meta{
		Total: 0,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Billing.ListHistory meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBillingServiceHandler_ListInvoicesEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/invoices", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"billing_invoices": [],
			"meta": {
				"total":0,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})

	invoices, meta,_, err := client.Billing.ListInvoices(ctx, nil)
	if err != nil {
		t.Errorf("Billing.ListInvoices returned error: %v", err)
	}

	expected := []Invoice{}

	if !reflect.DeepEqual(invoices, expected) {
		t.Errorf("Billing.ListInvoices returned %+v, expected %+v", invoices, expected)
	}

	expectedMeta := &Meta{
		Total: 0,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Billing.ListInvoices meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestBillingServiceHandler_GetInvoice(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/invoices/123456", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"billing_invoice": {
				"id": 123456,
				"date": "2018-04-01T00:30:05+00:00",
				"description": "Invoice #5317782",
				"amount": 2.35,
				"balance": -497.65
			  }
		}
		`

		fmt.Fprint(w, response)
	})

	invoice,_,err := client.Billing.GetInvoice(ctx, "123456")
	if err != nil {
		t.Errorf("Billing.GetInvoice returned error: %v", err)
	}

	expected := &Invoice{
		ID:          123456,
		Date:        "2018-04-01T00:30:05+00:00",
		Description: "Invoice #5317782",
		Amount:      2.35,
		Balance:     -497.65,
	}

	if !reflect.DeepEqual(invoice, expected) {
		t.Errorf("Billing.GetInvoice returned %+v, expected %+v", invoice, expected)
	}
}

func TestBillingServiceHandler_ListInvoiceItems(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/billing/invoices/123456/items", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"invoice_items": [
				{
					"description": "1.1.1.1 (1024 MB)",
					"product": "Vultr Cloud Compute",
					"start_date": "2018-03-18T21:57:58+00:00",
					"end_date": "2018-04-01T00:00:00+00:00",
					"units": 315,
					"unit_type": "hours",
					"unit_price": 0.0074,
					"total": 2.35
				}
			],
			"meta": {
				"total":1,
				"links": {
					"next":"",
					"prev":""
				}
			}
		}
		`

		fmt.Fprint(w, response)
	})
	invoices, meta,_,err := client.Billing.ListInvoiceItems(ctx, 123456, nil)
	if err != nil {
		t.Errorf("Billing.ListInvoiceItems returned error: %v", err)
	}

	expected := []InvoiceItem{
		{
			Description: "1.1.1.1 (1024 MB)",
			Product:     "Vultr Cloud Compute",
			StartDate:   "2018-03-18T21:57:58+00:00",
			EndDate:     "2018-04-01T00:00:00+00:00",
			Units:       315,
			UnitType:    "hours",
			UnitPrice:   0.0074,
			Total:       2.35,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(invoices, expected) {
		t.Errorf("Billing.ListInvoiceItems returned %+v, expected %+v", invoices, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Billing.ListInvoiceItems returned %+v, expected %+v", meta, expectedMeta)
	}
}
