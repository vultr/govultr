package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDomainRecordsServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/records", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"record":{"id":"dev-preview-abc123","type":"A","name":"www","data":"127.0.0.1","priority":0,"ttl":300}}`
		fmt.Fprint(writer, response)
	})

	r := &DomainRecordReq{
		Name:     "www",
		Type:     "A",
		Data:     "127.0.0.1",
		Priority: 300,
	}
	record, err := client.DomainRecord.Create(ctx, "vultr.com", r)
	if err != nil {
		t.Errorf("DomainRecord.Create returned %+v, expected %+v", err, nil)
	}

	expected := &DomainRecord{
		ID:       "dev-preview-abc123",
		Type:     "A",
		Name:     "www",
		Data:     "127.0.0.1",
		Priority: 0,
		TTL:      300,
	}

	if !reflect.DeepEqual(record, expected) {
		t.Errorf("DomainRecord.Create returned %+v, expected %+v", record, expected)
	}
}

func TestDomainRecordsServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/records/dev-preview-abc123", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"record":{"id":"dev-preview-abc123","type":"A","name":"www","data":"127.0.0.1","priority":0,"ttl":300}}`
		fmt.Fprint(writer, response)
	})

	record, err := client.DomainRecord.Get(ctx, "vultr.com", "dev-preview-abc123")
	if err != nil {
		t.Errorf("DomainRecord.Get returned %+v, expected %+v", err, nil)
	}

	expected := &DomainRecord{
		ID:       "dev-preview-abc123",
		Type:     "A",
		Name:     "www",
		Data:     "127.0.0.1",
		Priority: 0,
		TTL:      300,
	}

	if !reflect.DeepEqual(record, expected) {
		t.Errorf("DomainRecord.Get returned %+v, expected %+v", record, expected)
	}
}

func TestDomainRecordsServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/records/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})
	r := &DomainRecordReq{
		Name:     "*",
		Type:     "A",
		Data:     "127.0.0.1",
		TTL:      1200,
		Priority: 10,
	}
	err := client.DomainRecord.Update(ctx, "vultr.com", "abc123", r)
	if err != nil {
		t.Errorf("DNSRecord.Update returned %+v, expected %+v", err, nil)
	}
}

func TestDomainRecordsServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/records/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DomainRecord.Delete(ctx, "vultr.com", "abc123")
	if err != nil {
		t.Errorf("DomainRecord.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestDomainRecordsServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/records", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"records":[{"id":"abc123","type":"A","name":"test","data":"127.0.0.1","priority":0,"ttl":300}],"meta":{"total":1,"links":{"next":"thisismycursor","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	records, meta, err := client.DomainRecord.List(ctx, "vultr.com", options)
	if err != nil {
		t.Errorf("DomainRecord.List returned %+v, expected %+v", err, nil)
	}

	expectedRecords := []DomainRecord{
		{
			ID:       "abc123",
			Type:     "A",
			Name:     "test",
			Data:     "127.0.0.1",
			Priority: 0,
			TTL:      300,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycursor",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(records, expectedRecords) {
		t.Errorf("DomainRecord.List returned %+v, expected %+v", records, expectedRecords)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("DomainRecord.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

