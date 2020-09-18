package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {

	directory := "."
	createCmdOptions.name = testApp
	createCmdOptions.targetPort = testPortNumber
	createCmdOptions.externalPort = testPortNumber
	createCmdOptions.image = testImage
	err := create("app", directory)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}

	filename := constructDeploymentFilename(directory, createCmdOptions.name)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var d deployment
	err = yaml.Unmarshal([]byte(yamlFile), &d)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}

	if (*d.Spec.Template.Spec.Containers)[0].Image != createCmdOptions.image {
		t.Errorf("Image name in object dont match Desired:%s  Found:%s", createCmdOptions.image, (*d.Spec.Template.Spec.Containers)[0].Image)
	}

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", filename, err)
	}

	filename = constructServiceFilename(directory, createCmdOptions.name)

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", filename, err)
	}

}
