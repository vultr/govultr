package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLogsServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/logs", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"logs":[{"resource_id":"xb671a46-66ed-4dfb-b839-543f2c6c0b63","resource_type":"kubernetes","log_level":"Debug","message":"Success","timestamp":"2025-08-26T00:00:07+00:00","metadata":{"user_id":"765b8aa0-e134-4b62-88e1-22f40959ffe7","ip_address":"123.123.123.123","username":"","http_status_code":200,"method":"GET","request_path":"/v2/kubernetes/clusters/xb671a46-66ed-4dfb-b839-543f2c6c0b63/node-pools","request_body": "","query_parameters": ""}}],"meta":{"continue_time": "","returned_count":1,"unreturned_count":0,"total_count":1}}`
		fmt.Fprint(writer, response)
	})

	logs, meta, _, err := client.Logs.List(ctx, LogsOptions{
		StartTime:  "2025-08-26T00:00:00Z",
		EndTime:    "2025-08-26T00:00:10Z",
		ResourceID: "xb671a46-66ed-4dfb-b839-543f2c6c0b63",
	})
	if err != nil {
		t.Errorf("Logs.List returned error: %v", err)
	}

	expectedLogs := []Log{{
		ResourceID:   "xb671a46-66ed-4dfb-b839-543f2c6c0b63",
		ResourceType: "kubernetes",
		Level:        "Debug",
		Message:      "Success",
		Timestamp:    "2025-08-26T00:00:07+00:00",
		Metadata: LogMetadata{
			UserID:          "765b8aa0-e134-4b62-88e1-22f40959ffe7",
			IPAddress:       "123.123.123.123",
			UserName:        "",
			HTTPStatusCode:  200,
			Method:          "GET",
			RequestPath:     "/v2/kubernetes/clusters/xb671a46-66ed-4dfb-b839-543f2c6c0b63/node-pools",
			RequestBody:     "",
			QueryParameters: "",
		},
	}}

	expectedMeta := &LogsMeta{
		ContinueTime:    "",
		ReturnedCount:   1,
		UnreturnedCount: 0,
		TotalCount:      1,
	}

	if !reflect.DeepEqual(logs, expectedLogs) {
		t.Errorf("Logs.List returned %+v, expected %+v", logs, expectedLogs)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Logs.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
