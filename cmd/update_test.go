package cmd

import (
	"github.com/razielt77/kyml/cmd/schema"
	"github.com/razielt77/kyml/cmd/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

var (
	name_deployment string = "myapp_deployment"
	name_rollout    string = "myapp_rollout"
	app             string = "my_app"
	image           string = "my_image:0.1"
	replica         int    = 2
	port            int    = 8080
	path_deployment string = "temp_deployment.yaml"
	path_service string = "temp_service.yaml"
	path_rollout    string = "temp_rollout.yaml"
	new_image       string = "my_image:0.2"
)

func TestUpdate(t *testing.T) {
	var d schema.Deployment

	d.Init(name_deployment, app, image, replica, port)
	err := utils.MarshalAndSave(d, path_deployment)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = name_deployment
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = new_image
	updateCmdOptions.index = 0
	err = update("deployment", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(path_deployment)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var deployment schema.Deployment
	err = yaml.Unmarshal([]byte(yamlFile), &deployment)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*deployment.Spec.Template.Spec.Containers)[0].Image != new_image {
		t.Errorf("Update failed")
	}
	err = os.Remove(path_deployment)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", path_deployment, err)
	}
}

func TestRollout(t *testing.T) {
	var r schema.Rollout
	r.Init(name_rollout, app, image, replica, port)
	err := utils.MarshalAndSave(r, path_rollout)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = name_rollout
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = new_image
	updateCmdOptions.index = 0
	err = update("rollout", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(path_rollout)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var rollout schema.Rollout
	err = yaml.Unmarshal([]byte(yamlFile), &rollout)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*rollout.Spec.Template.Spec.Containers)[0].Image != new_image {
		t.Errorf("Update failed")
	}
	err = os.Remove(path_rollout)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", path_rollout, err)
	}
}
