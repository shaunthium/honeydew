package cmd

import (
	"github.com/shaunthium/honeydew/undeploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(undeployCmd)

	undeployCmd.Flags().StringVar(&targetDirectory, "target_directory", "", "Name of the target directory to unmount the volume from")
	undeployCmd.Flags().StringVar(&instanceHostname, "instance_hostname", "", "Hostname of instance to undeploy from")
}

var undeployCmd = &cobra.Command{
	Use:   "undeploy",
	Short: "Undeploys your app from your selected environment",
	Long:  "Undeploys your app from your selected environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return undeploy.Undeploy(targetDirectory, instanceHostname)
	},
}
