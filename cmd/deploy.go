package cmd

import (
	"github.com/shaunthium/honeydew/deploy"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys your app to the production env",
	Long:  "Deploys your app to the production env",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deploy.Deploy()
	},
}
