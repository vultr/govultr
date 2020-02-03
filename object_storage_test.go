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

	mux.HandleFunc("/v1/objectstorage/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"SUBID": 1234}`
		fmt.Fprint(writer, response)
	})

	id, err := client.ObjectStorage.Create(ctx, 1, "s3 label")
	if err != nil {
		t.Errorf("ObjectStorage.Create returned %+v", err)
	}

	expected := &struct {
		ID int `json:"SUBID"`
	}{ID: 1234}

	if !reflect.DeepEqual(id, expected) {
		t.Errorf("ObjectStorage.Create returned %+v, expected %+v", id, expected)
	}
}

func TestObjectStorageServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/objectstorage/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ObjectStorage.Delete(ctx, 1234)
	if err != nil {
		t.Errorf("ObjectStorage.Delete returned %+v", err)
	}
}

func TestObjectStorageServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/objectstorage/list", func(writer http.ResponseWriter, request *http.Request) {
		//response := `{"1314217": {"SUBID": 1314217,"date_created": "2019-04-17 17:33:00","OBJSTORECLUSTERID": 1,"DCID": 1,"location": "New Jersey","label": "object1","status": "active","s3_hostname": "nj1.vultrobjects.com","s3_access_key": "abc1234","s3_secret_key": "def5678"}}`
		response := `{"SUBID": 1314217,"date_created": "2019-04-17 17:33:00","OBJSTORECLUSTERID": 1,"DCID": 1,"location": "New Jersey","label": "object1","status": "active","s3_hostname": "nj1.vultrobjects.com","s3_access_key": "abc1234","s3_secret_key": "def5678"}`
		fmt.Fprint(writer, response)
	})

	s3, err := client.ObjectStorage.Get(ctx, 1314217)

	if err != nil {
		t.Errorf("ObjectStorage.Get returned %+v", err)
	}

	expected := &ObjectStorage{
		ID:                   1314217,
		DateCreated:          "2019-04-17 17:33:00",
		ObjectStoreClusterID: 1,
		RegionID:             1,
		Location:             "New Jersey",
		Label:                "object1",
		Status:               "active",
		S3Keys: S3Keys{
			S3Hostname:  "nj1.vultrobjects.com",
			S3AccessKey: "abc1234",
			S3SecretKey: "def5678",
		},
	}
	if !reflect.DeepEqual(s3, expected) {
		t.Errorf("ObjectStorage.Get returned %+v, expected %+v", s3, expected)
	}
}

func TestObjectStorageServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/objectstorage/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1314217": {"SUBID": 1314217,"date_created": "2019-04-17 17:33:00","OBJSTORECLUSTERID": 1,"DCID": 1,"location": "New Jersey","label": "object1","status": "active","s3_hostname": "nj1.vultrobjects.com","s3_access_key": "abc1234","s3_secret_key": "def5678"}}`
		fmt.Fprint(writer, response)
	})

	options := &ObjectListOptions{
		IncludeS3: false,
		Label:     "label",
	}
	s3s, err := client.ObjectStorage.List(ctx, options)

	if err != nil {
		t.Errorf("ObjectStorage.List returned %+v", err)
	}
	expected := []ObjectStorage{
		{
			ID:                   1314217,
			DateCreated:          "2019-04-17 17:33:00",
			ObjectStoreClusterID: 1,
			RegionID:             1,
			Location:             "New Jersey",
			Label:                "object1",
			Status:               "active",
			S3Keys: S3Keys{
				S3Hostname:  "nj1.vultrobjects.com",
				S3AccessKey: "abc1234",
				S3SecretKey: "def5678",
			},
		},
	}
	if !reflect.DeepEqual(s3s, expected) {
		t.Errorf("ObjectStorage.List returned %+v, expected %+v", s3s, expected)
	}
}

func TestObjectStorageServiceHandler_ListCluster(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/objectstorage/list_cluster", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1": {"OBJSTORECLUSTERID": 1,"DCID": 1,"location": "New Jersey","hostname": "nj1.vultrobjects.com","deploy": "yes"}}`
		fmt.Fprint(writer, response)
	})

	clusterList, err := client.ObjectStorage.ListCluster(ctx)

	if err != nil {
		t.Errorf("ObjectStorage.ListCluster returned %+v", err)
	}

	expected := []ObjectStorageCluster{
		{
			ObjectStoreClusterID: 1,
			RegionID:             1,
			Location:             "New Jersey",
			Hostname:             "nj1.vultrobjects.com",
			Deploy:               "yes",
		},
	}

	if !reflect.DeepEqual(clusterList, expected) {
		t.Errorf("ObjectStorage.ListCluster returned %+v, expected %+v", clusterList, expected)
	}
}

func TestObjectStorageServiceHandler_RegenerateKeys(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/objectstorage/s3key_regenerate", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"s3_hostname": "nj1.vultrobjects.com","s3_access_key": "abc1236","s3_secret_key": "def5679"}`
		fmt.Fprint(writer, response)
	})

	s3Keys, err := client.ObjectStorage.RegenerateKeys(ctx, 1234, "acv123")

	if err != nil {
		t.Errorf("ObjectStorage.RegenerateKeys returned %+v", err)
	}

	expected := &S3Keys{
		S3Hostname:  "nj1.vultrobjects.com",
		S3AccessKey: "abc1236",
		S3SecretKey: "def5679",
	}

	if !reflect.DeepEqual(s3Keys, expected) {
		t.Errorf("ObjectStorage.RegenerateKeys returned %+v, expected %+v", s3Keys, expected)
	}
}

func TestObjectStorageServiceHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/objectstorage/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.ObjectStorage.SetLabel(ctx, 1, "s3 label")
	if err != nil {
		t.Errorf("ObjectStorage.Create returned %+v", err)
	}
}
