package govultr

import (
	"context"
	"net/http"
	"net/url"
)

// ServerService is the interface to interact with the server endpoints on the Vultr API
// Link: https://www.vultr.com/api/#server
type ServerService interface {
	ChangeApp(ctx context.Context, vpsID, appID string) error
	ListApps(ctx context.Context, vpsID string) ([]Application, error)
	AppInfo(ctx context.Context, vpsID string) (*ServerAppInfo, error)
}

// ServerServiceHandler handles interaction with the server methods for the Vultr API
type ServerServiceHandler struct {
	client *Client
}

// ServerAppInfo represents information about the application on your VPS
type ServerAppInfo struct {
	AppInfo string `json:"app_info"`
}

// ChangeApp changes the VPS to a different application.
func (s *ServerServiceHandler) ChangeApp(ctx context.Context, vpsID, appID string) error {

	uri := "/v1/server/app_change"

	values := url.Values{
		"SUBID": {vpsID},
		"APPID": {appID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// ListApps retrieves a list of applications to which a virtual machine can be changed.
func (s *ServerServiceHandler) ListApps(ctx context.Context, vpsID string) ([]Application, error) {

	uri := "/v1/server/app_change_list"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", vpsID)
	req.URL.RawQuery = q.Encode()

	var appMap map[string]Application
	err = s.client.DoWithContext(ctx, req, &appMap)

	if err != nil {
		return nil, err
	}

	var appList []Application
	for _, a := range appMap {
		appList = append(appList, a)
	}

	return appList, nil
}

// AppInfo retrieves the application information for a given VPS ID
func (s *ServerServiceHandler) AppInfo(ctx context.Context, vpsID string) (*ServerAppInfo, error) {

	uri := "/v1/server/get_app_info"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", vpsID)
	req.URL.RawQuery = q.Encode()

	appInfo := new(ServerAppInfo)

	err = s.client.DoWithContext(ctx, req, appInfo)

	if err != nil {
		return nil, err
	}

	return appInfo, nil
}
