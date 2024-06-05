package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInferenceServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"subscriptions": [
				{
					"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
					"date_created": "2024-06-05T10:13:31+00:00",
					"label": "example-inference",
					"api_key": "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH"
				}
			]
		}`
		fmt.Fprint(writer, response)
	})

	inference, _, err := client.Inference.List(ctx)
	if err != nil {
		t.Errorf("Inference.List returned %+v", err)
	}

	expected := []Inference{
		{
			ID:          "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
			DateCreated: "2024-06-05T10:13:31+00:00",
			Label:       "example-inference",
			APIKey:      "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH",
		},
	}

	if !reflect.DeepEqual(inference, expected) {
		t.Errorf("Inference.List returned %+v, expected %+v", inference, expected)
	}
}

func TestInferenceServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"subscription": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2024-06-05T10:13:31+00:00",
				"label": "example-inference",
				"api_key": "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH"
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &InferenceCreateUpdateReq{
		Label: "example-inference",
	}

	inference, _, err := client.Inference.Create(ctx, options)
	if err != nil {
		t.Errorf("Inference.Create returned %+v", err)
	}

	expected := &Inference{
		ID:          "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated: "2024-06-05T10:13:31+00:00",
		Label:       "example-inference",
		APIKey:      "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH",
	}

	if !reflect.DeepEqual(inference, expected) {
		t.Errorf("Inference.Create returned %+v, expected %+v", inference, expected)
	}
}

func TestInferenceServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"subscription": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2024-06-05T10:13:31+00:00",
				"label": "example-inference",
				"api_key": "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH"
			}
		}`
		fmt.Fprint(writer, response)
	})

	inference, _, err := client.Inference.Get(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5")
	if err != nil {
		t.Errorf("Inference.Get returned %+v", err)
	}

	expected := &Inference{
		ID:          "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated: "2024-06-05T10:13:31+00:00",
		Label:       "example-inference",
		APIKey:      "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH",
	}

	if !reflect.DeepEqual(inference, expected) {
		t.Errorf("Inference.Get returned %+v, expected %+v", inference, expected)
	}
}

func TestInferenceServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"subscription": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2024-06-05T10:13:31+00:00",
				"label": "example-inference-updated",
				"api_key": "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH"
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &InferenceCreateUpdateReq{
		Label: "example-inference-updated",
	}

	inference, _, err := client.Inference.Update(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", options)
	if err != nil {
		t.Errorf("Inference.Update returned %+v", err)
	}

	expected := &Inference{
		ID:          "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated: "2024-06-05T10:13:31+00:00",
		Label:       "example-inference-updated",
		APIKey:      "ABCD7PQSLLGS6XDHQY4CMHUL55T5YO63EFGH",
	}

	if !reflect.DeepEqual(inference, expected) {
		t.Errorf("Inference.Update returned %+v, expected %+v", inference, expected)
	}
}

func TestInferenceServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Inference.Delete(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5")

	if err != nil {
		t.Errorf("Inference.Delete returned %+v", err)
	}
}

func TestInferenceServiceHandler_GetUsage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/inference/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5/usage", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"usage": {
				"chat": {
					"current_tokens": 654732,
					"monthly_allotment": 50000000,
					"overage": 0
				},
				"audio": {
					"tts_characters": 5678,
					"tts_sm_characters": 1234
				}
			}
		}`
		fmt.Fprint(writer, response)
	})

	usage, _, err := client.Inference.GetUsage(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5")
	if err != nil {
		t.Errorf("Inference.Get returned %+v", err)
	}

	chatUsage := InferenceChatUsage{
		CurrentTokens:    654732,
		MonthlyAllotment: 50000000,
		Overage:          0,
	}

	audioUsage := InferenceAudioUsage{
		TTSCharacters:   5678,
		TTSSMCharacters: 1234,
	}

	expected := &InferenceUsage{
		Chat:  chatUsage,
		Audio: audioUsage,
	}

	if !reflect.DeepEqual(usage, expected) {
		t.Errorf("Inference.GetUsage returned %+v, expected %+v", usage, expected)
	}
}
