package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/user/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"USERID": "564a1a88947b4",
			"api_key": "AAAAAAAA"
		}
		`

		fmt.Fprint(writer, response)
	})

	user, err := client.User.Create(ctx, "example@vultr.com", "Example User", "t0rbj0rn!", "no", []string{"support", "abuse", "alerts"})

	if err != nil {
		t.Errorf("User.Create returned %+v, expected %+v", err, nil)
	}

	expected := &User{
		UserID:     "564a1a88947b4",
		Name:       "Example User",
		Email:      "example@vultr.com",
		APIEnabled: "no",
		ACL:        []string{"support", "abuse", "alerts"},
		APIKey:     "AAAAAAAA",
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.Create returned %+v, expected %+v", user, expected)
	}
}

func TestUserServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/user/delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.User.Delete(ctx, "foo")

	if err != nil {
		t.Errorf("User.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestUserServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/user/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		[
			{
				"USERID": "564a1a7794d83",
				"name": "example user 1",
				"email": "example@vultr.com",
				"api_enabled": "yes",
				"acls": [
					"manage_users",
					"subscriptions",
					"billing",
					"provisioning"
				]
			},
			{
				"USERID": "564a1a88947b4",
				"name": "example user 2",
				"email": "example@vultr.com",
				"api_enabled": "no",
				"acls": [
					"support",
					"dns"
				]
			}
		]
		`
		fmt.Fprintf(writer, response)
	})

	Users, err := client.User.GetList(ctx)

	if err != nil {
		t.Errorf("User.List returned error: %v", err)
	}

	expected := []User{
		{
			UserID:     "564a1a7794d83",
			Name:       "example user 1",
			Email:      "example@vultr.com",
			APIEnabled: "yes",
			ACL:        []string{"manage_users", "subscriptions", "billing", "provisioning"},
		},
		{
			UserID:     "564a1a88947b4",
			Name:       "example user 2",
			Email:      "example@vultr.com",
			APIEnabled: "no",
			ACL:        []string{"support", "dns"},
		},
	}

	if !reflect.DeepEqual(Users, expected) {
		t.Errorf("User.GetList returned %+v, expected %+v", Users, expected)
	}
}

func TestUserServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/user/update", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writer)
	})

	user := &User{
		UserID: "2e35cc32f9923",
		Email:  "example@vultr.com",
		ACL:    []string{"support"},
	}

	err := client.User.Update(ctx, user)

	if err != nil {
		t.Errorf("User.Update returned error: %+v", err)
	}
}
