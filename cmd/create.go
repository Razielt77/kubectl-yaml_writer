package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	createCmdOptions struct {
		name         string
		image        string
		targetPort   int
		externalPort int
	}
)

var createCmd = &cobra.Command{
	Use:   "create KIND [flags] PATH",
	Short: "Create k8s resources yaml files",
	Long:  "Create k8s resources yaml files.\nCurrently supported resources are: services & deployments\n\nExample:\nkubectl yaml-writer create app -name APP_NAME -image IMAGE_NAME:0.1 -targetport CONTAINER_PORT_NUMBER -externalPort EXTERNAL_PORT.\n",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := create(args[0], args[1])
		dieOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&createCmdOptions.name, "name", "n", "", "Name of the resource to update. (Required)")
	createCmd.Flags().StringVarP(&createCmdOptions.image, "image", "i", "", "Name of the attribute to update. (Required)")
	createCmd.Flags().IntVarP(&createCmdOptions.targetPort, "targetport", "t", 80, "Desired value of the attribute to update. (Required)")
	createCmd.Flags().IntVarP(&createCmdOptions.externalPort, "externalport", "e", 80, "In case attribute is in array, use index to specify the array index. (Optional)")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("image")
}

func constructDeploymentFilename(directory, app string) string {
	return directory + "/" + app + "_deployment.yaml"
}
func constructServiceFilename(directory, app string) string {
	return directory + "/" + app + "_service.yaml"
}

func create(kind, directory string) error {

	var err error = nil

	var s service
	var d deployment

	switch kind {
	case "app":
		targetPort := strconv.Itoa(createCmdOptions.targetPort)
		externalPort := strconv.Itoa(createCmdOptions.externalPort)
		s.Init(createCmdOptions.name, targetPort, externalPort)
		d.Init(createCmdOptions.name+"_deployment", createCmdOptions.name, createCmdOptions.image, 1, createCmdOptions.targetPort)
		filename := constructServiceFilename(directory, createCmdOptions.name)
		err = marshalAndSave(s, filename)
		if err != nil {
			return fmt.Errorf("failed to save file %s: %w", filename, err)
		}

		filename = constructDeploymentFilename(directory, createCmdOptions.name)
		err = marshalAndSave(d, filename)
		if err != nil {
			return fmt.Errorf("failed to save file %s: %w", filename, err)
		}
	default:
		return fmt.Errorf("Kind %s is not supported yet", kind)
	}

	return err

}
