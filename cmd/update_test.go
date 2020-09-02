package cmd

import (
	"github.com/razielt77/kyml/cmd/schema"
	"github.com/razielt77/kyml/cmd/utils"
	"io/ioutil"
	"os"
	"testing"
	"gopkg.in/yaml.v2"
)

var (
	name string ="myapp_deployment"
	app string ="my_app"
	image string ="my_image:0.1"
	replica int =2
	port int =8080
	path string ="temp.yaml"
	new_image string ="my_image:0.2"
)


func TestUpdate(t *testing.T) {
	var d schema.Deployment
	d.Init(name,app,image,replica,port)
	err := utils.MarshalAndSave(d,path)
	if err != nil {
		t.Errorf("Error saving, got: %s.", err)
	}
	updateCmdOptions.name = name
	updateCmdOptions.attribute = "image"
	updateCmdOptions.value = new_image
	updateCmdOptions.index = 0
	err = update("deployment",".")
	if err != nil {
		t.Errorf("Error uppdating file, error: %s.", err)
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to read file: %w", err)
	}
	var deployment schema.Deployment
	err = yaml.Unmarshal([]byte(yamlFile), &deployment)
	if err != nil {
		t.Errorf("Failed to unmarshal: %w", err)
	}
	if (*deployment.Spec.Template.Spec.Containers)[0].Image != new_image{
		t.Errorf("Update failed")
	}
	err = os.Remove(path)
	if err != nil {
		t.Errorf("Failed to remove %s err: %w", path,err)
	}
}
