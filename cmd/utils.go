package cmd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func dieOnError(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}

func marshalAndSave(in interface{}, path string) error {
	data, err := yaml.Marshal(&in)
	if err == nil {
		err = ioutil.WriteFile(path, data, 0644)
	}
	return err
}
