package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const clusterPath = "/v2/clusters"

// ClusterService is the interface to interact with the cluster endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/clusters
type ClusterService interface {
	Create(ctx context.Context, clusterReq *ClusterCreate) (*InstanceCluster, *http.Response, error)
	Get(ctx context.Context, clusterID string) (*InstanceCluster, *http.Response, error)
	Update(ctx context.Context, clusterID string, clusterReq *ClusterUpdate) (*InstanceCluster, *http.Response, error)
	Delete(ctx context.Context, clusterID string) error
	List(ctx context.Context, options *ListOptions) ([]InstanceCluster, *Meta, *http.Response, error)

	AttachInstance(ctx context.Context, clusterID string, instanceID string) error
	AttachInstances(ctx context.Context, clusterID string, instanceIDs []string) error
	AttachHeadNode(ctx context.Context, clusterID string, instanceID string) error
	DetachInstance(ctx context.Context, clusterID string, instanceID string) error
	DetachInstances(ctx context.Context, clusterID string, instanceIDs []string) error

	GetMetrics(ctx context.Context, clusterID string, options *ClusterMetricsOpts) (*ClusterMetrics, *http.Response, error)
}

// ClusterServiceHandler handles interaction with the server methods for the Vultr API
type ClusterServiceHandler struct {
	client *Client
}

// InstanceCluster represents a cluster of Vultr instances or bare metal servers
type InstanceCluster struct {
	ID                         string                  `json:"id"`
	Region                     string                  `json:"region"`
	Label                      string                  `json:"label"`
	Plan                       string                  `json:"plan"`
	MinPoolCount               int                     `json:"min_pool_count"`
	DesiredPoolCount           int                     `json:"desired_pool_count"`
	Hostname                   string                  `json:"hostname"`
	Status                     string                  `json:"status"`
	State                      string                  `json:"state"`
	InstanceTemplate           ClusterInstanceTemplate `json:"instance_template"`
	DateCreated                string                  `json:"date_created"`
	ClusterType                string                  `json:"cluster_type"`
	Type                       string                  `json:"type"`
	HeadNodeInstanceTemplateID string                  `json:"head_node_instance_template_id"`
	VFS                        []ClusterVFS            `json:"vfs"`
	Instances                  []ClusterInstance       `json:"instances"`
	VPCNetworks                []ClusterVPC            `json:"vpc_networks"`
}

// ClusterInstanceTemplate represents the instance template information for a cluster
type ClusterInstanceTemplate struct {
	ID               string                             `json:"id"`
	Plan             string                             `json:"plan"`
	Label            string                             `json:"label"`
	OS               string                             `json:"os"`
	MarketplaceApp   string                             `json:"marketplace_app"`
	MarketplaceImage string                             `json:"marketplace_image"`
	Snapshot         string                             `json:"snapshot"`
	ISO              string                             `json:"iso"`
	SSHKeys          []InstanceTemplateSSHKey           `json:"ssh_keys"`
	StartupScript    string                             `json:"startup_script"`
	DiskConfig       string                             `json:"disk_config"`
	PlanDetails      ClusterInstanceTemplatePlanDetails `json:"plan_details"`
	UserData         string                             `json:"user_data"`
}

// ClusterInstanceTemplatePlanDetails represents the plan details for a cluster's instance template
type ClusterInstanceTemplatePlanDetails struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	CPUManufacturer string `json:"cpu_manufacturer"`
	CPUCount        int    `json:"cpu_count"`
	CPUThreads      int    `json:"cpu_threads"`
	CPUMhz          int    `json:"cpu_mhz"`
	CPUModel        string `json:"cpu_model"`
	MemoryMB        int    `json:"memory_mb"`
	DiskGB          int    `json:"disk_gb"`
	DiskType        string `json:"disk_type"`
	DiskNum         int    `json:"disk_num"`
	BandwidthTB     int    `json:"bandwidth_tb"`
	GPUBrand        string `json:"gpu_brand"`
	GPUModel        string `json:"gpu_model"`
	GPUCount        int    `json:"gpu_count"`
	GPUProfile      string `json:"gpu_profile"`
	GPUProfileCount int    `json:"gpu_profile_count"`
	Price           string `json:"price"`
	PriceHr         string `json:"price_hr"`
}

