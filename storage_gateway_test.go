package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStorageGatewayServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "storage_gateway": {
			"id": "string",
			"date_created": "string",
			"status": "string",
			"type": "string",
			"region": "string",
			"label": "string",
			"pending_charges": 0,
			"tags": [
			  "string"
			],
			"health": "string",
			"network_config": {
			  "primary": {
				"ipv4_public_enabled": true,
				"ipv6_public_enabled": true,
				"vpc": {
				  "vpc_ip_address": "string",
				  "vpc_uuid": "string",
				  "vpc_description": "string"
				}
			  }
			},
			"export_config": {
			  "label": "string",
			  "vfs_uuid": "string",
			  "pseudo_root_path": "string",
			  "allowed_ips": [
				"string"
			  ]
			}
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &StorageGatewayCreateReq{
		Label:  "string",
		Type:   "nfs4",
		Region: "string",
		ExportConfig: StorageGatewayExportConfig{
			Label:          "string",
			VfsUuid:        "string",
			PseudoRootPath: "string",
			AllowedIPs: []string{
				"string",
			},
		},
		NetworkConfig: StorageGatewayNetworkConfig{
			Primary: &StorageGatewayNetworkConfigPrimary{
				IPv4PublicEnabled: true,
				IPv6PublicEnabled: true,
				VPC: &StorageGatewayNetworkConfigVPC{
					IPAddress:   "string",
					UUID:        "string",
					Description: "string",
				},
			},
		},
	}

	gateway, _, err := client.StorageGateway.Create(ctx, req)
	if err != nil {
		t.Errorf("StorageGateway.Create returned %+v", err)
	}

	expected := &StorageGateway{
		ID:             "string",
		DateCreated:    "string",
		Status:         "string",
		Type:           "string",
		Region:         "string",
		Label:          "string",
		PendingCharges: 0,
		Tags: []string{
			"string",
		},
		Health: "string",
		NetworkConfig: &StorageGatewayNetworkConfig{
			Primary: &StorageGatewayNetworkConfigPrimary{
				IPv4PublicEnabled: true,
				IPv6PublicEnabled: true,
				VPC: &StorageGatewayNetworkConfigVPC{
					IPAddress:   "string",
					UUID:        "string",
					Description: "string",
				},
			},
		},
		ExportConfig: &StorageGatewayExportConfig{
			Label:          "string",
			VfsUuid:        "string",
			PseudoRootPath: "string",
			AllowedIPs: []string{
				"string",
			},
		},
	}

	if !reflect.DeepEqual(gateway, expected) {
		t.Errorf("StorageGateway.Create returned %+v, expected %+v", gateway, expected)
	}
}

func TestStorageGatewayServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "storage_gateway": {
			  "id": "string",
			  "date_created": "string",
			  "status": "string",
			  "type": "string",
			  "region": "string",
			  "label": "string",
			  "pending_charges": 0,
			  "tags": [
				"string"
			  ],
			  "health": "string",
			  "network_config": {
				"primary": {
				  "ipv4_public_enabled": true,
				  "ipv6_public_enabled": true,
				  "vpc": {
					"vpc_ip_address": "string",
					"vpc_uuid": "string",
					"vpc_description": "string"
				  }
				}
			  },
			  "export_config": {
				"label": "string",
				"vfs_uuid": "string",
				"pseudo_root_path": "string",
				"allowed_ips": [
				  "string"
				]
			  }
		  }
		}`
		fmt.Fprint(writer, response)
	})

	gateways, _, err := client.StorageGateway.Get(ctx, "abcdefg")
	if err != nil {
		t.Errorf("StorageGateway.Get returned %+v", err)
	}

	expected := &StorageGateway{
		ID:             "string",
		DateCreated:    "string",
		Status:         "string",
		Type:           "string",
		Region:         "string",
		Label:          "string",
		PendingCharges: 0,
		Tags: []string{
			"string",
		},
		Health: "string",
		NetworkConfig: &StorageGatewayNetworkConfig{
			Primary: &StorageGatewayNetworkConfigPrimary{
				IPv4PublicEnabled: true,
				IPv6PublicEnabled: true,
				VPC: &StorageGatewayNetworkConfigVPC{
					IPAddress:   "string",
					UUID:        "string",
					Description: "string",
				},
			},
		},
		ExportConfig: &StorageGatewayExportConfig{
			Label:          "string",
			VfsUuid:        "string",
			PseudoRootPath: "string",
			AllowedIPs: []string{
				"string",
			},
		},
	}

	if !reflect.DeepEqual(gateways, expected) {
		t.Errorf("StorageGateway.Get returned %+v, expected %+v", gateways, expected)
	}
}

func TestStorageGatewayServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "storage_gateway": {
			"id": "string",
			"date_created": "string",
			"status": "string",
			"type": "string",
			"region": "string",
			"label": "Updated Storage Gateway Label",
			"pending_charges": 0,
			"tags": [
			  "string"
			],
			"health": "string",
			"network_config": {
			  "primary": {
				"ipv4_public_enabled": true,
				"ipv6_public_enabled": true,
				"vpc": {
				  "vpc_ip_address": "string",
				  "vpc_uuid": "string",
				  "vpc_description": "string"
				}
			  }
			},
			"export_config": {
			  "label": "string",
			  "vfs_uuid": "string",
			  "pseudo_root_path": "string",
			  "allowed_ips": [
				"string"
			  ]
			}
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &StorageGatewayUpdateReq{
		Label: "Updated Storage Gateway Label",
	}

	gateway, _, err := client.StorageGateway.Update(ctx, "abcdefg", req)
	if err != nil {
		t.Errorf("StorageGateway.Update returned %+v", err)
	}

	expected := &StorageGateway{
		ID:             "string",
		DateCreated:    "string",
		Status:         "string",
		Type:           "string",
		Region:         "string",
		Label:          "Updated Storage Gateway Label",
		PendingCharges: 0,
		Tags: []string{
			"string",
		},
		Health: "string",
		NetworkConfig: &StorageGatewayNetworkConfig{
			Primary: &StorageGatewayNetworkConfigPrimary{
				IPv4PublicEnabled: true,
				IPv6PublicEnabled: true,
				VPC: &StorageGatewayNetworkConfigVPC{
					IPAddress:   "string",
					UUID:        "string",
					Description: "string",
				},
			},
		},
		ExportConfig: &StorageGatewayExportConfig{
			Label:          "string",
			VfsUuid:        "string",
			PseudoRootPath: "string",
			AllowedIPs: []string{
				"string",
			},
		},
	}

	if !reflect.DeepEqual(gateway, expected) {
		t.Errorf("StorageGateway.Update returned %+v, expected %+v", gateway, expected)
	}
}

func TestStorageGatewayServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "storage_gateways": [
			{
			  "id": "string",
			  "date_created": "string",
			  "status": "string",
			  "type": "string",
			  "region": "string",
			  "label": "string",
			  "pending_charges": 0,
			  "tags": [
				"string"
			  ],
			  "health": "string",
			  "network_config": {
				"primary": {
				  "ipv4_public_enabled": true,
				  "ipv6_public_enabled": true,
				  "vpc": {
					"vpc_ip_address": "string",
					"vpc_uuid": "string",
					"vpc_description": "string"
				  }
				}
			  },
			  "export_config": {
				"label": "string",
				"vfs_uuid": "string",
				"pseudo_root_path": "string",
				"allowed_ips": [
				  "string"
				]
			  }
			}
		  ],
		  "meta": {
			"total": 0,
			"links": {
			  "next": "string",
			  "prev": "string"
			}
		  }
		}`
		fmt.Fprint(writer, response)
	})

	gateways, meta, _, err := client.StorageGateway.List(ctx, nil)
	if err != nil {
		t.Errorf("StorageGateway.List returned %+v", err)
	}

	expected := []StorageGateway{
		{
			ID:             "string",
			DateCreated:    "string",
			Status:         "string",
			Type:           "string",
			Region:         "string",
			Label:          "string",
			PendingCharges: 0,
			Tags: []string{
				"string",
			},
			Health: "string",
			NetworkConfig: &StorageGatewayNetworkConfig{
				Primary: &StorageGatewayNetworkConfigPrimary{
					IPv4PublicEnabled: true,
					IPv6PublicEnabled: true,
					VPC: &StorageGatewayNetworkConfigVPC{
						IPAddress:   "string",
						UUID:        "string",
						Description: "string",
					},
				},
			},
			ExportConfig: &StorageGatewayExportConfig{
				Label:          "string",
				VfsUuid:        "string",
				PseudoRootPath: "string",
				AllowedIPs: []string{
					"string",
				},
			},
		},
	}

	if !reflect.DeepEqual(gateways, expected) {
		t.Errorf("StorageGateway.List returned %+v, expected %+v", gateways, expected)
	}

	expectedMeta := &Meta{
		Total: 0,
		Links: &Links{
			Next: "string",
			Prev: "string",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("StorageGateway.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestStorageGatewayServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.StorageGateway.Delete(ctx, "abcdefg")
	if err != nil {
		t.Errorf("StorageGateway.Delete returned %+v", err)
	}
}

func TestStorageGatewayServiceHandler_CreateExport(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg/exports", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "label": "string",
		  "vfs_uuid": "string",
		  "pseudo_root_path": "string",
		  "allowed_ips": [
			"string"
		  ]
		}`
		fmt.Fprint(writer, response)
	})

	req := &StorageGatewayExportConfig{
		Label:          "string",
		VfsUuid:        "string",
		PseudoRootPath: "string",
		AllowedIPs: []string{
			"string",
		},
	}

	export, _, err := client.StorageGateway.CreateExport(ctx, "abcdefg", req)
	if err != nil {
		t.Errorf("StorageGateway.CreateExport returned %+v", err)
	}

	expected := &StorageGatewayExportConfig{
		Label:          "string",
		VfsUuid:        "string",
		PseudoRootPath: "string",
		AllowedIPs: []string{
			"string",
		},
	}

	if !reflect.DeepEqual(export, expected) {
		t.Errorf("StorageGateway.CreateExport returned %+v, expected %+v", export, expected)
	}
}

func TestStorageGatewayServiceHandler_UpdateExport(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg/exports/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &StorageGatewayExportUpdateReq{
		AllowedIPs: []string{
			"string",
		},
	}

	err := client.StorageGateway.UpdateExport(ctx, "abcdefg", "abc123", req)
	if err != nil {
		t.Errorf("StorageGateway.UpdateExport returned %+v", err)
	}
}

func TestStorageGatewayServiceHandler_DeleteExport(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/storage-gateways/abcdefg/exports/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.StorageGateway.DeleteExport(ctx, "abcdefg", "abc123")
	if err != nil {
		t.Errorf("StorageGateway.DeleteExport returned %+v", err)
	}
}
