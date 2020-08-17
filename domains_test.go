package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDomainServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"domain": {"domain": "vultr.com","date_created": "2020-05-08 19:09:07"}}`
		fmt.Fprint(writer, response)
	})

	req := &DomainReq{
		Domain: "vultr.com",
	}
	domain, err := client.Domain.Create(ctx, req)
	if err != nil {
		t.Errorf("DNSDomain.Create returned %+v, expected %+v", err, nil)
	}

	expected := &Domain{
		Domain:      "vultr.com",
		DateCreated: "2020-05-08 19:09:07",
	}

	if !reflect.DeepEqual(domain, expected) {
		t.Errorf("Domain.Create returned %+v, expected %+v", domain, expected)
	}
}

func TestDomainServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"domain": {"domain": "vultr.com","date_created": "2020-05-08 19:09:07"}}`
		fmt.Fprint(writer, response)
	})

	domain, err := client.Domain.Get(ctx, "vultr.com")
	if err != nil {
		t.Errorf("DNSDomain.Create returned %+v, expected %+v", err, nil)
	}

	expected := &Domain{
		Domain:      "vultr.com",
		DateCreated: "2020-05-08 19:09:07",
	}

	if !reflect.DeepEqual(domain, expected) {
		t.Errorf("Domain.Create returned %+v, expected %+v", domain, expected)
	}
}

func TestDomainServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/domains/vultr.com", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Domain.Update(ctx, "vultr.com", "enabled")
	if err != nil {
		t.Errorf("Domain.Update returned %+v, expected %+v", err, nil)
	}
}

func TestDomainServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/domain.com", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Domain.Delete(ctx, "domain.com")
	if err != nil {
		t.Errorf("Domain.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestDNSDomainServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"domains":[{"domain":"vultr.com","date_created":"2020-05-0819:09:07"}],"meta":{"total":1,"links":{"next":"thisismycusror","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	domains, meta, err := client.Domain.List(ctx, options)
	if err != nil {
		t.Errorf("Domain.List returned %+v, expected %+v", err, nil)
	}

	expectedDomain := []Domain{
		{
			Domain:      "vultr.com",
			DateCreated: "2020-05-0819:09:07",
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(domains, expectedDomain) {
		t.Errorf("Domain.List returned %+v, expected %+v", domains, expectedDomain)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Domain.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestDomainServiceHandler_GetSoa(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/soa", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"dns_soa":{"nsprimary":"ns1.vultr.com","email":"dnsadm@vultr.com"}}`
		fmt.Fprint(writer, response)
	})

	soa, err := client.Domain.GetSoa(ctx, "vultr.com")
	if err != nil {
		t.Errorf("Domain.GetSoa returned %+v, expected %+v", err, nil)
	}

	expected := &Soa{NSPrimary: "ns1.vultr.com", Email: "dnsadm@vultr.com"}

	if !reflect.DeepEqual(soa, expected) {
		t.Errorf("DNSDomain.GetSoa returned %+v, expected %+v", soa, expected)
	}
}

func TestDNSDomainServiceHandler_UpdateSoa(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/domains/vultr.com/soa", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	r := &Soa{
		NSPrimary: "ns4.vultr.com",
		Email:     "vultr@vultr.com",
	}
	err := client.Domain.UpdateSoa(ctx, "vultr.com", r)

	if err != nil {
		t.Errorf("Domain.UpdateSoa returned %+v, expected %+v", err, nil)
	}
}

func TestDNSDomainServiceHandler_DNSSecInfo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/domains/vultr.com/dnssec", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"dns_sec":[
		"example.com IN DNSKEY 257 3 13 kRrxANp7YTGqVbaWtMy8hhsK0jcG4ajjICZKMb4fKv79Vx/RSn76vNjzIT7/Uo0BXil01Fk8RRQc4nWZctGJBA==",
		"example.com IN DS 27933 13 1 2d9ac457e5c11a104e25d971d0a6254562bddde7",
		"example.com IN DS 27933 13 2 8858e7b0dfb881280ce2ca1e0eafcd93d5b53687c21da284d4f8799ba82208a9"
]}`
		fmt.Fprint(writer, response)
	})

	dnsSec, err := client.Domain.GetDnsSec(ctx, "vultr.com")
	if err != nil {
		t.Errorf("Domain.GetDnsSec returned %+v, expected %+v", err, nil)
	}

	expected := []string{
		"example.com IN DNSKEY 257 3 13 kRrxANp7YTGqVbaWtMy8hhsK0jcG4ajjICZKMb4fKv79Vx/RSn76vNjzIT7/Uo0BXil01Fk8RRQc4nWZctGJBA==",
		"example.com IN DS 27933 13 1 2d9ac457e5c11a104e25d971d0a6254562bddde7",
		"example.com IN DS 27933 13 2 8858e7b0dfb881280ce2ca1e0eafcd93d5b53687c21da284d4f8799ba82208a9",

	}

	if !reflect.DeepEqual(dnsSec, expected) {
		t.Errorf("Domain.GetDnsSec returned %+v, expected %+v", dnsSec, expected)
	}
}

