package govultr

import (
	"context"
	"fmt"
	"net/http"
)

const templatePath = "/v2/instances/templates"

// InstanceTemplateService is the interface to interact with the instance template endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/instance-templates
type InstanceTemplateService interface {
	Create(ctx context.Context, templateReq *InstanceTemplateReq) (*InstanceTemplate, *http.Response, error)
	Get(ctx context.Context, templateID string) (*InstanceTemplate, *http.Response, error)
	Update(ctx context.Context, templateID string, templateReq *InstanceTemplateReq) (*InstanceTemplate, *http.Response, error)
	Delete(ctx context.Context, templateID string) error
	List(ctx context.Context) ([]InstanceTemplate, *http.Response, error)

	CreateFromInstance(ctx context.Context, instanceID string) (*InstanceTemplate, *http.Response, error)
}

// InstanceTemplateServiceHandler handles interaction with the server methods for the Vultr API
type InstanceTemplateServiceHandler struct {
	client *Client
}

type InstanceTemplate struct {
	ID               string                   `json:"id"`
	Plan             string                   `json:"plan"`
	Label            string                   `json:"label"`
	OS               string                   `json:"os"`
	MarketplaceApp   string                   `json:"marketplace_app"`
	MarketplaceImage string                   `json:"marketplace_image"`
	Snapshot         string                   `json:"snapshot"`
	ISO              string                   `json:"iso"`
	SSHKeys          []InstanceTemplateSSHKey `json:"ssh_keys"`
	StartupScript    string                   `json:"startup_script"`
	DiskConfig       string                   `json:"disk_config"`
}

type InstanceTemplateSSHKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InstanceTemplateReq struct {
	Plan               string            `json:"plan"`
	Label              string            `json:"label"`
	IsoID              string            `json:"iso_id"`
	MarketplaceAppID   int               `json:"marketplace_app_id"`
	MarketplaceImageID int               `json:"marketplace_image_id"`
	OsID               int               `json:"os_id"`
	SnapshotID         string            `json:"snapshot_id"`
	SSHKeyIDs          []string          `json:"ssh_key_ids"`
	Template           map[string]string `json:"template"`
	UserData           string            `json:"user_data"`
}

type instanceTemplateBase struct {
	InstanceTemplate *InstanceTemplate `json:"instance_template"`
}

type instanceTemplatesBase struct {
	InstanceTemplates []InstanceTemplate `json:"instance_templates"`
}

// Create will create the instance template with the given parameters
func (t *InstanceTemplateServiceHandler) Create(ctx context.Context, templateReq *InstanceTemplateReq) (*InstanceTemplate, *http.Response, error) { //nolint:lll
	req, err := t.client.NewRequest(ctx, http.MethodPost, templatePath, templateReq)
	if err != nil {
		return nil, nil, err
	}

	template := new(instanceTemplateBase)
	resp, err := t.client.DoWithContext(ctx, req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.InstanceTemplate, resp, nil
}

// Get will get the instance template with the given templateID
func (t *InstanceTemplateServiceHandler) Get(ctx context.Context, templateID string) (*InstanceTemplate, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", templatePath, templateID)

	req, err := t.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	template := new(instanceTemplateBase)
	resp, err := t.client.DoWithContext(ctx, req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.InstanceTemplate, resp, nil
}

// Update will update the instance template with the given parameters
func (t *InstanceTemplateServiceHandler) Update(ctx context.Context, templateID string, templateReq *InstanceTemplateReq) (*InstanceTemplate, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", templatePath, templateID)

	req, err := t.client.NewRequest(ctx, http.MethodPut, uri, templateReq)
	if err != nil {
		return nil, nil, err
	}

	template := new(instanceTemplateBase)
	resp, err := t.client.DoWithContext(ctx, req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.InstanceTemplate, resp, nil
}

// Delete an instance template
func (t *InstanceTemplateServiceHandler) Delete(ctx context.Context, templateID string) error {
	uri := fmt.Sprintf("%s/%s", templatePath, templateID)

	req, err := t.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = t.client.DoWithContext(ctx, req, nil)
	return err
}

// List all instance templates on your account.
func (t *InstanceTemplateServiceHandler) List(ctx context.Context) ([]InstanceTemplate, *http.Response, error) {
	req, err := t.client.NewRequest(ctx, http.MethodGet, templatePath, nil)
	if err != nil {
		return nil, nil, err
	}

	templates := new(instanceTemplatesBase)
	resp, err := t.client.DoWithContext(ctx, req, templates)
	if err != nil {
		return nil, resp, err
	}

	return templates.InstanceTemplates, resp, nil
}

// CreateFromInstance will create an instance template from an existing instance
func (t *InstanceTemplateServiceHandler) CreateFromInstance(ctx context.Context, instanceID string) (*InstanceTemplate, *http.Response, error) { //nolint:lll
	req, err := t.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/from-instance/%s", templatePath, instanceID), nil)
	if err != nil {
		return nil, nil, err
	}

	template := new(instanceTemplateBase)
	resp, err := t.client.DoWithContext(ctx, req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.InstanceTemplate, resp, nil
}
