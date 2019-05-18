package cmd

import (
	"github.com/spf13/cobra"

	"go.zenithar.org/spotigraph/pkg/grpc/v1/spotigraph/pb"
)

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
		pb.UserClientCommand,
		pb.SquadClientCommand,
		pb.GuildClientCommand,
		pb.ChapterClientCommand,
		pb.TribeClientCommand,
		pb.GraphClientCommand,
	)
}
