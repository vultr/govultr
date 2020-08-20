package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const instancePath = "/v2/instances"

// InstanceService is the interface to interact with the instance endpoints on the Vultr API
// Link: https://www.vultr.com/api/v2/#tag/instances
type InstanceService interface {
	//Create(ctx context.Context, regionID, vpsPlanID, osID int, options *ServerOptions) (*Server, error)
	Get(ctx context.Context, instanceID string) (*Instance, error)
	//Update
	Delete(ctx context.Context, instanceID string) error
	List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, error)

	Start(ctx context.Context, instanceID string) error
	Halt(ctx context.Context, instanceID string) error
	Reboot(ctx context.Context, instanceID string) error
	Reinstall(ctx context.Context, instanceID string) error

	Bandwidth(ctx context.Context, instanceID string) (*Bandwidth, error)

	ListPrivateNetworks(ctx context.Context, instanceID string) ([]PrivateNetwork, *Meta, error)
	AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error
	DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error

	ISOStatus(ctx context.Context, instanceID string) (*Iso, error)
	AttachISO(ctx context.Context, instanceID, isoID string) error
	DetachISO(ctx context.Context, instanceID string) error

	GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, error)
	SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) (*BackupSchedule, error)

	ListIPV4(ctx context.Context, instanceID string, option *ListOptions) ([]IPv4, *Meta, error)
	ListIPV6(ctx context.Context, instanceID string, option *ListOptions) ([]IPv6, *Meta, error)

	// todo reverse ipv6
	CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error
	ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, error)
	// Delete
	// default ipv6

	CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error

	GetUserData(ctx context.Context, instanceID string) (*UserData, error)
}

// ServerServiceHandler handles interaction with the server methods for the Vultr API
type InstanceServiceHandler struct {
	client *Client
}

// Instance represents a VPS
type Instance struct {
	ID               string   `json:"id"`
	Os               string   `json:"os"`
	Ram              int      `json:"ram"`
	Disk             int      `json:"disk"`
	MainIP           string   `json:"main_ip"`
	VCPUCount        int      `json:"vcpu_count"`
	Region           string   `json:"region"`
	DefaultPassword  string   `json:"default_password,omitempty"`
	DateCreated      string   `json:"date_created"`
	Status           string   `json:"status"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6NetworkSize    string   `json:"v6_network_size"`
	Label            string   `json:"label"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	Tag              string   `json:"tag"`
	OsID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	Features         []string `json:"features"`
}

type instanceBase struct {
	Instance *Instance `json:"instance"`
}

type instancesBase struct {
	Instances []Instance `json:"instances"`
	Meta      *Meta      `json:"meta"`
}

type Bandwidth struct {
	Bandwidth map[string]struct {
		IncomingBytes int `json:"incoming_bytes"`
		OutgoingBytes int `json:"outgoing_bytes"`
	} `json:"bandwidth"`
}

type privateNetworksBase struct {
	PrivateNetworks []PrivateNetwork `json:"private_networks"`
	Meta            *Meta            `json:"meta"`
}

type PrivateNetwork struct {
	NetworkID  string `json:"network_id"`
	MacAddress string `json:"mac_address"`
	IPAddress  string `json:"ip_address"`
}

type isoStatusBase struct {
	IsoStatus *Iso `json:"iso_status"`
}

type Iso struct {
	State string `json:"state"`
	IsoID string `json:"iso_id"`
}

type backupScheduleBase struct {
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
}
type BackupSchedule struct {
	Enabled             bool   `json:"enabled,omitempty"`
	Type                string `json:"type,omitempty"`
	NextScheduleTimeUTC string `json:"next_schedule_time_utc,omitempty"`
	Hour                int    `json:"hour,omitempty"`
	Dow                 int    `json:"dow,omitempty"`
	Dom                 int    `json:"dom,omitempty"`
}

type BackupScheduleReq struct {
	Type string `json:"type"`
	Hour int    `json:"hour,omitempty"`
	Dow  int    `json:"dow,omitempty"`
	Dom  int    `json:"dom,omitempty"`
}

// todo can we remove this list and return this data back in the list?
type reverseIPv6sBase struct {
	ReverseIPv6s []ReverseIP `json:"reverse_ipv6s"`
	// no meta?
}

type ReverseIP struct {
	IP      string `json:"ip"`
	Reverse string `json:"reverse"`
}

type userDataBase struct {
	UserData *UserData `json:"user_data"`
}

type UserData struct {
	Data string `json:"data"`
}

// GetServer will get the server with the given instanceID
func (i *InstanceServiceHandler) Get(ctx context.Context, instanceID string) (*Instance, error) {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	instance := new(instanceBase)
	if err = i.client.DoWithContext(ctx, req, instance); err != nil {
		return nil, err

	}

	return instance.Instance, nil
}

