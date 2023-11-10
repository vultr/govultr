package govultr

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestVCRServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(vcrPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "id": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
    "name": "govultrtest",
    "region": "sjc",
    "urn": "sjc.vultrcr.com/govultrtest",
    "storage": {
        "used": {
            "updated_at": "2023-11-09 13:37:12",
            "bytes": 0,
            "mb": 0,
            "gb": 0,
            "tb": 0
        },
        "allowed": {
            "bytes": 21474836480,
            "mb": 20480,
            "gb": 20,
            "tb": 0.02
        }
    },
    "date_created": "2023-11-09 13:37:12",
    "public": false,
    "root_user": {
        "id": 635,
        "username": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
        "password": "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
        "root": true,
        "added_at": "2023-11-09 13:37:12",
        "updated_at": "2023-11-09 13:37:12"
    },
    "metadata": {
        "region": {
            "id": 3,
            "name": "sjc",
            "urn": "sjc.vultrcr.com",
            "base_url": "https://sjc.vultrcr.com",
            "public": true,
            "added_at": "2023-09-14 09:09:16",
            "updated_at": "2023-09-14 09:09:16",
            "data_center": {
                "id": 12,
                "name": "Silicon Valley",
                "site_code": "SJC2",
                "region": "West",
                "country": "US",
                "continent": "North America",
                "description": "Silicon Valley, California",
                "airport": "SJC"
            }
        },
        "subscription": {
            "billing": {
                "monthly_price": 5,
                "pending_charges": 0
            }
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	req := &ContainerRegistryReq{
		Name:   "govultrtest",
		Public: false,
		Region: "sjc",
		Plan:   "business",
	}

	vcr, _, err := client.ContainerRegistry.Create(ctx, req)
	if err != nil {
		t.Errorf("ContainerRegistry.Create returned %v", err)
	}

	expected := &ContainerRegistry{
		ID:   "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
		Name: "govultrtest",
		URN:  "sjc.vultrcr.com/govultrtest",
		Storage: ContainerRegistryStorage{
			Used: ContainerRegistryStorageCount{
				Bytes:        0,
				MegaBytes:    0,
				GigaBytes:    0,
				TeraBytes:    0,
				DateModified: "2023-11-09 13:37:12",
			},
			Allowed: ContainerRegistryStorageCount{
				Bytes:        21474836480,
				MegaBytes:    20480,
				GigaBytes:    20,
				TeraBytes:    0.02,
				DateModified: "",
			},
		},
		DateCreated: "2023-11-09 13:37:12",
		Public:      false,
		RootUser: ContainerRegistryUser{
			ID:           635,
			UserName:     "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
			Password:     "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
			Root:         true,
			DateCreated:  "2023-11-09 13:37:12",
			DateModified: "2023-11-09 13:37:12",
		},
		Metadata: ContainerRegistryMetadata{
			Region: ContainerRegistryRegion{
				ID:           3,
				Name:         "sjc",
				URN:          "sjc.vultrcr.com",
				BaseURL:      "https://sjc.vultrcr.com",
				Public:       true,
				DateCreated:  "2023-09-14 09:09:16",
				DateModified: "2023-09-14 09:09:16",
				DataCenter: ContainerRegistryRegionDataCenter{
					ID:          12,
					Name:        "Silicon Valley",
					SiteCode:    "SJC2",
					Region:      "West",
					Country:     "US",
					Continent:   "North America",
					Description: "Silicon Valley, California",
					Airport:     "SJC",
				},
			},
			Subscription: ContainerRegistrySubscription{
				Billing: ContainerRegistrySubscriptionBilling{
					MonthlyPrice:   5,
					PendingCharges: 0,
				},
			},
		},
	}

	if !reflect.DeepEqual(vcr, expected) {
		t.Errorf("ContainerRegistry.Create returned %+v, expected %+v", vcr, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.ContainerRegistry.Create(c, req)
	if err == nil {
		t.Error("ContainerRegistry.Create returned nil")
	}
}

func TestVCRServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(vcrListPath, func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "registries": [
        {
            "id": "297cfaff-cb5a-4b7c-ac2e-407fcdbd643e",
            "name": "govultrtest",
            "region": "sjc",
            "urn": "sjc.vultrcr.com/govultrtest",
            "storage": {
                "used": {
                    "updated_at": "2023-11-09 23:09:24",
                    "bytes": 0,
                    "mb": 0,
                    "gb": 0,
                    "tb": 0
                },
                "allowed": {
                    "bytes": 21474836480,
                    "mb": 20480,
                    "gb": 20,
                    "tb": 0.02
                }
            },
            "date_created": "2023-11-09 17:32:17",
            "public": true,
            "root_user": {
                "id": 639,
                "username": "297cfaff-cb5a-4b7c-ac2e-407fcdbd643e",
                "password": "Je2S9SkjrowwMtP933SSaxZG4BPR7D8Au33P",
                "root": true,
                "added_at": "2023-11-09 17:32:17",
                "updated_at": "2023-11-09 17:32:17"
            },
            "metadata": {
                "region": {
                    "id": 3,
                    "name": "sjc",
                    "urn": "sjc.vultrcr.com",
                    "base_url": "https://sjc.vultrcr.com",
                    "public": true,
                    "added_at": "2023-09-14 09:09:16",
                    "updated_at": "2023-09-14 09:09:16",
                    "data_center": {
                        "id": 12,
                        "name": "Silicon Valley",
                        "site_code": "SJC2",
                        "region": "West",
                        "country": "US",
                        "continent": "North America",
                        "description": "Silicon Valley, California",
                        "airport": "SJC"
                    }
                },
                "subscription": {
                    "billing": {
                        "monthly_price": 5,
                        "pending_charges": 0.01
                    }
                }
            }
        },
        {
            "id": "c247a5d7-b3e1-468c-bcba-c23d0716ffc8",
            "name": "govultrtest2",
            "region": "sjc",
            "urn": "sjc.vultrcr.com/govultrtest2",
            "storage": {
                "used": {
                    "updated_at": "2023-11-09 23:09:24",
                    "bytes": 0,
                    "mb": 0,
                    "gb": 0,
                    "tb": 0
                },
                "allowed": {
                    "bytes": 21474836480,
                    "mb": 20480,
                    "gb": 20,
                    "tb": 0.02
                }
            },
            "date_created": "2023-11-09 17:33:13",
            "public": true,
            "root_user": {
                "id": 640,
                "username": "c247a5d7-b3e1-468c-bcba-c23d0716ffc8",
                "password": "c9NhkeH7aeF7zj3cRFHbMFizxEik4rhWYGdW",
                "root": true,
                "added_at": "2023-11-09 17:33:13",
                "updated_at": "2023-11-09 17:33:13"
            },
            "metadata": {
                "region": {
                    "id": 3,
                    "name": "sjc",
                    "urn": "sjc.vultrcr.com",
                    "base_url": "https://sjc.vultrcr.com",
                    "public": true,
                    "added_at": "2023-09-14 09:09:16",
                    "updated_at": "2023-09-14 09:09:16",
                    "data_center": {
                        "id": 12,
                        "name": "Silicon Valley",
                        "site_code": "SJC2",
                        "region": "West",
                        "country": "US",
                        "continent": "North America",
                        "description": "Silicon Valley, California",
                        "airport": "SJC"
                    }
                },
                "subscription": {
                    "billing": {
                        "monthly_price": 5,
                        "pending_charges": 0.01
                    }
                }
            }
        }
    ],
    "meta": {
        "total": 2,
        "links": {
            "next": "",
            "prev": ""
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vcrs, meta, _, err := client.ContainerRegistry.List(ctx, nil)
	if err != nil {
		t.Errorf("ContainerRegistry.List returned %+v", err)
	}

	expected := []ContainerRegistry{
		{
			ID:   "297cfaff-cb5a-4b7c-ac2e-407fcdbd643e",
			Name: "govultrtest",
			URN:  "sjc.vultrcr.com/govultrtest",
			Storage: ContainerRegistryStorage{
				Used: ContainerRegistryStorageCount{
					Bytes:        0,
					MegaBytes:    0,
					GigaBytes:    0,
					TeraBytes:    0,
					DateModified: "2023-11-09 23:09:24",
				},
				Allowed: ContainerRegistryStorageCount{
					Bytes:        21474836480,
					MegaBytes:    20480,
					GigaBytes:    20,
					TeraBytes:    0.02,
					DateModified: "",
				},
			},
			DateCreated: "2023-11-09 17:32:17",
			Public:      true,
			RootUser: ContainerRegistryUser{
				ID:           639,
				UserName:     "297cfaff-cb5a-4b7c-ac2e-407fcdbd643e",
				Password:     "Je2S9SkjrowwMtP933SSaxZG4BPR7D8Au33P",
				Root:         true,
				DateCreated:  "2023-11-09 17:32:17",
				DateModified: "2023-11-09 17:32:17",
			},
			Metadata: ContainerRegistryMetadata{
				Region: ContainerRegistryRegion{
					ID:           3,
					Name:         "sjc",
					URN:          "sjc.vultrcr.com",
					BaseURL:      "https://sjc.vultrcr.com",
					Public:       true,
					DateCreated:  "2023-09-14 09:09:16",
					DateModified: "2023-09-14 09:09:16",
					DataCenter: ContainerRegistryRegionDataCenter{
						ID:          12,
						Name:        "Silicon Valley",
						SiteCode:    "SJC2",
						Region:      "West",
						Country:     "US",
						Continent:   "North America",
						Description: "Silicon Valley, California",
						Airport:     "SJC",
					},
				},
				Subscription: ContainerRegistrySubscription{
					Billing: ContainerRegistrySubscriptionBilling{
						MonthlyPrice:   5,
						PendingCharges: 0.01,
					},
				},
			},
		},
		{
			ID:   "c247a5d7-b3e1-468c-bcba-c23d0716ffc8",
			Name: "govultrtest2",
			URN:  "sjc.vultrcr.com/govultrtest2",
			Storage: ContainerRegistryStorage{
				Used: ContainerRegistryStorageCount{
					Bytes:        0,
					MegaBytes:    0,
					GigaBytes:    0,
					TeraBytes:    0,
					DateModified: "2023-11-09 23:09:24",
				},
				Allowed: ContainerRegistryStorageCount{
					Bytes:        21474836480,
					MegaBytes:    20480,
					GigaBytes:    20,
					TeraBytes:    0.02,
					DateModified: "",
				},
			},
			DateCreated: "2023-11-09 17:33:13",
			Public:      true,
			RootUser: ContainerRegistryUser{
				ID:           640,
				UserName:     "c247a5d7-b3e1-468c-bcba-c23d0716ffc8",
				Password:     "c9NhkeH7aeF7zj3cRFHbMFizxEik4rhWYGdW",
				Root:         true,
				DateCreated:  "2023-11-09 17:33:13",
				DateModified: "2023-11-09 17:33:13",
			},
			Metadata: ContainerRegistryMetadata{
				Region: ContainerRegistryRegion{
					ID:           3,
					Name:         "sjc",
					URN:          "sjc.vultrcr.com",
					BaseURL:      "https://sjc.vultrcr.com",
					Public:       true,
					DateCreated:  "2023-09-14 09:09:16",
					DateModified: "2023-09-14 09:09:16",
					DataCenter: ContainerRegistryRegionDataCenter{
						ID:          12,
						Name:        "Silicon Valley",
						SiteCode:    "SJC2",
						Region:      "West",
						Country:     "US",
						Continent:   "North America",
						Description: "Silicon Valley, California",
						Airport:     "SJC",
					},
				},
				Subscription: ContainerRegistrySubscription{
					Billing: ContainerRegistrySubscriptionBilling{
						MonthlyPrice:   5,
						PendingCharges: 0.01,
					},
				},
			},
		},
	}

	expectedMeta := &Meta{
		Total: 2,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(vcrs, expected) {
		t.Errorf("ContainerRegistry.List returned %+v, expected %+v", vcrs, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ContainerRegistry.List meta returned %+v, expected %+v", meta, expectedMeta)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	if _, _, _, err = client.ContainerRegistry.List(c, nil); err == nil {
		t.Error("ContainerRegistry.List returned nil")
	}
}

func TestVCRServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"

	mux.HandleFunc(fmt.Sprintf("%s/%s", vcrPath, vcrID), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "id": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
    "name": "govultrtest",
    "region": "sjc",
    "urn": "sjc.vultrcr.com/govultrtest",
    "storage": {
        "used": {
            "updated_at": "2023-11-09 13:37:12",
            "bytes": 0,
            "mb": 0,
            "gb": 0,
            "tb": 0
        },
        "allowed": {
            "bytes": 21474836480,
            "mb": 20480,
            "gb": 20,
            "tb": 0.02
        }
    },
    "date_created": "2023-11-09 13:37:12",
    "public": false,
    "root_user": {
        "id": 635,
        "username": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
        "password": "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
        "root": true,
        "added_at": "2023-11-09 13:37:12",
        "updated_at": "2023-11-09 13:37:12"
    },
    "metadata": {
        "region": {
            "id": 3,
            "name": "sjc",
            "urn": "sjc.vultrcr.com",
            "base_url": "https://sjc.vultrcr.com",
            "public": true,
            "added_at": "2023-09-14 09:09:16",
            "updated_at": "2023-09-14 09:09:16",
            "data_center": {
                "id": 12,
                "name": "Silicon Valley",
                "site_code": "SJC2",
                "region": "West",
                "country": "US",
                "continent": "North America",
                "description": "Silicon Valley, California",
                "airport": "SJC"
            }
        },
        "subscription": {
            "billing": {
                "monthly_price": 5,
                "pending_charges": 0
            }
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vcr, _, err := client.ContainerRegistry.Get(ctx, vcrID)
	if err != nil {
		t.Errorf("ContainerRegistry.Get returned %v", err)
	}

	expected := &ContainerRegistry{
		ID:   "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
		Name: "govultrtest",
		URN:  "sjc.vultrcr.com/govultrtest",
		Storage: ContainerRegistryStorage{
			Used: ContainerRegistryStorageCount{
				Bytes:        0,
				MegaBytes:    0,
				GigaBytes:    0,
				TeraBytes:    0,
				DateModified: "2023-11-09 13:37:12",
			},
			Allowed: ContainerRegistryStorageCount{
				Bytes:        21474836480,
				MegaBytes:    20480,
				GigaBytes:    20,
				TeraBytes:    0.02,
				DateModified: "",
			},
		},
		DateCreated: "2023-11-09 13:37:12",
		Public:      false,
		RootUser: ContainerRegistryUser{
			ID:           635,
			UserName:     "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
			Password:     "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
			Root:         true,
			DateCreated:  "2023-11-09 13:37:12",
			DateModified: "2023-11-09 13:37:12",
		},
		Metadata: ContainerRegistryMetadata{
			Region: ContainerRegistryRegion{
				ID:           3,
				Name:         "sjc",
				URN:          "sjc.vultrcr.com",
				BaseURL:      "https://sjc.vultrcr.com",
				Public:       true,
				DateCreated:  "2023-09-14 09:09:16",
				DateModified: "2023-09-14 09:09:16",
				DataCenter: ContainerRegistryRegionDataCenter{
					ID:          12,
					Name:        "Silicon Valley",
					SiteCode:    "SJC2",
					Region:      "West",
					Country:     "US",
					Continent:   "North America",
					Description: "Silicon Valley, California",
					Airport:     "SJC",
				},
			},
			Subscription: ContainerRegistrySubscription{
				Billing: ContainerRegistrySubscriptionBilling{
					MonthlyPrice:   5,
					PendingCharges: 0,
				},
			},
		},
	}

	if !reflect.DeepEqual(vcr, expected) {
		t.Errorf("ContainerRegistry.Get returned %+v, expected %+v", vcr, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.ContainerRegistry.Get(c, vcrID)
	if err == nil {
		t.Error("ContainerRegistry.Get returned nil")
	}
}

func TestVCRServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"

	mux.HandleFunc(fmt.Sprintf("%s/%s", vcrPath, vcrID), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "id": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
    "name": "govultrtest",
    "region": "sjc",
    "urn": "sjc.vultrcr.com/govultrtest",
    "storage": {
        "used": {
            "updated_at": "2023-11-09 13:37:12",
            "bytes": 0,
            "mb": 0,
            "gb": 0,
            "tb": 0
        },
        "allowed": {
            "bytes": 21474836480,
            "mb": 20480,
            "gb": 20,
            "tb": 0.02
        }
    },
    "date_created": "2023-11-09 13:37:12",
    "public": false,
    "root_user": {
        "id": 635,
        "username": "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
        "password": "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
        "root": true,
        "added_at": "2023-11-09 13:37:12",
        "updated_at": "2023-11-09 13:37:12"
    },
    "metadata": {
        "region": {
            "id": 3,
            "name": "sjc",
            "urn": "sjc.vultrcr.com",
            "base_url": "https://sjc.vultrcr.com",
            "public": true,
            "added_at": "2023-09-14 09:09:16",
            "updated_at": "2023-09-14 09:09:16",
            "data_center": {
                "id": 12,
                "name": "Silicon Valley",
                "site_code": "SJC2",
                "region": "West",
                "country": "US",
                "continent": "North America",
                "description": "Silicon Valley, California",
                "airport": "SJC"
            }
        },
        "subscription": {
            "billing": {
                "monthly_price": 5,
                "pending_charges": 0
            }
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	req := &ContainerRegistryUpdateReq{
		Public: BoolToBoolPtr(true),
	}

	vcr, _, err := client.ContainerRegistry.Update(ctx, vcrID, req)
	if err != nil {
		t.Errorf("ContainerRegistry.Update returned %v", err)
	}

	expected := &ContainerRegistry{
		ID:   "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
		Name: "govultrtest",
		URN:  "sjc.vultrcr.com/govultrtest",
		Storage: ContainerRegistryStorage{
			Used: ContainerRegistryStorageCount{
				Bytes:        0,
				MegaBytes:    0,
				GigaBytes:    0,
				TeraBytes:    0,
				DateModified: "2023-11-09 13:37:12",
			},
			Allowed: ContainerRegistryStorageCount{
				Bytes:        21474836480,
				MegaBytes:    20480,
				GigaBytes:    20,
				TeraBytes:    0.02,
				DateModified: "",
			},
		},
		DateCreated: "2023-11-09 13:37:12",
		Public:      false,
		RootUser: ContainerRegistryUser{
			ID:           635,
			UserName:     "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0",
			Password:     "5dEr33smkSYauMmRWusNFk9HpweL5CevpnFr",
			Root:         true,
			DateCreated:  "2023-11-09 13:37:12",
			DateModified: "2023-11-09 13:37:12",
		},
		Metadata: ContainerRegistryMetadata{
			Region: ContainerRegistryRegion{
				ID:           3,
				Name:         "sjc",
				URN:          "sjc.vultrcr.com",
				BaseURL:      "https://sjc.vultrcr.com",
				Public:       true,
				DateCreated:  "2023-09-14 09:09:16",
				DateModified: "2023-09-14 09:09:16",
				DataCenter: ContainerRegistryRegionDataCenter{
					ID:          12,
					Name:        "Silicon Valley",
					SiteCode:    "SJC2",
					Region:      "West",
					Country:     "US",
					Continent:   "North America",
					Description: "Silicon Valley, California",
					Airport:     "SJC",
				},
			},
			Subscription: ContainerRegistrySubscription{
				Billing: ContainerRegistrySubscriptionBilling{
					MonthlyPrice:   5,
					PendingCharges: 0,
				},
			},
		},
	}

	if !reflect.DeepEqual(vcr, expected) {
		t.Errorf("ContainerRegistry.Update returned %+v \n\n expected %+v", vcr, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.ContainerRegistry.Update(c, vcrID, req)
	if err == nil {
		t.Error("ContainerRegistry.Update returned nil")
	}
}

func TestVCRServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"

	mux.HandleFunc(fmt.Sprintf("%s/%s", vcrPath, vcrID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ContainerRegistry.Delete(ctx, vcrID)
	if err != nil {
		t.Errorf("ContainerRegistry.Delete returned %+v", err)
	}
}

func TestVCRServiceHandler_GetRepository(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"
	vcrImage := "vultr-csi"

	mux.HandleFunc(fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, vcrImage), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "name": "govultrtest/vultr-csi",
    "image": "vultr-csi",
    "description": "",
    "added_at": "2023-10-05T18:22:17.041Z",
    "updated_at": "2023-10-27T23:26:09.369Z",
    "pull_count": 9,
    "artifact_count": 7
}`
		fmt.Fprint(writer, response)
	})

	vcrRepo, _, err := client.ContainerRegistry.GetRepository(ctx, vcrID, vcrImage)
	if err != nil {
		t.Errorf("ContainerRegistry.GetRepository returned %+v", err)
	}

	expected := &ContainerRegistryRepo{
		Name:          "govultrtest/vultr-csi",
		Image:         "vultr-csi",
		Description:   "",
		DateCreated:   "2023-10-05T18:22:17.041Z",
		DateModified:  "2023-10-27T23:26:09.369Z",
		PullCount:     9,
		ArtifactCount: 7,
	}

	if !reflect.DeepEqual(vcrRepo, expected) {
		t.Errorf("ContainerRegistry.GetRepository returned %+v, expected %+v", vcrRepo, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	if _, _, err = client.ContainerRegistry.GetRepository(c, vcrID, vcrImage); err == nil {
		t.Error("ContainerRegistry.GetRepository returned nil")
	}
}

func TestVCRServiceHandler_ListRepositories(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"

	mux.HandleFunc(fmt.Sprintf("%s/%s/repositories", vcrPath, vcrID), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "repositories": [
        {
            "name": "govultrtest/vultr-csi",
            "image": "vultr-csi",
            "description": "",
            "added_at": "2023-10-05T18:22:17.041Z",
            "updated_at": "2023-10-27T23:26:09.369Z",
            "pull_count": 9,
            "artifact_count": 7
        }
    ],
    "meta": {
        "total": 1,
        "links": {
            "next": "",
            "prev": ""
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vcrRepos, meta, _, err := client.ContainerRegistry.ListRepositories(ctx, vcrID, nil)
	if err != nil {
		t.Errorf("ContainerRegistry.ListRepositories returned %+v", err)
	}

	expected := []ContainerRegistryRepo{
		{
			Name:          "govultrtest/vultr-csi",
			Image:         "vultr-csi",
			Description:   "",
			DateCreated:   "2023-10-05T18:22:17.041Z",
			DateModified:  "2023-10-27T23:26:09.369Z",
			PullCount:     9,
			ArtifactCount: 7,
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(vcrRepos, expected) {
		t.Errorf("ContainerRegistry.ListRepositories returned %+v, expected %+v", vcrRepos, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ContainerRegistry.ListRepositories meta returned %+v, expected %+v", meta, expectedMeta)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	if _, _, _, err = client.ContainerRegistry.ListRepositories(c, vcrID, nil); err == nil {
		t.Error("ContainerRegistry.ListRepositories returned nil")
	}
}

func TestVCRServiceHandler_UpdateRepository(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"
	vcrImage := "vultr-csi"

	mux.HandleFunc(fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, vcrImage), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "name": "govultrtest/vultr-csi",
    "image": "vultr-csi",
    "description": "test",
    "added_at": "2023-10-05T18:22:17.041Z",
    "updated_at": "2023-10-27T23:26:09.369Z",
    "pull_count": 9,
    "artifact_count": 7
}`
		fmt.Fprint(writer, response)
	})

	req := &ContainerRegistryRepoUpdateReq{
		Description: "test",
	}

	vcrRepo, _, err := client.ContainerRegistry.UpdateRepository(ctx, vcrID, vcrImage, req)
	if err != nil {
		t.Errorf("ContainerRegistry.UpdateRepository returned %+v", err)
	}

	expected := &ContainerRegistryRepo{
		Name:          "govultrtest/vultr-csi",
		Image:         "vultr-csi",
		Description:   "test",
		DateCreated:   "2023-10-05T18:22:17.041Z",
		DateModified:  "2023-10-27T23:26:09.369Z",
		PullCount:     9,
		ArtifactCount: 7,
	}

	if !reflect.DeepEqual(vcrRepo, expected) {
		t.Errorf("ContainerRegistry.UpdateRepository returned %+v, expected %+v", vcrRepo, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	if _, _, err = client.ContainerRegistry.UpdateRepository(c, vcrID, vcrImage, req); err == nil {
		t.Error("ContainerRegistry.UpdateRepository returned nil")
	}
}

func TestVCRServiceHandler_DeleteRepository(t *testing.T) {
	setup()
	defer teardown()

	vcrID := "e1d6be16-2b0c-4d76-a3eb-f28bf6ea5fe0"
	vcrImage := "vultr-csi"

	mux.HandleFunc(fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, vcrImage), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ContainerRegistry.DeleteRepository(ctx, vcrID, vcrImage)
	if err != nil {
		t.Errorf("ContainerRegistry.DeleteRepository returned %+v", err)
	}
}

func TestVCRServiceHandler_ListRegions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/region/list", vcrPath), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "regions": [
        {
            "id": 3,
            "name": "sjc",
            "urn": "sjc.vultrcr.com",
            "base_url": "https://sjc.vultrcr.com",
            "public": true,
            "added_at": "2023-09-14 09:09:16",
            "updated_at": "2023-09-14 09:09:16",
            "data_center": {
                "id": 12,
                "name": "Silicon Valley",
                "site_code": "SJC2",
                "region": "West",
                "country": "US",
                "continent": "North America",
                "description": "Silicon Valley, California",
                "airport": "SJC"
            }
        }
    ],
    "meta": {
        "total": 1,
        "links": {
            "next": "",
            "prev": ""
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vcrRegions, meta, _, err := client.ContainerRegistry.ListRegions(ctx, nil)
	if err != nil {
		t.Errorf("ContainerRegistry.ListRegions returned %v", err)
	}

	expected := []ContainerRegistryRegion{
		{
			ID:           3,
			Name:         "sjc",
			URN:          "sjc.vultrcr.com",
			BaseURL:      "https://sjc.vultrcr.com",
			Public:       true,
			DateCreated:  "2023-09-14 09:09:16",
			DateModified: "2023-09-14 09:09:16",
			DataCenter: ContainerRegistryRegionDataCenter{
				ID:          12,
				Name:        "Silicon Valley",
				SiteCode:    "SJC2",
				Region:      "West",
				Country:     "US",
				Continent:   "North America",
				Description: "Silicon Valley, California",
				Airport:     "SJC",
			},
		},
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(vcrRegions, expected) {
		t.Errorf("ContainerRegistry.ListRegions returned %+v, expected %+v", vcrRegions, expected)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ContainerRegistry.ListRegions meta returned %+v, expected %+v", meta, expectedMeta)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, _, err = client.ContainerRegistry.ListRegions(c, nil)
	if err == nil {
		t.Error("ContainerRegistry.ListRegions returned nil")
	}
}

func TestVCRServiceHandler_ListPlans(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/plan/list", vcrPath), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
    "plans": {
        "start_up": {
            "vanity_name": "Start Up",
            "max_storage_mb": 10240,
            "monthly_price": 0
        },
        "business": {
            "vanity_name": "Business",
            "max_storage_mb": 20480,
            "monthly_price": 5
        },
        "premium": {
            "vanity_name": "Premium",
            "max_storage_mb": 51200,
            "monthly_price": 10
        },
        "enterprise": {
            "vanity_name": "Enterprise",
            "max_storage_mb": 1048576,
            "monthly_price": 20
        }
    }
}`
		fmt.Fprint(writer, response)
	})

	vcrPlans, _, err := client.ContainerRegistry.ListPlans(ctx)
	if err != nil {
		t.Errorf("ContainerRegistry.ListPlans returned %v", err)
	}

	expected := &ContainerRegistryPlans{
		Plans: ContainerRegistryPlanTypes{
			StartUp: ContainerRegistryPlan{
				VanityName:   "Start Up",
				MaxStorageMB: 10240,
				MonthlyPrice: 0,
			},
			Business: ContainerRegistryPlan{
				VanityName:   "Business",
				MaxStorageMB: 20480,
				MonthlyPrice: 5,
			},
			Premium: ContainerRegistryPlan{
				VanityName:   "Premium",
				MaxStorageMB: 51200,
				MonthlyPrice: 10,
			},
			Enterprise: ContainerRegistryPlan{
				VanityName:   "Enterprise",
				MaxStorageMB: 1048576,
				MonthlyPrice: 20,
			},
		},
	}

	if !reflect.DeepEqual(vcrPlans, expected) {
		t.Errorf("ContainerRegistry.ListPlans returned %+v, expected %+v", vcrPlans, expected)
	}

	c, can := context.WithTimeout(ctx, 1*time.Microsecond)
	defer can()
	_, _, err = client.ContainerRegistry.ListPlans(c)
	if err == nil {
		t.Error("ContainerRegistry.ListPlans returned nil")
	}
}
