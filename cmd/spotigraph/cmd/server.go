package cmd

import (
	"github.com/spf13/cobra"
)

// -----------------------------------------------------------------------------

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the spotigraph server",
}
