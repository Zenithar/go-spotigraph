// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package http

import (
	"context"
	"crypto/tls"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/pkg/db/adapter/rethinkdb"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/cli/spotigraph/internal/core"
	"go.zenithar.org/spotigraph/cli/spotigraph/internal/dispatchers/http/handlers/v1"
	mongodb2 "go.zenithar.org/spotigraph/internal/repositories/pkg/mongodb"
	postgresql2 "go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	rethinkdb2 "go.zenithar.org/spotigraph/internal/repositories/pkg/rethinkdb"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
	"net/http"
	"time"
)

// Injectors from wire.go:

func setupLocalMongoDB(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	configuration := core.MongoDBConfig(cfg)
	wrappedClient, err := mongodb.Connection(ctx, configuration)
	if err != nil {
		return nil, err
	}
	repositoriesUser := mongodb2.NewUserRepository(configuration, wrappedClient)
	servicesUser := user.New(repositoriesUser)
	repositoriesSquad := mongodb2.NewSquadRepository(configuration, wrappedClient)
	servicesSquad := squad.New(repositoriesSquad, repositoriesUser)
	repositoriesChapter := mongodb2.NewChapterRepository(configuration, wrappedClient)
	servicesChapter := chapter.New(repositoriesChapter)
	repositoriesGuild := mongodb2.NewGuildRepository(configuration, wrappedClient)
	servicesGuild := guild.New(repositoriesGuild)
	repositoriesTribe := mongodb2.NewTribeRepository(configuration, wrappedClient)
	servicesTribe := tribe.New(repositoriesTribe)
	server, err := httpServer(ctx, cfg, servicesUser, servicesSquad, servicesChapter, servicesGuild, servicesTribe)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func setupLocalRethinkDB(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	configuration := core.RethinkDBConfig(cfg)
	session, err := rethinkdb.Connection(ctx, configuration)
	if err != nil {
		return nil, err
	}
	repositoriesUser := rethinkdb2.NewUserRepository(configuration, session)
	servicesUser := user.New(repositoriesUser)
	repositoriesSquad := rethinkdb2.NewSquadRepository(configuration, session)
	servicesSquad := squad.New(repositoriesSquad, repositoriesUser)
	repositoriesChapter := rethinkdb2.NewChapterRepository(configuration, session)
	servicesChapter := chapter.New(repositoriesChapter)
	repositoriesGuild := rethinkdb2.NewGuildRepository(configuration, session)
	servicesGuild := guild.New(repositoriesGuild)
	repositoriesTribe := rethinkdb2.NewTribeRepository(configuration, session)
	servicesTribe := tribe.New(repositoriesTribe)
	server, err := httpServer(ctx, cfg, servicesUser, servicesSquad, servicesChapter, servicesGuild, servicesTribe)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func setupLocalPostgreSQL(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	configuration := core.PosgreSQLConfig(cfg)
	db, err := postgresql.Connection(ctx, configuration)
	if err != nil {
		return nil, err
	}
	repositoriesUser := postgresql2.NewUserRepository(configuration, db)
	servicesUser := user.New(repositoriesUser)
	repositoriesSquad := postgresql2.NewSquadRepository(configuration, db)
	servicesSquad := squad.New(repositoriesSquad, repositoriesUser)
	repositoriesChapter := postgresql2.NewChapterRepository(configuration, db)
	servicesChapter := chapter.New(repositoriesChapter)
	repositoriesGuild := postgresql2.NewGuildRepository(configuration, db)
	servicesGuild := guild.New(repositoriesGuild)
	repositoriesTribe := postgresql2.NewTribeRepository(configuration, db)
	servicesTribe := tribe.New(repositoriesTribe)
	server, err := httpServer(ctx, cfg, servicesUser, servicesSquad, servicesChapter, servicesGuild, servicesTribe)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

func httpServer(ctx context.Context, cfg *config.Configuration, users services.User, squads services.Squad, chapters services.Chapter, guilds services.Guild, tribes services.Tribe) (*http.Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", ochttp.WithRouteTag(v1.UserRoutes(users), "/api/v1/users"))
	})

	server := &http.Server{
		Handler: &ochttp.Handler{
			Handler:     r,
			Propagation: &b3.HTTPFormat{},
		},
	}

	if cfg.Server.HTTP.UseTLS {

		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.HTTP.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

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

		server.TLSConfig = tlsConfig
	} else {
		log.For(ctx).Info("No transport encryption enabled for HTTP server")
	}

	err := view.Register(ochttp.ServerRequestCountView, ochttp.ServerRequestBytesView, ochttp.ServerResponseBytesView, ochttp.ServerLatencyView, ochttp.ServerRequestCountByMethod, ochttp.ServerResponseCountByStatusCode)
	if err != nil {
		log.For(ctx).Fatal("Unable to register stat views", zap.Error(err))
	}

	return server, nil
}
