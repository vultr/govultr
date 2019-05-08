package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDNSRecordsServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/create_record", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSRecord.Create(ctx, "domain.com", "A", "api", "127.0.0.1", 300, 0)

	if err != nil {
		t.Errorf("DNSRecord.Create returned %+v, expected %+v", err, nil)
	}
}

func TestDNSRecordsServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/delete_record", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSRecord.Delete(ctx, "domain.com", "12345678")

	if err != nil {
		t.Errorf("DNSRecord.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestDNSRecordsServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/records", func(writer http.ResponseWriter, request *http.Request) {
		response := `[{"type": "A","name": "", "data": "127.0.0.1","priority": 0,"RECORDID": 1265276,"ttl": 300},{"type": "A","name": "", "data": "127.0.0.1","priority": 0,"RECORDID": 1265276,"ttl": 300}]`

		fmt.Fprint(writer, response)
	})

	records, err := client.DNSRecord.GetList(ctx, "domain.com")
	if err != nil {
		t.Errorf("DNSRecord.GetList returned %+v, expected %+v", err, nil)
	}

	expected := []DNSRecord{
		{Type: "A", Name: "", Data: "127.0.0.1", Priority: 0, RecordID: 1265276, TTL: 300},
		{Type: "A", Name: "", Data: "127.0.0.1", Priority: 0, RecordID: 1265276, TTL: 300},
	}

	if !reflect.DeepEqual(records, expected) {
		t.Errorf("DNSRecord.GetList returned %+v, expected %+v", records, expected)
	}
}

func TestDNSRecordsServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/update_record", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writer)
	})

	params := &DNSRecord{
		RecordID: 14283638,
		Name:     "api",
		Data:     "turnip.data",
		TTL:      120,
		Priority: 10,
	}

	err := client.DNSRecord.Update(ctx, "turnip.services", params)

	if err != nil {
		t.Errorf("DNSRecord.Update returned %+v, expected %+v", err, nil)
	}
}
