package cmd

import (
	"context"
	"net/http"

	"github.com/cloudflare/tableflip"
	"github.com/dchest/uniuri"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/core"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/grpc"
	"go.zenithar.org/spotigraph/internal/version"
)

// -----------------------------------------------------------------------------

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Starts the spotigraph gRPC server",
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
			AppName:   "spotigraph-grpc",
			AppID:     appID,
			Version:   version.Version,
			Revision:  version.Revision,
			SentryDSN: conf.Instrumentation.Logs.SentryDSN,
		})

		// Starting banner
		log.For(ctx).Info("Starting spotigraph gRPC server ...")

		// Preparing instrumentation
		instrumentationRouter := http.NewServeMux()

		// Register common features
		core.RegisterDiagnostic(ctx, conf, instrumentationRouter)
		core.RegisterPrometheusExporter(ctx, conf, instrumentationRouter)
		core.RegisterJaegerExporter(ctx, conf)

		// Start goroutine group
		err := core.Run(ctx, conf, instrumentationRouter, func(upg *tableflip.Upgrader, group run.Group) {
			ln, err := upg.Fds.Listen(conf.Server.GRPC.Network, conf.Server.GRPC.Listen)
			if err != nil {
				log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
			}

			server, err := grpc.New(ctx, conf)
			if err != nil {
				log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
			}

			group.Add(
				func() error {
					log.For(ctx).Info("Starting GRPC server", zap.Stringer("address", ln.Addr()))
					return server.Serve(ln)
				},
				func(e error) {
					log.For(ctx).Info("Shutting GRPC server down")
					server.GracefulStop()
				},
			)
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
