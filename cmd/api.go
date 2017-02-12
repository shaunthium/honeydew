package cmd

import (
	"github.com/shaunthium/honeydew/api"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts honeydew in API mode",
	Long:  "Starts honeydew in API mode",
	Run: func(cmd *cobra.Command, args []string) {
		api.RunServer()
	},
}
