package govultr

import (
	"context"
)

// IsoService is the interface to interact with the ISO endpoints on the Vultr API
// Link: https://www.vultr.com/api/#iso
type IsoService interface {
	CreateFromURL(ctx context.Context, url string) (*Iso, error)
	Delete(ctx context.Context, isoID int) error
	GetList(ctx context.Context) ([]Iso, error)
	GetPublicList(ctx context.Context) ([]Iso, error)
}

// IsoServiceHandler handles interaction with the ISO methods for the Vultr API
type IsoServiceHandler struct {
	Client *Client
}

// PublicIso represents public ISOs offered in the Vultr ISO library.
type PublicIso struct {
	IsoID       int    `json:"ISOID"`
	DateCreated string `json:"date_created"`
	FileName    string `json:"filename"`
	Size        int    `json:"size"`
	MD5Sum      string `json:"md5sum"`
	SHA512Sum   string `json:"sha512sum"`
	Status      string `json:"status"`
}

// Iso represents ISOs currently available on this account.
type Iso struct {
	IsoID       int    `json:"ISOID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateFromUrl will create a new ISO image on your account
func (i *IsoServiceHandler) CreateFromURL(ctx context.Context, url string) (*Iso, error) {

	return nil, nil
}

// Delete will delete an ISO image from your account
func (i *IsoServiceHandler) Delete(ctx context.Context, isoID int) error {
	return nil
}

// GetList will list all ISOs currently available on your account
func (i *IsoServiceHandler) GetList(ctx context.Context) ([]Iso, error) {

	return nil, nil
}

// GetPublicList will list public ISOs offered in the Vultr ISO library.
func (i *IsoServiceHandler) GetPublicList(ctx context.Context) ([]Iso, error) {

	return nil, nil
}
