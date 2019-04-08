package cmd

import (
	"context"

	"github.com/dchest/uniuri"
	"github.com/google/gops/agent"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"

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

		// Generate an instance identifier
		appID := uniuri.NewLen(64)

		// Initialize config
		initConfig()

		// Prepare logger
		log.Setup(ctx, &log.Options{
			Debug:     conf.Debug.Enable,
			AppName:   "spotigraph-gql",
			AppID:     appID,
			Version:   version.Version,
			Revision:  version.Revision,
			SentryDSN: conf.Observability.Logs.SentryDSN,
		})

		// gops debug
		if conf.Debug.Enable {
			if conf.Debug.RemoteURL != "" {
				log.For(ctx).Info("Starting gops agent", zap.String("url", conf.Debug.RemoteURL))
				if err := agent.Listen(agent.Options{Addr: conf.Debug.RemoteURL}); err != nil {
					log.For(ctx).Error("Error on starting gops agent", zap.Error(err))
				}
			} else {
				log.For(ctx).Info("Starting gops agent locally")
				if err := agent.Listen(agent.Options{}); err != nil {
					log.For(ctx).Error("Error on starting gops agent locally", zap.Error(err))
				}
			}
		}

		// Starting banner
		log.For(ctx).Info("Starting spotigraph GraphQL server ...")

		// Start server
		graphql.WaitForShutdown(ctx, conf)
	},
}
