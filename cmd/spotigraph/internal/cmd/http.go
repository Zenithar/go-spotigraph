package cmd

import (
	"context"

	"github.com/dchest/uniuri"
	"github.com/google/gops/agent"
	"github.com/spf13/cobra"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform/jaeger"
	"go.zenithar.org/pkg/platform/prometheus"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/http"
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

		// Start prometheus
		if conf.Instrumentation.Prometheus.Enabled {
			log.For(ctx).Info("prometheus exporter enabled")

			exporter, err := prometheus.NewExporter(conf.Instrumentation.Prometheus.Config)
			log.For(ctx).Fatal("Unable to register prometheus exporter", zap.Error(err))

			view.RegisterExporter(exporter)
		}

		// Start tracing
		if conf.Instrumentation.Jaeger.Enabled {
			log.For(ctx).Info("jaeger exporter enabled")

			exporter, err := jaeger.NewExporter(conf.Instrumentation.Jaeger.Config)
			log.For(ctx).Fatal("Unable to register jaeger exporter", zap.Error(err))

			trace.RegisterExporter(exporter)

			// Trace everything when debugging is enabled
			if conf.Debug.Enable {
				trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
			}
		}

		// Starting banner
		log.For(ctx).Info("Starting spotigraph HTTP server ...")

		// Start server
		http.WaitForShutdown(ctx, conf)
	},
}
