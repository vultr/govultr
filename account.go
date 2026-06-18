package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// AccountService is the interface to interact with Accounts endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/account
type AccountService interface {
	Get(ctx context.Context) (*Account, *http.Response, error)

	GetBandwidth(ctx context.Context) (*AccountBandwidth, *http.Response, error)

	SetupBGP(ctx context.Context, setup *AccountBGPSetup) error
	GetBGP(ctx context.Context) (*AccountBGP, *http.Response, error)

	AddBGPPrefixes(ctx context.Context, prefixes []string) error
	ListBGPPrefixes(ctx context.Context, options *ListOptions) ([]AccountBGPPrefix, *Meta, *http.Response, error)

	ListCustomSubscriptions(ctx context.Context, options *ListOptions) ([]AccountCustomSubscription, *Meta, *http.Response, error)
}

// AccountServiceHandler handles interaction with the account methods for the Vultr API
type AccountServiceHandler struct {
	client *Client
}

type accountBase struct {
	Account *Account `json:"account"`
}

// Account represents a Vultr account
type Account struct {
	Balance           float32  `json:"balance"`
	PendingCharges    float32  `json:"pending_charges"`
	LastPaymentDate   string   `json:"last_payment_date"`
	LastPaymentAmount float32  `json:"last_payment_amount"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	ACL               []string `json:"acls"`
}

type accountBandwidthBase struct {
	Bandwidth *AccountBandwidth `json:"bandwidth"`
}

// Bandwidth represents a Vultr account bandwidth
type AccountBandwidth struct {
	PreviousMonth         AccountBandwidthPeriod `json:"previous_month"`
	CurrentMonthToDate    AccountBandwidthPeriod `json:"current_month_to_date"`
	CurrentMonthProjected AccountBandwidthPeriod `json:"current_month_projected"`
}

// AccountBandwidthPeriod represents a Vultr account bandwidth period
type AccountBandwidthPeriod struct {
	TimestampStart            string  `json:"timestamp_start"`
	TimestampEnd              string  `json:"timestamp_end"`
	GBIn                      int     `json:"gb_in"`
	GBOut                     int     `json:"gb_out"`
	TotalInstanceHours        int     `json:"total_instance_hours"`
	TotalInstanceCount        int     `json:"total_instance_count"`
	InstanceBandwidthCredits  int     `json:"instance_bandwidth_credits"`
	FreeBandwidthCredits      int     `json:"free_bandwidth_credits"`
	PurchasedBandwidthCredits int     `json:"purchased_bandwidth_credits"`
	Overage                   float32 `json:"overage"`
	OverageUnitCost           float32 `json:"overage_unit_cost"`
	OverageCost               float32 `json:"overage_cost"`
}

// AccountBGP represents a Vultr account BGP information
type AccountBGP struct {
	Enabled  bool   `json:"enabled"`
	ASN      int    `json:"asn"`
	Password string `json:"password"`
}

// AccountBGPSetup struct is used for setting up BGP in a Vultr account
type AccountBGPSetup struct {
	Prefixes              []string `json:"prefixes"`
	ASN                   int      `json:"asn,omitempty"`
	Password              string   `json:"password,omitempty"`
	LetterOfAuthorization string   `json:"letter_of_authorization,omitempty"`
	RequestedRoutes       string   `json:"requested_routes"`
	UseCase               string   `json:"use_case"`
}

// AccountBGPPrefix represents a BGP prefix configured in a Vultr account
type AccountBGPPrefix struct {
	Prefix     string `json:"prefix"`
	DateAdded  string `json:"date_added"`
	RPKIStatus string `json:"rpki_status"`
}

type accountBGPPrefixBase struct {
	Prefixes []AccountBGPPrefix `json:"prefixes"`
	Meta     *Meta              `json:"meta"`
}

// AccountBGPPrefixAdd struct is used for adding BGP prefixes to a Vultr account
type AccountBGPPrefixAdd struct {
	Prefixes []string `json:"prefixes"`
}

// AccountCustomSubscription represents a custom subscription in a Vultr account
type AccountCustomSubscription struct {
	ID             string  `json:"id"`
	Label          string  `json:"label"`
	Description    string  `json:"description"`
	Type           string  `json:"type"`
	Region         string  `json:"region"`
	Status         string  `json:"status"`
	DateCreated    string  `json:"date_created"`
	Cost           float32 `json:"cost"`
	PendingCharges float32 `json:"pending_charges"`
}

type accountCustomSubscriptionBase struct {
	CustomSubscriptions []AccountCustomSubscription `json:"custom_subscriptions"`
	Meta                *Meta                       `json:"meta"`
}

// Get Vultr account info
func (a *AccountServiceHandler) Get(ctx context.Context) (*Account, *http.Response, error) {
	uri := "/v2/account"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(accountBase)
	resp, err := a.client.DoWithContext(ctx, req, account)
	if err != nil {
		return nil, resp, err
	}

	return account.Account, resp, nil
}

// Get Vultr account bandwidth info
func (a *AccountServiceHandler) GetBandwidth(ctx context.Context) (*AccountBandwidth, *http.Response, error) {
	uri := "/v2/account/bandwidth"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	accountBandwidth := new(accountBandwidthBase)
	resp, err := a.client.DoWithContext(ctx, req, accountBandwidth)
	if err != nil {
		return nil, resp, err
	}

	return accountBandwidth.Bandwidth, resp, nil
}

// Get Vultr account BGP info
func (a *AccountServiceHandler) GetBGP(ctx context.Context) (*AccountBGP, *http.Response, error) {
	uri := "/v2/account/bgp"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	accountBGP := new(AccountBGP)
	resp, err := a.client.DoWithContext(ctx, req, accountBGP)
	if err != nil {
		return nil, resp, err
	}

	return accountBGP, resp, nil
}

// Setup BGP in a Vultr account
func (a *AccountServiceHandler) SetupBGP(ctx context.Context, setup *AccountBGPSetup) error {
	uri := "/v2/account/bgp/setup"

	req, err := a.client.NewRequest(ctx, http.MethodPost, uri, setup)
	if err != nil {
		return err
	}
	_, err = a.client.DoWithContext(ctx, req, nil)
	return err
}

// Add BGP prefixes to a Vultr account
func (a *AccountServiceHandler) AddBGPPrefixes(ctx context.Context, prefixes []string) error {
	uri := "/v2/account/bgp/prefixes"

	prefixAdd := &AccountBGPPrefixAdd{Prefixes: prefixes}
	req, err := a.client.NewRequest(ctx, http.MethodPost, uri, prefixAdd)
	if err != nil {
		return err
	}
	_, err = a.client.DoWithContext(ctx, req, nil)
	return err
}

// List BGP prefixes in a Vultr account
func (a *AccountServiceHandler) ListBGPPrefixes(ctx context.Context, options *ListOptions) ([]AccountBGPPrefix, *Meta, *http.Response, error) {
	uri := "/v2/account/bgp/prefixes"

	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	prefixes := new(accountBGPPrefixBase)
	resp, err := a.client.DoWithContext(ctx, req, prefixes)
	if err != nil {
		return nil, nil, resp, err
	}

	return prefixes.Prefixes, prefixes.Meta, resp, nil
}

func (a *AccountServiceHandler) ListCustomSubscriptions(ctx context.Context, options *ListOptions) ([]AccountCustomSubscription, *Meta, *http.Response, error) {
	uri := "/v2/account/custom-subscriptions"

	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	subscriptions := new(accountCustomSubscriptionBase)
	resp, err := a.client.DoWithContext(ctx, req, subscriptions)
	if err != nil {
		return nil, nil, resp, err
	}

	return subscriptions.CustomSubscriptions, subscriptions.Meta, resp, nil
}