// ClusterVFS represents the VFS information for a cluster
type ClusterVFS struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// ClusterInstance represents the instance information for a cluster
type ClusterInstance struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	IPAddress   string `json:"ip_address"`
	Hostname    string `json:"hostname"`
	IsHeadNode  bool   `json:"is_head_node"`
}

// ClusterVPC represents the VPC information for a cluster
type ClusterVPC struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type clusterBase struct {
	Cluster *InstanceCluster `json:"cluster"`
}

type clustersBase struct {
	Clusters []InstanceCluster `json:"clusters"`
	Meta     *Meta             `json:"meta"`
}

// ClusterCreate struct used to create an instance
type ClusterCreate struct {
	Region           string          `json:"region"`
	Plan             string          `json:"plan,omitempty"`
	InstanceTemplate string          `json:"instance_template,omitempty"`
	Label            string          `json:"label,omitempty"`
	MinPoolCount     int             `json:"min_pool_count,omitempty"`
	DesiredPoolCount int             `json:"desired_pool_count,omitempty"`
	Hostname         string          `json:"hostname,omitempty"`
	NotifyActivate   bool            `json:"notify_activate,omitempty"`
	OsID             int             `json:"os_id,omitempty"`
	AppID            int             `json:"app_id,omitempty"`
	ImageID          int             `json:"image_id,omitempty"`
	GPUFabric        bool            `json:"gpu_fabric,omitempty"`
	VFS              ClusterVFSInput `json:"vfs,omitempty"`
	UseHeadNode      bool            `json:"use_head_node,omitempty"`
	HeadNodePlan     string          `json:"head_node_plan,omitempty"`
	HeadNodeTemplate string          `json:"head_node_template,omitempty"`
	VpcIDs           []string        `json:"vpc_ids,omitempty"`
}

// ClusterUpdate struct used to update a cluster
type ClusterUpdate struct {
	Label            string          `json:"label,omitempty"`
	MinPoolCount     int             `json:"min_pool_count,omitempty"`
	Hostname         string          `json:"hostname,omitempty"`
	DesiredPoolCount int             `json:"desired_pool_count,omitempty"`
	VpcIDs           []string        `json:"vpc_ids,omitempty"`
	VFS              ClusterVFSInput `json:"vfs,omitempty"`
}

// ClusterVFSInput represents the attachments for a cluster
type ClusterVFSInput struct {
	IDs    []string                          `json:"ids,omitempty"`
	Create VirtualFileSystemStorageUpdateReq `json:"create,omitempty"`
}

type clusterMassUpdate struct {
	Action    string   `json:"action,omitempty"`
	Instances []string `json:"instances,omitempty"`
}

type ClusterMetrics struct {
	Instances map[string]ClusterInstanceMetrics `json:"instances"`
}

type ClusterInstanceMetrics struct {
	GPU    ClusterGPUMetrics    `json:"gpu"`
	Fabric ClusterFabricMetrics `json:"fabric"`
}

type ClusterGPUMetrics struct {
	Temperature ClusterGPUTemperatureMetrics `json:"temperature"`
	Utilization ClusterGPUUtilizationMetrics `json:"utilization"`
	Power       ClusterGPUPowerMetrics       `json:"power"`
}

type ClusterGPUTemperatureMetrics struct {
	GPU []ClusterMetric `json:"gpu"`
}

type ClusterGPUUtilizationMetrics struct {
	GPU    []ClusterMetric         `json:"gpu"`
	Memory ClusterGPUMemoryMetrics `json:"memory"`
}

type ClusterGPUMemoryMetrics struct {
	Total []ClusterMetric `json:"total"`
	Free  []ClusterMetric `json:"free"`
	Used  []ClusterMetric `json:"used"`
}

type ClusterGPUPowerMetrics struct {
	Used []ClusterMetric `json:"used"`
	Max  []ClusterMetric `json:"max"`
}

type ClusterFabricMetrics struct {
	RawThroughput        ClusterRxTxMetrics `json:"raw_throughput"`
	Retries              ClusterRxTxMetrics `json:"retries"`
	OpticalPowerStrength ClusterRxTxMetrics `json:"optical_power_strength"`
}

type ClusterRxTxMetrics struct {
	Rx []ClusterMetric `json:"rx"`
	Tx []ClusterMetric `json:"tx"`
}

