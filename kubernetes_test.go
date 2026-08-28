package govultr

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestKubernetesHandler_CreateCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/clusters", testJSONResponseHandlerFunc(http.StatusOK, `
{
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
}`))

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

	mux.HandleFunc("/v2/kubernetes/clusters/014da059-21e3-47eb-acb5-91bf697c31aa", testJSONResponseHandlerFunc(http.StatusOK, `
{
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
}`))

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

	mux.HandleFunc("/v2/kubernetes/clusters/", testJSONResponseHandlerFunc(http.StatusOK, `
{
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
}`))

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
		t.Errorf("Kubernetes.List meta returned %+v, expected %+v", meta, expectedMeta)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	if _, _, _, err = client.Kubernetes.ListClusters(c, nil); err == nil {
		t.Error("Kubernetes.ListClusters returned nil")
	}
}

func TestKubernetesHandler_UpdateCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	update := ClusterReqUpdate{Label: "new label"}
	err := client.Kubernetes.UpdateCluster(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", &update)

	if err != nil {
		t.Errorf("Kubernetes.UpdateCluster returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteCluster(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	err := client.Kubernetes.DeleteCluster(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Kubernetes.DeleteCluster returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteClusterWithResources(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/delete-with-linked-resources",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	err := client.Kubernetes.DeleteClusterWithResources(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33")
	if err != nil {
		t.Errorf("Kubernetes.DeleteClusterWithResources returned %+v", err)
	}
}

func TestKubernetesHandler_CreateNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/node-pools",
		testJSONResponseHandlerFunc(http.StatusCreated, `
{
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
		"vpc_only": true,
		"labels": {
			"vultr.com/label1": "value1",
			"vultr.com/label2": "value2"
		},
		"taints": [
			{
				"key": "gpu",
				"value": "test",
				"effect": "NoSchedule"
			}
		],
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "pending"
            }
        ]
    }
}`))

	createReq := &NodePoolReq{
		NodeQuantity: 1,
		Label:        "nodepool-48959140",
		Plan:         "vc2-1c-2gb",
		Tag:          "mytag",
		VPCOnly:      BoolToBoolPtr(true),
		Labels: map[string]string{
			"vultr.com/label1": "value1",
			"vultr.com/label2": "value2",
		},
		Taints: []Taint{
			{
				Key:    "gpu",
				Value:  "test",
				Effect: "NoSchedule",
			},
		},
	}
	np, _, err := client.Kubernetes.CreateNodePool(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", createReq)
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
		VPCOnly:      true,
		Labels: map[string]string{
			"vultr.com/label1": "value1",
			"vultr.com/label2": "value2",
		},
		Taints: []Taint{
			{
				Key:    "gpu",
				Value:  "test",
				Effect: "NoSchedule",
			},
		},
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

	mux.HandleFunc(
		"/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools/13e25ca6-bce9-4703-b8dd-d593b1254f89",
		testJSONResponseHandlerFunc(http.StatusOK, `
{
    "node_pool": {
        "id": "13e25ca6-bce9-4703-b8dd-d593b1254f89",
        "date_created": "2021-07-13T15:42:21+00:00",
        "label": "nodepool-48959140",
        "plan": "vc2-2c-4gb",
        "status": "active",
        "node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
		"taints": [
			{
				"key": "gpu",
				"value": "test",
				"effect": "NoSchedule"
			}
		],
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "active"
            }
        ]
    }
}`))

	np, _, err := client.Kubernetes.GetNodePool(ctx, "e97bdee9-2781-4f31-be03-60fc75f399ae", "13e25ca6-bce9-4703-b8dd-d593b1254f89")
	if err != nil {
		t.Errorf("Kubernetes.GetNodePool returned %v", err)
	}

	expected := &NodePool{
		ID:           "13e25ca6-bce9-4703-b8dd-d593b1254f89",
		DateCreated:  "2021-07-13T15:42:21+00:00",
		Label:        "nodepool-48959140",
		Plan:         "vc2-2c-4gb",
		Status:       "active",
		Tag:          "mytag",
		NodeQuantity: 1,
		MinNodes:     1,
		MaxNodes:     2,
		AutoScaler:   true,
		Taints: []Taint{
			{
				Key:    "gpu",
				Value:  "test",
				Effect: "NoSchedule",
			},
		},
		Nodes: []Node{
			{
				ID:          "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
				Label:       "nodepool-48959140-74a60edb45de0",
				DateCreated: "2021-07-13T15:42:21+00:00",
				Status:      "active",
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

	mux.HandleFunc("/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools", testJSONResponseHandlerFunc(http.StatusOK, `
{
    "node_pools": [{
        "id": "554e7248-705a-5862-516f-4f4a6735346a",
        "date_created": "2021-07-13T15:42:21+00:00",
        "label": "nodepool-48959140",
        "plan": "vc2-2c-4gb",
        "status": "active",
        "node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
		"taints": [
			{
				"key": "gpu",
				"value": "test",
				"effect": "NoSchedule"
			}
		],
        "nodes": [
            {
                "id": "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
                "label": "nodepool-48959140-74a60edb45de0",
                "date_created": "2021-07-13T15:42:21+00:00",
                "status": "active"
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
}`))

	np, meta, _, err := client.Kubernetes.ListNodePools(ctx, "e97bdee9-2781-4f31-be03-60fc75f399ae", nil)
	if err != nil {
		t.Errorf("Kubernetes.ListNodePools returned %v", err)
	}

	expected := []NodePool{
		{
			ID:           "554e7248-705a-5862-516f-4f4a6735346a",
			DateCreated:  "2021-07-13T15:42:21+00:00",
			Label:        "nodepool-48959140",
			Plan:         "vc2-2c-4gb",
			Status:       "active",
			Tag:          "mytag",
			NodeQuantity: 1,
			MinNodes:     1,
			MaxNodes:     2,
			AutoScaler:   true,
			Taints: []Taint{
				{
					Key:    "gpu",
					Value:  "test",
					Effect: "NoSchedule",
				},
			},
			Nodes: []Node{
				{
					ID:          "3e1ca1e0-25be-4977-907a-3dee42b9bb15",
					Label:       "nodepool-48959140-74a60edb45de0",
					DateCreated: "2021-07-13T15:42:21+00:00",
					Status:      "active",
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
	if _, _, _, err = client.Kubernetes.ListNodePools(c, "1", nil); err == nil {
		t.Error("Kubernetes.ListNodePools returned nil")
	}
}

func TestKubernetesHandler_UpdateNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools/554e7248-705a-5862-516f-4f4a6735346a",
		testJSONResponseHandlerFunc(http.StatusAccepted, `
{
	"node_pool": {
		"id": "554e7248-705a-5862-516f-4f4a6735346a",
		"date_created": "2021-07-07T23:27:08+00:00",
		"date_updated": "2021-07-08T12:12:44+00:00",
		"label": "my-label-48770703",
		"plan": "vc2-2c-4gb",
		"status": "active",
		"node_quantity": 1,
		"min_nodes": 1,
		"max_nodes": 2,
		"auto_scaler": true,
		"tag": "mytag",
		"taints": [
			{
			"key": "gpu",
			"value": "updated-test",
			"effect": "NoSchedule"
			}
		],
		"nodes": [
			{
			"id": "f2e11430-76e5-4dc6-a1c9-ef5682c21ddf",
			"label": "my-label-48770703-44060e6384c45",
			"date_created": "2021-07-07T23:27:08+00:00",
			"status": "active"
			}
		]
	}
}`))

	taints := []Taint{
		{
			Key:    "gpu",
			Value:  "updated-test",
			Effect: "NoSchedule",
		},
	}

	update := NodePoolReqUpdate{
		NodeQuantity: 1,
		Taints:       taints,
	}

	response, _, err := client.Kubernetes.UpdateNodePool(
		ctx,
		"e97bdee9-2781-4f31-be03-60fc75f399ae",
		"554e7248-705a-5862-516f-4f4a6735346a",
		&update,
	)
	if err != nil {
		t.Errorf("Kubernetes.UpdateNodePool returned %+v", err)
	}

	expected := &NodePool{
		ID:           "554e7248-705a-5862-516f-4f4a6735346a",
		DateCreated:  "2021-07-07T23:27:08+00:00",
		DateUpdated:  "2021-07-08T12:12:44+00:00",
		Label:        "my-label-48770703",
		Plan:         "vc2-2c-4gb",
		Status:       "active",
		NodeQuantity: 1,
		MinNodes:     1,
		MaxNodes:     2,
		AutoScaler:   true,
		Tag:          "mytag",
		Taints:       taints,
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
}

func TestKubernetesHandler_DeleteNodePool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools/554e7248-705a-5862-516f-4f4a6735346a",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	err := client.Kubernetes.DeleteNodePool(ctx, "e97bdee9-2781-4f31-be03-60fc75f399ae", "554e7248-705a-5862-516f-4f4a6735346a")
	if err != nil {
		t.Errorf("Kubernetes.DeleteNodePool returned %+v", err)
	}
}

func TestKubernetesHandler_DeleteNodePoolInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools/554e7248-705a-5862-516f-4f4a6735346a/nodes/73e3c9c5-ba24-45a5-ab42-a63c254c5e44",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	err := client.Kubernetes.DeleteNodePoolInstance(
		ctx,
		"e97bdee9-2781-4f31-be03-60fc75f399ae",
		"554e7248-705a-5862-516f-4f4a6735346a",
		"73e3c9c5-ba24-45a5-ab42-a63c254c5e44",
	)
	if err != nil {
		t.Errorf("Kubernetes.DeleteNodePoolInstance returned %+v", err)
	}
}

func TestKubernetesHandler_RecycleNodePoolInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/node-pools/554e7248-705a-5862-516f-4f4a6735346a/nodes/73e3c9c5-ba24-45a5-ab42-a63c254c5e44/recycle",
		testJSONResponseHandlerFunc(http.StatusNoContent, ""),
	)

	err := client.Kubernetes.RecycleNodePoolInstance(
		ctx,
		"e97bdee9-2781-4f31-be03-60fc75f399ae",
		"554e7248-705a-5862-516f-4f4a6735346a",
		"73e3c9c5-ba24-45a5-ab42-a63c254c5e44",
	)
	if err != nil {
		t.Errorf("Kubernetes.RecycleNodePoolInstance returned %+v", err)
	}
}

func TestKubernetesHandler_GetKubeConfig(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/clusters/e97bdee9-2781-4f31-be03-60fc75f399ae/config", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"kube_config": "YXBpdmVyc2lvbjogdjEKY2x1c3RlcnM6Ci0gY2x1c3RlcjoKICAgIGNlcnRpZmljYXRlLWF1dGhvcml0eS1kYXRhOiBMUzB0TFMxQ1JVZEpUaUJEUlZKVVNVWkpRMEZVUlMwdExTMHRDazFKU1VSblZFTkRRVzF0WjBGM1NVSkJaMGxKU2psdFN6bEViRk5pY0RSM1JGRlpTa3R2V2tsb2RtTk9RVkZGVEVKUlFYZFVla1ZNVFVGclIwRXhWVVVLUW1oTlExWldUWGhHYWtGVlFtZE9Wa0pCWTFSRVZrNW9ZbWxDUjJOdFJuVlpNbXg2V1RJNGVFVjZRVkpDWjA1V1FrRnZWRU5yZERGWmJWWjVZbTFXTUFwYVdFMTRSWHBCVWtKblRsWkNRVTFVUTJ0ME1WbHRWbmxpYlZZd1dsaE5kMGhvWTA1TmFrVjNUbnBCZVUxVVNYaE5la0Y1VjJoalRrMXFTWGRPZWtGNUNrMVVTWGhOZWtGNVYycENVRTFSYzNkRFVWbEVWbEZSUjBWM1NsWlZla1ZYVFVKUlIwRXhWVVZDZUUxT1ZUSkdkVWxGV25sWlZ6VnFZVmhPYW1KNlJWUUtUVUpGUjBFeFZVVkRhRTFMVXpOV2FWcFlTblZhV0ZKc1kzcEZWRTFDUlVkQk1WVkZRWGhOUzFNelZtbGFXRXAxV2xoU2JHTjZRME5CVTBsM1JGRlpTZ3BMYjFwSmFIWmpUa0ZSUlVKQ1VVRkVaMmRGVUVGRVEwTkJVVzlEWjJkRlFrRk1laTlITXpOaVlWZG5TMU5GVmpKQ2RsQlhZbWd6WkhZclYybEhOVlJqQ2s1bllVTlZNMlJWVm5KdGNtaHVXbVJPYWtkTVl5OUJTR3RIWm1OaVIxQlRXbkJ2UVZCbWFuaGtWRTA0WlVOTFlXdGxkR0Z6YkRsdFNDOVhlVTlETXpnS1pGcEZVWGRSZWpseFIzWnpaa3BTT0RKQ01WWTBWM3AxUVdRMEwxSmtaVGxqU3psaVdIWktkRUZMU2xrNVF6aG9VM2RtTDNNM1drRlNabGxYYTIxb1R3cHZkSHBFUnpaR2JtaFljSFJtUkRZdmRXNXNXRWhyYTNveFVHSjZhR1Z2ZG5adU9GUkNUamR2UWpkTVdUaG9kRE5tVTBJeFEwSlRTMGxxV1hsaGJEaHJDbU5XZVU1R1MyUndVRVoxV0ZvelkyYzRaMHN2ZUROS1VXZHBLMGxqVTA0Mk9GZDJaa3gxYXpjM2JXOXNWWE5IY25neWNWRkZTa2RwU2k5SGEzUm5kMndLWlVnNU1FbHpMMkZDZERjd1dsaG9aM2cyTnpkSmVuTnVWMnN3UWpWSFlVeGpaRk5oYUN0WlNraGthbkpMU0N0R01qVTNWekpvTUVOQmQwVkJRV0ZPYUFwTlJqaDNSR2RaUkZaU01GQkJVVWd2UWtGUlJFRm5TMFZOUWpCSFFURlZaRXBSVVZkTlFsRkhRME56UjBGUlZVWkNkMDFEUW1kbmNrSm5SVVpDVVdORUNrRlVRVkJDWjA1V1NGSk5Ra0ZtT0VWQ1ZFRkVRVkZJTDAxQ01FZEJNVlZrUkdkUlYwSkNVME5EVWtoSmFERm1XbnBzU210MFMwVmtOalpITVZWWGRqQUtNRVJCVGtKbmEzRm9hMmxIT1hjd1FrRlJjMFpCUVU5RFFWRkZRV0V3VG1SUVlYa3dPREp0YVcxWllUa3ZOVVpMY1hWa1YwSmpabVpHVkVScFdrTmljUXA1YVUxNFZXeEVTQzl6Tm1Od1YzbEJORXRuY0ZGRWMySXdiM0pzYTNwTk1ERjNieTlsTUc1clUxTTFVVkIyWVZvNU9FaFNObFlyTUV4a0swZzViM1JCQ2xZM2VUbEdlQzlJVUhCdldGWTJhVWswYWpCaVpFdFBNMHQ0VUZKVWRsaDFRMUZETTNRd2FHc3pjVnBRSzFSalNEaHhWRTV6VkVwb1JGTnlSMWRLUjNvS1dqZ3liMGwwY0c5RVRsaEJZVUpqYmxSRmNUUkNXRzFoTTJVNVJHSkpVMU5SZW5aaGFIYzBWMkZwVTFWNGVYUllVakJ5Um1oaFpFUnpkbFJuVVhZNGF3cEdlbkV5TjNkS2RUaHZUV2hIWTJWb2VGUlRXVUpyWjNGWVUzYzNPR2xsTVZadk1XVlBMMGxTYlhsM1ZtMWlhM2M1TWswcldtZFdOV0Z3VERCNlNYRnNDbFEzWmtkekszWTViREkwYkM4eGFIbExVekZCU1ZKTmVrRkljMGw1YVdWdE1GUkZUM0Z6WVVVNVFYWjBlWEZZZEZKblBUMEtMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0KICAgIHNlcnZlcjogaHR0cHM6Ly9jOTA3ZTgzMi0zMDgwLTQ4YTYtYTU0ZC03Mzc5ZTY0NWMwYjcudnVsdHItazhzLmNvbTo2NDQzCiAgbmFtZTogdmtlCmNvbnRleHRzOgotIGNvbnRleHQ6CiAgICBjbHVzdGVyOiB2a2UKICAgIHVzZXI6IGFkbWluCiAgbmFtZTogdmtlCmN1cnJlbnQtY29udGV4dDogdmtlCmtpbmQ6IENvbmZpZwpwcmVmZXJlbmNlczoge30KdXNlcnM6Ci0gbmFtZTogYWRtaW4KICB1c2VyOgogICAgY2xpZW50LWNlcnRpZmljYXRlLWRhdGE6IExTMHRMUzFDUlVkSlRpQkRSVkpVU1VaSlEwRlVSUzB0TFMwdENrMUpTVVJWUkVORFFXcHBaMEYzU1VKQlowbEpUVmg0VTFOSGRFRnliR2QzUkZGWlNrdHZXa2xvZG1OT1FWRkZURUpSUVhkVWVrVk1UVUZyUjBFeFZVVUtRbWhOUTFaV1RYaEdha0ZWUW1kT1ZrSkJZMVJFVms1b1ltbENSMk50Um5WWk1teDZXVEk0ZUVWNlFWSkNaMDVXUWtGdlZFTnJkREZaYlZaNVltMVdNQXBhV0UxNFJYcEJVa0puVGxaQ1FVMVVRMnQwTVZsdFZubGliVll3V2xoTmQwaG9ZMDVOYWtWM1RucEJlVTFVU1hoTmVrRjVWMmhqVGsxcVNYZE9la0Y1Q2sxVVNYaE5la0Y1VjJwQ1QwMVJjM2REVVZsRVZsRlJSMFYzU2xaVmVrVlhUVUpSUjBFeFZVVkNlRTFPVlRKR2RVbEZXbmxaVnpWcVlWaE9hbUo2UlZnS1RVSlZSMEV4VlVWRGFFMVBZek5zZW1SSFZuUlBiVEZvWXpOU2JHTnVUWGhFYWtGTlFtZE9Wa0pCVFZSQ1YwWnJZbGRzZFUxSlNVSkpha0ZPUW1kcmNRcG9hMmxIT1hjd1FrRlJSVVpCUVU5RFFWRTRRVTFKU1VKRFowdERRVkZGUVhselRIVndNSHBvYXpsUFVHODVWa05TTUZSbmJ6UTFORThyV0hOTVEyUXhDbE5CWVdKNmFtMVJaM1pEVVZKeFdEaEZUa0Z0VW5kbVdFUjNaRkJMWTFkbmFtcHpRaTlQU2pSR2F6TmpWWFZIVVdkNmFrRkRXVVJYVjNBM1RWaG1TM1VLVm5GeVNGTmtZMnhQWVV0dEwwbGpNMEkxWVd0a1pYcGxRVFJ4UzFGRlRrbFVSbXR1VkdSWVJ6RTFVV3MxU2tNMGNIWXpaa3M1ZUhVMldqZHhjVmRXVlFwdmVFMXdjR2huV1hGWFVsUkNSMnByT0hSRk5sbDZOazVZZGs5NkwxVXpNWEprV0ZOVFluYzRWakpxTUdnNU1FTlRMMkZLVkN0U01sRmxNRWh3YkZNeUNsSjBWek0yYlRjMFVGaHpXRGQ2Ym1aTVZWZEpaMGQxYjBvNVdYTkJNRFphUTFSVllrdFNTekV2V0haRmFGZHVPSGRtWTFCblRHTXlRWEJRTnpsMVlYa0taV0phZVV4SmFXOWFXRXRNVERWQ05tcEZaVkZWV2pGWlRFTjNSV0pCTXpWYVdYSm1lRTVCUmsxcFUwcDFTMnhhUlRWSGNYRlJTVVJCVVVGQ2IzcEZkd3BNZWtGUFFtZE9Wa2hST0VKQlpqaEZRa0ZOUTBGdlVYZElVVmxFVmxJd2JFSkNXWGRHUVZsSlMzZFpRa0pSVlVoQmQwbEhRME56UjBGUlZVWkNkMDFDQ2sxQk1FZERVM0ZIVTBsaU0wUlJSVUpEZDFWQlFUUkpRa0ZSUWpWbUwwdHJVVGxRV1d4WE1uQllUek13V1dSYVZHMUlhbWRhTm10RlFUUmhVelJvVWs4S2NqSldSbHBwUjBoUVluZGFZMjVuZFc1UVRXTnJaRmh2UWs5a1dsVkhkelpoYkUxaVFVOUZhRlpIUVVOSU5IcEhkM2RUUlZrMk5HRTJVV0ZsVFVaSWF3cHZkalU1UW1GclJIZFJkVlprTVdoMk1rcFZkMXB3WTFsTVZUZE5PWGRLWTI5a09FODBNM0EyVGxwTmNrVjBObHB2YmtsSWJGbEpkMGhFTWxWaGVYcHZDamhUVkhWeWNXVm5jakJvYzAwd1ltWlFRbkZzY25CdE9VTXZOV2hVVjJVemJ6STJiRFpNUTBabWFFdzBaamN5VURSaWFYWnNkVTVoYVc5UFp6QXZXVVlLZFVwd09WUjZkMnRuUWtSVE9DOU5hVTFUVDFwSFpVdHlia2hWYlhKa2FGbHpSbTFCVVRCRVRYWlJiMnh1TWtwVlRYSXlkWE4yU0VGcFJGWm9PVkZMWlFwM1lrSlRMMlJ3UW04M09UbEZRWHBpWkdaclpIcG5iVWhDU2k5WU4wVjNNR3B4Tm5Nek5YTkRNMUpqY0dNNFJrd0tMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0KICAgIGNsaWVudC1rZXktZGF0YTogTFMwdExTMUNSVWRKVGlCU1UwRWdVRkpKVmtGVVJTQkxSVmt0TFMwdExRcE5TVWxGYjJkSlFrRkJTME5CVVVWQmVYTk1kWEF3ZW1ock9VOVFiemxXUTFJd1ZHZHZORFUwVHl0WWMweERaREZUUVdGaWVtcHRVV2QyUTFGU2NWZzRDa1ZPUVcxU2QyWllSSGRrVUV0alYyZHFhbk5DTDA5S05FWnJNMk5WZFVkUlozcHFRVU5aUkZkWGNEZE5XR1pMZFZaeGNraFRaR05zVDJGTGJTOUpZek1LUWpWaGEyUmxlbVZCTkhGTFVVVk9TVlJHYTI1VVpGaEhNVFZSYXpWS1F6Undkak5tU3psNGRUWmFOM0Z4VjFaVmIzaE5jSEJvWjFseFYxSlVRa2RxYXdvNGRFVTJXWG8yVGxoMlQzb3ZWVE14Y21SWVUxTmlkemhXTW1vd2FEa3dRMU12WVVwVUsxSXlVV1V3U0hCc1V6SlNkRmN6Tm0wM05GQlljMWczZW01bUNreFZWMGxuUjNWdlNqbFpjMEV3TmxwRFZGVmlTMUpMTVM5WWRrVm9WMjQ0ZDJaalVHZE1ZekpCY0ZBM09YVmhlV1ZpV25sTVNXbHZXbGhMVEV3MVFqWUtha1ZsVVZWYU1WbE1RM2RGWWtFek5WcFpjbVo0VGtGR1RXbFRTblZMYkZwRk5VZHhjVkZKUkVGUlFVSkJiMGxDUVVKaWN6VXpUQzlJUm5CTmFESmpjd3A1Tm5WdVVFRmpRMHQwU1VzNGVVMXBObll6VkRCWVdVWjVSRTFzTlVGdk5EVnJSVFJhTjNWTlVsZExjbTUxV0VsT1NtdG5WSFE0Tmpndk1FSnVURWMyQ2xVd05tazRaMDlvUkRWME4ySlFkMHRZYlM5eFN6RktUVUY1WkRkSWIzQmhPVE4yYVV0dlNYa3pMemxwWjB4Tk0yRkZkRXB2Vlc5S2NUaDJaMDFxVDNRS2MxWk5aMVJWVmpKVVVYZGFUR056ZEdFNU5YTlphamh1V214S2QyczNhSHBFTmtFemNUSTBhRVJ0YUU1a2FUZ3dSSEJEVDJjMk1IbFpTaXQ2Y0dab1dBcHZORkJPTlhsTVZGaFhkSG80SzJ0UllqaDZaR3B0Wms5a1pHVnJaeTlOTTA1T1pUY3JPRVZHV2tJMWExWkRkbEV3UmxoSVIxRlZPREUzV1dNdlNEZHZDamhpVFZsM2QySlZRbEppTkdobmVWWnZkemxVV2s5YVkwMXVMM2xoYVd3M1JtVmljblpGU2s5dVJtdG9Wa3BoUlRKdVlWSlJjbEJ1TmtORWRVdEdXRGdLUTJNM2JHZGhSVU5uV1VWQmVuZEVURlpFY1UxWlYxcHNiWFJUYzJGdWVGaEJhMnBwUW5GVWFreFhTbGRIZVZSTk0xZEtUM2RzZG1GWk9FTlRiSHB5TkFwelMyOTNOMXByT0doQ1ZXSm1WRFJCV1ZWTlpVRlBSV0p2UkhCTVpYQTBURmxYU0hSUmJuQkpWakY1ZURSMVdWZHlSM28yZVhSemVGbE1MMnRvYTBSMENtUnVRVTFDVTBOdlZXOW1VVWczSzBWa2NHdFpZbGxrYW05bFZuWXdlalpPVG1SQ1ZYSk9VbFpQWlc1NFR6WkllR1JIZDNwSFUxVkRaMWxGUVN0elJXVUtVMk5wU0RSbWIybDVkRlZNWW5VMVJVVmtaM0YwTDFseUszTmtVMjl6TURGTWVtOXpSVU4wZGpFMGJrNHZWVlpIYTI5WE5UUTRVVEpOYmtkeUswSTFLd3BvUkRCME1XTXpXQ3RPT1dOc2FYRXdWVzVGZVhad09DOXhWblJUZGxSYWFscEdOVTFKTXpadWJqWXhVVkkxV0hOcFNEVmhWbWQyUlRoYWNFNTBSbHBsQ21sWlRVNHpRM2R6VjNCb1drMUJTV2hVZDI1S1RVOVZlVXhHVlhWT05rTjNhelJFUlhacVZVTm5XVUZ2UWs5R1MxVldPV1ZZVTNRM1pWYzBlakJCU1VzS1VWRnFhR1V2VFc1cVVVWlZhMmhNUWtsbldsUTRUMjlTY1hWTmMwNVplSEZ4ZUhneVkzTXJUMUZZTld4QmFESlZjME5OVjNwSE5Va3hZbmhPTkd0M1ZRbzRXbVZ4TmtoVlpqTndNMDVZWlc4MFVEUkdhM2R0WlZSaGMxaEdXbkozUW5vM2RXcExhVTFuWjFnd2RFSm9kVmg1YUZVNE5UVllUbU5OZG1weVVUUnFDblo2YWk5cFRGWktWWFkzSzBwc2NrWjRWREZFZFZGTFFtZEhZbkJrVmxoUk1FUjJiRmhtZDJrellrNXJXa1pzZVdZeU0wdDNXbWRHTUVKMGVUTjJORWtLUVU5bFRpOUVVRmx2WTNKMFkwbzFXRGxqWms0d2JsSXphRW95Y2xocU9VWnZTbU5qUTFOU2FDOHllU3RGUVRJMU1scHhNRmgyZEZKMVN6Vm9UVkFyVkFwV0szQkdXU3RSTjNwRVVHeERlRTlzZWpoME5uWjRVblJSVEVSbWRHRk1ORkF5Y2paVmFWaEpWM1Z3UTBwWmREZDBOaXQ1YTFKdWVYZzJiMkUzTHpGS0NtNTJWakZCYjBkQlRtRkZSV1VyZERoNGFXUlFaeXRIUm5FemN6VkhUa1pPU0VKQmMzcG1NV3BxVFdwdmF6ZEtiMmxyWWxaVU1FNUJWVmhLZEdWclNra0taekZPVUd0UVQxSjRWakZWY0dKUGVsUXJhMWhyTjBwVmExTTNiVE0zV0VneWRuUm1abVIxTlhFNVEzWkhObkZaYW1OcVpVOHZSRWhzYUZWRmNrMUpXUXBsVHpSQk9FOUNWRWgwUVRkM01XTjVVa2hIVUZWUE1VRlFiRXRTTWtWVU9XMXVOeXROV0hSRU1GQmxjRWh2TVV0UU9VRTlDaTB0TFMwdFJVNUVJRkpUUVNCUVVrbFdRVlJGSUV0RldTMHRMUzB0Q2c9PQo"
}`))

	config, _, err := client.Kubernetes.GetKubeConfig(ctx, "e97bdee9-2781-4f31-be03-60fc75f399ae")
	if err != nil {
		t.Errorf("Kubernetes.GetKubeConfig returned %+v", err)
	}

	expected := &KubeConfig{
		KubeConfig: "YXBpdmVyc2lvbjogdjEKY2x1c3RlcnM6Ci0gY2x1c3RlcjoKICAgIGNlcnRpZmljYXRlLWF1dGhvcml0eS1kYXRhOiBMUzB0TFMxQ1JVZEpUaUJEUlZKVVNVWkpRMEZVUlMwdExTMHRDazFKU1VSblZFTkRRVzF0WjBGM1NVSkJaMGxKU2psdFN6bEViRk5pY0RSM1JGRlpTa3R2V2tsb2RtTk9RVkZGVEVKUlFYZFVla1ZNVFVGclIwRXhWVVVLUW1oTlExWldUWGhHYWtGVlFtZE9Wa0pCWTFSRVZrNW9ZbWxDUjJOdFJuVlpNbXg2V1RJNGVFVjZRVkpDWjA1V1FrRnZWRU5yZERGWmJWWjVZbTFXTUFwYVdFMTRSWHBCVWtKblRsWkNRVTFVUTJ0ME1WbHRWbmxpYlZZd1dsaE5kMGhvWTA1TmFrVjNUbnBCZVUxVVNYaE5la0Y1VjJoalRrMXFTWGRPZWtGNUNrMVVTWGhOZWtGNVYycENVRTFSYzNkRFVWbEVWbEZSUjBWM1NsWlZla1ZYVFVKUlIwRXhWVVZDZUUxT1ZUSkdkVWxGV25sWlZ6VnFZVmhPYW1KNlJWUUtUVUpGUjBFeFZVVkRhRTFMVXpOV2FWcFlTblZhV0ZKc1kzcEZWRTFDUlVkQk1WVkZRWGhOUzFNelZtbGFXRXAxV2xoU2JHTjZRME5CVTBsM1JGRlpTZ3BMYjFwSmFIWmpUa0ZSUlVKQ1VVRkVaMmRGVUVGRVEwTkJVVzlEWjJkRlFrRk1laTlITXpOaVlWZG5TMU5GVmpKQ2RsQlhZbWd6WkhZclYybEhOVlJqQ2s1bllVTlZNMlJWVm5KdGNtaHVXbVJPYWtkTVl5OUJTR3RIWm1OaVIxQlRXbkJ2UVZCbWFuaGtWRTA0WlVOTFlXdGxkR0Z6YkRsdFNDOVhlVTlETXpnS1pGcEZVWGRSZWpseFIzWnpaa3BTT0RKQ01WWTBWM3AxUVdRMEwxSmtaVGxqU3psaVdIWktkRUZMU2xrNVF6aG9VM2RtTDNNM1drRlNabGxYYTIxb1R3cHZkSHBFUnpaR2JtaFljSFJtUkRZdmRXNXNXRWhyYTNveFVHSjZhR1Z2ZG5adU9GUkNUamR2UWpkTVdUaG9kRE5tVTBJeFEwSlRTMGxxV1hsaGJEaHJDbU5XZVU1R1MyUndVRVoxV0ZvelkyYzRaMHN2ZUROS1VXZHBLMGxqVTA0Mk9GZDJaa3gxYXpjM2JXOXNWWE5IY25neWNWRkZTa2RwU2k5SGEzUm5kMndLWlVnNU1FbHpMMkZDZERjd1dsaG9aM2cyTnpkSmVuTnVWMnN3UWpWSFlVeGpaRk5oYUN0WlNraGthbkpMU0N0R01qVTNWekpvTUVOQmQwVkJRV0ZPYUFwTlJqaDNSR2RaUkZaU01GQkJVVWd2UWtGUlJFRm5TMFZOUWpCSFFURlZaRXBSVVZkTlFsRkhRME56UjBGUlZVWkNkMDFEUW1kbmNrSm5SVVpDVVdORUNrRlVRVkJDWjA1V1NGSk5Ra0ZtT0VWQ1ZFRkVRVkZJTDAxQ01FZEJNVlZrUkdkUlYwSkNVME5EVWtoSmFERm1XbnBzU210MFMwVmtOalpITVZWWGRqQUtNRVJCVGtKbmEzRm9hMmxIT1hjd1FrRlJjMFpCUVU5RFFWRkZRV0V3VG1SUVlYa3dPREp0YVcxWllUa3ZOVVpMY1hWa1YwSmpabVpHVkVScFdrTmljUXA1YVUxNFZXeEVTQzl6Tm1Od1YzbEJORXRuY0ZGRWMySXdiM0pzYTNwTk1ERjNieTlsTUc1clUxTTFVVkIyWVZvNU9FaFNObFlyTUV4a0swZzViM1JCQ2xZM2VUbEdlQzlJVUhCdldGWTJhVWswYWpCaVpFdFBNMHQ0VUZKVWRsaDFRMUZETTNRd2FHc3pjVnBRSzFSalNEaHhWRTV6VkVwb1JGTnlSMWRLUjNvS1dqZ3liMGwwY0c5RVRsaEJZVUpqYmxSRmNUUkNXRzFoTTJVNVJHSkpVMU5SZW5aaGFIYzBWMkZwVTFWNGVYUllVakJ5Um1oaFpFUnpkbFJuVVhZNGF3cEdlbkV5TjNkS2RUaHZUV2hIWTJWb2VGUlRXVUpyWjNGWVUzYzNPR2xsTVZadk1XVlBMMGxTYlhsM1ZtMWlhM2M1TWswcldtZFdOV0Z3VERCNlNYRnNDbFEzWmtkekszWTViREkwYkM4eGFIbExVekZCU1ZKTmVrRkljMGw1YVdWdE1GUkZUM0Z6WVVVNVFYWjBlWEZZZEZKblBUMEtMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0KICAgIHNlcnZlcjogaHR0cHM6Ly9jOTA3ZTgzMi0zMDgwLTQ4YTYtYTU0ZC03Mzc5ZTY0NWMwYjcudnVsdHItazhzLmNvbTo2NDQzCiAgbmFtZTogdmtlCmNvbnRleHRzOgotIGNvbnRleHQ6CiAgICBjbHVzdGVyOiB2a2UKICAgIHVzZXI6IGFkbWluCiAgbmFtZTogdmtlCmN1cnJlbnQtY29udGV4dDogdmtlCmtpbmQ6IENvbmZpZwpwcmVmZXJlbmNlczoge30KdXNlcnM6Ci0gbmFtZTogYWRtaW4KICB1c2VyOgogICAgY2xpZW50LWNlcnRpZmljYXRlLWRhdGE6IExTMHRMUzFDUlVkSlRpQkRSVkpVU1VaSlEwRlVSUzB0TFMwdENrMUpTVVJWUkVORFFXcHBaMEYzU1VKQlowbEpUVmg0VTFOSGRFRnliR2QzUkZGWlNrdHZXa2xvZG1OT1FWRkZURUpSUVhkVWVrVk1UVUZyUjBFeFZVVUtRbWhOUTFaV1RYaEdha0ZWUW1kT1ZrSkJZMVJFVms1b1ltbENSMk50Um5WWk1teDZXVEk0ZUVWNlFWSkNaMDVXUWtGdlZFTnJkREZaYlZaNVltMVdNQXBhV0UxNFJYcEJVa0puVGxaQ1FVMVVRMnQwTVZsdFZubGliVll3V2xoTmQwaG9ZMDVOYWtWM1RucEJlVTFVU1hoTmVrRjVWMmhqVGsxcVNYZE9la0Y1Q2sxVVNYaE5la0Y1VjJwQ1QwMVJjM2REVVZsRVZsRlJSMFYzU2xaVmVrVlhUVUpSUjBFeFZVVkNlRTFPVlRKR2RVbEZXbmxaVnpWcVlWaE9hbUo2UlZnS1RVSlZSMEV4VlVWRGFFMVBZek5zZW1SSFZuUlBiVEZvWXpOU2JHTnVUWGhFYWtGTlFtZE9Wa0pCVFZSQ1YwWnJZbGRzZFUxSlNVSkpha0ZPUW1kcmNRcG9hMmxIT1hjd1FrRlJSVVpCUVU5RFFWRTRRVTFKU1VKRFowdERRVkZGUVhselRIVndNSHBvYXpsUFVHODVWa05TTUZSbmJ6UTFORThyV0hOTVEyUXhDbE5CWVdKNmFtMVJaM1pEVVZKeFdEaEZUa0Z0VW5kbVdFUjNaRkJMWTFkbmFtcHpRaTlQU2pSR2F6TmpWWFZIVVdkNmFrRkRXVVJYVjNBM1RWaG1TM1VLVm5GeVNGTmtZMnhQWVV0dEwwbGpNMEkxWVd0a1pYcGxRVFJ4UzFGRlRrbFVSbXR1VkdSWVJ6RTFVV3MxU2tNMGNIWXpaa3M1ZUhVMldqZHhjVmRXVlFwdmVFMXdjR2huV1hGWFVsUkNSMnByT0hSRk5sbDZOazVZZGs5NkwxVXpNWEprV0ZOVFluYzRWakpxTUdnNU1FTlRMMkZLVkN0U01sRmxNRWh3YkZNeUNsSjBWek0yYlRjMFVGaHpXRGQ2Ym1aTVZWZEpaMGQxYjBvNVdYTkJNRFphUTFSVllrdFNTekV2V0haRmFGZHVPSGRtWTFCblRHTXlRWEJRTnpsMVlYa0taV0phZVV4SmFXOWFXRXRNVERWQ05tcEZaVkZWV2pGWlRFTjNSV0pCTXpWYVdYSm1lRTVCUmsxcFUwcDFTMnhhUlRWSGNYRlJTVVJCVVVGQ2IzcEZkd3BNZWtGUFFtZE9Wa2hST0VKQlpqaEZRa0ZOUTBGdlVYZElVVmxFVmxJd2JFSkNXWGRHUVZsSlMzZFpRa0pSVlVoQmQwbEhRME56UjBGUlZVWkNkMDFDQ2sxQk1FZERVM0ZIVTBsaU0wUlJSVUpEZDFWQlFUUkpRa0ZSUWpWbUwwdHJVVGxRV1d4WE1uQllUek13V1dSYVZHMUlhbWRhTm10RlFUUmhVelJvVWs4S2NqSldSbHBwUjBoUVluZGFZMjVuZFc1UVRXTnJaRmh2UWs5a1dsVkhkelpoYkUxaVFVOUZhRlpIUVVOSU5IcEhkM2RUUlZrMk5HRTJVV0ZsVFVaSWF3cHZkalU1UW1GclJIZFJkVlprTVdoMk1rcFZkMXB3WTFsTVZUZE5PWGRLWTI5a09FODBNM0EyVGxwTmNrVjBObHB2YmtsSWJGbEpkMGhFTWxWaGVYcHZDamhUVkhWeWNXVm5jakJvYzAwd1ltWlFRbkZzY25CdE9VTXZOV2hVVjJVemJ6STJiRFpNUTBabWFFdzBaamN5VURSaWFYWnNkVTVoYVc5UFp6QXZXVVlLZFVwd09WUjZkMnRuUWtSVE9DOU5hVTFUVDFwSFpVdHlia2hWYlhKa2FGbHpSbTFCVVRCRVRYWlJiMnh1TWtwVlRYSXlkWE4yU0VGcFJGWm9PVkZMWlFwM1lrSlRMMlJ3UW04M09UbEZRWHBpWkdaclpIcG5iVWhDU2k5WU4wVjNNR3B4Tm5Nek5YTkRNMUpqY0dNNFJrd0tMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0KICAgIGNsaWVudC1rZXktZGF0YTogTFMwdExTMUNSVWRKVGlCU1UwRWdVRkpKVmtGVVJTQkxSVmt0TFMwdExRcE5TVWxGYjJkSlFrRkJTME5CVVVWQmVYTk1kWEF3ZW1ock9VOVFiemxXUTFJd1ZHZHZORFUwVHl0WWMweERaREZUUVdGaWVtcHRVV2QyUTFGU2NWZzRDa1ZPUVcxU2QyWllSSGRrVUV0alYyZHFhbk5DTDA5S05FWnJNMk5WZFVkUlozcHFRVU5aUkZkWGNEZE5XR1pMZFZaeGNraFRaR05zVDJGTGJTOUpZek1LUWpWaGEyUmxlbVZCTkhGTFVVVk9TVlJHYTI1VVpGaEhNVFZSYXpWS1F6Undkak5tU3psNGRUWmFOM0Z4VjFaVmIzaE5jSEJvWjFseFYxSlVRa2RxYXdvNGRFVTJXWG8yVGxoMlQzb3ZWVE14Y21SWVUxTmlkemhXTW1vd2FEa3dRMU12WVVwVUsxSXlVV1V3U0hCc1V6SlNkRmN6Tm0wM05GQlljMWczZW01bUNreFZWMGxuUjNWdlNqbFpjMEV3TmxwRFZGVmlTMUpMTVM5WWRrVm9WMjQ0ZDJaalVHZE1ZekpCY0ZBM09YVmhlV1ZpV25sTVNXbHZXbGhMVEV3MVFqWUtha1ZsVVZWYU1WbE1RM2RGWWtFek5WcFpjbVo0VGtGR1RXbFRTblZMYkZwRk5VZHhjVkZKUkVGUlFVSkJiMGxDUVVKaWN6VXpUQzlJUm5CTmFESmpjd3A1Tm5WdVVFRmpRMHQwU1VzNGVVMXBObll6VkRCWVdVWjVSRTFzTlVGdk5EVnJSVFJhTjNWTlVsZExjbTUxV0VsT1NtdG5WSFE0Tmpndk1FSnVURWMyQ2xVd05tazRaMDlvUkRWME4ySlFkMHRZYlM5eFN6RktUVUY1WkRkSWIzQmhPVE4yYVV0dlNYa3pMemxwWjB4Tk0yRkZkRXB2Vlc5S2NUaDJaMDFxVDNRS2MxWk5aMVJWVmpKVVVYZGFUR056ZEdFNU5YTlphamh1V214S2QyczNhSHBFTmtFemNUSTBhRVJ0YUU1a2FUZ3dSSEJEVDJjMk1IbFpTaXQ2Y0dab1dBcHZORkJPTlhsTVZGaFhkSG80SzJ0UllqaDZaR3B0Wms5a1pHVnJaeTlOTTA1T1pUY3JPRVZHV2tJMWExWkRkbEV3UmxoSVIxRlZPREUzV1dNdlNEZHZDamhpVFZsM2QySlZRbEppTkdobmVWWnZkemxVV2s5YVkwMXVMM2xoYVd3M1JtVmljblpGU2s5dVJtdG9Wa3BoUlRKdVlWSlJjbEJ1TmtORWRVdEdXRGdLUTJNM2JHZGhSVU5uV1VWQmVuZEVURlpFY1UxWlYxcHNiWFJUYzJGdWVGaEJhMnBwUW5GVWFreFhTbGRIZVZSTk0xZEtUM2RzZG1GWk9FTlRiSHB5TkFwelMyOTNOMXByT0doQ1ZXSm1WRFJCV1ZWTlpVRlBSV0p2UkhCTVpYQTBURmxYU0hSUmJuQkpWakY1ZURSMVdWZHlSM28yZVhSemVGbE1MMnRvYTBSMENtUnVRVTFDVTBOdlZXOW1VVWczSzBWa2NHdFpZbGxrYW05bFZuWXdlalpPVG1SQ1ZYSk9VbFpQWlc1NFR6WkllR1JIZDNwSFUxVkRaMWxGUVN0elJXVUtVMk5wU0RSbWIybDVkRlZNWW5VMVJVVmtaM0YwTDFseUszTmtVMjl6TURGTWVtOXpSVU4wZGpFMGJrNHZWVlpIYTI5WE5UUTRVVEpOYmtkeUswSTFLd3BvUkRCME1XTXpXQ3RPT1dOc2FYRXdWVzVGZVhad09DOXhWblJUZGxSYWFscEdOVTFKTXpadWJqWXhVVkkxV0hOcFNEVmhWbWQyUlRoYWNFNTBSbHBsQ21sWlRVNHpRM2R6VjNCb1drMUJTV2hVZDI1S1RVOVZlVXhHVlhWT05rTjNhelJFUlhacVZVTm5XVUZ2UWs5R1MxVldPV1ZZVTNRM1pWYzBlakJCU1VzS1VWRnFhR1V2VFc1cVVVWlZhMmhNUWtsbldsUTRUMjlTY1hWTmMwNVplSEZ4ZUhneVkzTXJUMUZZTld4QmFESlZjME5OVjNwSE5Va3hZbmhPTkd0M1ZRbzRXbVZ4TmtoVlpqTndNMDVZWlc4MFVEUkdhM2R0WlZSaGMxaEdXbkozUW5vM2RXcExhVTFuWjFnd2RFSm9kVmg1YUZVNE5UVllUbU5OZG1weVVUUnFDblo2YWk5cFRGWktWWFkzSzBwc2NrWjRWREZFZFZGTFFtZEhZbkJrVmxoUk1FUjJiRmhtZDJrellrNXJXa1pzZVdZeU0wdDNXbWRHTUVKMGVUTjJORWtLUVU5bFRpOUVVRmx2WTNKMFkwbzFXRGxqWms0d2JsSXphRW95Y2xocU9VWnZTbU5qUTFOU2FDOHllU3RGUVRJMU1scHhNRmgyZEZKMVN6Vm9UVkFyVkFwV0szQkdXU3RSTjNwRVVHeERlRTlzZWpoME5uWjRVblJSVEVSbWRHRk1ORkF5Y2paVmFWaEpWM1Z3UTBwWmREZDBOaXQ1YTFKdWVYZzJiMkUzTHpGS0NtNTJWakZCYjBkQlRtRkZSV1VyZERoNGFXUlFaeXRIUm5FemN6VkhUa1pPU0VKQmMzcG1NV3BxVFdwdmF6ZEtiMmxyWWxaVU1FNUJWVmhLZEdWclNra0taekZPVUd0UVQxSjRWakZWY0dKUGVsUXJhMWhyTjBwVmExTTNiVE0zV0VneWRuUm1abVIxTlhFNVEzWkhObkZaYW1OcVpVOHZSRWhzYUZWRmNrMUpXUXBsVHpSQk9FOUNWRWgwUVRkM01XTjVVa2hIVUZWUE1VRlFiRXRTTWtWVU9XMXVOeXROV0hSRU1GQmxjRWh2TVV0UU9VRTlDaTB0TFMwdFJVNUVJRkpUUVNCUVVrbFdRVlJGSUV0RldTMHRMUzB0Q2c9PQo",
	}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetKubeConfig  returned %+v, expected %+v", config, expected)
	}
}

