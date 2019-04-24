package govultr

// DNSDomainService is the interface to interact with the DNS endpoints on the Vultr API
// Link: https://www.vultr.com/api/#dns
type DNSDomainService interface {
	//Create(ctx context.Context, domain string, vpsIP int) (string, error)
	//Delete(ctx context.Context, domain string) error
	//EnableDnssec(ctx context.Context, domain, enabled string) error
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
