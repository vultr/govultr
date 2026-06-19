package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClusterServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "cluster": {
			"id": "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
			"region": "ewr",
			"label": "My GPU Cluster",
			"plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			"min_pool_count": 1,
			"desired_pool_count": 2,
			"hostname": "my-cluster",
			"status": "pending",
			"state": "provisioning",
			"date_created": "2026-01-15T10:30:00+00:00",
			"cluster_type": "fabric",
			"type": "cluster"
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &ClusterCreate{
		Region:           "ewr",
		Plan:             "vbm-256c-3072gb-8-mi325x-aac-gpu",
		Label:            "My GPU Cluster",
		MinPoolCount:     1,
		DesiredPoolCount: 2,
		Hostname:         "my-cluster",
	}

	cluster, _, err := client.Cluster.Create(ctx, options)
	if err != nil {
		t.Errorf("Cluster.Create returned %+v", err)
	}

	expected := &InstanceCluster{
		ID:               "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
		Region:           "ewr",
		Label:            "My GPU Cluster",
		Plan:             "vbm-256c-3072gb-8-mi325x-aac-gpu",
		MinPoolCount:     1,
		DesiredPoolCount: 2,
		Hostname:         "my-cluster",
		Status:           "pending",
		State:            "provisioning",
		DateCreated:      "2026-01-15T10:30:00+00:00",
		ClusterType:      "fabric",
		Type:             "cluster",
	}

	if !reflect.DeepEqual(cluster, expected) {
		t.Errorf("Cluster.Create returned %+v, expected %+v", cluster, expected)
	}
}

func TestClusterServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/7a1b2c3d-4e5f-6789-abcd-ef0123456789", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "cluster": {
			"id": "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
			"region": "ewr",
			"label": "My GPU Cluster",
			"plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			"min_pool_count": 1,
			"desired_pool_count": 2,
			"hostname": "my-cluster",
			"status": "active",
			"state": "running",
			"date_created": "2026-01-15T10:30:00+00:00",
			"cluster_type": "fabric",
			"type": "fabric",
			"head_node_instance_template_id": "98765432-10ab-cdef-1234-567890abcdef",
			"instance_template": {
			  "id": "98765432-10ab-cdef-1234-567890abcdef",
			  "plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			  "label": "cluster-template",
			  "os": "Ubuntu 22.04 x64"
			},
			"instances": [
			  {
				"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
				"label": "cluster-node-1",
				"date_created": "2026-01-15T10:32:00+00:00",
				"status": "active",
				"ip_address": "192.0.2.10",
				"hostname": "my-cluster-1"
			  },
			  {
				"id": "dc787b57-77fe-5ecc-c940-554f3f7d1c71",
				"label": "cluster-head-1",
				"date_created": "2026-01-15T10:32:00+00:00",
				"status": "active",
				"ip_address": "192.0.2.11",
				"hostname": "my-cluster-2",
				"is_head_node": true
			  }
			],
			"vpc_networks": [
			  {
				"id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
				"description": "default-vpc"
			  }
			]
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	cluster, _, err := client.Cluster.Get(ctx, "7a1b2c3d-4e5f-6789-abcd-ef0123456789")
	if err != nil {
		t.Errorf("Cluster.Get returned %+v", err)
	}

	expected := &InstanceCluster{
		ID:                         "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
		Region:                     "ewr",
		Label:                      "My GPU Cluster",
		Plan:                       "vbm-256c-3072gb-8-mi325x-aac-gpu",
		MinPoolCount:               1,
		DesiredPoolCount:           2,
		Hostname:                   "my-cluster",
		Status:                     "active",
		State:                      "running",
		DateCreated:                "2026-01-15T10:30:00+00:00",
		ClusterType:                "fabric",
		Type:                       "fabric",
		HeadNodeInstanceTemplateID: "98765432-10ab-cdef-1234-567890abcdef",
		InstanceTemplate: ClusterInstanceTemplate{
			ID:    "98765432-10ab-cdef-1234-567890abcdef",
			Plan:  "vbm-256c-3072gb-8-mi325x-aac-gpu",
			Label: "cluster-template",
			OS:    "Ubuntu 22.04 x64",
		},
		Instances: []ClusterInstance{
			{
				ID:          "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
				Label:       "cluster-node-1",
				DateCreated: "2026-01-15T10:32:00+00:00",
				Status:      "active",
				IPAddress:   "192.0.2.10",
				Hostname:    "my-cluster-1",
			},
			{
				ID:          "dc787b57-77fe-5ecc-c940-554f3f7d1c71",
				Label:       "cluster-head-1",
				DateCreated: "2026-01-15T10:32:00+00:00",
				Status:      "active",
				IPAddress:   "192.0.2.11",
				Hostname:    "my-cluster-2",
				IsHeadNode:  true,
			},
		},
		VPCNetworks: []ClusterVPC{
			{
				ID:          "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
				Description: "default-vpc",
			},
		},
	}

	if !reflect.DeepEqual(cluster, expected) {
		t.Errorf("Cluster.Get returned %+v, expected %+v", cluster, expected)
	}
}

func TestClusterServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/7a1b2c3d-4e5f-6789-abcd-ef0123456789", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "cluster": {
			"id": "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
			"region": "ewr",
			"label": "my new label",
			"plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			"min_pool_count": 1,
			"desired_pool_count": 2,
			"hostname": "my-cluster",
			"status": "active",
			"state": "running",
			"date_created": "2026-01-15T10:30:00+00:00",
			"cluster_type": "fabric",
			"type": "fabric",
			"head_node_instance_template_id": "98765432-10ab-cdef-1234-567890abcdef",
			"instance_template": {
			  "id": "98765432-10ab-cdef-1234-567890abcdef",
			  "plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			  "label": "cluster-template",
			  "os": "Ubuntu 22.04 x64"
			},
			"instances": [
			  {
				"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
				"label": "cluster-node-1",
				"date_created": "2026-01-15T10:32:00+00:00",
				"status": "active",
				"ip_address": "192.0.2.10",
				"hostname": "my-cluster-1"
			  },
			  {
				"id": "dc787b57-77fe-5ecc-c940-554f3f7d1c71",
				"label": "cluster-head-1",
				"date_created": "2026-01-15T10:32:00+00:00",
				"status": "active",
				"ip_address": "192.0.2.11",
				"hostname": "my-cluster-2",
				"is_head_node": true
			  }
			],
			"vpc_networks": [
			  {
				"id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
				"description": "default-vpc"
			  }
			]
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &ClusterUpdate{
		Label: "my new label",
	}

	cluster, _, err := client.Cluster.Update(ctx, "7a1b2c3d-4e5f-6789-abcd-ef0123456789", options)
	if err != nil {
		t.Errorf("BareMetal.Update returned %+v, expected %+v", err, nil)
	}

	expected := &InstanceCluster{
		ID:                         "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
		Region:                     "ewr",
		Label:                      "my new label",
		Plan:                       "vbm-256c-3072gb-8-mi325x-aac-gpu",
		MinPoolCount:               1,
		DesiredPoolCount:           2,
		Hostname:                   "my-cluster",
		Status:                     "active",
		State:                      "running",
		DateCreated:                "2026-01-15T10:30:00+00:00",
		ClusterType:                "fabric",
		Type:                       "fabric",
		HeadNodeInstanceTemplateID: "98765432-10ab-cdef-1234-567890abcdef",
		InstanceTemplate: ClusterInstanceTemplate{
			ID:    "98765432-10ab-cdef-1234-567890abcdef",
			Plan:  "vbm-256c-3072gb-8-mi325x-aac-gpu",
			Label: "cluster-template",
			OS:    "Ubuntu 22.04 x64",
		},
		Instances: []ClusterInstance{
			{
				ID:          "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
				Label:       "cluster-node-1",
				DateCreated: "2026-01-15T10:32:00+00:00",
				Status:      "active",
				IPAddress:   "192.0.2.10",
				Hostname:    "my-cluster-1",
			},
			{
				ID:          "dc787b57-77fe-5ecc-c940-554f3f7d1c71",
				Label:       "cluster-head-1",
				DateCreated: "2026-01-15T10:32:00+00:00",
				Status:      "active",
				IPAddress:   "192.0.2.11",
				Hostname:    "my-cluster-2",
				IsHeadNode:  true,
			},
		},
		VPCNetworks: []ClusterVPC{
			{
				ID:          "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
				Description: "default-vpc",
			},
		},
	}

	if !reflect.DeepEqual(cluster, expected) {
		t.Errorf("Cluster.Get returned %+v, expected %+v", cluster, expected)
	}
}

func TestClusterServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/7a1b2c3d-4e5f-6789-abcd-ef0123456789", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Cluster.Delete(ctx, "7a1b2c3d-4e5f-6789-abcd-ef0123456789")

	if err != nil {
		t.Errorf("Cluster.Delete returned %+v", err)
	}
}

func TestClusterServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "clusters": [
			{
			  "id": "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
			  "region": "ewr",
			  "label": "My GPU Cluster",
			  "plan": "vbm-256c-3072gb-8-mi325x-aac-gpu",
			  "min_pool_count": 1,
			  "desired_pool_count": 2,
			  "hostname": "my-cluster",
			  "status": "active",
			  "state": "running",
			  "date_created": "2026-01-15T10:30:00+00:00",
			  "cluster_type": "fabric",
			  "type": "cluster",
			  "head_node_instance_template_id": "98765432-10ab-cdef-1234-567890abcdef",
			  "instances": [
				{
				  "id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
				  "label": "cluster-node-1",
				  "date_created": "2026-01-15T10:32:00+00:00",
				  "status": "active"
				}
			  ],
			  "vpc_networks": [
				{
				  "id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
				  "description": "default-vpc"
				}
			  ]
			}
		  ],
		  "meta": {
			"total": 1,
			"links": {
			  "next": "",
			  "prev": ""
			}
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	clusters, meta, _, err := client.Cluster.List(ctx, nil)
	if err != nil {
		t.Errorf("Cluster.List returned %+v", err)
	}

	expected := []InstanceCluster{
		{
			ID:                         "7a1b2c3d-4e5f-6789-abcd-ef0123456789",
			Region:                     "ewr",
			Label:                      "My GPU Cluster",
			Plan:                       "vbm-256c-3072gb-8-mi325x-aac-gpu",
			MinPoolCount:               1,
			DesiredPoolCount:           2,
			Hostname:                   "my-cluster",
			Status:                     "active",
			State:                      "running",
			DateCreated:                "2026-01-15T10:30:00+00:00",
			ClusterType:                "fabric",
			Type:                       "cluster",
			HeadNodeInstanceTemplateID: "98765432-10ab-cdef-1234-567890abcdef",
			Instances: []ClusterInstance{
				{
					ID:          "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
					Label:       "cluster-node-1",
					DateCreated: "2026-01-15T10:32:00+00:00",
					Status:      "active",
				},
			},
			VPCNetworks: []ClusterVPC{
				{
					ID:          "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
					Description: "default-vpc",
				},
			},
		},
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Cluster.List returned %+v, expected %+v", clusters, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Cluster.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestClusterServiceHandler_AttachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/attach/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Cluster.AttachInstance(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Cluster.AttachInstance returned %+v", err)
	}
}

func TestClusterServiceHandler_AttachInstances(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Cluster.AttachInstances(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Cluster.AttachInstances returned %+v", err)
	}
}

func TestClusterServiceHandler_AttachHeadNode(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/attach_head_node/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Cluster.AttachHeadNode(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Cluster.AttachHeadNode returned %+v", err)
	}
}

func TestClusterServiceHandler_DetachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/detach/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	if err := client.Cluster.DetachInstance(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", "14b3e7d6-ffb5-4994-8502-57fcd9db3b33"); err != nil {
		t.Errorf("Cluster.DetachInstance returned %+v", err)
	}
}

func TestClusterServiceHandler_DetachInstances(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})
	if err := client.Cluster.DetachInstances(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", []string{"14b3e7d6-ffb5-4994-8502-57fcd9db3b33"}); err != nil {
		t.Errorf("Cluster.DetachInstances returned %+v", err)
	}
}

