package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestObjectStorageServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/object-storage", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"object_storage":{"id":"39239784","date_created":"2020-07-1414:07:28","cluster_id":2,"region":"ewr","location":"New Jersey","label":"api-obj-storage2","status":"pending","s3_hostname":"","s3_access_key":"","s3_secret_key":""}}`
		fmt.Fprint(writer, response)
	})

	objectStorage,_,err := client.ObjectStorage.Create(ctx, 2, "api-obj-storage2")
	if err != nil {
		t.Errorf("ObjectStorage.Create returned %+v", err)
	}

	expected := &ObjectStorage{
		ID:                   "39239784",
		DateCreated:          "2020-07-1414:07:28",
		ObjectStoreClusterID: 2,
		Region:               "ewr",
		Location:             "New Jersey",
		Label:                "api-obj-storage2",
		Status:               "pending",
		S3Keys:               S3Keys{},
	}

	if !reflect.DeepEqual(objectStorage, expected) {
		t.Errorf("ObjectStorage.Create returned %+v, expected %+v", objectStorage, expected)
	}
}

func TestObjectStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/object-storage/39239784", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"object_storage":{"id":"39239784","date_created":"2020-07-1414:07:28","cluster_id":2,"region":"ewr","label":"","status":"active","s3_hostname":"ewr1.vultrobjects.com","s3_access_key":"F123","s3_secret_key":"F1234"}}`
		fmt.Fprint(writer, response)
	})

	s3,_, err := client.ObjectStorage.Get(ctx, "39239784")

	if err != nil {
		t.Errorf("ObjectStorage.Get returned %+v", err)
	}

	expected := &ObjectStorage{
		ID:                   "39239784",
		DateCreated:          "2020-07-1414:07:28",
		ObjectStoreClusterID: 2,
		Region:               "ewr",
		Label:                "",
		Status:               "active",
		S3Keys: S3Keys{
			S3Hostname:  "ewr1.vultrobjects.com",
			S3AccessKey: "F123",
			S3SecretKey: "F1234",
		},
	}
	if !reflect.DeepEqual(s3, expected) {
		t.Errorf("ObjectStorage.Get returned %+v, expected %+v", s3, expected)
	}
}

func TestObjectStorageServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/object-storage/1234", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ObjectStorage.Update(ctx, "1234", "s3 label")
	if err != nil {
		t.Errorf("ObjectStorage.Create returned %+v", err)
	}
}

func TestObjectStorageServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/object-storage/1234", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ObjectStorage.Delete(ctx, "1234")
	if err != nil {
		t.Errorf("ObjectStorage.Delete returned %+v", err)
	}
}

func TestObjectStorageServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/object-storage", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"object_storages":[{"id":"39240368","date_created":"2020-07-1414:22:38","cluster_id":2,"region":"ewr","label":"govultr","status":"active","s3_hostname":"ewr1.vultrobjects.com","s3_access_key":"n1234","s3_secret_key":"b1234"}],"meta":{"total":1,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	s3s, meta,_, err := client.ObjectStorage.List(ctx, nil)
	if err != nil {
		t.Errorf("ObjectStorage.List returned %+v", err)
	}

	expectedObject := []ObjectStorage{
		{
			ID:                   "39240368",
			DateCreated:          "2020-07-1414:22:38",
			ObjectStoreClusterID: 2,
			Region:               "ewr",
			Label:                "govultr",
			Status:               "active",
			S3Keys: S3Keys{
				S3Hostname:  "ewr1.vultrobjects.com",
				S3AccessKey: "n1234",
				S3SecretKey: "b1234",
			},
		},
	}

	if !reflect.DeepEqual(s3s, expectedObject) {
		t.Errorf("ObjectStorage.List object returned %+v, expected %+v", s3s, expectedObject)
	}

	expectedmeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(meta, expectedmeta) {
		t.Errorf("ObjectStorage.List meta object returned %+v, expected %+v", meta, expectedmeta)
	}
}

func TestObjectStorageServiceHandler_ListCluster(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/object-storage/clusters", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"clusters":[{"id":2,"region":"ewr","hostname":"ewr1.vultrobjects.com","deploy":"yes"}],"meta":{"total":1,"links":{"next":"","prev":""}}}`
		fmt.Fprint(writer, response)
	})

	clusters, meta,_, err := client.ObjectStorage.ListCluster(ctx, nil)

	if err != nil {
		t.Errorf("ObjectStorage.ListCluster returned %+v", err)
	}

	expected := []ObjectStorageCluster{
		{
			ID:       2,
			Region:   "ewr",
			Hostname: "ewr1.vultrobjects.com",
			Deploy:   "yes",
		},
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("ObjectStorage.ListCluster clusters returned %+v, expected %+v", clusters, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("ObjectStorage.List meta object returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestObjectStorageServiceHandler_RegenerateKeys(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v2/object-storage/1234/regenerate-keys", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"s3_credentials":{"s3_hostname":"ewr1.vultrobjects.com","s3_access_key":"f1234","s3_secret_key":"g1234"}}`
		fmt.Fprint(writer, response)
	})

	s3Keys,_, err := client.ObjectStorage.RegenerateKeys(ctx, "1234")

	if err != nil {
		t.Errorf("ObjectStorage.RegenerateKeys returned %+v", err)
	}

	expected := &S3Keys{
		S3Hostname:  "ewr1.vultrobjects.com",
		S3AccessKey: "f1234",
		S3SecretKey: "g1234",
	}

	if !reflect.DeepEqual(s3Keys, expected) {
		t.Errorf("ObjectStorage.RegenerateKeys returned %+v, expected %+v", s3Keys, expected)
	}
}
