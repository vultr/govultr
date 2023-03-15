package govultr

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestKubernetesHandler_CreateCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(vkePath, func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "vke_cluster": {
        "id": "014da059-21e3-47eb-acb5-91bf697c31aa",
        "label": "vke",
        "date_created": "2021-07-13T14:20:16+00:00",
        "cluster_subnet": "10.244.0.0/16",
        "service_subnet": "10.96.0.0/12",
        "ip": "0.0.0.0",
        "endpoint": "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
        "version": "1.20",
        "region": "lax",
        "status": "pending",
        "node_pools": [
            {
                "id": "e1c7a313-e42d-43bb-82ef-4f287639b303",
                "date_created": "2021-07-13T14:20:16+00:00",
                "label": "my-label-48957292",
                "plan": "vc2-1c-2gb",
                "status": "pending",
                "node_quantity": 1,
				"min_nodes": 1,
				"max_nodes": 2,
				"auto_scaler": true,
                "nodes": [
                    {
                        "id": "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
                        "label": "my-label-48957292-fef60eda12071",
                        "date_created": "2021-07-13T14:20:16+00:00",
                        "status": "pending"
                    }
                ]
            }
        ]
    }
}`
		fmt.Fprint(writer, response)
	})

	createReq := &ClusterReq{
		Label:     "vke",
		Region:    "lax",
		Version:   "1.20",
		NodePools: nil,
	}
	vke, _, err := client.Kubernetes.CreateCluster(ctx, createReq)
	if err != nil {
		t.Errorf("Kubernetes.CreateCluster returned %v", err)
	}

	expected := &Cluster{
		ID:            "014da059-21e3-47eb-acb5-91bf697c31aa",
		Label:         "vke",
		DateCreated:   "2021-07-13T14:20:16+00:00",
		ClusterSubnet: "10.244.0.0/16",
		ServiceSubnet: "10.96.0.0/12",
		IP:            "0.0.0.0",
		Endpoint:      "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
		Version:       "1.20",
		Region:        "lax",
		Status:        "pending",
		NodePools: []NodePool{
			{
				ID:           "e1c7a313-e42d-43bb-82ef-4f287639b303",
				DateCreated:  "2021-07-13T14:20:16+00:00",
				Label:        "my-label-48957292",
				Plan:         "vc2-1c-2gb",
				Status:       "pending",
				NodeQuantity: 1,
				MinNodes:     1,
				MaxNodes:     2,
				AutoScaler:   true,
				Nodes: []Node{
					{
						ID:          "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
						DateCreated: "2021-07-13T14:20:16+00:00",
						Label:       "my-label-48957292-fef60eda12071",
						Status:      "pending",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(vke, expected) {
		t.Errorf("Kubernetes.CreateCluster returned %+v, expected %+v", vke, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.CreateCluster(c, createReq)
	if err == nil {
		t.Error("Kubernetes.CreateCluster returned nil")
	}
}

func TestKubernetesHandler_GetCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s", vkePath, "014da059-21e3-47eb-acb5-91bf697c31aa"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "vke_cluster": {
        "id": "014da059-21e3-47eb-acb5-91bf697c31aa",
        "label": "vke",
        "date_created": "2021-07-13T14:20:16+00:00",
        "cluster_subnet": "10.244.0.0/16",
        "service_subnet": "10.96.0.0/12",
        "ip": "0.0.0.0",
        "endpoint": "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
        "version": "1.20",
        "region": "lax",
        "status": "pending",
        "node_pools": [
            {
                "id": "e1c7a313-e42d-43bb-82ef-4f287639b303",
                "date_created": "2021-07-13T14:20:16+00:00",
                "label": "my-label-48957292",
                "plan": "vc2-1c-2gb",
                "status": "pending",
                "node_quantity": 1,
				"min_nodes": 1,
				"max_nodes": 2,
				"auto_scaler": true,
                "nodes": [
                    {
                        "id": "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
                        "label": "my-label-48957292-fef60eda12071",
                        "date_created": "2021-07-13T14:20:16+00:00",
                        "status": "pending"
                    }
                ]
            }
        ]
    }
}`
		fmt.Fprint(writer, response)
	})

	vke, _, err := client.Kubernetes.GetCluster(ctx, "014da059-21e3-47eb-acb5-91bf697c31aa")
	if err != nil {
		t.Errorf("Kubernetes.GetCluster returned %v", err)
	}

	expected := &Cluster{
		ID:            "014da059-21e3-47eb-acb5-91bf697c31aa",
		Label:         "vke",
		DateCreated:   "2021-07-13T14:20:16+00:00",
		ClusterSubnet: "10.244.0.0/16",
		ServiceSubnet: "10.96.0.0/12",
		IP:            "0.0.0.0",
		Endpoint:      "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
		Version:       "1.20",
		Region:        "lax",
		Status:        "pending",
		NodePools: []NodePool{
			{
				ID:           "e1c7a313-e42d-43bb-82ef-4f287639b303",
				DateCreated:  "2021-07-13T14:20:16+00:00",
				Label:        "my-label-48957292",
				Plan:         "vc2-1c-2gb",
				Status:       "pending",
				NodeQuantity: 1,
				MinNodes:     1,
				MaxNodes:     2,
				AutoScaler:   true,
				Nodes: []Node{
					{
						ID:          "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
						DateCreated: "2021-07-13T14:20:16+00:00",
						Label:       "my-label-48957292-fef60eda12071",
						Status:      "pending",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(vke, expected) {
		t.Errorf("Kubernetes.GetCluster returned %+v, expected %+v", vke, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.GetCluster(c, "014da059-21e3-47eb-acb5-91bf697c31aa")
	if err == nil {
		t.Error("Kubernetes.GetCluster returned nil")
	}
}

func TestKubernetesHandler_ListClusters(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(vkePath, func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "vke_clusters": [{
        "id": "014da059-21e3-47eb-acb5-91bf697c31aa",
        "label": "vke",
        "date_created": "2021-07-13T14:20:16+00:00",
        "cluster_subnet": "10.244.0.0/16",
        "service_subnet": "10.96.0.0/12",
        "ip": "0.0.0.0",
        "endpoint": "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
        "version": "1.20",
        "region": "lax",
        "status": "pending",
        "node_pools": [
            {
                "id": "e1c7a313-e42d-43bb-82ef-4f287639b303",
                "date_created": "2021-07-13T14:20:16+00:00",
                "label": "my-label-48957292",
                "plan": "vc2-1c-2gb",
                "status": "pending",
				"tag": "mytag",
                "node_quantity": 1,
				"min_nodes": 1,
				"max_nodes": 2,
				"auto_scaler": true,
                "nodes": [
                    {
                        "id": "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
                        "label": "my-label-48957292-fef60eda12071",
                        "date_created": "2021-07-13T14:20:16+00:00",
                        "status": "pending"
                    }
                ]
            }
        ]
    }
],
    "meta": {
        "total": 1,
        "links": {
            "next": "thisismycusror",
            "prev": ""
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vke, meta, _, err := client.Kubernetes.ListClusters(ctx, nil)
	if err != nil {
		t.Errorf("Kubernetes.ListClusters returned %v", err)
	}

	expected := []Cluster{
		{
			ID:            "014da059-21e3-47eb-acb5-91bf697c31aa",
			Label:         "vke",
			DateCreated:   "2021-07-13T14:20:16+00:00",
			ClusterSubnet: "10.244.0.0/16",
			ServiceSubnet: "10.96.0.0/12",
			IP:            "0.0.0.0",
			Endpoint:      "014da059-21e3-47eb-acb5-91bf697c31aa.vultr-k8s.com",
			Version:       "1.20",
			Region:        "lax",
			Status:        "pending",
			NodePools: []NodePool{
				{
					ID:           "e1c7a313-e42d-43bb-82ef-4f287639b303",
					DateCreated:  "2021-07-13T14:20:16+00:00",
					Label:        "my-label-48957292",
					Plan:         "vc2-1c-2gb",
					Status:       "pending",
					Tag:          "mytag",
					NodeQuantity: 1,
					MinNodes:     1,
					MaxNodes:     2,
					AutoScaler:   true,
					Nodes: []Node{
						{
							ID:          "38364f79-17e3-4f1f-b7df-d9494bce0e4a",
							DateCreated: "2021-07-13T14:20:16+00:00",
							Label:       "my-label-48957292-fef60eda12071",
							Status:      "pending",
						},
					},
				},
			},
		},
	}
	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}
	if !reflect.DeepEqual(vke, expected) {
		t.Errorf("Kubernetes.List returned %+v, expected %+v", vke, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Kubernetes.List meta returned %+v, expected %+v", vke, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, _, err = client.Kubernetes.ListClusters(c, nil)
	if err == nil {
		t.Error("Kubernetes.ListClusters returned nil")
	}
}

func TestKubernetesHandler_UpdateCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s", vkePath, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})
	update := ClusterReqUpdate{Label: "new label"}
	err := client.Kubernetes.UpdateCluster(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", &update)

	if err != nil {
		t.Errorf("Kubernetes.UpdateCluster returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s", vkePath, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Kubernetes.DeleteCluster(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Kubernetes.DeleteCluster returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteClusterWithResources(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/delete-with-linked-resources", vkePath, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Kubernetes.DeleteClusterWithResources(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Kubernetes.DeleteClusterWithResources returned %+v", err)
	}
}

func TestKubernetesHandler_CreateNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/node-pools", vkePath, "1"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "node_pool": {
        "id": "554e7248-705a-5862-516f-4f4a6735346a",
        "date_created": "2021-07-13T15:42:21+00:00",
        "label": "nodepool-48959140",
        "plan": "vc2-1c-2gb",
        "status": "pending",
        "node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "pending"
            }
        ]
    }
}`
		fmt.Fprint(writer, response)
	})

	createReq := &NodePoolReq{
		NodeQuantity: 1,
		Label:        "nodepool-48959140",
		Plan:         "vc2-1c-2gb",
		Tag:          "mytag",
	}
	np, _, err := client.Kubernetes.CreateNodePool(ctx, "1", createReq)
	if err != nil {
		t.Errorf("Kubernetes.CreateNodePool returned %v", err)
	}

	expected := &NodePool{
		ID:           "554e7248-705a-5862-516f-4f4a6735346a",
		DateCreated:  "2021-07-13T15:42:21+00:00",
		Label:        "nodepool-48959140",
		Plan:         "vc2-1c-2gb",
		Status:       "pending",
		NodeQuantity: 1,
		MinNodes:     1,
		MaxNodes:     2,
		AutoScaler:   true,
		Tag:          "mytag",
		Nodes: []Node{
			{
				ID:          "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
				Label:       "nodepool-48959140-74a60edb45de0",
				DateCreated: "2021-07-13T15:42:21+00:00",
				Status:      "pending",
			},
		},
	}

	if !reflect.DeepEqual(np, expected) {
		t.Errorf("Kubernetes.CreateNodePool returned %+v, expected %+v", np, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.CreateNodePool(c, "1", createReq)
	if err == nil {
		t.Error("Kubernetes.CreateNodePool returned nil")
	}
}

func TestKubernetesHandler_GetNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/node-pools/%s", vkePath, "1", "2"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "node_pool": {
        "id": "554e7248-705a-5862-516f-4f4a6735346a",
        "date_created": "2021-07-13T15:42:21+00:00",
        "label": "nodepool-48959140",
        "plan": "vc2-1c-2gb",
        "status": "pending",
        "node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "pending"
            }
        ]
    }
}`
		fmt.Fprint(writer, response)
	})

	np, _, err := client.Kubernetes.GetNodePool(ctx, "1", "2")
	if err != nil {
		t.Errorf("Kubernetes.GetNodePool returned %v", err)
	}

	expected := &NodePool{
		ID:           "554e7248-705a-5862-516f-4f4a6735346a",
		DateCreated:  "2021-07-13T15:42:21+00:00",
		Label:        "nodepool-48959140",
		Plan:         "vc2-1c-2gb",
		Status:       "pending",
		Tag:          "mytag",
		NodeQuantity: 1,
		MinNodes:     1,
		MaxNodes:     2,
		AutoScaler:   true,
		Nodes: []Node{
			{
				ID:          "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
				Label:       "nodepool-48959140-74a60edb45de0",
				DateCreated: "2021-07-13T15:42:21+00:00",
				Status:      "pending",
			},
		},
	}

	if !reflect.DeepEqual(np, expected) {
		t.Errorf("Kubernetes.GetNodePool returned %+v, expected %+v", np, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.GetNodePool(c, "1", "2")
	if err == nil {
		t.Error("Kubernetes.GetNodePool returned nil")
	}
}

func TestKubernetesHandler_ListNodePools(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/node-pools", vkePath, "1"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "node_pools": [{
        "id": "554e7248-705a-5862-516f-4f4a6735346a",
        "date_created": "2021-07-13T15:42:21+00:00",
        "label": "nodepool-48959140",
        "plan": "vc2-1c-2gb",
        "status": "pending",
        "node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "pending"
            }
        ]
    }],
    "meta": {
        "total": 1,
        "links": {
            "next": "thisismycusror",
            "prev": ""
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	np, meta, _, err := client.Kubernetes.ListNodePools(ctx, "1", nil)
	if err != nil {
		t.Errorf("Kubernetes.ListNodePools returned %v", err)
	}

	expected := []NodePool{
		{
			ID:           "554e7248-705a-5862-516f-4f4a6735346a",
			DateCreated:  "2021-07-13T15:42:21+00:00",
			Label:        "nodepool-48959140",
			Plan:         "vc2-1c-2gb",
			Status:       "pending",
			Tag:          "mytag",
			NodeQuantity: 1,
			MinNodes:     1,
			MaxNodes:     2,
			AutoScaler:   true,
			Nodes: []Node{
				{
					ID:          "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
					Label:       "nodepool-48959140-74a60edb45de0",
					DateCreated: "2021-07-13T15:42:21+00:00",
					Status:      "pending",
				},
			},
		},
	}

	if !reflect.DeepEqual(np, expected) {
		t.Errorf("Kubernetes.ListNodePools returned %+v, expected %+v", np, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Kubernetes.ListNodePools meta returned %+v, expected %+v", meta, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, _, err = client.Kubernetes.ListNodePools(c, "1", nil)
	if err == nil {
		t.Error("Kubernetes.ListNodePools returned nil")
	}
}

func TestKubernetesHandler_UpdateNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/node-pools/%s", vkePath, "1", "2"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
  "node_pool": {
    "id": "e97bdee9-2781-4f31-be03-60fc75f399ae",
    "date_created": "2021-07-07T23:27:08+00:00",
    "date_updated": "2021-07-08T12:12:44+00:00",
    "label": "my-label-48770703",
    "plan": "vc2-1c-2gb",
    "status": "active",
    "node_quantity": 1,
	"min_nodes": 1,
	"max_nodes": 2,
	"auto_scaler": true,
	"tag": "mytag",
    "nodes": [
      {
        "id": "f2e11430-76e5-4dc6-a1c9-ef5682c21ddf",
        "label": "my-label-48770703-44060e6384c45",
        "date_created": "2021-07-07T23:27:08+00:00",
        "status": "active"
      }
    ]
  }
}`
		fmt.Fprint(writer, response)
	})
	update := NodePoolReqUpdate{NodeQuantity: 1}
	response, _, err := client.Kubernetes.UpdateNodePool(ctx, "1", "2", &update)
	if err != nil {
		t.Errorf("Kubernetes.UpdateNodePool returned %+v", err)
	}

	expected := &NodePool{
		ID:           "e97bdee9-2781-4f31-be03-60fc75f399ae",
		DateCreated:  "2021-07-07T23:27:08+00:00",
		DateUpdated:  "2021-07-08T12:12:44+00:00",
		Label:        "my-label-48770703",
		Plan:         "vc2-1c-2gb",
		Status:       "active",
		NodeQuantity: 1,
		MinNodes:     1,
		MaxNodes:     2,
		AutoScaler:   true,
		Tag:          "mytag",
		Nodes: []Node{
			{
				ID:          "f2e11430-76e5-4dc6-a1c9-ef5682c21ddf",
				DateCreated: "2021-07-07T23:27:08+00:00",
				Label:       "my-label-48770703-44060e6384c45",
				Status:      "active",
			},
		},
	}

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("Kubernetes.UpdateNodePool meta returned %+v, expected %+v", response, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.UpdateNodePool(c, "1", "2", &update)
	if err == nil {
		t.Error("Kubernetes.UpdateNodePool returned nil")
	}
}

func TestKubernetesHandler_DeleteNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/node-pools/%s", vkePath, "1", "2"), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Kubernetes.DeleteNodePool(ctx, "1", "2")
	if err != nil {
		t.Errorf("Kubernetes.DeleteNodePool returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteNodePoolInstance(t *testing.T) {
	setup()
	defer teardown()
	path := fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s", vkePath, "1", "2", "3")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Kubernetes.DeleteNodePoolInstance(ctx, "1", "2", "3")
	if err != nil {
		t.Errorf("Kubernetes.DeleteNodePoolInstance returned %+v", err)
	}
}

func TestKubernetesHandler_RecycleNodePoolInstance(t *testing.T) {
	setup()
	defer teardown()
	path := fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s/recycle", vkePath, "1", "2", "3")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Kubernetes.RecycleNodePoolInstance(ctx, "1", "2", "3")
	if err != nil {
		t.Errorf("Kubernetes.RecycleNodePoolInstance returned %+v", err)
	}
}

func TestKubernetesHandler_GetKubeConfig(t *testing.T) {
	setup()
	defer teardown()
	path := fmt.Sprintf("%s/%s/config", vkePath, "1")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `{"kube_config": "config="}`
		fmt.Fprint(writer, response)
	})

	config, _, err := client.Kubernetes.GetKubeConfig(ctx, "1")
	if err != nil {
		t.Errorf("Kubernetes.GetKubeConfig returned %+v", err)
	}

	expected := &KubeConfig{KubeConfig: "config="}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetKubeConfig  returned %+v, expected %+v", config, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.GetKubeConfig(c, "1")
	if err == nil {
		t.Error("Kubernetes.GetKubeConfig returned nil")
	}
}

func TestKubernetesHandler_GetVersions(t *testing.T) {
	setup()
	defer teardown()
	path := "/v2/kubernetes/versions"
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `{"versions": ["v1.20.0+1"]}`
		fmt.Fprint(writer, response)
	})

	config, _, err := client.Kubernetes.GetVersions(ctx)
	if err != nil {
		t.Errorf("Kubernetes.GetVersions returned %+v", err)
	}

	expected := &Versions{Versions: []string{"v1.20.0+1"}}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetVersions returned %+v, expected %+v", config, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.GetVersions(c)
	if err == nil {
		t.Error("Kubernetes.GetVersions returned nil")
	}
}

func TestKubernetesHandler_GetUpgrades(t *testing.T) {
	setup()
	defer teardown()
	path := fmt.Sprintf("%s/%s/available-upgrades", vkePath, "1")
	mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		response := `{"available_upgrades": ["v1.20.0+1"]}`
		fmt.Fprint(writer, response)
	})

	config, _, err := client.Kubernetes.GetUpgrades(ctx, "1")
	if err != nil {
		t.Errorf("Kubernetes.GetVersions returned %+v", err)
	}

	expected := []string{"v1.20.0+1"}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetVersions returned %+v, expected %+v", config, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.Kubernetes.GetUpgrades(c, "1")
	if err == nil {
		t.Error("Kubernetes.GetUpgradeVersions returned nil")
	}
}

func TestKubernetesHandler_Upgrade(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/%s/upgrades", vkePath, "1"), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &ClusterUpgradeReq{UpgradeVersion: "2"}
	err := client.Kubernetes.Upgrade(ctx, "1", req)
	if err != nil {
		t.Errorf("Kubernetes.StartUpgrade returned %+v", err)
	}
}
