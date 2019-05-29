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

	mux.HandleFunc("/v1/sshkey/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SSHKEYID": "541b4960f23bd"
		}
		`

		fmt.Fprint(writer, response)
	})

	key, err := client.SSHKey.Create(ctx, "foo", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname")

	if err != nil {
		t.Errorf("SSHKey.Create returned %+v, expected %+v", err, nil)
	}

	expected := &SSHKey{
		SSHKeyID:    "541b4960f23bd",
		Name:        "foo",
		Key:         "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname",
		DateCreated: "",
	}

	if !reflect.DeepEqual(key, expected) {
		t.Errorf("SSHKey.Create returned %+v, expected %+v", key, expected)
	}
}

func TestSSHKeyServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sshkey/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.SSHKey.Delete(ctx, "foo")

	if err != nil {
		t.Errorf("SSHKey.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestSSHKeyServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sshkey/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"541b4960f23bd": {
				"SSHKEYID": "541b4960f23bd",
				"date_created": null,
				"name": "test",
				"ssh_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname"
			}
		}
		`
		fmt.Fprintf(writer, response)
	})

	sshKeys, err := client.SSHKey.List(ctx)

	if err != nil {
		t.Errorf("SSHKey.List returned error: %v", err)
	}

	expected := []SSHKey{
		{
			SSHKeyID:    "541b4960f23bd",
			Name:        "test",
			Key:         "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname",
			DateCreated: "",
		},
	}

	if !reflect.DeepEqual(sshKeys, expected) {
		t.Errorf("SSHKey.List returned %+v, expected %+v", sshKeys, expected)
	}
}

func TestSSHKeyServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sshkey/update", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writer)
	})

	sshKey := &SSHKey{
		SSHKeyID:    "561b4960f23cc",
		Name:        "foo",
		Key:         "ssh-rsa CCCCB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname",
		DateCreated: "",
	}

	err := client.SSHKey.Update(ctx, sshKey)

	if err != nil {
		t.Errorf("SSHKey.Update returned error: %+v", err)
	}
}
