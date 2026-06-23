package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInstanceTemplateServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "instance_template": {
			"id": "98765432-10ab-cdef-1234-567890abcdef",
			"plan": "vc2-6c-16gb",
			"label": "my-template",
			"os": "Ubuntu 22.04 x64",
			"marketplace_app": "",
			"marketplace_image": "",
			"snapshot": "",
			"iso": "",
			"ssh_keys": [],
			"startup_script": "",
			"disk_config": ""
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &InstanceTemplateReq{
		Plan:  "vc2-6c-16gb",
		Label: "my-template",
		OsID:  1743,
	}

	template, _, err := client.InstanceTemplate.Create(ctx, options)
	if err != nil {
		t.Errorf("InstanceTemplate.Create returned %+v", err)
	}

	expected := &InstanceTemplate{
		ID:      "98765432-10ab-cdef-1234-567890abcdef",
		Plan:    "vc2-6c-16gb",
		Label:   "my-template",
		OS:      "Ubuntu 22.04 x64",
		SSHKeys: []InstanceTemplateSSHKey{},
	}

	if !reflect.DeepEqual(template, expected) {
		t.Errorf("InstanceTemplate.Create returned %+v, expected %+v", template, expected)
	}
}

func TestInstanceTemplateServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates/98765432-10ab-cdef-1234-567890abcdef", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "instance_template": {
			"id": "98765432-10ab-cdef-1234-567890abcdef",
			"plan": "vc2-6c-16gb",
			"label": "my-template",
			"os": "Ubuntu 22.04 x64",
			"marketplace_app": "",
			"marketplace_image": "",
			"snapshot": "",
			"iso": "",
			"ssh_keys": [],
			"startup_script": "",
			"disk_config": ""
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	template, _, err := client.InstanceTemplate.Get(ctx, "98765432-10ab-cdef-1234-567890abcdef")
	if err != nil {
		t.Errorf("InstanceTemplate.Get returned %+v", err)
	}

	expected := &InstanceTemplate{
		ID:      "98765432-10ab-cdef-1234-567890abcdef",
		Plan:    "vc2-6c-16gb",
		Label:   "my-template",
		OS:      "Ubuntu 22.04 x64",
		SSHKeys: []InstanceTemplateSSHKey{},
	}

	if !reflect.DeepEqual(template, expected) {
		t.Errorf("InstanceTemplate.Get returned %+v, expected %+v", template, expected)
	}
}

func TestInstanceTemplateServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates/98765432-10ab-cdef-1234-567890abcdef", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "instance_template": {
			"id": "98765432-10ab-cdef-1234-567890abcdef",
			"plan": "vc2-6c-16gb",
			"label": "renamed-template",
			"os": "Ubuntu 22.04 x64",
			"marketplace_app": "",
			"marketplace_image": "",
			"snapshot": "",
			"iso": "",
			"ssh_keys": [],
			"startup_script": "",
			"disk_config": ""
		  }
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &InstanceTemplateReq{
		Label: "renamed-template",
	}

	template, _, err := client.InstanceTemplate.Update(ctx, "98765432-10ab-cdef-1234-567890abcdef", options)
	if err != nil {
		t.Errorf("InstanceTemplate.Update returned %+v, expected %+v", err, nil)
	}

	expected := &InstanceTemplate{
		ID:      "98765432-10ab-cdef-1234-567890abcdef",
		Plan:    "vc2-6c-16gb",
		Label:   "renamed-template",
		OS:      "Ubuntu 22.04 x64",
		SSHKeys: []InstanceTemplateSSHKey{},
	}

	if !reflect.DeepEqual(template, expected) {
		t.Errorf("InstanceTemplate.Get returned %+v, expected %+v", template, expected)
	}
}

func TestInstanceTemplateServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates/98765432-10ab-cdef-1234-567890abcdef", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.InstanceTemplate.Delete(ctx, "98765432-10ab-cdef-1234-567890abcdef")

	if err != nil {
		t.Errorf("InstanceTemplate.Delete returned %+v", err)
	}
}

func TestInstanceTemplateServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "instance_templates": [
			{
			  "id": "98765432-10ab-cdef-1234-567890abcdef",
			  "plan": "vc2-6c-16gb",
			  "label": "my-template",
			  "os": "Ubuntu 22.04 x64",
			  "marketplace_app": "",
			  "marketplace_image": "",
			  "snapshot": "",
			  "iso": "",
			  "ssh_keys": [
				{
				  "id": "3a7d56e1-8c2a-4721-84dc-77fa4e6e4f91",
				  "name": "my-ssh-key"
				}
			  ],
			  "startup_script": "",
			  "disk_config": ""
			}
		  ]
		}		
		`

		fmt.Fprint(writer, response)
	})

	templates, _, err := client.InstanceTemplate.List(ctx)
	if err != nil {
		t.Errorf("InstanceTemplate.List returned %+v", err)
	}

	expected := []InstanceTemplate{
		{
			ID:    "98765432-10ab-cdef-1234-567890abcdef",
			Plan:  "vc2-6c-16gb",
			Label: "my-template",
			OS:    "Ubuntu 22.04 x64",
			SSHKeys: []InstanceTemplateSSHKey{
				{
					ID:   "3a7d56e1-8c2a-4721-84dc-77fa4e6e4f91",
					Name: "my-ssh-key",
				},
			},
		},
	}

	if !reflect.DeepEqual(templates, expected) {
		t.Errorf("InstanceTemplate.List returned %+v, expected %+v", templates, expected)
	}
}

func TestInstanceTemplateServiceHandler_CreateFromInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/instances/templates/from-instance/a1b2c3d4", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
		  "instance_template": {
			"id": "98765432-10ab-cdef-1234-567890abcdef",
			"plan": "vc2-6c-16gb",
			"label": "Template: my-instance",
			"os": "Snapshot",
			"marketplace_app": "",
			"marketplace_image": "",
			"snapshot": "a1b2c3d4",
			"iso": "",
			"ssh_keys": [],
			"startup_script": "",
			"disk_config": ""
		  }
		}
`

		fmt.Fprint(writer, response)
	})

	template, _, err := client.InstanceTemplate.CreateFromInstance(ctx, "a1b2c3d4")
	if err != nil {
		t.Errorf("InstanceTemplate.CreateFromInstance returned %+v", err)
	}

	expected := &InstanceTemplate{
		ID:       "98765432-10ab-cdef-1234-567890abcdef",
		Plan:     "vc2-6c-16gb",
		Label:    "Template: my-instance",
		OS:       "Snapshot",
		Snapshot: "a1b2c3d4",
		SSHKeys:  []InstanceTemplateSSHKey{},
	}

	if !reflect.DeepEqual(template, expected) {
		t.Errorf("InstanceTemplate.CreateFromInstance returned %+v, expected %+v", template, expected)
	}
}
