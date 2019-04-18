package cmd

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudflare/tableflip"
	"github.com/dchest/uniuri"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/core"
	httpDispatcher "go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/http"
	"go.zenithar.org/spotigraph/internal/version"
)

// -----------------------------------------------------------------------------

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts the spotigraph HTTP server",
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
			AppName:   "spotigraph-http",
			AppID:     appID,
			Version:   version.Version,
			Revision:  version.Revision,
			SentryDSN: conf.Instrumentation.Logs.SentryDSN,
		})

		// Starting banner
		log.For(ctx).Info("Starting spotigraph HTTP server ...")

		// Preparing instrumentation
		instrumentationRouter := http.NewServeMux()

		// Register common features
		core.RegisterDiagnostic(ctx, conf, instrumentationRouter)
		core.RegisterPrometheusExporter(ctx, conf, instrumentationRouter)
		core.RegisterJaegerExporter(ctx, conf)

		// Start goroutine group
		err := core.Run(ctx, conf, instrumentationRouter, func(upg *tableflip.Upgrader, group run.Group) {
			ln, err := upg.Fds.Listen(conf.Server.HTTP.Network, conf.Server.HTTP.Listen)
			if err != nil {
				log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
			}

			server, err := httpDispatcher.New(ctx, conf)
			if err != nil {
				log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
			}

			group.Add(
				func() error {
					log.For(ctx).Info("Starting HTTP server", zap.Stringer("address", ln.Addr()))
					return server.Serve(ln)
				},
				func(e error) {
					log.For(ctx).Info("Shutting HTTP server down")

					ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
					defer cancel()

					log.CheckErrCtx(ctx, "Error raised while shutting down the server", server.Shutdown(ctx))
					log.SafeClose(server, "Unable to close HTTP server")
				},
			)
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
