package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql"
	"go.zenithar.org/spotigraph/internal/version"
)

// -----------------------------------------------------------------------------

var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Starts the spotigraph GraphQL server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize config
		initConfig()

		// Starting banner
		log.For(ctx).Info("Starting spotigraph GraphQL server ...")

		// Start goroutine group
		err := platform.Run(ctx, &platform.Application{
			Debug:           conf.Debug.Enable,
			Name:            "spotigraph-graphql",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group run.Group) {
				ln, err := upg.Fds.Listen(conf.Server.GraphQL.Network, conf.Server.GraphQL.Listen)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GraphQL server", zap.Error(err))
				}

				server, err := graphql.New(ctx, conf)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GraphQL server", zap.Error(err))
				}

				group.Add(
					func() error {
						log.For(ctx).Info("Starting GraphQL server", zap.Stringer("address", ln.Addr()))
						return server.Serve(ln)
					},
					func(e error) {
						log.For(ctx).Info("Shutting GraphQL server down")
						log.SafeClose(server, "Unable to close HTTP server")
					},
				)
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
