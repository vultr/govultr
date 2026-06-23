package govultr

// ListOptions are the available query params
type ListOptions struct {
	// These query params are used for all list calls that support pagination
	PerPage int    `url:"per_page,omitempty"`
	Cursor  string `url:"cursor,omitempty"`

	// Query params that can be used on the list applications call
	// https://www.vultr.com/api/#operation/list-applications
	Type string `url:"type,omitempty"`

	// Query params that can be used on the list backups call
	// https://www.vultr.com/api/#operation/list-backups
	InstanceID string `url:"instance_id,omitempty"`

	// Query params that can be used on the list calls in the container registry endpoints
	// https://www.vultr.com/api/#tag/Container-Registry
	RegistryID string `url:"registry-id,omitempty"`

	// Query params that can be used on the list artifacts call
	// https://www.vultr.com/api/#operation/list-registry-repository-artifacts
	RepositoryImage string `url:"repository-image,omitempty"`

	// These three query params are currently used for the list instance call
	// These may be extended to other list calls
	// https://www.vultr.com/api/#operation/list-instances
	MainIP             string `url:"main_ip,omitempty"`
	Label              string `url:"label,omitempty"`
	Tag                string `url:"tag,omitempty"`
	Region             string `url:"region,omitempty"`
	FirewallGroupID    string `url:"firewall_group_id,omitempty"`
	Hostname           string `url:"hostname,omitempty"`
	ShowPendingCharges bool   `url:"show_pending_charges,omitempty"`

	// Query params that can be used on the list instance IPv4 call
	// https://www.vultr.com/api/#operation/get-instance-ipv4
	PublicNetwork bool `url:"public_network,omitempty"`

	// Query params that can be used on the list snapshots call
	// https://www.vultr.com/api/#operation/list-snapshots
	Description string `url:"description,omitempty"`
}
