package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "honeydew",
	Short: "Deploy your app using NetApp, Docker and more.",
	Long:  `Deploy your app using NetApp, Docker and more.`,
}
