package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAPIServiceHandler_GetInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/auth/info", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
    	"acls": [
        	"subscriptions",
        	"billing",
        	"support",
        	"provisioning"
    	],
    	"email": "example@vultr.com",
    	"name": "Example Account"
		}`
		fmt.Fprint(writer, response)
	})

	apiAuth, err := client.API.GetInfo(ctx)

	if err != nil {
		t.Errorf("Account.GetInfo returned error: %v", err)
	}

	expected := &API{ACL: []string{"subscriptions", "billing", "support", "provisioning"}, Email: "example@vultr.com", Name: "Example Account"}
	if !reflect.DeepEqual(apiAuth, expected) {
		t.Errorf("API.GetInfo returned %+v, expected %+v", apiAuth, expected)
	}
}
