package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd describes root command of the tool
var RootCmd = &cobra.Command{
	Use:   "spotigraph",
	Short: "Spotify Agile methodology graph microservice",
}
