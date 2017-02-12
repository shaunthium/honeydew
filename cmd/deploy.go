package cmd

import (
	"github.com/shaunthium/honeydew/deploy"

	"github.com/spf13/cobra"
)

var volumeMountHostname string
var volumeName string
var targetDirectory string

func init() {
	RootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVar(&volumeMountHostname, "volume_mount_hostname", "", "Hostname of the volume mount")
	deployCmd.Flags().StringVar(&volumeName, "volume_name", "", "Name of the volume")
	deployCmd.Flags().StringVar(&targetDirectory, "target_folder", "", "Name of the target folder to mount the volume into")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys your app to the production env",
	Long:  "Deploys your app to the production env",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deploy.Deploy(volumeMountHostname, volumeName, targetDirectory)
	},
}
