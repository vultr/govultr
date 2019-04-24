package govultr

import (
	"context"
	"net/http"
	"net/url"
)

// DNSDomainService is the interface to interact with the DNS endpoints on the Vultr API
// Link: https://www.vultr.com/api/#dns
type DNSDomainService interface {
	Create(ctx context.Context, domain, vpsIP string) error
	Delete(ctx context.Context, domain string) error
	ToggleDNSSec(ctx context.Context, domain string, enabled bool) error
	//DnssecInfo(ctx context.Context, domain string) (string, error)
	//GetList(ctx context.Context) ([]DNSDomain, error)
	//GetSoa(ctx context.Context, domain string) ([]Soa, error)
	//UpdateSoa(ctx context.Context, domain, nsPrimary, email string) error
}

// DNSDomainServiceHandler handles interaction with the DNS methods for the Vultr API
type DNSDomainServiceHandler struct {
	client *Client
}

// DNSDomain represents a DNS Domain entry on Vultr
type DNSDomain struct {
	Domain      string `json:"domain"`
	DateCreated string `json:"date_created"`
}

// Soa represents record information for a domain on Vultr
type Soa struct {
	NsPrimary string `json:"nsprimary"`
	Email     string `json:"email"`
}

// Create will create a DNS Domain entry on Vultr
func (d *DNSDomainServiceHandler) Create(ctx context.Context, domain, vpsIP string) error {

	uri := "/v1/dns/create_domain"

	values := url.Values{
		"domain":   {domain},
		"serverip": {vpsIP},
	}

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = d.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

func (d *DNSDomainServiceHandler) Delete(ctx context.Context, domain string) error {
	uri := "/v1/dns/delete_domain"

	values := url.Values{
		"domain": {domain},
	}

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = d.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// ToggleDNSSec will enable or disable DNSSEC for a domain on Vultr
func (d *DNSDomainServiceHandler) ToggleDNSSec(ctx context.Context, domain string, enabled bool) error {

	uri := "/v1/dns/dnssec_enable"

	enable := "no"
	if enabled == true {
		enable = "yes"
	}

	values := url.Values{
		"domain": {domain},
		"enable": {enable},
	}

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = d.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}
