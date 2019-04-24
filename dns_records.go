package govultr

// DNSRecordService is the interface to interact with the DNS Records endpoints on the Vultr API
// Link: https://www.vultr.com/api/#dns
type DNSRecordsService interface {
	//Create(ctx context.Context, domain string, dnsRecord *DNSRecord) (string, error)
	//Delete(ctx context.Context, domain, recordID string) error
	//GetList(ctx context.Context, domain string) ([]DNSRecord, error)
	//Update (ctx context.Context, domain string, dnsRecord *DNSRecord) error
}

// DNSRecordsServiceHandler handles interaction with the DNS Records methods for the Vultr API
type DNSRecordsServiceHandler struct {
	client *Client
}

// DNSRecord represents a DNS record on Vultr
type DNSRecord struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	RecordID int    `json:"RECORDID"`
	TTL      int    `json:"ttl"`
}
