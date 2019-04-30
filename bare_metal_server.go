package govultr

import (
	"context"
	"net/http"
	"net/url"
)

// BareMetalServerService is the interface to interact with the bare metal endpoints on the Vultr API
// Link: https://www.vultr.com/api/#baremetal
type BareMetalServerService interface {
	Create(ctx context.Context, regionID, planID, osID string, options *BareMetalServerOptions) (*BareMetalServer, error)
	Destroy(ctx context.Context, serverID string) error
	GetList(ctx context.Context, serverID, tag, label, mainIP string) ([]BareMetalServer, error)
	Halt(ctx context.Context, serverID string) error
	Reboot(ctx context.Context, serverID string) error
	Reinstall(ctx context.Context, serverID string) error
}

// BareMetalServerServiceHandler handles interaction with the bare metal methods for the Vultr API
type BareMetalServerServiceHandler struct {
	client *Client
}

// BareMetalServer represents a bare metal server on Vultr
type BareMetalServer struct {
	BareMetalServerID string      `json:"SUBID"`
	Os                string      `json:"os"`
	RAM               string      `json:"ram"`
	Disk              string      `json:"disk"`
	MainIP            string      `json:"main_ip"`
	CPUCount          int         `json:"cpu_count"`
	Location          string      `json:"location"`
	RegionID          string      `json:"DCID"`
	DefaultPassword   string      `json:"default_password"`
	DateCreated       string      `json:"date_created"`
	Status            string      `json:"status"`
	NetmaskV4         string      `json:"netmask_v4"`
	GatewayV4         string      `json:"gateway_v4"`
	BareMetalPlanID   int         `json:"METALPLANID"`
	V6Networks        []V6Network `json:"v6_networks"`
	Label             string      `json:"label"`
	Tag               string      `json:"tag"`
	OsID              string      `json:"OSID"`
	AppID             string      `json:"APPID"`
}

// V6Network represents a IPv6 network on Vultr
type V6Network struct {
	Network     string `json:"v6_network"`
	MainIP      string `json:"v6_main_ip"`
	NetworkSize int    `json:"v6_network_size"`
}

// BareMetalServerOptions represents the optional parameters that can be set when creating a bare metal server
type BareMetalServerOptions struct {
	StartupScriptID string
	SnapshotID      string
	EnableIPV6      string
	Label           string
	SSHKeyID        string
	AppID           string
	UserData        string
	NotifyActivate  string
	Hostname        string
	Tag             string
	ReservedIPV4    string
}

// Create a new bare metal server.
func (b *BareMetalServerServiceHandler) Create(ctx context.Context, regionID, planID, osID string, options *BareMetalServerOptions) (*BareMetalServer, error) {
	uri := "/v1/baremetal/create"

	values := url.Values{
		"DCID":        {regionID},
		"METALPLANID": {planID},
		"OSID":        {osID},
	}

	if options != nil {
		if options.StartupScriptID != "" {
			values.Add("SCRIPTID", options.StartupScriptID)
		}
		if options.SnapshotID != "" {
			values.Add("SNAPSHOTID", options.SnapshotID)
		}
		if options.EnableIPV6 != "" {
			values.Add("enable_ipv6", options.EnableIPV6)
		}
		if options.Label != "" {
			values.Add("label", options.Label)
		}
		if options.SSHKeyID != "" {
			values.Add("SSHKEYID", options.SSHKeyID)
		}
		if options.AppID != "" {
			values.Add("APPID", options.AppID)
		}
		if options.UserData != "" {
			values.Add("userdata", options.UserData)
		}
		if options.NotifyActivate != "" {
			values.Add("notify_activate", options.NotifyActivate)
		}
		if options.Hostname != "" {
			values.Add("hostname", options.Hostname)
		}
		if options.Tag != "" {
			values.Add("tag", options.Tag)
		}
		if options.ReservedIPV4 != "" {
			values.Add("reserved_ip_v4", options.ReservedIPV4)
		}
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return nil, err
	}

	bm := new(BareMetalServer)

	err = b.client.DoWithContext(ctx, req, bm)

	if err != nil {
		return nil, err
	}

	return bm, nil
}

// Destroy (delete) a bare metal server.
// All data will be permanently lost, and the IP address will be released. There is no going back from this call.
func (b *BareMetalServerServiceHandler) Destroy(ctx context.Context, serverID string) error {
	uri := "/v1/baremetal/destroy"

	values := url.Values{
		"SUBID": {serverID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetList lists all bare metal servers on the current account. This includes both pending and active servers.
// If you need to filter the list, review the parameters for this function.
// Currently, only one filter at a time may be applied (serverID, tag, label, mainIP).
func (b *BareMetalServerServiceHandler) GetList(ctx context.Context, serverID, tag, label, mainIP string) ([]BareMetalServer, error) {
	uri := "/v1/baremetal/list"

	values := url.Values{}

	if serverID != "" {
		values.Add("SUBID", serverID)
	}
	if tag != "" {
		values.Add("tag", tag)
	}
	if label != "" {
		values.Add("label", label)
	}
	if mainIP != "" {
		values.Add("main_ip", mainIP)
	}

	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, values)

	if err != nil {
		return nil, err
	}

	bmsMap := make(map[string]BareMetalServer)
	err = b.client.DoWithContext(ctx, req, &bmsMap)
	if err != nil {
		return nil, err
	}

	var bms []BareMetalServer
	for _, bm := range bmsMap {
		bms = append(bms, bm)
	}

	return bms, nil
}

// Halt a bare metal server.
// This is a hard power off, meaning that the power to the machine is severed.
// The data on the machine will not be modified, and you will still be billed for the machine.
func (b *BareMetalServerServiceHandler) Halt(ctx context.Context, serverID string) error {
	uri := "/v1/baremetal/halt"

	values := url.Values{
		"SUBID": {serverID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// Reboot a bare metal server. This is a hard reboot, which means that the server is powered off, then back on.
func (b *BareMetalServerServiceHandler) Reboot(ctx context.Context, serverID string) error {
	uri := "/v1/baremetal/reboot"

	values := url.Values{
		"SUBID": {serverID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

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
	uri := "/v1/baremetal/reinstall"

	values := url.Values{
		"SUBID": {serverID},
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = b.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}
