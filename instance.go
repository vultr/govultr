package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const instancePath = "/v2/instances"

// InstanceService is the interface to interact with the instance endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/instances
type InstanceService interface {
	Create(ctx context.Context, instanceReq *InstanceCreateReq) (*Instance, *http.Response, error)
	Get(ctx context.Context, instanceID string) (*Instance, *http.Response, error)
	Update(ctx context.Context, instanceID string, instanceReq *InstanceUpdateReq) (*Instance, *http.Response, error)
	Delete(ctx context.Context, instanceID string) error
	List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, *http.Response, error)

	Start(ctx context.Context, instanceID string) error
	Halt(ctx context.Context, instanceID string) error
	Reboot(ctx context.Context, instanceID string) error
	Reinstall(ctx context.Context, instanceID string, reinstallReq *ReinstallReq) (*Instance, *http.Response, error)

	MassStart(ctx context.Context, instanceList []string) error
	MassHalt(ctx context.Context, instanceList []string) error
	MassReboot(ctx context.Context, instanceList []string) error

	Restore(ctx context.Context, instanceID string, restoreReq *RestoreReq) (*http.Response, error)

	GetBandwidth(ctx context.Context, instanceID string) (*Bandwidth, *http.Response, error)
	GetNeighbors(ctx context.Context, instanceID string) (*Neighbors, *http.Response, error)

	// Deprecated: ListPrivateNetworks should no longer be used. Instead, use ListVPCInfo.
	ListPrivateNetworks(ctx context.Context, instanceID string, options *ListOptions) ([]PrivateNetwork, *Meta, *http.Response, error)
	// Deprecated: AttachPrivateNetwork should no longer be used. Instead, use AttachVPC.
	AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error
	// Deprecated: DetachPrivateNetwork should no longer be used. Instead, use DetachVPC.
	DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error

	ListVPCInfo(ctx context.Context, instanceID string, options *ListOptions) ([]VPCInfo, *Meta, *http.Response, error)
	AttachVPC(ctx context.Context, instanceID, vpcID string) error
	DetachVPC(ctx context.Context, instanceID, vpcID string) error

	ListVPC2Info(ctx context.Context, instanceID string, options *ListOptions) ([]VPC2Info, *Meta, *http.Response, error)
	AttachVPC2(ctx context.Context, instanceID string, vpc2Req *AttachVPC2Req) error
	DetachVPC2(ctx context.Context, instanceID, vpcID string) error

	ISOStatus(ctx context.Context, instanceID string) (*Iso, *http.Response, error)
	AttachISO(ctx context.Context, instanceID, isoID string) (*http.Response, error)
	DetachISO(ctx context.Context, instanceID string) (*http.Response, error)

	GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, *http.Response, error)
	SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) (*http.Response, error)

	CreateIPv4(ctx context.Context, instanceID string, reboot *bool) (*IPv4, *http.Response, error)
	ListIPv4(ctx context.Context, instanceID string, option *ListOptions) ([]IPv4, *Meta, *http.Response, error)
	DeleteIPv4(ctx context.Context, instanceID, ip string) error
	ListIPv6(ctx context.Context, instanceID string, option *ListOptions) ([]IPv6, *Meta, *http.Response, error)

	CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error
	ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, *http.Response, error)
	DeleteReverseIPv6(ctx context.Context, instanceID, ip string) error

	CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error
	DefaultReverseIPv4(ctx context.Context, instanceID, ip string) error

	GetUserData(ctx context.Context, instanceID string) (*UserData, *http.Response, error)

	GetUpgrades(ctx context.Context, instanceID string) (*Upgrades, *http.Response, error)
}

// InstanceServiceHandler handles interaction with the server methods for the Vultr API
type InstanceServiceHandler struct {
	client *Client
}

