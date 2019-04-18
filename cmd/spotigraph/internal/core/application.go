package core

import (
	"context"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
	"github.com/google/gops/agent"
	"github.com/oklog/run"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform/jaeger"
	"go.zenithar.org/pkg/platform/prometheus"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
)

// RegisterDiagnostic adds diagnostic tools to main process
func RegisterDiagnostic(ctx context.Context, conf *config.Configuration, r *http.ServeMux) {
	// Start diagnostic handler
	if conf.Instrumentation.Diagnostic.Enabled {
		log.For(ctx).Info("Diagnostic endpoint enabled")

		if conf.Instrumentation.Diagnostic.RemoteURL != "" {
			log.For(ctx).Info("Starting gops agent", zap.String("url", conf.Instrumentation.Diagnostic.RemoteURL))
			if err := agent.Listen(agent.Options{Addr: conf.Instrumentation.Diagnostic.RemoteURL}); err != nil {
				log.For(ctx).Error("Error on starting gops agent", zap.Error(err))
			}
		} else {
			log.For(ctx).Info("Starting gops agent locally")
			if err := agent.Listen(agent.Options{}); err != nil {
				log.For(ctx).Error("Error on starting gops agent locally", zap.Error(err))
			}
		}

		r.HandleFunc("/diag/pprof", pprof.Index)
		r.HandleFunc("/diag/cmdline", pprof.Cmdline)
		r.HandleFunc("/diag/profile", pprof.Profile)
		r.HandleFunc("/diag/symbol", pprof.Symbol)
		r.HandleFunc("/diag/trace", pprof.Trace)
		r.Handle("/diag/goroutine", pprof.Handler("goroutine"))
		r.Handle("/diag/heap", pprof.Handler("heap"))
		r.Handle("/diag/threadcreate", pprof.Handler("threadcreate"))
		r.Handle("/diag/block", pprof.Handler("block"))
	}
}

// RegisterPrometheusExporter adds prometheus exporter
func RegisterPrometheusExporter(ctx context.Context, conf *config.Configuration, r *http.ServeMux) {
	// Start prometheus
	if conf.Instrumentation.Prometheus.Enabled {
		log.For(ctx).Info("Prometheus exporter enabled")

		exporter, err := prometheus.NewExporter(conf.Instrumentation.Prometheus.Config)
		if err != nil {
			log.For(ctx).Fatal("Unable to register prometheus exporter", zap.Error(err))
		}

		// Add exporter
		view.RegisterExporter(exporter)

		// Add metrics handler
		r.Handle("/metrics", exporter)
	}
}

// RegisterJaegerExporter add jaeger as trace exporter
func RegisterJaegerExporter(ctx context.Context, conf *config.Configuration) {
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
}

// -----------------------------------------------------------------------------

// Run the dispatcher
func Run(ctx context.Context, conf *config.Configuration, r *http.ServeMux, builder func(upg *tableflip.Upgrader, group run.Group)) error {
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
			Handler: r,
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

	// Initialize the component
	builder(upg, group)

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

	// Run goroutine group
	return group.Run()
}
