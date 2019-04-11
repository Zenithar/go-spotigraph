package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
)

type application struct {
	cfg    *config.Configuration
	server *http.Server
}

var healthy int32

// -----------------------------------------------------------------------------

// WaitForShutdown starts the server and wait for shutdown signal
func WaitForShutdown(ctx context.Context, cfg *config.Configuration) {
	// Initialize application
	app := &application{}

	// Apply configuration
	if err := app.ApplyConfiguration(cfg); err != nil {
		log.For(ctx).Fatal("Unable to initialize server settings", zap.Error(err))
	}

	// Fork the cancellation handler
	go func() {
		<-ctx.Done()
		err := app.Shutdown(context.Background())
		if err != nil {
			log.For(ctx).Error("Unable to shutdown spotigraph http service", zap.Error(err))
		}
	}()

	// Start the server
	if err := app.Serve(ctx); err != nil {
		log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
	}
}

// -----------------------------------------------------------------------------

// ApplyConfiguration apply the configuration after checking it
func (s *application) ApplyConfiguration(cfg interface{}) error {
	// Check configuration validity
	if err := s.checkConfiguration(cfg); err != nil {
		return err
	}

	// Apply to current component (type assertion done if check)
	s.cfg, _ = cfg.(*config.Configuration)

	// No error
	return nil
}

// Serve starts the component
func (s *application) Serve(ctx context.Context) error {
	log.For(ctx).Info("Starting spotigraph HTTP service ...")

	// Setup the gRPC server
	if err := s.setup(ctx); err != nil {
		log.For(ctx).Error("Unable to initialize spotigraph HTTP service", zap.Error(err))
		return err
	}

	// Initialize a listener
	lis, err := net.Listen(s.cfg.Server.HTTP.Network, s.cfg.Server.HTTP.Listen)
	if err != nil {
		return err
	}

	// Return no error
	log.For(ctx).Info("Spotigraph HTTP service listening ...", zap.String("listen", s.cfg.Server.HTTP.Listen))
	atomic.StoreInt32(&healthy, 1)
	return s.server.Serve(lis)
}

// Shutdown the component
func (s *application) Shutdown(ctx context.Context) error {
	log.For(ctx).Info("Try to gracefully shutdown spotigraph HTTP server...")
	atomic.StoreInt32(&healthy, 0)

	if s.server != nil {
		s.server.SetKeepAlivesEnabled(false)
		return s.server.Shutdown(ctx)
	}
	return nil
}

// -----------------------------------------------------------------------------

func (s *application) checkConfiguration(cfg interface{}) error {
	// Check via type assertion
	config, ok := cfg.(*config.Configuration)
	if !ok {
		return fmt.Errorf("server: invalid configuration")
	}

	switch config.Core.Mode {
	case "local":
		switch config.Core.Local.Type {
		case "rethinkdb":
			if config.Core.Local.Hosts == "" {
				return fmt.Errorf("server: database hosts list is mandatory")
			}
		case "mongodb":
			if config.Core.Local.Hosts == "" {
				return fmt.Errorf("server: database hosts list is mandatory")
			}
		case "postgresql":
			if config.Core.Local.Hosts == "" {
				return fmt.Errorf("server: database hosts list is mandatory")
			}
		default:
			return fmt.Errorf("server: invalid type (mongodb/rethinkdb/postgresql)")
		}
	default:
		return fmt.Errorf("server: invalid core mode, grpc only support 'local'")
	}

	// No error
	return nil
}

func (s *application) setup(ctx context.Context) error {
	var err error

	switch s.cfg.Core.Mode {
	case "local":
		switch s.cfg.Core.Local.Type {
		case "mongodb":
			s.server, err = setupLocalMongoDB(ctx, s.cfg)
		case "rethinkdb":
			s.server, err = setupLocalRethinkDB(ctx, s.cfg)
		case "postgresql":
			s.server, err = setupLocalPostgreSQL(ctx, s.cfg)
		}
	case "remote":
	default:
		log.For(ctx).Fatal("Invalid core mode, use 'remote' or 'local'.")
	}

	return err
}

// -----------------------------------------------------------------------------

func healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