// Instance represents a VPS
type Instance struct {
	PowerStatus      string   `json:"power_status"`
	Tag              string   `json:"tag"`
	ServerStatus     string   `json:"server_status"`
	ID               string   `json:"id"`
	Plan             string   `json:"plan"`
	MainIP           string   `json:"main_ip"`
	Hostname         string   `json:"hostname"`
	Region           string   `json:"region"`
	DefaultPassword  string   `json:"default_password,omitempty"`
	DateCreated      string   `json:"date_created"`
	Status           string   `json:"status"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	ImageID          string   `json:"image_id"`
	Os               string   `json:"os"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	KVM              string   `json:"kvm"`
	Label            string   `json:"label"`
	InternalIP       string   `json:"internal_ip"`
	Tags             []string `json:"tags"`
	Features         []string `json:"features"`
	V6NetworkSize    int      `json:"v6_network_size"`
	RAM              int      `json:"ram"`
	OsID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	VCPUCount        int      `json:"vcpu_count"`
	Disk             int      `json:"disk"`
}

type instanceBase struct {
	Instance *Instance `json:"instance"`
}

type ipv4Base struct {
	IPv4 *IPv4 `json:"ipv4"`
}

type instancesBase struct {
	Meta      *Meta      `json:"meta"`
	Instances []Instance `json:"instances"`
}

// Neighbors that might exist on the same host.
type Neighbors struct {
	Neighbors []string `json:"neighbors"`
}

// Bandwidth used on a given instance.
type Bandwidth struct {
	Bandwidth map[string]struct {
		IncomingBytes int `json:"incoming_bytes"`
		OutgoingBytes int `json:"outgoing_bytes"`
	} `json:"bandwidth"`
}

type privateNetworksBase struct {
	Meta            *Meta            `json:"meta"`
	PrivateNetworks []PrivateNetwork `json:"private_networks"`
}

// PrivateNetwork information for a given instance.
// Deprecated: PrivateNetwork should no longer be used. Instead, use VPCInfo.
type PrivateNetwork struct {
	NetworkID  string `json:"network_id"`
	MacAddress string `json:"mac_address"`
	IPAddress  string `json:"ip_address"`
}

type vpcInfoBase struct {
	Meta *Meta     `json:"meta"`
	VPCs []VPCInfo `json:"vpcs"`
}

// VPCInfo information for a given instance.
type VPCInfo struct {
	ID         string `json:"id"`
	MacAddress string `json:"mac_address"`
	IPAddress  string `json:"ip_address"`
}

type vpc2InfoBase struct {
	Meta *Meta      `json:"meta"`
	VPCs []VPC2Info `json:"vpcs"`
}

// VPC2Info information for a given instance.
type VPC2Info struct {
	ID         string `json:"id"`
	MacAddress string `json:"mac_address"`
	IPAddress  string `json:"ip_address"`
}

// AttachVPC2Req parameters for attaching a VPC 2.0 network
type AttachVPC2Req struct {
	IPAddress *string `json:"ip_address,omitempty"`
	VPCID     string  `json:"vpc_id,omitempty"`
}

type isoStatusBase struct {
	IsoStatus *Iso `json:"iso_status"`
}

// Iso information for a given instance.
type Iso struct {
	State string `json:"state"`
	IsoID string `json:"iso_id"`
}

type backupScheduleBase struct {
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
}

// BackupSchedule information for a given instance.
type BackupSchedule struct {
	Enabled             *bool  `json:"enabled,omitempty"`
	Type                string `json:"type,omitempty"`
	NextScheduleTimeUTC string `json:"next_scheduled_time_utc,omitempty"`
	Hour                int    `json:"hour,omitempty"`
	Dow                 int    `json:"dow,omitempty"`
	Dom                 int    `json:"dom,omitempty"`
}

// BackupScheduleReq struct used to create a backup schedule for an instance.
type BackupScheduleReq struct {
	Hour *int   `json:"hour,omitempty"`
	Dow  *int   `json:"dow,omitempty"`
	Type string `json:"type"`
	Dom  int    `json:"dom,omitempty"`
}

// RestoreReq struct used to supply whether a restore should be from a backup or snapshot.
type RestoreReq struct {
	BackupID   string `json:"backup_id,omitempty"`
	SnapshotID string `json:"snapshot_id,omitempty"`
}