type ClusterMetric struct {
	Unit       string            `json:"unit"`
	Target     string            `json:"target"`
	Datapoints [][]float64       `json:"datapoints"`
	Tags       map[string]string `json:"tags"`
}

type clusterMetricsBase struct {
	Metrics *ClusterMetrics `json:"metrics"`
}

type ClusterMetricsOpts struct {
	Period *string
}

// Create will create the cluster with the given parameters
func (c *ClusterServiceHandler) Create(ctx context.Context, clusterReq *ClusterCreate) (*InstanceCluster, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, clusterPath, clusterReq)
	if err != nil {
		return nil, nil, err
	}

	cluster := new(clusterBase)
	resp, err := c.client.DoWithContext(ctx, req, cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster.Cluster, resp, nil
}

// Get will get the cluster with the given clusterID
func (c *ClusterServiceHandler) Get(ctx context.Context, clusterID string) (*InstanceCluster, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", clusterPath, clusterID)

	req, err := c.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	cluster := new(clusterBase)
	resp, err := c.client.DoWithContext(ctx, req, cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster.Cluster, resp, nil
}

// Update will update the cluster with the given parameters
func (c *ClusterServiceHandler) Update(ctx context.Context, clusterID string, clusterReq *ClusterUpdate) (*InstanceCluster, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", clusterPath, clusterID)

	req, err := c.client.NewRequest(ctx, http.MethodPut, uri, clusterReq)
	if err != nil {
		return nil, nil, err
	}

	cluster := new(clusterBase)
	resp, err := c.client.DoWithContext(ctx, req, cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster.Cluster, resp, nil
}

// Delete a cluster
func (c *ClusterServiceHandler) Delete(ctx context.Context, clusterID string) error {
	uri := fmt.Sprintf("%s/%s", clusterPath, clusterID)

	req, err := c.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// List all clusters on your account.
func (c *ClusterServiceHandler) List(ctx context.Context, options *ListOptions) ([]InstanceCluster, *Meta, *http.Response, error) { //nolint:dupl,lll
	req, err := c.client.NewRequest(ctx, http.MethodGet, clusterPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	clusters := new(clustersBase)
	resp, err := c.client.DoWithContext(ctx, req, clusters)
	if err != nil {
		return nil, nil, resp, err
	}

	return clusters.Clusters, clusters.Meta, resp, nil
}

// AttachInstance will attach an instance to a cluster.
func (c *ClusterServiceHandler) AttachInstance(ctx context.Context, clusterID, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/attach/%s", clusterPath, clusterID, instanceID)

	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// AttachInstances will attach multiple instances to a cluster.
func (c *ClusterServiceHandler) AttachInstances(ctx context.Context, clusterID string, instanceIDs []string) error {
	uri := fmt.Sprintf("%s/%s", clusterPath, clusterID)

	reqBody := clusterMassUpdate{Action: "attach", Instances: instanceIDs}
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// AttachHeadNode will attach an instance to a cluster as a head node.
func (c *ClusterServiceHandler) AttachHeadNode(ctx context.Context, clusterID, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/attach_head_node/%s", clusterPath, clusterID, instanceID)

	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}
	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

func (c *ClusterServiceHandler) DetachInstance(ctx context.Context, clusterID, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/detach/%s", clusterPath, clusterID, instanceID)

	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}
	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachInstances will detach multiple instances from a cluster.
func (c *ClusterServiceHandler) DetachInstances(ctx context.Context, clusterID string, instanceIDs []string) error {
	uri := fmt.Sprintf("%s/%s", clusterPath, clusterID)

	reqBody := clusterMassUpdate{Action: "detach", Instances: instanceIDs}
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

func (c *ClusterServiceHandler) GetMetrics(ctx context.Context, clusterID string, options *ClusterMetricsOpts) (*ClusterMetrics, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/metrics", clusterPath, clusterID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	queryParam := req.URL.Query()
	if options.Period != nil {
		queryParam.Add("period", *options.Period)
	}

	req.URL.RawQuery = queryParam.Encode()

	metrics := new(clusterMetricsBase)
	resp, err := c.client.DoWithContext(ctx, req, metrics)
	if err != nil {
		return nil, resp, err
	}

	return metrics.Metrics, resp, nil
}
