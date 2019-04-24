package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDNSDomainServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/create_domain", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSDomain.Create(ctx, "domain.com", "123456")

	if err != nil {
		t.Errorf("DNSDomain.Create returned %+v, expected %+v", err, nil)
	}
}

func TestDNSDomainServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/delete_domain", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSDomain.Delete(ctx, "domain.com")

	if err != nil {
		t.Errorf("DNSDomain.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestDNSDomainServiceHandler_ToggleDNSSec(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/dns/dnssec_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSDomain.ToggleDNSSec(ctx, "domain.com", true)

	if err != nil {
		t.Errorf("DNSDomain.ToggleDNSSec returned %+v, expected %+v", err, nil)
	}
}

func TestDNSDomainServiceHandler_DNSSecInfo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/dns/dnssec_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		[
    	"example.com IN DNSKEY 257 3 13 kRrxANp7YTGqVbaWtMy8hhsK0jcG4ajjICZKMb4fKv79Vx/RSn76vNjzIT7/Uo0BXil01Fk8RRQc4nWZctGJBA==",
    	"example.com IN DS 27933 13 1 2d9ac457e5c11a104e25d971d0a6254562bddde7",
    	"example.com IN DS 27933 13 2 8858e7b0dfb881280ce2ca1e0eafcd93d5b53687c21da284d4f8799ba82208a9"
		]
		`

		fmt.Fprint(writer, response)
	})

	dnsSec, err := client.DNSDomain.DNSSecInfo(ctx, "example.com")

	if err != nil {
		t.Errorf("DNSDomain.DNSSecInfo returned %+v, expected %+v", err, nil)
	}

	expected := []string{
		"example.com IN DNSKEY 257 3 13 kRrxANp7YTGqVbaWtMy8hhsK0jcG4ajjICZKMb4fKv79Vx/RSn76vNjzIT7/Uo0BXil01Fk8RRQc4nWZctGJBA==",
		"example.com IN DS 27933 13 1 2d9ac457e5c11a104e25d971d0a6254562bddde7",
		"example.com IN DS 27933 13 2 8858e7b0dfb881280ce2ca1e0eafcd93d5b53687c21da284d4f8799ba82208a9",
	}

	if !reflect.DeepEqual(dnsSec, expected) {
		t.Errorf("DNSDomain.ToggleDNSSec returned %+v, expected %+v", dnsSec, expected)
	}
}

func TestDNSDomainServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `[{"domain":"example.com","date_created":"2014-12-11 16:20:59"},{"domain":"example.com","date_created":"2014-12-11 16:20:59"}]`
		fmt.Fprint(writer, response)
	})

	domainList, err := client.DNSDomain.GetList(ctx)

	if err != nil {
		t.Errorf("DNSDomain.GetList returned %+v, expected %+v", err, nil)
	}

	expected := []DNSDomain{
		{Domain: "example.com", DateCreated: "2014-12-11 16:20:59"},
		{Domain: "example.com", DateCreated: "2014-12-11 16:20:59"},
	}

	if !reflect.DeepEqual(domainList, expected) {
		t.Errorf("DNSDomain.GetList returned %+v, expected %+v", domainList, expected)
	}
}

func TestDNSDomainServiceHandler_GetSoa(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/soa_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    	"nsprimary": "ns1.vultr.com",
    	"email": "dnsadm@vultr.com"
		}`
		fmt.Fprint(writer, response)
	})

	soa, err := client.DNSDomain.GetSoa(ctx, "example.com")

	if err != nil {
		t.Errorf("DNSDomain.GetSoa returned %+v, expected %+v", err, nil)
	}

	expected := &Soa{NsPrimary: "ns1.vultr.com", Email: "dnsadm@vultr.com"}

	if !reflect.DeepEqual(soa, expected) {
		t.Errorf("DNSDomain.GetSoa returned %+v, expected %+v", soa, expected)
	}
}

func TestDNSDomainServiceHandler_UpdateSoa(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/soa_update", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.DNSDomain.UpdateSoa(ctx, "domain.com", "ns1.domain.com", "example@vultr.com")

	if err != nil {
		t.Errorf("DNSDomain.UpdateSoa returned %+v, expected %+v", err, nil)
	}
}
