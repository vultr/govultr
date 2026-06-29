package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMarketplaceServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/marketplace/apps", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "app": {
			"id": 42
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &MarketplaceAppCreate{
		Name:         "My App",
		NameIDFormat: "my-app",
		Description:  "Description of my app",
		OS:           "Linux",
		RepoURL:      "https://github.com/example/app",
		SupportURL:   "https://example.com/support",
		SupportEmail: "support@example.com",
		Readme:       "# My App\n\nDescription here.",
	}

	appID, _, err := client.Marketplace.Create(ctx, req)
	if err != nil {
		t.Errorf("Marketplace.Create returned %+v", err)
	}

	expected := 42

	if !reflect.DeepEqual(appID, expected) {
		t.Errorf("Marketplace.Create returned %+v, expected %+v", appID, expected)
	}
}

func TestMarketplaceServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/marketplace/apps", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "apps": [
			{
			  "id": 1001,
			  "app_name": "WordPress",
			  "app_desc": "WordPress is a content management system based on PHP and MySQL.",
			  "visibility": true
			},
			{
			  "id": 1002,
			  "app_name": "Internal Tool",
			  "app_desc": "An internal application for deployment.",
			  "visibility": false
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

	apps, _, err := client.Marketplace.List(ctx)
	if err != nil {
		t.Errorf("Marketplace.List returned %+v", err)
	}

	expected := []MarketplaceApp{
		{
			ID:          1001,
			Name:        "WordPress",
			Description: "WordPress is a content management system based on PHP and MySQL.",
			Visibility:  true,
		},
		{
			ID:          1002,
			Name:        "Internal Tool",
			Description: "An internal application for deployment.",
			Visibility:  false,
		},
	}

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("Marketplace.List returned %+v, expected %+v", apps, expected)
	}
}

func TestMarketplaceServiceHandler_ListAppVariables(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%s/variables", "testimage"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"variables": [
				{
					"name": "some_required_variable",
					"description": "This is an example of a required user-supplied variable for this Marketplace app.",
					"required": true
				},
				{
					"name": "some_optional_variable",
					"description": "This is an example of an optional user-supplied variable for this Marketplace app.",
					"required": false
				}
			]
		}`
		fmt.Fprint(writer, response)
	})

	variables, _, err := client.Marketplace.ListAppVariables(ctx, "testimage")
	if err != nil {
		t.Errorf("Marketplace.ListAppVariables returned %+v", err)
	}

	expected := []MarketplaceAppVariable{
		{
			Name:        "some_required_variable",
			Description: "This is an example of a required user-supplied variable for this Marketplace app.",
			Required:    BoolToBoolPtr(true),
		},
		{
			Name:        "some_optional_variable",
			Description: "This is an example of an optional user-supplied variable for this Marketplace app.",
			Required:    BoolToBoolPtr(false),
		},
	}

	if !reflect.DeepEqual(variables, expected) {
		t.Errorf("Marketplace.ListAppVariables returned %+v, expected %+v", variables, expected)
	}
}

func TestMarketplaceServiceHandler_CreateAppVariableV2(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%d/manage-variables", 468), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "app_variable": {
			"id": 121,
			"app_id": 468,
			"name": "server_name",
			"label": "Server Name",
			"type": "user_supplied",
			"password_length": null,
			"input_type": "text_field",
			"input_required": true
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &MarketplaceAppVariableV2Create{
		Name:          "server_name",
		Label:         "Server Name",
		Type:          "user_supplied",
		InputType:     StringToStringPtr("text_field"),
		InputRequired: BoolToBoolPtr(true),
	}

	variable, _, err := client.Marketplace.CreateAppVariableV2(ctx, 468, req)
	if err != nil {
		t.Errorf("Marketplace.CreateAppVariableV2 returned %+v", err)
	}

	expected := &MarketplaceAppVariableV2{
		ID:            121,
		AppID:         468,
		Name:          "server_name",
		Label:         "Server Name",
		Type:          "user_supplied",
		InputType:     StringToStringPtr("text_field"),
		InputRequired: BoolToBoolPtr(true),
	}

	if !reflect.DeepEqual(variable, expected) {
		t.Errorf("Marketplace.CreateAppVariableV2 returned %+v, expected %+v", variable, expected)
	}
}

func TestMarketplaceServiceHandler_UpdateAppVariableV2(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%d/manage-variables/%d", 468, 121), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "app_variable": {
			"id": 121,
			"app_id": 468,
			"name": "server_name",
			"label": "Public Server Name",
			"type": "user_supplied",
			"password_length": null,
			"input_type": "text_field",
			"input_required": true
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &MarketplaceAppVariableV2Update{
		Label: "Public Server Name",
	}

	variable, _, err := client.Marketplace.UpdateAppVariableV2(ctx, 468, 121, req)
	if err != nil {
		t.Errorf("Marketplace.UpdateAppVariableV2 returned %+v", err)
	}

	expected := &MarketplaceAppVariableV2{
		ID:            121,
		AppID:         468,
		Name:          "server_name",
		Label:         "Public Server Name",
		Type:          "user_supplied",
		InputType:     StringToStringPtr("text_field"),
		InputRequired: BoolToBoolPtr(true),
	}

	if !reflect.DeepEqual(variable, expected) {
		t.Errorf("Marketplace.UpdateAppVariableV2 returned %+v, expected %+v", variable, expected)
	}
}

func TestMarketplaceServiceHandler_ListAppVariablesV2(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%d/manage-variables", 468), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "app_variables": [
			{
			  "id": 121,
			  "app_id": 468,
			  "name": "server_name",
			  "label": "Server Name",
			  "type": "user_supplied",
			  "password_length": null,
			  "input_type": "text_field",
			  "input_required": true
			}
		  ]
		}`
		fmt.Fprint(writer, response)
	})

	variables, _, err := client.Marketplace.ListAppVariablesV2(ctx, 468)
	if err != nil {
		t.Errorf("Marketplace.ListAppVariablesV2 returned %+v", err)
	}

	expected := []MarketplaceAppVariableV2{
		{
			ID:            121,
			AppID:         468,
			Name:          "server_name",
			Label:         "Server Name",
			Type:          "user_supplied",
			InputType:     StringToStringPtr("text_field"),
			InputRequired: BoolToBoolPtr(true),
		},
	}

	if !reflect.DeepEqual(variables, expected) {
		t.Errorf("Marketplace.ListAppVariablesV2 returned %+v, expected %+v", variables, expected)
	}
}

