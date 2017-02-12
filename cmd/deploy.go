package cmd

import (
	"github.com/shaunthium/honeydew/deploy"

	"github.com/spf13/cobra"
)

var volumeMountHostname string
var volumeName string
var targetDirectory string
var imageTagName string
var instanceHostname string

func init() {
	RootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVar(&volumeMountHostname, "volume_mount_hostname", "", "Hostname of the volume mount")
	deployCmd.Flags().StringVar(&volumeName, "volume_name", "", "Name of the volume")
	deployCmd.Flags().StringVar(&targetDirectory, "target_directory", "", "Name of the target directory to mount the volume into")
	deployCmd.Flags().StringVar(&imageTagName, "image_tag_name", "", "Tag name of the built container")
	deployCmd.Flags().StringVar(&instanceHostname, "instance_hostname", "", "Hostname of instance to deploy to")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys your app to your selected environment",
	Long:  "Deploys your app to your selected environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deploy.Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname)
	},
}
