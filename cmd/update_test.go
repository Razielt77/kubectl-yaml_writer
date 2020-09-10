package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

var (
	nameDeployment string = "myapp_deployment"
	nameRollout    string = "myapp_rollout"
	app             string = "my_app"
	image           string = "my_image:0.1"
	replica         int    = 2
	port            int    = 8080
	pathDeployment string = "temp_deployment.yaml"
	pathService    string = "temp_service.yaml"
	pathRollout    string = "temp_rollout.yaml"
	newImage       string = "my_image:0.2"
)

func TestUpdate(t *testing.T) {
	var d deployment

	d.Init(nameDeployment, app, image, replica, port)
	err := marshalAndSave(d, pathDeployment)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = nameDeployment
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = newImage
	updateCmdOptions.index = 0
	err = update("deployment", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(pathDeployment)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var dp deployment
	err = yaml.Unmarshal([]byte(yamlFile), &dp)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*dp.Spec.Template.Spec.Containers)[0].Image != newImage {
		t.Errorf("Update failed")
	}
	err = os.Remove(pathDeployment)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", pathDeployment, err)
	}
}

func TestRollout(t *testing.T) {
	var r rollout
	r.Init(nameRollout, app, image, replica, port)
	err := marshalAndSave(r, pathRollout)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = nameRollout
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = newImage
	updateCmdOptions.index = 0
	err = update("rollout", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(pathRollout)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var ro rollout
	err = yaml.Unmarshal([]byte(yamlFile), &ro)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*ro.Spec.Template.Spec.Containers)[0].Image != newImage {
		t.Errorf("Update failed")
	}
	err = os.Remove(pathRollout)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", pathRollout, err)
	}
}
