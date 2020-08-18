package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	updateCmdOptions struct {
		file      string
		kind      string
		name      string
		attribute string
		value     string
		index     int
	}
)

var updateCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(updateCmdOptions.file, updateCmdOptions.kind, updateCmdOptions.name, updateCmdOptions.attribute, updateCmdOptions.value, updateCmdOptions.index)
		dieOnError("Failed to update", err)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&updateCmdOptions.file, "file", "", "File to update")
	updateCmd.Flags().StringVar(&updateCmdOptions.kind, "kind", "", "Kind of resource to update.")
	updateCmd.Flags().StringVar(&updateCmdOptions.name, "name", "", "Name of the resource to update.")
	updateCmd.Flags().StringVar(&updateCmdOptions.attribute, "att", "", "Name of the attribute to update.")
	updateCmd.Flags().StringVar(&updateCmdOptions.value, "value", "", "Desired value of the attribute to update.")
	updateCmd.Flags().IntVar(&updateCmdOptions.index, "index", 0, "In case attribute is in array, use index to specify the array index. (Optional)")

	updateCmd.MarkFlagRequired("kind")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("att")
	updateCmd.MarkFlagRequired("value")
}

func update(txtContext, txtKind, txtName, txtAtt, txtVal string, intIndex int) error {
	return filepath.Walk(txtContext, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		pathMatched, err := regexp.MatchString(`\.yaml$`, path)
		if err != nil {
			return fmt.Errorf("Failed to compile regexp: %w", err)
		}
		if !pathMatched {
			return nil
		}
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("Failed to read file: %w", err)
		}
		resourceMatched, err := matchResource(txtKind, txtName, []byte(yamlFile))
		if err != nil {
			return fmt.Errorf("Failed to match resource: %w", err)
		}
		if !resourceMatched {
			return nil
		}
		switch txtKind {
		case "deployment":
			var deployment Deployment
			err = yaml.Unmarshal([]byte(yamlFile), &deployment)
			if err != nil {
				return fmt.Errorf("Failed to unmarshal: %w", err)
			}
			if txtAtt == "image" {
				fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n", deployment.Kind, deployment.Meta.Name, deployment.Spec.Template.Spec.Containers[intIndex].Image, txtVal)
				deployment.Spec.Template.Spec.Containers[intIndex].Image = txtVal
			}
			data, err := yaml.Marshal(&deployment)
			if err != nil {
				return fmt.Errorf("Failed to marshal: %w", err)
			}
			err = ioutil.WriteFile(path, data, 0644)
			if err != nil {
				return fmt.Errorf("Failed write file: %w", err)
			}
		case "rollout":

			var rollout Rollout
			err = yaml.Unmarshal([]byte(yamlFile), &rollout)
			if err != nil {
				return fmt.Errorf("Failed to unmarshal: %w", err)
			}
			if txtAtt == "image" {
				fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n", rollout.Kind, rollout.Meta.Name, rollout.Spec.Template.Spec.Containers[intIndex].Image, txtVal)
				rollout.Spec.Template.Spec.Containers[intIndex].Image = txtVal
			}
			data, err := yaml.Marshal(&rollout)
			if err != nil {
				return fmt.Errorf("Failed to marshal: %w", err)
			}
			err = ioutil.WriteFile(path, data, 0644)
			if err != nil {
				return fmt.Errorf("Failed write file: %w", err)
			}
		default:
			return fmt.Errorf("Kind %s is not supported yet", txtKind)
		}
		return nil
	})

}

func matchResource(txtKind, txtName string, data []byte) (bool, error) {
	var base BaseInfo
	err := yaml.Unmarshal(data, &base)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal: %w", err)
	}
	if strings.EqualFold(base.Kind, txtKind) && strings.EqualFold(base.Meta.Name, txtName) {
		return true, nil
	}
	return false, nil
}
