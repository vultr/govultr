package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const path = "/v2/baremetal"

// BareMetalServerService is the interface to interact with the bare metal endpoints on the Vultr API
type BareMetalServerService interface {
	Create(ctx context.Context, options *BareMetalServerOptions) (*BareMetalServer, error)
	Get(ctx context.Context, serverID string) (*BareMetalServer, error)
	Update(ctx context.Context, serverID string, bmCreate *BareMetalServerOptions) error
	Delete(ctx context.Context, serverID string) error
	List(ctx context.Context, options *ListOptions) ([]BareMetalServer, *Meta, error)
	Bandwidth(ctx context.Context, serverID string) (map[string]map[string]BareMetalServerBandwidth, error)
	EnableIPV6(ctx context.Context, serverID string) error
	Halt(ctx context.Context, serverID string) error
	IPV4Info(ctx context.Context, serverID string, options *ListOptions) ([]BareMetalServerIPV4, *Meta, error)
	IPV6Info(ctx context.Context, serverID string, options *ListOptions) ([]BareMetalServerIPV6, *Meta, error)
	Reboot(ctx context.Context, serverID string) error
	Reinstall(ctx context.Context, serverID string) error
}

// BareMetalServerServiceHandler handles interaction with the bare metal methods for the Vultr API
type BareMetalServerServiceHandler struct {
	client *Client
}

// BareMetalServer represents a bare metal server on Vultr
type BareMetalServer struct {
	ID              int         `json:"id"`
	Os              string      `json:"os"`
	RAM             string      `json:"ram"`
	Disk            string      `json:"disk"`
	MainIP          string      `json:"main_ip"`
	CPUs            int         `json:"cpu_count"`
	RegionID        string      `json:"region"`
	DefaultPassword string      `json:"default_password"`
	DateCreated     string      `json:"date_created"`
	Status          string      `json:"status"`
	NetmaskV4       string      `json:"netmask_v4"`
	GatewayV4       string      `json:"gateway_v4"`
	Plan            string      `json:"plan"`
	V6Networks      []V6Network `json:"v6_networks"`
	Label           string      `json:"label"`
	Tag             string      `json:"tag"`
	OsID            int         `json:"os_id"`
	AppID           int         `json:"app_id"`
	UserData        string      `json:"user_data"`
}

// BareMetalServerOptions represents the optional parameters that can be set when creating or updating a bare metal server
type BareMetalServerOptions struct {
	Region          string   `json:"region,omitempty"`
	Plan            string   `json:"plan,omitempty"`
	OsID            int      `json:"os_id,omitempty"`
	StartupScriptID string   `json:"script_id,omitempty"`
	SnapshotID      string   `json:"snapshot_id,omitempty"`
	EnableIPV6      string   `json:"enable_ipv6,omitempty"`
	Label           string   `json:"label,omitempty"`
	SSHKeyIDs       []string `json:"sshkey_id,omitempty"`
	AppID           int      `json:"app_id,omitempty"`
	UserData        string   `json:"user_data,omitempty"`
	NotifyActivate  string   `json:"notify_activate,omitempty"`
	Hostname        string   `json:"hostname,omitempty"`
	Tag             string   `json:"tag,omitempty"`
	ReservedIPV4    string   `json:"reserved_ip_v4,omitempty"`
}

// BareMetalServerIPV4 represents IPV4 information for a bare metal server
type BareMetalServerIPV4 struct {
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	Type    string `json:"type"`
}

// BareMetalServerIPV6 represents IPV6 information for a bare metal server
type BareMetalServerIPV6 struct {
	IP          string `json:"ip"`
	Network     string `json:"network"`
	NetworkSize int    `json:"network_size"`
	Type        string `json:"type"`
}

// BareMetalServerBandwidth represents bandwidth information for a bare metal server
type BareMetalServerBandwidth struct {
	IncomingBytes int `json:"incoming_bytes"`
	OutgoingBytes int `json:"outgoing_bytes"`
}

type bareMetalsBase struct {
	BareMetals []BareMetalServer `json:"baremetals"`
	Meta       *Meta             `json:"meta"`
}

type bareMetalBase struct {
	BareMetal *BareMetalServer `json:"baremetal"`
}

type bareMetalIPv4sBase struct {
	BareMetalIPs []BareMetalServerIPV4 `json:"baremetal_ipv4s"`
	Meta         *Meta                 `json:"meta"`
}

