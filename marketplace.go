package govultr

import (
	"context"
	"fmt"
	"net/http"
)

const marketplacePath = "/v2/marketplace"

// MarketplaceService is the interface to interact with the Marketplace endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/marketplace
type MarketplaceService interface {
	Create(ctx context.Context, appReq *MarketplaceAppCreate) (int, *http.Response, error)
	List(ctx context.Context) ([]MarketplaceApp, *http.Response, error)

	CreateAppVariableV2(ctx context.Context, appID int, variableReq *MarketplaceAppVariableV2Create) (*MarketplaceAppVariableV2, *http.Response, error)                 //nolint:lll
	UpdateAppVariableV2(ctx context.Context, appID int, variableID int, variableReq *MarketplaceAppVariableV2Update) (*MarketplaceAppVariableV2, *http.Response, error) //nolint:lll
	DeleteAppVariableV2(ctx context.Context, appID int, variableID int) error
	ListAppVariablesV2(ctx context.Context, appID int) ([]MarketplaceAppVariableV2, *http.Response, error)

	CreateAppImage(ctx context.Context, appID int, imageReq *MarketplaceAppImageCreate) (*MarketplaceAppImage, *http.Response, error)

	CreateVendorUser(ctx context.Context, vendorReq *MarketplaceVendorUserReq) error
	UpdateVendorUser(ctx context.Context, vendorReq *MarketplaceVendorUserReq) error

	// Deprecated: Use ListAppVariablesV2 instead
	ListAppVariables(ctx context.Context, imageID string) ([]MarketplaceAppVariable, *http.Response, error)
}

// MarketplaceServiceHandler handles interaction with the server methods for the Vultr API
type MarketplaceServiceHandler struct {
	client *Client
}

// MarketplaceApp represents a Marketplace app
type MarketplaceApp struct {
	ID          int    `json:"id"`
	Name        string `json:"app_name"`
	Description string `json:"app_desc"`
	Visibility  bool   `json:"visibility"`
}

// marketplaceAppBase holds the API response for creating a Marketplace app
type marketplaceAppBase struct {
	App *MarketplaceApp `json:"app"`
}

// marketplaceAppsBase holds the API response for retrieving a list of Marketplace apps
type marketplaceAppsBase struct {
	Apps []MarketplaceApp `json:"apps"`
}

// MarketplaceAppCreate is used for creating a new Marketplace app
type MarketplaceAppCreate struct {
	Name         string `json:"name"`
	NameIDFormat string `json:"name_id_format,omitempty"`
	Description  string `json:"description"`
	OS           string `json:"os,omitempty"`
	RepoURL      string `json:"repo_url,omitempty"`
	SupportURL   string `json:"support_url,omitempty"`
	SupportEmail string `json:"support_email,omitempty"`
	Readme       string `json:"readme,omitempty"`
}

// MarketplaceAppVariableV2 represents a marketplace app variable
type MarketplaceAppVariableV2 struct {
	ID             int     `json:"id"`
	AppID          int     `json:"app_id"`
	Name           string  `json:"name"`
	Label          string  `json:"label"`
	Type           string  `json:"type"`
	PasswordLength *int    `json:"password_length"`
	InputType      *string `json:"input_type"`
	InputRequired  *bool   `json:"input_required"`
}

// marketplaceAppVariableV2Base holds the response from the API for a marketplace app variable
type marketplaceAppVariableV2Base struct {
	AppVariable *MarketplaceAppVariableV2 `json:"app_variable"`
}

// marketplaceAppVariablesV2Base holds the response from the API for a list of marketplace app variables
type marketplaceAppVariablesV2Base struct {
	AppVariables []MarketplaceAppVariableV2 `json:"app_variables"`
}

// MarketplaceAppVariableV2Create is used to create a marketplace app variable
type MarketplaceAppVariableV2Create struct {
	Name           string  `json:"name"`
	Label          string  `json:"label"`
	Type           string  `json:"type"`
	PasswordLength *int    `json:"password_length,omitempty"`
	InputType      *string `json:"input_type,omitempty"`
	InputRequired  *bool   `json:"input_required,omitempty"`
}

// MarketplaceAppVariableV2Update is used to update a marketplace app variable
type MarketplaceAppVariableV2Update struct {
	Label string `json:"label"`
}

// MarketplaceAppImageCreate is used to create a marketplace app image
type MarketplaceAppImageCreate struct {
	Image []byte `json:"b64_image"`
}

