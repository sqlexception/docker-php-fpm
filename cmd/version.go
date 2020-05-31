package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of php-fpm-healthcheck",
	Long:  `Print the version number of php-fpm-healthcheck`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("php-fpm-healthcheck %v\n", Version)

	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}