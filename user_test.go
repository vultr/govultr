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

	mux.HandleFunc("/v2/users", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"user": {"id": "564a1a88947b4","name": "Example User","email": "example@vultr.com","api_key": "aaavvvvvvbbbbbb","api_enabled": true,"acls": []}}`

		fmt.Fprint(writer, response)
	})
	api := true
	userReq := &UserReq{
		Email:      "example@vultr.com",
		Name:       "Example User",
		APIEnabled: &api,
		Password:   "password",
	}

	user, err := client.User.Create(ctx, userReq)

	if err != nil {
		t.Errorf("User.Create returned %+v, expected %+v", err, nil)
	}

	expected := &User{
		ID:         "564a1a88947b4",
		Name:       "Example User",
		Email:      "example@vultr.com",
		APIEnabled: BoolToBoolPtr(true),
		APIKey:     "aaavvvvvvbbbbbb",
		ACL:        []string{},
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.Create returned %+v, expected %+v", user, expected)
	}
}

func TestUserServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/123abc", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.User.Delete(ctx, "123abc")

	if err != nil {
		t.Errorf("User.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestUserServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users", func(writer http.ResponseWriter, request *http.Request) {
		response := `
	{
    "users": [
        {
            "id": "f255efc9700d9",
            "name": "test api",
            "email": "newmanapi@vultr.com",
            "api_enabled": true,
            "acls": []
        }
    ],
    "meta": {
        "total": 1,
        "links": {
            "next": "thisismycusror",
            "prev": ""
        }
    }
}
		`
		fmt.Fprintf(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	users, meta, err := client.User.List(ctx, options)

	if err != nil {
		t.Errorf("User.List returned error: %v", err)
	}

	expected := []User{
		{
			ID:         "f255efc9700d9",
			Name:       "test api",
			Email:      "newmanapi@vultr.com",
			APIEnabled: BoolToBoolPtr(true),
			ACL:        []string{},
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("User.List users returned %+v, expected %+v", users, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "thisismycusror",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("User.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestUserServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	api := true
	user := &UserReq{
		Name:       "Example User",
		Password:   "w1a4dcnst0n!",
		Email:      "example@vultr.com",
		APIEnabled: &api,
		ACL:        []string{"support"},
	}

	err := client.User.Update(ctx, "abc123", user)

	if err != nil {
		t.Errorf("User.Update returned error: %+v", err)
	}
}

func TestUserServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"user": {"id": "f255efc9c69ac","name": "Unit Test","email": "test@vultr.com","api_enabled": true,"acls": []}}`
		fmt.Fprint(writer, response)
	})

	user, err := client.User.Get(ctx, "abc123")
	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}
	expected := &User{
		ID:         "f255efc9c69ac",
		Name:       "Unit Test",
		Email:      "test@vultr.com",
		APIEnabled: BoolToBoolPtr(true),
		ACL:        []string{},
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.List users returned %+v, expected %+v", user, expected)
	}
}
