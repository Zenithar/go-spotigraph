package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opencensus.io/plugin/ochttp"

	"github.com/cloudflare/tableflip"
	"github.com/dchest/uniuri"
	"github.com/google/gops/agent"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform/jaeger"
	"go.zenithar.org/pkg/platform/prometheus"
	httpServer "go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/http"
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

		// Preparing instrumentation
		instrumentationRouter := http.NewServeMux()

		// Start prometheus
		if conf.Instrumentation.Prometheus.Enabled {
			log.For(ctx).Info("Prometheus exporter enabled")

			exporter, err := prometheus.NewExporter(conf.Instrumentation.Prometheus.Config)
			if err != nil {
				log.For(ctx).Fatal("Unable to register prometheus exporter", zap.Error(err))
			}

			view.RegisterExporter(exporter)
			instrumentationRouter.Handle("/metrics", exporter)
		}

		// Start tracing
		if conf.Instrumentation.Jaeger.Enabled {
			log.For(ctx).Info("Jaeger exporter enabled")

			exporter, err := jaeger.NewExporter(conf.Instrumentation.Jaeger.Config)
			if err != nil {
				log.For(ctx).Fatal("Unable to register jaeger exporter", zap.Error(err))
			}

			trace.RegisterExporter(exporter)

			// Trace everything when debugging is enabled
			if conf.Debug.Enable {
				trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
			}
		}

		// Configure graceful restart
		upg, err := tableflip.New(tableflip.Options{})
		if err != nil {
			log.For(ctx).Fatal("Unable to register graceful restart handler", zap.Error(err))
		}

		// Do an upgrade on SIGHUP
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGHUP)
			for range ch {
				log.For(ctx).Info("Graceful reloading")

				_ = upg.Upgrade()
			}
		}()

		var group run.Group

		// Instrumentation server
		{
			ln, err := upg.Fds.Listen(conf.Instrumentation.Network, conf.Instrumentation.Listen)
			if err != nil {
				log.For(ctx).Fatal("Unable to start instrumentation server", zap.Error(err))
			}

			server := &http.Server{
				Handler: instrumentationRouter,
			}

			group.Add(
				func() error {
					log.For(ctx).Info("Starting instrumentation server", zap.Stringer("address", ln.Addr()))
					return server.Serve(ln)
				},
				func(e error) {
					log.For(ctx).Info("Shutting instrumentation server down")

					ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
					defer cancel()

					log.CheckErrCtx(ctx, "Error raised while shutting down the server", server.Shutdown(ctx))
					log.SafeClose(server, "Unable to close instrumentation server")
				},
			)
		}

		// Register stat views
		err = view.Register(
			// HTTP
			ochttp.ServerRequestCountView,
			ochttp.ServerRequestBytesView,
			ochttp.ServerResponseBytesView,
			ochttp.ServerLatencyView,
			ochttp.ServerRequestCountByMethod,
			ochttp.ServerResponseCountByStatusCode,
		)
		if err != nil {
			log.For(ctx).Fatal("Unable to register stat views", zap.Error(err))
		}

		// HTTP server
		{
			ln, err := upg.Fds.Listen(conf.Server.HTTP.Network, conf.Server.HTTP.Listen)
			if err != nil {
				log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
			}

			server, err := httpServer.New(ctx, conf)
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
					log.SafeClose(server, "Unable to close instrumentation server")
				},
			)
		}

		// Setup signal handler
		{
			var (
				cancelInterrupt = make(chan struct{})
				ch              = make(chan os.Signal, 2)
			)
			defer close(ch)

			group.Add(
				func() error {
					signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

					select {
					case sig := <-ch:
						log.For(ctx).Info("Captured signal", zap.Any("signal", sig))
					case <-cancelInterrupt:
					}

					return nil
				},
				func(e error) {
					close(cancelInterrupt)
					signal.Stop(ch)
				},
			)
		}

		// Final handler
		{
			group.Add(
				func() error {
					// Tell the parent we are ready
					_ = upg.Ready()

					// Wait for children to be ready
					// (or application shutdown)
					<-upg.Exit()

					return nil
				},
				func(e error) {
					upg.Stop()
				},
			)
		}

		// Start goroutine group
		log.CheckErrCtx(ctx, "Unable to run application", group.Run())
	},
}
