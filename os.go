package govultr

import (
	"context"
	"net/http"
)

// OSService is the interface to interact with the operating system endpoint on the Vultr API
// Link: https://www.vultr.com/api/#os
type OSService interface {
	GetList(ctx context.Context) ([]OS, error)
}

// OSServiceHandler handles interaction with the operating system methods for the Vultr API
type OSServiceHandler struct {
	client *Client
}

// OS represents a Vultr operating system
type OS struct {
	OsID    string `json:"OSID"`
	Name    string `json:"name"`
	Arch    string `json:"arch"`
	Family  string `json:"family"`
	Windows bool   `json:"windows"`
}

// GetList retrieves a list of available operating systems.
// If the Windows flag is true, a Windows license will be included with the instance, which will increase the cost.
func (o *OSServiceHandler) GetList(ctx context.Context) ([]OS, error) {
	uri := "/v1/os/list"
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	osMap := make(map[string]OS)

	err = o.client.DoWithContext(ctx, req, &osMap)
	if err != nil {
		return nil, err
	}

	var oses []OS
	for _, os := range osMap {
		oses = append(oses, os)
	}

	return oses, nil
}
