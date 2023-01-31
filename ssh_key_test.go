package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSSHKeyServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/ssh-keys", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"ssh_key": {"id": "5f05d5a71fe28","date_created": "2020-07-08 14:18:15","name": "api-test-ssh","ssh_key": "ssh-rsa AF+LbfYYw== test@admin.com"}}`
		fmt.Fprint(writer, response)
	})

	sshKey := &SSHKeyReq{
		Name:   "api-test-ssh",
		SSHKey: "ssh-rsa AF+LbfYYw== test@admin.com",
	}

	key,_, err := client.SSHKey.Create(ctx, sshKey)
	if err != nil {
		t.Errorf("SSHKey.Create returned %+v, expected %+v", err, nil)
	}

	expected := &SSHKey{
		ID:          "5f05d5a71fe28",
		Name:        "api-test-ssh",
		SSHKey:      "ssh-rsa AF+LbfYYw== test@admin.com",
		DateCreated: "2020-07-08 14:18:15",
	}

	if !reflect.DeepEqual(key, expected) {
		t.Errorf("SSHKey.Create returned %+v, expected %+v", key, expected)
	}
}

func TestSSHKeyServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/ssh-keys/abc123", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"ssh_key": {"id": "5f05d5a71fe28","date_created": "2020-07-08 14:18:15","name": "api-test-ssh","ssh_key": "ssh-rsa AF+LbfYYw== test@admin.com"}}`
		fmt.Fprint(writer, response)
	})

	key,_, err := client.SSHKey.Get(ctx, "abc123")
	if err != nil {
		t.Errorf("SSHKey.Get returned %+v, expected %+v", err, nil)
	}

	expected := &SSHKey{
		ID:          "5f05d5a71fe28",
		Name:        "api-test-ssh",
		SSHKey:      "ssh-rsa AF+LbfYYw== test@admin.com",
		DateCreated: "2020-07-08 14:18:15",
	}

	if !reflect.DeepEqual(key, expected) {
		t.Errorf("SSHKey.Create returned %+v, expected %+v", key, expected)
	}
}

func TestSSHKeyServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/ssh-keys/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	sshKey := &SSHKeyReq{
		Name:   "foo",
		SSHKey: "ssh-rsa CCCCB3NzaC1yc your_username@hostname",
	}

	err := client.SSHKey.Update(ctx, "abc123", sshKey)

	if err != nil {
		t.Errorf("SSHKey.Update returned error: %+v", err)
	}
}

func TestSSHKeyServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/ssh-keys/abc123", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.SSHKey.Delete(ctx, "abc123")

	if err != nil {
		t.Errorf("SSHKey.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestSSHKeyServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/ssh-keys", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"ssh_keys": [{"id": "5ed139d1890db","date_created": "2020-05-29 16:35:29","name": "api-test-ssh","ssh_key": "ssh-rsa AAAAB3NzaC1ycYYw== test@admin.com"}],"meta": {"total": 8,"links": {"next": "","prev": ""}}}`
		fmt.Fprint(writer, response)
	})

	sshKeys, meta,_, err := client.SSHKey.List(ctx, nil)
	if err != nil {
		t.Errorf("SSHKey.List returned error: %v", err)
	}

	expectedSSH := []SSHKey{
		{
			ID:          "5ed139d1890db",
			Name:        "api-test-ssh",
			SSHKey:      "ssh-rsa AAAAB3NzaC1ycYYw== test@admin.com",
			DateCreated: "2020-05-29 16:35:29",
		},
	}

	expectedMeta := &Meta{
		Total: 8,
		Links: &Links{},
	}

	if !reflect.DeepEqual(sshKeys, expectedSSH) {
		t.Errorf("SSHKey.List ssh-keys returned %+v, expected %+v", sshKeys, expectedSSH)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("SSHKey.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}