// todo can we remove this list and return this data back in the list?
type reverseIPv6sBase struct {
	ReverseIPv6s []ReverseIP `json:"reverse_ipv6s"`
	// no meta?
}

// ReverseIP information for a given instance.
type ReverseIP struct {
	IP      string `json:"ip"`
	Reverse string `json:"reverse"`
}

type userDataBase struct {
	UserData *UserData `json:"user_data"`
}

// UserData information for a given struct.
type UserData struct {
	Data string `json:"data"`
}

type upgradeBase struct {
	Upgrades *Upgrades `json:"upgrades"`
}

// Upgrades that are available for a given Instance.
type Upgrades struct {
	Applications []Application `json:"applications,omitempty"`
	OS           []OS          `json:"os,omitempty"`
	Plans        []string      `json:"plans,omitempty"`
}

// InstanceCreateReq struct used to create an instance.
type InstanceCreateReq struct {
	DisablePublicIPv4    *bool             `json:"disable_public_ipv4,omitempty"`
	EnableIPv6           *bool             `json:"enable_ipv6,omitempty"`
	EnableVPC2           *bool             `json:"enable_vpc2,omitempty"`
	DDOSProtection       *bool             `json:"ddos_protection,omitempty"`
	AppVariables         map[string]string `json:"app_variables,omitempty"`
	ActivationEmail      *bool             `json:"activation_email,omitempty"`
	EnableVPC            *bool             `json:"enable_vpc,omitempty"`
	EnablePrivateNetwork *bool             `json:"enable_private_network,omitempty"`
	ISOID                string            `json:"iso_id,omitempty"`
	Backups              string            `json:"backups,omitempty"`
	Region               string            `json:"region,omitempty"`
	Label                string            `json:"label,omitempty"`
	ImageID              string            `json:"image_id,omitempty"`
	SnapshotID           string            `json:"snapshot_id,omitempty"`
	ScriptID             string            `json:"script_id,omitempty"`
	FirewallGroupID      string            `json:"firewall_group_id,omitempty"`
	ReservedIPv4         string            `json:"reserved_ipv4,omitempty"`
	UserData             string            `json:"user_data,omitempty"`
	Plan                 string            `json:"plan,omitempty"`
	Tag                  string            `json:"tag,omitempty"`
	IPXEChainURL         string            `json:"ipxe_chain_url,omitempty"`
	Hostname             string            `json:"hostname,omitempty"`
	AttachVPC2           []string          `json:"attach_vpc2,omitempty"`
	SSHKeys              []string          `json:"sshkey_id,omitempty"`
	AttachVPC            []string          `json:"attach_vpc,omitempty"`
	AttachPrivateNetwork []string          `json:"attach_private_network,omitempty"`
	Tags                 []string          `json:"tags"`
	AppID                int               `json:"app_id,omitempty"`
	OsID                 int               `json:"os_id,omitempty"`
}

// InstanceUpdateReq struct used to update an instance.
type InstanceUpdateReq struct {
	EnableVPC2           *bool    `json:"enable_vpc2,omitempty"`
	Tag                  *string  `json:"tag,omitempty"`
	DDOSProtection       *bool    `json:"ddos_protection"`
	EnableIPv6           *bool    `json:"enable_ipv6,omitempty"`
	EnablePrivateNetwork *bool    `json:"enable_private_network,omitempty"`
	EnableVPC            *bool    `json:"enable_vpc,omitempty"`
	Label                string   `json:"label,omitempty"`
	FirewallGroupID      string   `json:"firewall_group_id,omitempty"`
	UserData             string   `json:"user_data,omitempty"`
	ImageID              string   `json:"image_id,omitempty"`
	Backups              string   `json:"backups,omitempty"`
	Plan                 string   `json:"plan,omitempty"`
	AttachVPC2           []string `json:"attach_vpc2,omitempty"`
	DetachVPC            []string `json:"detach_vpc,omitempty"`
	AttachVPC            []string `json:"attach_vpc,omitempty"`
	DetachPrivateNetwork []string `json:"detach_private_network,omitempty"`
	DetachVPC2           []string `json:"detach_vpc2,omitempty"`
	AttachPrivateNetwork []string `json:"attach_private_network,omitempty"`
	Tags                 []string `json:"tags"`
	AppID                int      `json:"app_id,omitempty"`
	OsID                 int      `json:"os_id,omitempty"`
}