type bareMetalIPv6sBase struct {
	BareMetalIPs []BareMetalServerIPV6 `json:"baremetal_ipv6s"`
	Meta         *Meta                 `json:"meta"`
}

type bandwidthBase struct {
	BareMetalBandwidth []BareMetalServerBandwidth `json:"bandwidth"`
}

// Create a new bare metal server.
func (b *BareMetalServerServiceHandler) Create(ctx context.Context, bmCreate *BareMetalServerOptions) (*BareMetalServer, error) {
	req, err := b.client.NewRequest(ctx, http.MethodPost, path, bmCreate)

	if err != nil {
		return nil, err
	}

	bm := new(bareMetalBase)

	if err = b.client.DoWithContext(ctx, req, bm); err != nil {
		return nil, err
	}

	return bm.BareMetal, nil
}

// Get gets the server with the given ID
func (b *BareMetalServerServiceHandler) Get(ctx context.Context, serverID string) (*BareMetalServer, error) {
	uri := fmt.Sprintf("%s/%s", path, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	bms := new(bareMetalBase)
	err = b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, err
	}

	return bms.BareMetal, nil
}

// Update will update the given bare metal. Empty values are ignored
func (b *BareMetalServerServiceHandler) Update(ctx context.Context, serverID string, bmReq *BareMetalServerOptions) error {
	uri := fmt.Sprintf("%s/%s", path, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPatch, uri, bmReq)

	if err != nil {
		return err
	}

	if err = b.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Delete a bare metal server.
// All data will be permanently lost, and the IP address will be released. There is no going back from this call.
func (b *BareMetalServerServiceHandler) Delete(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s", path, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// List lists all bare metal servers on the current account. This includes both pending and active servers.
func (b *BareMetalServerServiceHandler) List(ctx context.Context, options *ListOptions) ([]BareMetalServer, *Meta, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, path, nil)

	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	bms := new(bareMetalsBase)
	err = b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, nil, err
	}

	return bms.BareMetals, bms.Meta, nil
}

// Bandwidth will get the bandwidth used by a bare metal server
func (b *BareMetalServerServiceHandler) Bandwidth(ctx context.Context, serverID string) (map[string]map[string]BareMetalServerBandwidth, error) {
	uri := fmt.Sprintf("%s/%s/bandwidth", path, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	bms := make(map[string]map[string]BareMetalServerBandwidth)

	err = b.client.DoWithContext(ctx, req, &bms)

	if err != nil {
		return nil, err
	}

	return bms, nil
}

// EnableIPV6 enables IPv6 networking on a bare metal server by assigning an IPv6 subnet to it.
// The server will not be rebooted when the subnet is assigned.
func (b *BareMetalServerServiceHandler) EnableIPV6(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/enable-ipv6", path, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// Halt a bare metal server.
// This is a hard power off, meaning that the power to the machine is severed.
// The data on the machine will not be modified, and you will still be billed for the machine.
func (b *BareMetalServerServiceHandler) Halt(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/halt", path, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// IPV4Info will List the IPv4 information of a bare metal server.
// IP information is only available for bare metal servers in the "active" state.
func (b *BareMetalServerServiceHandler) IPV4Info(ctx context.Context, serverID string, options *ListOptions) ([]BareMetalServerIPV4, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv4", path, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	ipv4 := new(bareMetalIPv4sBase)
	err = b.client.DoWithContext(ctx, req, ipv4)
	if err != nil {
		return nil, nil, err
	}

	return ipv4.BareMetalIPs, ipv4.Meta, nil
}

// IPV6Info ists the IPv6 information of a bare metal server.
// IP information is only available for bare metal servers in the "active" state.
// If the bare metal server does not have IPv6 enabled, then an empty array is returned.
func (b *BareMetalServerServiceHandler) IPV6Info(ctx context.Context, serverID string, options *ListOptions) ([]BareMetalServerIPV6, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv6", path, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	ipv6 := new(bareMetalIPv6sBase)
	err = b.client.DoWithContext(ctx, req, ipv6)
	if err != nil {
		return nil, nil, err
	}

	// fmt.Print(req)

	return ipv6.BareMetalIPs, ipv6.Meta, nil
}

// Reboot a bare metal server. This is a hard reboot, which means that the server is powered off, then back on.
func (b *BareMetalServerServiceHandler) Reboot(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/reboot", path, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// Reinstall the operating system on a bare metal server.
// All data will be permanently lost, but the IP address will remain the same. There is no going back from this call.
func (b *BareMetalServerServiceHandler) Reinstall(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/reinstall", path, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}