// Delete an instance. All data will be permanently lost, and the IP address will be released
func (i *InstanceServiceHandler) Delete(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// List all instances on your account.
func (i *InstanceServiceHandler) List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, instancePath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	instances := new(instancesBase)
	if err = i.client.DoWithContext(ctx, req, instances); err != nil {
		return nil, nil, err
	}

	return instances.Instances, instances.Meta, nil
}

// Start will start a vps instance the machine is already running, it will be restarted.
func (i *InstanceServiceHandler) Start(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/start", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Halt will pause an instance.
func (i *InstanceServiceHandler) Halt(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/halt", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Reboot an instance.
func (i *InstanceServiceHandler) Reboot(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/reboot", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Reinstall will reinstall the operating system on a instance.
func (i *InstanceServiceHandler) Reinstall(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/reinstall", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (i *InstanceServiceHandler) Bandwidth(ctx context.Context, instanceID string) (*Bandwidth, error) {
	uri := fmt.Sprintf("%s/%s/bandwidth", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	bandwidth := new(Bandwidth)
	if err = i.client.DoWithContext(ctx, req, bandwidth); err != nil {
		return nil, err
	}

	return bandwidth, nil
}

func (i *InstanceServiceHandler) ListPrivateNetworks(ctx context.Context, instanceID string) ([]PrivateNetwork, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/private-networks", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	networks := new(privateNetworksBase)
	if err = i.client.DoWithContext(ctx, req, networks); err != nil {
		return nil, nil, err
	}

	return networks.PrivateNetworks, networks.Meta, nil
}

// AttachPrivateNetwork to an instance
func (i *InstanceServiceHandler) AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-networks/attach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DetachedPrivateNetwork from an instance.
func (i *InstanceServiceHandler) DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-network/detach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

//// IsoStatus retrieves the current ISO state for a given VPS.
//// The returned state may be one of: ready | isomounting | isomounted.
func (i *InstanceServiceHandler) ISOStatus(ctx context.Context, instanceID string) (*Iso, error) {
	uri := fmt.Sprintf("%s/%s/iso", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	iso := new(isoStatusBase)
	if err = i.client.DoWithContext(ctx, req, iso); err != nil {
		return nil, err
	}
	return iso.IsoStatus, nil
}

// AttachISO will attach an ISO to the given instance and reboot it
func (i *InstanceServiceHandler) AttachISO(ctx context.Context, instanceID, isoID string) error {
	uri := fmt.Sprintf("%s/%s/iso/attach", instancePath, instanceID)
	body := RequestBody{"iso_id": isoID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

//// DetachISO will detach the currently mounted ISO and reboot the instance.
func (i *InstanceServiceHandler) DetachISO(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/iso/detach", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// GetBackupSchedule retrieves the backup schedule for a given instance - all time values are in UTC
func (i *InstanceServiceHandler) GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, error) {
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	backup := new(backupScheduleBase)
	if err = i.client.DoWithContext(ctx, req, backup); err != nil {
		return nil, err
	}

	return backup.BackupSchedule, nil
}

// SetBackupSchedule sets the backup schedule for a given instance - all time values are in UTC.
func (i *InstanceServiceHandler) SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) (*BackupSchedule, error) {
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, backup)
	if err != nil {
		return nil, err
	}

	b := new(backupScheduleBase)
	if err = i.client.DoWithContext(ctx, req, backup); err != nil {
		return nil, err
	}

	return b.BackupSchedule, nil
}

func (i *InstanceServiceHandler) ListIPV4(ctx context.Context, instanceID string, options *ListOptions) ([]IPv4, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv4", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	if err = i.client.DoWithContext(ctx, req, ips); err != nil {
		return nil, nil, err
	}

	return ips.IPv4S, ips.Meta, nil
}

func (i *InstanceServiceHandler) ListIPV6(ctx context.Context, instanceID string, options *ListOptions) ([]IPv6, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv6", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	if err = i.client.DoWithContext(ctx, req, ips); err != nil {
		return nil, nil, err
	}

	return ips.IPv6S, ips.Meta, nil
}

func (i *InstanceServiceHandler) CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (i *InstanceServiceHandler) ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, error) {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	reverse := new(reverseIPv6sBase)
	if err = i.client.DoWithContext(ctx, req, reverse); err != nil {
		return nil, err
	}

	return reverse.ReverseIPv6s, nil
}

func (i *InstanceServiceHandler) CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv4/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}

	if err = i.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (i *InstanceServiceHandler) GetUserData(ctx context.Context, instanceID string) (*UserData, error) {
	uri := fmt.Sprintf("%s/%s/user-data", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	userData := new(userDataBase)
	if err = i.client.DoWithContext(ctx, req, userData); err != nil {
		return nil, err
	}

	return userData.UserData, nil
}