func TestMarketplaceServiceHandler_DeleteAppVariableV2(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%d/manage-variables/%d", 468, 121), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Marketplace.DeleteAppVariableV2(ctx, 468, 121)
	if err != nil {
		t.Errorf("Marketplace.DeleteAppVariableV2 returned %+v", err)
	}
}

func TestMarketplaceServiceHandler_CreateAppImage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%d/images", 468), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "id": 150,
		  "app_id": "1234",
		  "filename": "testvendor-gallery_image-12345ac-6789-01c2-3c4c-56a789bf012a-3456789012.png",
		  "url": "http://www.vultr.com/marketplace-assets-dev/testvendor-gallery_image-12345ac-6789-01c2-3c4c-56a789bf012a-3456789012.png"
		}`
		fmt.Fprint(writer, response)
	})

	req := &MarketplaceAppImageCreate{
		Image: []byte{},
	}

	img, _, err := client.Marketplace.CreateAppImage(ctx, 468, req)
	if err != nil {
		t.Errorf("Marketplace.CreateAppImage returned %+v", err)
	}

	expected := &MarketplaceAppImage{
		ID:       150,
		AppID:    "1234",
		Filename: "testvendor-gallery_image-12345ac-6789-01c2-3c4c-56a789bf012a-3456789012.png",
		URL:      "http://www.vultr.com/marketplace-assets-dev/testvendor-gallery_image-12345ac-6789-01c2-3c4c-56a789bf012a-3456789012.png",
	}

	if !reflect.DeepEqual(img, expected) {
		t.Errorf("Marketplace.CreateAppImage returned %+v, expected %+v", img, expected)
	}
}

func TestMarketplaceServiceHandler_CreateVendorUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/marketplace/vendor", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &MarketplaceVendorUserReq{
		VendorName: "testVendor",
		XHandle:    "mycompany",
		WebsiteURL: "https://example.com",
		GitURL:     "https://github.com/example",
		SlackURL:   "https://example.slack.com",
	}

	err := client.Marketplace.CreateVendorUser(ctx, req)
	if err != nil {
		t.Errorf("Marketplace.CreateVendorUser returned %+v", err)
	}
}

func TestMarketplaceServiceHandler_UpdateVendorUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/marketplace/vendor", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &MarketplaceVendorUserReq{
		VendorName: "testVendor",
		XHandle:    "mycompany",
		WebsiteURL: "https://example.com",
		GitURL:     "https://github.com/example",
		SlackURL:   "https://example.slack.com",
	}

	err := client.Marketplace.UpdateVendorUser(ctx, req)
	if err != nil {
		t.Errorf("Marketplace.UpdateVendorUser returned %+v", err)
	}
}
