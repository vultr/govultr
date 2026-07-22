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

	user, _, err := client.User.Create(ctx, userReq)

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
			"first_name": "Example",
			"last_name": "User",
            "email": "newmanapi@vultr.com",
            "api_enabled": true,
            "acls": [],
			"invited_by": "admin@example.com",
			"invited_on": "2026-01-01T12:00:00+00:00",
			"invite_accepted": true,
			"status": "active",
			"last_activity": "2024-10-15T12:34:56+00:00"
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
		fmt.Fprint(writer, response)
	})

	options := &ListOptions{
		PerPage: 1,
	}
	users, meta, _, err := client.User.List(ctx, options)

	if err != nil {
		t.Errorf("User.List returned error: %v", err)
	}

	expected := []User{
		{
			ID:             "f255efc9700d9",
			Name:           "test api",
			FirstName:      "Example",
			LastName:       "User",
			Email:          "newmanapi@vultr.com",
			APIEnabled:     BoolToBoolPtr(true),
			ACL:            []string{},
			InvitedBy:      "admin@example.com",
			InvitedOn:      "2026-01-01T12:00:00+00:00",
			InviteAccepted: true,
			Status:         "active",
			LastActivity:   "2024-10-15T12:34:56+00:00",
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

	user, _, err := client.User.Get(ctx, "abc123")
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
func TestUserServiceHandler_CreateAPIKey(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123/apikeys", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "api_key": {
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"api_key": "*******************************00000",
			"name": "Default",
			"expire": true,
			"date_expire": "2030-01-01T00:00:00Z"
		  }
		}`
		fmt.Fprint(writer, response)
	})

	req := &APIKeyCreate{
		Name:       "Default",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}

	apiKey, _, err := client.User.CreateAPIKey(ctx, "abc123", req)
	if err != nil {
		t.Errorf("User.CreateAPIKey returned error: %v", err)
	}
	expected := &APIKey{
		ID:         "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		APIKey:     "*******************************00000",
		Name:       "Default",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}

	if !reflect.DeepEqual(apiKey, expected) {
		t.Errorf("User.CreateAPIKey returned %+v, expected %+v", apiKey, expected)
	}
}

func TestUserServiceHandler_GetAPIKey(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123/apikeys/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "api_key": {
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"api_key": "*******************************00000",
			"name": "Default",
			"expire": true,
			"date_expire": "2030-01-01T00:00:00Z"
		  }
		}`
		fmt.Fprint(writer, response)
	})

	apiKey, _, err := client.User.GetAPIKey(ctx, "abc123", "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("User.GetAPIKey returned error: %v", err)
	}
	expected := &APIKey{
		ID:         "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		APIKey:     "*******************************00000",
		Name:       "Default",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}

	if !reflect.DeepEqual(apiKey, expected) {
		t.Errorf("User.GetAPIKey returned %+v, expected %+v", apiKey, expected)
	}
}

func TestUserServiceHandler_DeleteAPIKey(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123/apikeys/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.User.DeleteAPIKey(ctx, "abc123", "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("User.DeleteAPIKey returned error: %v", err)
	}
}

func TestUserServiceHandler_ListAPIKeys(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/users/abc123/apikeys", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "api_keys": [
			{
			  "id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			  "api_key": "*******************************00000",
			  "name": "Default",
			  "expire": true,
			  "date_expire": "2030-01-01T00:00:00Z"
			}
		  ]
		}`
		fmt.Fprint(writer, response)
	})

	apiKeys, _, err := client.User.ListAPIKeys(ctx, "abc123")
	if err != nil {
		t.Errorf("User.ListAPIKeys returned error: %v", err)
	}
	expected := []APIKey{
		{
			ID:         "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			APIKey:     "*******************************00000",
			Name:       "Default",
			Expire:     true,
			DateExpire: "2030-01-01T00:00:00Z",
		},
	}

	if !reflect.DeepEqual(apiKeys, expected) {
		t.Errorf("User.ListAPIKeys returned %+v, expected %+v", apiKeys, expected)
	}
}
