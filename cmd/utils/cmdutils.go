package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
)

func DieOnError(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}

func MarshalAndSave(in interface{}, path string) error{
	data, err := yaml.Marshal(&in)
	if err == nil {
		err = ioutil.WriteFile(path, data, 0644)
	}
	return err
}