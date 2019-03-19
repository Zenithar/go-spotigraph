package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.zenithar.org/spotigraph/internal/version"
)

// -----------------------------------------------------------------------------

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display service version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s-%s (%s)", version.Version, version.Revision, version.GoVersion)
	},
}
