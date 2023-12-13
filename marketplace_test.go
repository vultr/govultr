package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMarketplaceServiceHandler_ListAppVariables(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/v2/marketplace/apps/%s/variables", "testimage"), func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"variables": [
				{
					"name": "some_required_variable",
					"description": "This is an example of a required user-supplied variable for this Marketplace app.",
					"required": true
				},
				{
					"name": "some_optional_variable",
					"description": "This is an example of an optional user-supplied variable for this Marketplace app.",
					"required": false
				}
			]
		}`
		fmt.Fprint(writer, response)
	})

	variables, _, err := client.Marketplace.ListAppVariables(ctx, "testimage")
	if err != nil {
		t.Errorf("Marketplace.ListAppVariables returned %+v", err)
	}

	expected := []MarketplaceAppVariable{
		{
			Name:        "some_required_variable",
			Description: "This is an example of a required user-supplied variable for this Marketplace app.",
			Required:    BoolToBoolPtr(true),
		},
		{
			Name:        "some_optional_variable",
			Description: "This is an example of an optional user-supplied variable for this Marketplace app.",
			Required:    BoolToBoolPtr(false),
		},
	}

	if !reflect.DeepEqual(variables, expected) {
		t.Errorf("Marketplace.ListAppVariables returned %+v, expected %+v", variables, expected)
	}
}
