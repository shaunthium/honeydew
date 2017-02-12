package cmd

import (
	"github.com/shaunthium/honeydew/setup"
	"github.com/spf13/cobra"
)

var newVolumeName string

func init() {
	RootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVar(&newVolumeName, "volume_name", "new-volume", "New volume name")
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a new production env",
	Long:  "Sets up a new production env",
	RunE: func(cmd *cobra.Command, args []string) error {
		return setup.Setup(newVolumeName)
	},
}
