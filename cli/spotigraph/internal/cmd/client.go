package cmd

import (
	"github.com/spf13/cobra"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
		chapterv1.ChapterAPIClientCommand,
		personv1.PersonAPIClientCommand,
		squadv1.SquadAPIClientCommand,
	)
}
