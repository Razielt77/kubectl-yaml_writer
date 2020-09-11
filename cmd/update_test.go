package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

var (
	testDeploymentName string = "myapp_deployment"
	testRolloutName    string = "myapp_rollout"
	testApp            string = "my_app"
	testImage          string = "my_image:0.1"
	testReplica        int    = 2
	testPortNumber     int    = 8080
	testDeploymentPath string = "temp_deployment.yaml"
	testRolloutPath    string = "temp_rollout.yaml"
	testNewImage       string = "my_image:0.2"
)

func TestUpdate(t *testing.T) {
	var d deployment

	d.Init(testDeploymentName, testApp, testImage, testReplica, testPortNumber)
	err := marshalAndSave(d, testDeploymentPath)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = testDeploymentName
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = testNewImage
	updateCmdOptions.index = 0
	err = update("deployment", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(testDeploymentPath)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var dp deployment
	err = yaml.Unmarshal([]byte(yamlFile), &dp)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*dp.Spec.Template.Spec.Containers)[0].Image != testNewImage {
		t.Errorf("Update failed")
	}
	err = os.Remove(testDeploymentPath)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", testDeploymentPath, err)
	}
}

func TestRollout(t *testing.T) {
	var r rollout
	r.Init(testRolloutName, testApp, testImage, testReplica, testPortNumber)
	err := marshalAndSave(r, testRolloutPath)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = testRolloutName
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = testNewImage
	updateCmdOptions.index = 0
	err = update("rollout", ".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(testRolloutPath)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var ro rollout
	err = yaml.Unmarshal([]byte(yamlFile), &ro)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*ro.Spec.Template.Spec.Containers)[0].Image != testNewImage {
		t.Errorf("Update failed")
	}
	err = os.Remove(testRolloutPath)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", testRolloutPath, err)
	}
}
