package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vkePath = "/v2/kubernetes/clusters"

type KubernetesService interface {
	CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, error)
	GetCluster(ctx context.Context, id string) (*Cluster, error)
	ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, error)
	UpdateCluster()
	DeleteCluster(ctx context.Context, id string) error

	CreateNodePool()
	ListNodePools()
	GetNodePool()
	UpdateNodePool()
	DeleteNodePool()

	DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error
	RecycleNodePoolInstance()

	GetKubeConfig()
}

type KubernetesHandler struct {
	client *Client
}

type Cluster struct {
	ID            string     `json:"id"`
	Label         string     `json:"label"`
	DateCreated   string     `json:"date_created"`
	ClusterSubnet string     `json:"cluster_subnet"`
	ServiceSubnet string     `json:"service_subnet"`
	IP            string     `json:"ip"`
	Endpoint      string     `json:"endpoint"`
	Version       string     `json:"version"`
	Region        string     `json:"region"`
	Status        string     `json:"status"`
	NodePools     []NodePool `json:"node_pools"`
}

type NodePool struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Label       string `json:"label"`
	PlanID      string `json:"plan_id"`
	Status      string `json:"status"`
	Count       int    `json:"count"`
	Nodes       []Node `json:"nodes"`
}

type Node struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type ClusterReq struct {
	Label     string        `json:"label"`
	Region    string        `json:"region"`
	Version   string        `json:"version"`
	NodePools []NodePoolReq `json:"node_pools"`
}

type NodePoolReq struct {
	NodeQuantity int    `json:"node_quantity"`
	Label        string `json:"label"`
	Plan         string `json:"plan"`
}

type vkeClustersBase struct {
	VKEClusters []Cluster `json:"vke_clusters"`
	Meta        *Meta     `json:"meta"`
}

type vkeClusterBase struct {
	VKECluster *Cluster `json:"vke_cluster"`
}

// CreateCluster will create a Kubernetes cluster.
func (k *KubernetesHandler) CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPost, vkePath, createReq)
	if err != nil {
		return nil, err
	}

	var k8 = new(vkeClusterBase)
	if err = k.client.DoWithContext(ctx, req, &k8); err != nil {
		return nil, err
	}

	return k8.VKECluster, nil
}

// GetCluster will return a Kubernetes cluster.
func (k *KubernetesHandler) GetCluster(ctx context.Context, id string) (*Cluster, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return nil, err
	}

	k8 := new(vkeClusterBase)
	if err = k.client.DoWithContext(ctx, req, &k8); err != nil {
		return nil, err
	}

	return k8.VKECluster, nil
}

// ListClusters will return all kubernetes clusters.
func (k *KubernetesHandler) ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, vkePath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	k8s := new(vkeClustersBase)
	if err = k.client.DoWithContext(ctx, req, &k8s); err != nil {
		return nil, nil, err
	}

	return k8s.VKEClusters, k8s.Meta, nil
}

func (k *KubernetesHandler) UpdateCluster() {
	panic("implement me")
}

// DeleteCluster will delete a Kubernetes cluster.
func (k *KubernetesHandler) DeleteCluster(ctx context.Context, id string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

func (k *KubernetesHandler) CreateNodePool() {
	panic("implement me")
}

func (k *KubernetesHandler) ListNodePools() {
	panic("implement me")
}

func (k *KubernetesHandler) GetNodePool() {
	panic("implement me")
}

func (k *KubernetesHandler) UpdateNodePool() {
	panic("implement me")
}

func (k *KubernetesHandler) DeleteNodePool() {
	panic("implement me")
}

// DeleteNodePoolInstance will remove a specified node from a nodepool
func (k *KubernetesHandler) DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s", vkePath, vkeID, nodePoolID, nodeID), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// RecycleNodePoolInstance will recycle (destroy + redeploy) a given node on a nodepool
func (k *KubernetesHandler) RecycleNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s/recycle", vkePath, vkeID, nodePoolID, nodeID), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

func (k *KubernetesHandler) GetKubeConfig() {
	panic("implement me")
}
