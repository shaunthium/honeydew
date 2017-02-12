package cmd

import (
	"errors"

	"github.com/shaunthium/honeydew/setup"
	"github.com/spf13/cobra"
)

var newVolumeName string

func init() {
	RootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVar(&newVolumeName, "volume_name", "", "New volume name")
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a new production env",
	Long:  "Sets up a new production env",
	RunE: func(cmd *cobra.Command, args []string) error {
		if newVolumeName == "" {
			return errors.New("New volume name must be specified")
		}
		return setup.Setup(newVolumeName)
	},
}
