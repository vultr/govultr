package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAPIKeyServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/apikeys", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "api_key": {
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"api_key": "000000000000000000000000000000012345",
			"name": "Production",
			"expire": true,
			"date_expire": "2030-01-01T00:00:00Z"
		  }
		}
		`

		fmt.Fprint(writer, response)
	})
	apiKeyReq := &APIKeyCreate{
		Name:       "Production",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}
	apiKey, _, err := client.APIKey.Create(ctx, apiKeyReq)
	if err != nil {
		t.Errorf("APIKey.Create returned error: %v", err)
	}

	expected := &APIKey{
		ID:         "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		APIKey:     "000000000000000000000000000000012345",
		Name:       "Production",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}

	if !reflect.DeepEqual(apiKey, expected) {
		t.Errorf("APIKey.Create returned %+v, expected %+v", apiKey, expected)
	}
}

func TestAPIKeyServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/apikeys/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "api_key": {
			"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			"api_key": "*******************************00000",
			"name": "Default",
			"expire": true,
			"date_expire": "2030-01-01T00:00:00Z"
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	apiKey, _, err := client.APIKey.Get(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("APIKey.Get returned error: %v", err)
	}

	expected := &APIKey{
		ID:         "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		APIKey:     "*******************************00000",
		Name:       "Default",
		Expire:     true,
		DateExpire: "2030-01-01T00:00:00Z",
	}

	if !reflect.DeepEqual(apiKey, expected) {
		t.Errorf("APIKey.Get returned %+v, expected %+v", apiKey, expected)
	}
}

func TestAPIKeyServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/apikeys/cb676a46-66fd-4dfb-b839-443f2e6c0b60", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.APIKey.Delete(ctx, "cb676a46-66fd-4dfb-b839-443f2e6c0b60")
	if err != nil {
		t.Errorf("APIKey.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestAPIKeyServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/apikeys", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "api_keys": [
			{
			  "id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			  "api_key": "*******************************00000",
			  "name": "Default",
			  "expire": true,
			  "date_expire": "2030-01-01T00:00:00Z"
			}
		  ]
		}
		`

		fmt.Fprint(writer, response)
	})

	apiKeys, _, err := client.APIKey.List(ctx, nil)
	if err != nil {
		t.Errorf("APIKey.List returned error: %v", err)
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
		t.Errorf("APIKey.List returned %+v, expected %+v", apiKeys, expected)
	}
}
