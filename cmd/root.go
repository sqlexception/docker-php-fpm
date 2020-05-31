package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Version that is being reported by the CLI
var Version string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "php-fpm-healthcheck",
	Short: "Exports php-fpm metrics for prometheus",
	Long:  `php-fpm_exporter exports prometheus compatible metrics from php-fpm.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}