// ReinstallReq struct used to allow changes during a reinstall
type ReinstallReq struct {
	Hostname string `json:"hostname,omitempty"`
}

// Create will create the server with the given parameters
func (i *InstanceServiceHandler) Create(ctx context.Context, instanceReq *InstanceCreateReq) (*Instance, *http.Response, error) {
	req, err := i.client.NewRequest(ctx, http.MethodPost, instancePath, instanceReq)
	if err != nil {
		return nil, nil, err
	}

	instance := new(instanceBase)
	resp, err := i.client.DoWithContext(ctx, req, instance)
	if err != nil {
		return nil, resp, err
	}

	return instance.Instance, resp, nil
}

// Get will get the server with the given instanceID
func (i *InstanceServiceHandler) Get(ctx context.Context, instanceID string) (*Instance, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	instance := new(instanceBase)
	resp, err := i.client.DoWithContext(ctx, req, instance)
	if err != nil {
		return nil, resp, err
	}

	return instance.Instance, resp, nil
}

// Update will update the server with the given parameters
func (i *InstanceServiceHandler) Update(ctx context.Context, instanceID string, instanceReq *InstanceUpdateReq) (*Instance, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPatch, uri, instanceReq)
	if err != nil {
		return nil, nil, err
	}

	instance := new(instanceBase)
	resp, err := i.client.DoWithContext(ctx, req, instance)
	if err != nil {
		return nil, resp, err
	}

	return instance.Instance, resp, nil
}

