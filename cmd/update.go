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
		name      string
		attribute string
		value     string
		index     int
	}
)

var updateCmd = &cobra.Command{
	Use: "update KIND [flags] PATH",
	Short: "Update k8s resources yaml files",
	Long: "Update k8s resources yaml files\n\nExample:\nkyml update deployment -n my_deployment -a image -v myimage:0.1 .\n",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := update(args[0],args[1])
		dieOnError("Failed to update", err)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateCmdOptions.name, "name", "n", "","Name of the resource to update. (Required)")
	updateCmd.Flags().StringVarP(&updateCmdOptions.attribute, "att", "a", "","Name of the attribute to update. (Required)")
	updateCmd.Flags().StringVarP(&updateCmdOptions.value, "value", "v", "","Desired value of the attribute to update. (Required)")
	updateCmd.Flags().IntVarP(&updateCmdOptions.index, "index", "i",0, "In case attribute is in array, use index to specify the array index. (Optional)")


	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("att")
	updateCmd.MarkFlagRequired("value")
}

func update(kind, directory string) error {

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

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
		resourceMatched, err := matchResource(kind, updateCmdOptions.name, []byte(yamlFile))
		if err != nil {
			return fmt.Errorf("Failed to match resource: %w", err)
		}
		if !resourceMatched {
			return nil
		}



		switch kind {
		case "deployment":
			var deployment Deployment
			err = yaml.Unmarshal([]byte(yamlFile), &deployment)
			if err != nil {
				return fmt.Errorf("Failed to unmarshal: %w", err)
			}
			if updateCmdOptions.attribute == "image" {
				fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n", deployment.Kind, deployment.Meta.Name, deployment.Spec.Template.Spec.Containers[updateCmdOptions.index].Image, updateCmdOptions.value)
				deployment.Spec.Template.Spec.Containers[updateCmdOptions.index].Image = updateCmdOptions.value
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
			if updateCmdOptions.attribute == "image" {
				fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n", rollout.Kind, rollout.Meta.Name, rollout.Spec.Template.Spec.Containers[updateCmdOptions.index].Image, updateCmdOptions.value)
				rollout.Spec.Template.Spec.Containers[updateCmdOptions.index].Image = updateCmdOptions.value
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
			return fmt.Errorf("Kind %s is not supported yet", kind)
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
