package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const path = "/v2/users"

// UserService is the interface to interact with the user management endpoints on the Vultr API
type UserService interface {
	Create(ctx context.Context, userCreate *UserReq) (*User, error)
	Get(ctx context.Context, userID string) (*User, error)
	Update(ctx context.Context, userID string, userReq *UserReq) error
	Delete(ctx context.Context, userID string) error
	List(ctx context.Context, options *ListOptions) ([]User, *Meta, error)
}

var _ UserService = &UserServiceHandler{}

// UserServiceHandler handles interaction with the user methods for the Vultr API
type UserServiceHandler struct {
	client *Client
}

// User represents an user on Vultr
type User struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	APIEnabled string   `json:"api_enabled"`
	APIKey     string   `json:"api_key,omitempty"`
	ACL        []string `json:"acls,omitempty"`
}

// UserReq is the user struct for create and update calls
type UserReq struct {
	Email      string   `json:"email"`
	Name       string   `json:"name"`
	APIEnabled string   `json:"api_enabled"`
	ACL        []string `json:"acl"`
	Password   string   `json:"password"`
}

type usersBase struct {
	Users []User `json:"users"`
	Meta  *Meta  `json:"meta"`
}

type userBase struct {
	User *User `json:"user"`
}

// Create will add the specified user to your Vultr account
func (u *UserServiceHandler) Create(ctx context.Context, userCreate *UserReq) (*User, error) {
	req, err := u.client.NewRequest(ctx, http.MethodPost, path, userCreate)

	if err != nil {
		return nil, err
	}

	user := new(userBase)
	if err = u.client.DoWithContext(ctx, req, user); err != nil {
		return nil, err
	}

	return user.User, nil
}

// Get will retrieve a specific user account
func (u *UserServiceHandler) Get(ctx context.Context, userID string) (*User, error) {
	uri := fmt.Sprintf("%s/%s", path, userID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	user := new(userBase)
	if err = u.client.DoWithContext(ctx, req, user); err != nil {
		return nil, err
	}

	return user.User, nil
}

// Update will update the given user. Empty strings will be ignored.
func (u *UserServiceHandler) Update(ctx context.Context, userID string, userReq *UserReq) error {
	uri := fmt.Sprintf("%s/%s", path, userID)
	req, err := u.client.NewRequest(ctx, http.MethodPatch, uri, userReq)

	if err != nil {
		return err
	}

	if err = u.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

//Delete will remove the specified user from your Vultr account
func (u *UserServiceHandler) Delete(ctx context.Context, userID string) error {
	uri := fmt.Sprintf("%s/%s", path, userID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	err = u.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// List will list all the users associated with your Vultr account
func (u *UserServiceHandler) List(ctx context.Context, options *ListOptions) ([]User, *Meta, error) {
	req, err := u.client.NewRequest(ctx, http.MethodGet, path, nil)

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	if err != nil {
		return nil, nil, err
	}

	users := new(usersBase)
	err = u.client.DoWithContext(ctx, req, &users)
	if err != nil {
		return nil, nil, err
	}

	return users.Users, users.Meta, nil
}
