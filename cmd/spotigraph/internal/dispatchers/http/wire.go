//+build wireinject

package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/wire"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/core"
	v1 "go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/http/handlers/v1"
	"go.zenithar.org/spotigraph/internal/services"
)

func httpServer(ctx context.Context, cfg *config.Configuration, users services.User, squads services.Squad, chapters services.Chapter, guilds services.Guild, tribes services.Tribe) (*http.Server, error) {
	r := chi.NewRouter()

	// middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// timeout before request cancelation
	r.Use(middleware.Timeout(60 * time.Second))

	// API endpoint
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("users", v1.UserRoutes(users))
		r.Mount("squads", v1.SquadRoutes(squads))
		r.Mount("chapters", v1.ChapterRoutes(chapters))
		r.Mount("guilds", v1.GuildRoutes(guilds))
		r.Mount("tribes", v1.TribeRoutes(tribes))
	})

	// Health checking status endpoint
	r.Handle("/healthz", http.HandlerFunc(healthz()))

	// Assign router to server
	server := &http.Server{
		Handler: r,
	}

	// Enable TLS if requested
	if cfg.Server.HTTP.UseTLS {
		// Client authentication enabled but not required
		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.HTTP.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		// Generate TLS configuration
		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.HTTP.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.HTTP.TLS.CertificatePath,
			CAFile:     cfg.Server.HTTP.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", zap.Error(err))
			return nil, err
		}

		// Create the TLS credentials
		server.TLSConfig = tlsConfig
	} else {
		log.For(ctx).Info("No transport authentication enabled")
	}

	// Return result
	return server, nil
}

// -----------------------------------------------------------------------------

func setupLocalMongoDB(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	wire.Build(
		core.LocalMongoDBSet,
		httpServer,
	)
	return &http.Server{}, nil
}

func setupLocalRethinkDB(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	wire.Build(
		core.LocalRethinkDBSet,
		httpServer,
	)
	return &http.Server{}, nil
}

func setupLocalPostgreSQL(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	wire.Build(
		core.LocalPostgreSQLSet,
		httpServer,
	)
	return &http.Server{}, nil
}