func TestKubernetesHandler_GetVersions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/versions", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"versions": [
		"v1.20.0+1"
	]
}`))

	config, _, err := client.Kubernetes.GetVersions(ctx)
	if err != nil {
		t.Errorf("Kubernetes.GetVersions returned %+v", err)
	}

	expected := &Versions{Versions: []string{"v1.20.0+1"}}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetVersions returned %+v, expected %+v", config, expected)
	}
}

func TestKubernetesHandler_GetUpgrades(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/kubernetes/clusters/79832407-63ff-4a67-bedb-f3dfe6b8f75f/available-upgrades", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"available_upgrades": [
		"v1.20.0+1"
	]
}`))

	config, _, err := client.Kubernetes.GetUpgrades(ctx, "79832407-63ff-4a67-bedb-f3dfe6b8f75f")
	if err != nil {
		t.Errorf("Kubernetes.GetVersions returned %+v", err)
	}

	expected := []string{"v1.20.0+1"}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Kubernetes.GetVersions returned %+v, expected %+v", config, expected)
	}
}

func TestKubernetesHandler_Upgrade(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(
		"/v2/kubernetes/clusters/79832407-63ff-4a67-bedb-f3dfe6b8f75f/upgrades",
		testJSONResponseHandlerFunc(http.StatusAccepted, ""),
	)

	req := &ClusterUpgradeReq{UpgradeVersion: "v1.22.8+3"}
	err := client.Kubernetes.Upgrade(ctx, "79832407-63ff-4a67-bedb-f3dfe6b8f75f", req)
	if err != nil {
		t.Errorf("Kubernetes.StartUpgrade returned %+v", err)
	}
}
