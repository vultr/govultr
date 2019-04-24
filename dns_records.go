package govultr

import (
	"context"
	"net/http"
)

// DNSRecordService is the interface to interact with the DNS Records endpoints on the Vultr API
// Link: https://www.vultr.com/api/#dns
type DNSRecordService interface {
	//Create(ctx context.Context, domain string, dnsRecord *DNSRecord) (string, error)
	//Delete(ctx context.Context, domain, recordID string) error
	GetList(ctx context.Context, domain string) ([]DNSRecord, error)
	//Update (ctx context.Context, domain string, dnsRecord *DNSRecord) error
}

// DNSRecordsServiceHandler handles interaction with the DNS Records methods for the Vultr API
type DNSRecordsServiceHandler struct {
	client *Client
}

// DNSRecord represents a DNS record on Vultr
type DNSRecord struct {
	RecordID int    `json:"RECORDID"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

// GetList will list all the records associated with a particular domain on Vultr
func (d *DNSRecordsServiceHandler) GetList(ctx context.Context, domain string) ([]DNSRecord, error) {

	uri := "/v1/dns/records"

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("domain", domain)
	req.URL.RawQuery = q.Encode()

	var dnsRecord []DNSRecord
	err = d.client.DoWithContext(ctx, req, &dnsRecord)

	if err != nil {
		return nil, err
	}

	return dnsRecord, nil
}
