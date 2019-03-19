package cmd

import "github.com/spf13/cobra"

// -----------------------------------------------------------------------------

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Service Configuration",
}

// -----------------------------------------------------------------------------

func init() {
	configNewCmd.Flags().BoolVar(&configNewAsEnvFlag, "env", false, "Print configuration as environment variable")
	configCmd.AddCommand(configNewCmd)
}
