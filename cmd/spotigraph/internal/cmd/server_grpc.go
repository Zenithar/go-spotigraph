package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
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

		// Initialize config
		initConfig()

		// Starting banner
		log.For(ctx).Info("Starting spotigraph gRPC server ...")

		// Start goroutine group
		err := platform.Run(ctx, &platform.Application{
			Debug:           conf.Debug.Enable,
			Name:            "spotigraph-grpc",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group run.Group) {
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
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
