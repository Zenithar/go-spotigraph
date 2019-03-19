package cmd

import (
	"context"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/server"

	"github.com/dchest/uniuri"
	"github.com/google/gops/agent"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"

	"go.zenithar.org/spotigraph/internal/version"
)

// -----------------------------------------------------------------------------

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the spotigraph server",
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
			AppName:   "spotigraph",
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
		log.For(ctx).Info("Starting spotigraph server ...")

		// Start server
		server.WaitForShutdown(ctx, conf)
	},
}
