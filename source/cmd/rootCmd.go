package cmd

import (
	"fmt"
	"npm-registry/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "serve",
	Short:   "serve npm registry",
	Long:    "This command is used to serve npm registry",
	Version: "1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Serve()
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Unable to serve registry: ", err)
		os.Exit(1)
	}
}