// Delete an instance. All data will be permanently lost, and the IP address will be released
func (i *InstanceServiceHandler) Delete(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// List all instances on your account.
func (i *InstanceServiceHandler) List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, *http.Response, error) { //nolint:dupl
	req, err := i.client.NewRequest(ctx, http.MethodGet, instancePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	instances := new(instancesBase)
	resp, err := i.client.DoWithContext(ctx, req, instances)
	if err != nil {
		return nil, nil, resp, err
	}

	return instances.Instances, instances.Meta, resp, nil
}

// Start will start a vps instance the machine is already running, it will be restarted.
func (i *InstanceServiceHandler) Start(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/start", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// Halt will pause an instance.
func (i *InstanceServiceHandler) Halt(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/halt", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// Reboot an instance.
func (i *InstanceServiceHandler) Reboot(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/reboot", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// Reinstall an instance.
func (i *InstanceServiceHandler) Reinstall(ctx context.Context, instanceID string, reinstallReq *ReinstallReq) (*Instance, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/reinstall", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reinstallReq)
	if err != nil {
		return nil, nil, err
	}

	instance := new(instanceBase)
	resp, err := i.client.DoWithContext(ctx, req, instance)
	if err != nil {
		return nil, resp, err
	}
	return instance.Instance, resp, nil
}

// MassStart will start a list of instances the machine is already running, it will be restarted.
func (i *InstanceServiceHandler) MassStart(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/start", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// MassHalt will pause a list of instances.
func (i *InstanceServiceHandler) MassHalt(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/halt", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// MassReboot reboots a list of instances.
func (i *InstanceServiceHandler) MassReboot(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/reboot", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// Restore an instance.
func (i *InstanceServiceHandler) Restore(ctx context.Context, instanceID string, restoreReq *RestoreReq) (*http.Response, error) {
	uri := fmt.Sprintf("%s/%s/restore", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, restoreReq)
	if err != nil {
		return nil, err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// GetBandwidth for a given instance.
func (i *InstanceServiceHandler) GetBandwidth(ctx context.Context, instanceID string) (*Bandwidth, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/bandwidth", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	bandwidth := new(Bandwidth)
	resp, err := i.client.DoWithContext(ctx, req, bandwidth)
	if err != nil {
		return nil, resp, err
	}

	return bandwidth, resp, nil
}

// GetNeighbors gets a list of other instances in the same location as this Instance.
func (i *InstanceServiceHandler) GetNeighbors(ctx context.Context, instanceID string) (*Neighbors, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/neighbors", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	neighbors := new(Neighbors)
	resp, err := i.client.DoWithContext(ctx, req, neighbors)
	if err != nil {
		return nil, resp, err
	}

	return neighbors, resp, nil
}

// ListPrivateNetworks currently attached to an instance.
// Deprecated: ListPrivateNetworks should no longer be used. Instead, use ListVPCInfo
func (i *InstanceServiceHandler) ListPrivateNetworks(ctx context.Context, instanceID string, options *ListOptions) ([]PrivateNetwork, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/private-networks", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	networks := new(privateNetworksBase)
	resp, err := i.client.DoWithContext(ctx, req, networks)
	if err != nil {
		return nil, nil, resp, err
	}

	return networks.PrivateNetworks, networks.Meta, resp, nil
}

// AttachPrivateNetwork to an instance
// Deprecated: AttachPrivateNetwork should no longer be used. Instead, use AttachVPC
func (i *InstanceServiceHandler) AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-networks/attach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachPrivateNetwork from an instance.
// Deprecated: DetachPrivateNetwork should no longer be used. Instead, use DetachVPC
func (i *InstanceServiceHandler) DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-networks/detach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// ListVPCInfo currently attached to an instance.
func (i *InstanceServiceHandler) ListVPCInfo(ctx context.Context, instanceID string, options *ListOptions) ([]VPCInfo, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/vpcs", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vpcs := new(vpcInfoBase)
	resp, err := i.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, nil, resp, err
	}

	return vpcs.VPCs, vpcs.Meta, resp, nil
}

// AttachVPC to an instance
func (i *InstanceServiceHandler) AttachVPC(ctx context.Context, instanceID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpcs/attach", instancePath, instanceID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachVPC from an instance.
func (i *InstanceServiceHandler) DetachVPC(ctx context.Context, instanceID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpcs/detach", instancePath, instanceID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// ListVPC2Info currently attached to an instance.
func (i *InstanceServiceHandler) ListVPC2Info(ctx context.Context, instanceID string, options *ListOptions) ([]VPC2Info, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/vpc2", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vpcs := new(vpc2InfoBase)
	resp, err := i.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, nil, resp, err
	}

	return vpcs.VPCs, vpcs.Meta, resp, nil
}

// AttachVPC2 to an instance
func (i *InstanceServiceHandler) AttachVPC2(ctx context.Context, instanceID string, vpc2Req *AttachVPC2Req) error {
	uri := fmt.Sprintf("%s/%s/vpc2/attach", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, vpc2Req)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachVPC2 from an instance.
func (i *InstanceServiceHandler) DetachVPC2(ctx context.Context, instanceID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpc2/detach", instancePath, instanceID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// ISOStatus retrieves the current ISO state for a given VPS.
// The returned state may be one of: ready | isomounting | isomounted.
func (i *InstanceServiceHandler) ISOStatus(ctx context.Context, instanceID string) (*Iso, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/iso", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	iso := new(isoStatusBase)
	resp, err := i.client.DoWithContext(ctx, req, iso)
	if err != nil {
		return nil, resp, err
	}
	return iso.IsoStatus, resp, nil
}

// AttachISO will attach an ISO to the given instance and reboot it
func (i *InstanceServiceHandler) AttachISO(ctx context.Context, instanceID, isoID string) (*http.Response, error) {
	uri := fmt.Sprintf("%s/%s/iso/attach", instancePath, instanceID)
	body := RequestBody{"iso_id": isoID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// DetachISO will detach the currently mounted ISO and reboot the instance.
func (i *InstanceServiceHandler) DetachISO(ctx context.Context, instanceID string) (*http.Response, error) {
	uri := fmt.Sprintf("%s/%s/iso/detach", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// GetBackupSchedule retrieves the backup schedule for a given instance - all time values are in UTC
func (i *InstanceServiceHandler) GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	backup := new(backupScheduleBase)
	resp, err := i.client.DoWithContext(ctx, req, backup)
	if err != nil {
		return nil, resp, err
	}

	return backup.BackupSchedule, resp, nil
}

// SetBackupSchedule sets the backup schedule for a given instance - all time values are in UTC.
func (i *InstanceServiceHandler) SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) (*http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, backup)
	if err != nil {
		return nil, err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// CreateIPv4 an additional IPv4 address for given instance.
func (i *InstanceServiceHandler) CreateIPv4(ctx context.Context, instanceID string, reboot *bool) (*IPv4, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/ipv4", instancePath, instanceID)

	body := RequestBody{"reboot": reboot}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, nil, err
	}

	ip := new(ipv4Base)
	resp, err := i.client.DoWithContext(ctx, req, ip)
	if err != nil {
		return nil, resp, err
	}

	return ip.IPv4, resp, nil
}

// ListIPv4 addresses that are currently assigned to a given instance.
func (i *InstanceServiceHandler) ListIPv4(ctx context.Context, instanceID string, options *ListOptions) ([]IPv4, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/ipv4", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	resp, err := i.client.DoWithContext(ctx, req, ips)
	if err != nil {
		return nil, nil, resp, err
	}

	return ips.IPv4s, ips.Meta, resp, nil
}

// DeleteIPv4 address from a given instance.
func (i *InstanceServiceHandler) DeleteIPv4(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv4/%s", instancePath, instanceID, ip)
	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// ListIPv6 addresses that are currently assigned to a given instance.
func (i *InstanceServiceHandler) ListIPv6(ctx context.Context, instanceID string, options *ListOptions) ([]IPv6, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/ipv6", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	resp, err := i.client.DoWithContext(ctx, req, ips)
	if err != nil {
		return nil, nil, resp, err
	}

	return ips.IPv6s, ips.Meta, resp, nil
}

// CreateReverseIPv6 for a given instance.
func (i *InstanceServiceHandler) CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// ListReverseIPv6 currently assigned to a given instance.
func (i *InstanceServiceHandler) ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	reverse := new(reverseIPv6sBase)
	resp, err := i.client.DoWithContext(ctx, req, reverse)
	if err != nil {
		return nil, resp, err
	}

	return reverse.ReverseIPv6s, resp, nil
}

// DeleteReverseIPv6 a given reverse IPv6.
func (i *InstanceServiceHandler) DeleteReverseIPv6(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse/%s", instancePath, instanceID, ip)
	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// CreateReverseIPv4 for a given IP on a given instance.
func (i *InstanceServiceHandler) CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv4/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// DefaultReverseIPv4 will set the IPs reverse setting back to the original one supplied by Vultr.
func (i *InstanceServiceHandler) DefaultReverseIPv4(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv4/reverse/default", instancePath, instanceID)
	reqBody := RequestBody{"ip": ip}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// GetUserData from given instance. The userdata returned will be in base64 encoding.
func (i *InstanceServiceHandler) GetUserData(ctx context.Context, instanceID string) (*UserData, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/user-data", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	userData := new(userDataBase)
	resp, err := i.client.DoWithContext(ctx, req, userData)
	if err != nil {
		return nil, resp, err
	}

	return userData.UserData, resp, nil
}

// GetUpgrades that are available for a given instance.
func (i *InstanceServiceHandler) GetUpgrades(ctx context.Context, instanceID string) (*Upgrades, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/upgrades", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	upgrades := new(upgradeBase)
	resp, err := i.client.DoWithContext(ctx, req, upgrades)
	if err != nil {
		return nil, resp, err
	}

	return upgrades.Upgrades, resp, nil
}