// MarketplaceAppImage represents a marketplace app image
type MarketplaceAppImage struct {
	ID       int    `json:"id"`
	AppID    string `json:"app_id"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

// MarketplaceVendorUserReq is used to create a marketplace vendor user
type MarketplaceVendorUserReq struct {
	VendorName string `json:"vendor_name"`
	XHandle    string `json:"x_handle,omitempty"`
	WebsiteURL string `json:"website_url,omitempty"`
	GitURL     string `json:"git_url,omitempty"`
	SlackURL   string `json:"slack_url,omitempty"`
}

// MarketplaceAppVariable represents a user-supplied variable for a Marketplace app
//
// Deprecated: Use MarketplaceAppVariableV2 instead
type MarketplaceAppVariable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    *bool  `json:"required"`
}

// marketplaceAppVariablesBase holds the API response for retrieving a list of user-supplied variables for a Marketplace app
//
// Deprecated: Use marketplaceAppVariableV2Base instead
type marketplaceAppVariablesBase struct {
	MarketplaceAppVariables []MarketplaceAppVariable `json:"variables"`
}

// Create a new marketplace app for the authenticated vendor account
func (d *MarketplaceServiceHandler) Create(ctx context.Context, appReq *MarketplaceAppCreate) (int, *http.Response, error) {
	uri := fmt.Sprintf("%s/apps", marketplacePath)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, appReq)
	if err != nil {
		return 0, nil, err
	}

	app := new(marketplaceAppBase)
	resp, err := d.client.DoWithContext(ctx, req, app)
	if err != nil {
		return 0, nil, err
	}

	return app.App.ID, resp, nil
}

// List all Marketplace apps owned by the authenticated vendor account
func (d *MarketplaceServiceHandler) List(ctx context.Context) ([]MarketplaceApp, *http.Response, error) {
	uri := fmt.Sprintf("%s/apps", marketplacePath)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	apps := new(marketplaceAppsBase)
	resp, err := d.client.DoWithContext(ctx, req, apps)
	if err != nil {
		return nil, nil, err
	}

	return apps.Apps, resp, nil
}

// CreateAppVariableV2 adds a new variable to a Marketplace App
func (d *MarketplaceServiceHandler) CreateAppVariableV2(ctx context.Context, appID int, variableReq *MarketplaceAppVariableV2Create) (*MarketplaceAppVariableV2, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/apps/%d/manage-variables", marketplacePath, appID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, variableReq)
	if err != nil {
		return nil, nil, err
	}

	variable := new(marketplaceAppVariableV2Base)
	resp, err := d.client.DoWithContext(ctx, req, variable)
	if err != nil {
		return nil, nil, err
	}

	return variable.AppVariable, resp, nil
}

// UpdateAppVariableV2 updates an existing Marketplace App Variable
func (d *MarketplaceServiceHandler) UpdateAppVariableV2(ctx context.Context, appID, variableID int, variableReq *MarketplaceAppVariableV2Update) (*MarketplaceAppVariableV2, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/apps/%d/manage-variables/%d", marketplacePath, appID, variableID)

	req, err := d.client.NewRequest(ctx, http.MethodPatch, uri, variableReq)
	if err != nil {
		return nil, nil, err
	}

	variable := new(marketplaceAppVariableV2Base)
	resp, err := d.client.DoWithContext(ctx, req, variable)
	if err != nil {
		return nil, nil, err
	}

	return variable.AppVariable, resp, nil
}

// ListAppVariablesV2 list all variables for a Marketplace App
func (d *MarketplaceServiceHandler) ListAppVariablesV2(ctx context.Context, appID int) ([]MarketplaceAppVariableV2, *http.Response, error) {
	uri := fmt.Sprintf("%s/apps/%d/manage-variables", marketplacePath, appID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	variables := new(marketplaceAppVariablesV2Base)
	resp, err := d.client.DoWithContext(ctx, req, variables)
	if err != nil {
		return nil, nil, err
	}

	return variables.AppVariables, resp, nil
}

// DeleteAppVariableV2 deletes a Marketplace App variable
func (d *MarketplaceServiceHandler) DeleteAppVariableV2(ctx context.Context, appID, variableID int) error {
	uri := fmt.Sprintf("%s/apps/%d/manage-variables/%d", marketplacePath, appID, variableID)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// CreateAppImage uploads an image to a marketplace app
func (d *MarketplaceServiceHandler) CreateAppImage(ctx context.Context, appID int, imageReq *MarketplaceAppImageCreate) (*MarketplaceAppImage, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/apps/%d/images", marketplacePath, appID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, imageReq)
	if err != nil {
		return nil, nil, err
	}

	img := new(MarketplaceAppImage)
	resp, err := d.client.DoWithContext(ctx, req, img)
	if err != nil {
		return nil, nil, err
	}

	return img, resp, nil
}

// CreateVendorUser creates a vendor user for the current vendor account
func (d *MarketplaceServiceHandler) CreateVendorUser(ctx context.Context, vendorUserReq *MarketplaceVendorUserReq) error {
	uri := fmt.Sprintf("%s/vendor", marketplacePath)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, vendorUserReq)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// UpdateVendorUser updates marketplace vendor user settings for the current vendor account
func (d *MarketplaceServiceHandler) UpdateVendorUser(ctx context.Context, vendorReq *MarketplaceVendorUserReq) error {
	uri := fmt.Sprintf("%s/vendor", marketplacePath)

	req, err := d.client.NewRequest(ctx, http.MethodPatch, uri, vendorReq)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListAppVariables retrieves all user-supplied variables for a Marketplace app
//
// Deprecated: Use ListAppVariablesV2 instead
func (d *MarketplaceServiceHandler) ListAppVariables(ctx context.Context, imageID string) ([]MarketplaceAppVariable, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/apps/%s/variables", marketplacePath, imageID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	marketplaceAppVariables := new(marketplaceAppVariablesBase)
	resp, err := d.client.DoWithContext(ctx, req, marketplaceAppVariables)
	if err != nil {
		return nil, nil, err
	}

	return marketplaceAppVariables.MarketplaceAppVariables, resp, nil
}