func TestClusterServiceHandler_GetMetrics(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/clusters/14b3e7d6-ffb5-4994-8502-57fcd9db3b33/metrics", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "metrics": {
			"instances": {
			  "cb676a46-66fd-4dfb-b839-443f2e6c0b60": {
				"gpu": {
				  "temperature": {
					"gpu": [
					  {
						"unit": "celsius",
						"target": "0",
						"datapoints": [
						  [
							42.5,
							1705318800
						  ],
						  [
							43.1,
							1705319100
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					]
				  },
				  "utilization": {
					"gpu": [
					  {
						"unit": "percent",
						"target": "0",
						"datapoints": [
						  [
							78.2,
							1705318800
						  ],
						  [
							85.5,
							1705319100
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					],
					"memory": {
					  "total": [
						{
						  "unit": "megabytes",
						  "target": "0",
						  "datapoints": [
							[
							  81920,
							  1705318800
							]
						  ],
						  "tags": {
							"name": "0"
						  }
						}
					  ],
					  "used": [
						{
						  "unit": "megabytes",
						  "target": "0",
						  "datapoints": [
							[
							  65536,
							  1705318800
							]
						  ],
						  "tags": {
							"name": "0"
						  }
						}
					  ],
					  "free": [
						{
						  "unit": "megabytes",
						  "target": "0",
						  "datapoints": [
							[
							  16384,
							  1705318800
							]
						  ],
						  "tags": {
							"name": "0"
						  }
						}
					  ]
					}
				  },
				  "power": {
					"used": [
					  {
						"unit": "watts",
						"target": "0",
						"datapoints": [
						  [
							320.5,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					],
					"max": [
					  {
						"unit": "watts",
						"target": "0",
						"datapoints": [
						  [
							400,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					]
				  }
				},
				"fabric": {
				  "raw_throughput": {
					"tx": [
					  {
						"unit": "bytes",
						"target": "0",
						"datapoints": [
						  [
							125829120,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					],
					"rx": [
					  {
						"unit": "bytes",
						"target": "0",
						"datapoints": [
						  [
							130023424,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					]
				  },
				  "retries": {
					"tx": [
					  {
						"unit": "count",
						"target": "0",
						"datapoints": [
						  [
							3,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					],
					"rx": [
					  {
						"unit": "count",
						"target": "0",
						"datapoints": [
						  [
							1,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					]
				  },
				  "optical_power_strength": {
					"tx": [
					  {
						"unit": "milliwatts",
						"target": "0",
						"datapoints": [
						  [
							1.02,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					],
					"rx": [
					  {
						"unit": "milliwatts",
						"target": "0",
						"datapoints": [
						  [
							0.98,
							1705318800
						  ]
						],
						"tags": {
						  "name": "0"
						}
					  }
					]
				  }
				}
			  }
			}
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &ClusterMetricsOpts{
		Period: StringToStringPtr("-1days"),
	}

	metrics, _, err := client.Cluster.GetMetrics(ctx, "14b3e7d6-ffb5-4994-8502-57fcd9db3b33", options)
	if err != nil {
		t.Errorf("Cluster.GetMetrics returned %+v", err)
	}

	expected := &ClusterMetrics{
		Instances: map[string]ClusterInstanceMetrics{
			"cb676a46-66fd-4dfb-b839-443f2e6c0b60": {
				GPU: ClusterGPUMetrics{
					Temperature: ClusterGPUTemperatureMetrics{
						GPU: []ClusterMetric{{
							Unit:   "celsius",
							Target: "0",
							Datapoints: [][]float64{
								{42.5, 1705318800},
								{43.1, 1705319100},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
					},
					Utilization: ClusterGPUUtilizationMetrics{
						GPU: []ClusterMetric{{
							Unit:   "percent",
							Target: "0",
							Datapoints: [][]float64{
								{78.2, 1705318800},
								{85.5, 1705319100},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
						Memory: ClusterGPUMemoryMetrics{
							Total: []ClusterMetric{{
								Unit:   "megabytes",
								Target: "0",
								Datapoints: [][]float64{
									{81920, 1705318800},
								},
								Tags: map[string]string{
									"name": "0",
								},
							}},
							Used: []ClusterMetric{{
								Unit:   "megabytes",
								Target: "0",
								Datapoints: [][]float64{
									{65536, 1705318800},
								},
								Tags: map[string]string{
									"name": "0",
								},
							}},
							Free: []ClusterMetric{{
								Unit:   "megabytes",
								Target: "0",
								Datapoints: [][]float64{
									{16384, 1705318800},
								},
								Tags: map[string]string{
									"name": "0",
								},
							}},
						},
					},
					Power: ClusterGPUPowerMetrics{
						Used: []ClusterMetric{{
							Unit:   "watts",
							Target: "0",
							Datapoints: [][]float64{
								{320.5, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
						Max: []ClusterMetric{{
							Unit:   "watts",
							Target: "0",
							Datapoints: [][]float64{
								{400, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
					},
				},
				Fabric: ClusterFabricMetrics{
					RawThroughput: ClusterRxTxMetrics{
						Tx: []ClusterMetric{{
							Unit:   "bytes",
							Target: "0",
							Datapoints: [][]float64{
								{125829120, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
						Rx: []ClusterMetric{{
							Unit:   "bytes",
							Target: "0",
							Datapoints: [][]float64{
								{130023424, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
					},
					Retries: ClusterRxTxMetrics{
						Tx: []ClusterMetric{{
							Unit:   "count",
							Target: "0",
							Datapoints: [][]float64{
								{3, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
						Rx: []ClusterMetric{{
							Unit:   "count",
							Target: "0",
							Datapoints: [][]float64{
								{1, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
					},
					OpticalPowerStrength: ClusterRxTxMetrics{
						Tx: []ClusterMetric{{
							Unit:   "milliwatts",
							Target: "0",
							Datapoints: [][]float64{
								{1.02, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
						Rx: []ClusterMetric{{
							Unit:   "milliwatts",
							Target: "0",
							Datapoints: [][]float64{
								{0.98, 1705318800},
							},
							Tags: map[string]string{
								"name": "0",
							},
						}},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(metrics, expected) {
		t.Errorf("Cluster.GetMetrics returned %+v, expected %+v", metrics, expected)
	}
}
