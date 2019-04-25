package govultr

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// IsoService is the interface to interact with the ISO endpoints on the Vultr API
// Link: https://www.vultr.com/api/#iso
type IsoService interface {
	CreateFromURL(ctx context.Context, isoUrl string) (*Iso, error)
	Delete(ctx context.Context, isoID int) error
	GetList(ctx context.Context) ([]Iso, error)
	GetPublicList(ctx context.Context) ([]PublicIso, error)
}

// IsoServiceHandler handles interaction with the ISO methods for the Vultr API
type IsoServiceHandler struct {
	Client *Client
}

// Iso represents ISOs currently available on this account.
type Iso struct {
	IsoID       int    `json:"ISOID"`
	DateCreated string `json:"date_created"`
	FileName    string `json:"filename"`
	Size        int    `json:"size"`
	MD5Sum      string `json:"md5sum"`
	SHA512Sum   string `json:"sha512sum"`
	Status      string `json:"status"`
}

// PublicIso represents public ISOs offered in the Vultr ISO library.
type PublicIso struct {
	IsoID       int    `json:"ISOID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateFromUrl will create a new ISO image on your account
func (i *IsoServiceHandler) CreateFromURL(ctx context.Context, isoUrl string) (*Iso, error) {

	uri := "/v1/iso/create_from_url"

	values := url.Values{
		"url": {isoUrl},
	}

	req, err := i.Client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return nil, err
	}

	iso := new(Iso)
	err = i.Client.DoWithContext(ctx, req, iso)

	if err != nil {
		return nil, err
	}

	return iso, nil
}

// Delete will delete an ISO image from your account
func (i *IsoServiceHandler) Delete(ctx context.Context, isoID int) error {

	uri := "/v1/iso/destroy"

	values := url.Values{
		"ISOID": {strconv.Itoa(isoID)},
	}

	req, err := i.Client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = i.Client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetList will list all ISOs currently available on your account
func (i *IsoServiceHandler) GetList(ctx context.Context) ([]Iso, error) {

	uri := "/v1/iso/list"

	req, err := i.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	var isoMap map[string]Iso
	err = i.Client.DoWithContext(ctx, req, &isoMap)

	if err != nil {
		return nil, err
	}

	var iso []Iso
	for _, i := range isoMap {
		iso = append(iso, i)
	}

	return iso, nil
}

// GetPublicList will list public ISOs offered in the Vultr ISO library.
func (i *IsoServiceHandler) GetPublicList(ctx context.Context) ([]PublicIso, error) {

	uri := "/v1/iso/list_public"

	req, err := i.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	var isoMap map[string]PublicIso
	err = i.Client.DoWithContext(ctx, req, &isoMap)

	if err != nil {
		return nil, err
	}

	var publicIso []PublicIso

	for _, p := range isoMap {
		publicIso = append(publicIso, p)
	}

	return publicIso, nil
}
