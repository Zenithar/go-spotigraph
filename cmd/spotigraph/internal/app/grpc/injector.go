//+build wireinject

package grpc

import (
	"context"
	"crypto/tls"
	"strings"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/mongodb"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/rethinkdb"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/graph"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
	"go.zenithar.org/spotigraph/pkg/grpc/v1/spotigraph/pb"

	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	mdb "go.zenithar.org/pkg/db/adapter/mongodb"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"
	rdb "go.zenithar.org/pkg/db/adapter/rethinkdb"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// -----------------------------------------------------------------------------

var serviceSet = wire.NewSet(
	user.New,
	chapter.New,
	squad.New,
	guild.New,
	tribe.New,
	graph.New,
)

// rdbConfig declares a Database configuration provider for Wire
func rdbConfig(cfg *config.Configuration) *rdb.Configuration {
	return &rdb.Configuration{
		AutoMigrate: cfg.Server.Database.AutoMigrate,
		Addresses:   strings.Split(cfg.Server.Database.Hosts, ","),
		Database:    cfg.Server.Database.Database,
		AuthKey:     cfg.Server.Database.Password,
	}
}

// mgoConfig declares a Database configuration provider for Wire
func mgoConfig(cfg *config.Configuration) *mdb.Configuration {
	return &mdb.Configuration{
		AutoMigrate:      cfg.Server.Database.AutoMigrate,
		ConnectionString: cfg.Server.Database.Hosts,
		DatabaseName:     cfg.Server.Database.Database,
		Username:         cfg.Server.Database.Username,
		Password:         cfg.Server.Database.Password,
	}
}

// pgConfig declares a Database configuration provider for Wire
func pgConfig(cfg *config.Configuration) *pgdb.Configuration {
	return &pgdb.Configuration{
		AutoMigrate:      cfg.Server.Database.AutoMigrate,
		ConnectionString: cfg.Server.Database.Hosts,
		Username:         cfg.Server.Database.Username,
		Password:         cfg.Server.Database.Password,
	}
}

func grpcServer(ctx context.Context, cfg *config.Configuration, users services.User, chapters services.Chapter, guilds services.Guild, squads services.Squad, tribes services.Tribe, graph services.Graph) (*grpc.Server, error) {
	// gRPC middlewares
	sopts := []grpc.ServerOption{}

	// Replace gRPC logger
	grpc_zap.ReplaceGrpcLogger(zap.L())

	// gRPC middlewares
	sopts = append(sopts, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(zap.L()),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(zap.L()),
			),
		))

	// Enable TLS if requested
	if cfg.Server.GRPC.UseTLS {
		// Client authentication enabled but not required
		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.GRPC.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		// Generate TLS configuration
		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.GRPC.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.GRPC.TLS.CertificatePath,
			CAFile:     cfg.Server.GRPC.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", zap.Error(err))
			return nil, err
		}

		// Create the TLS credentials
		sopts = append(sopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	} else {
		log.For(ctx).Info("No transport authentication enabled")
	}

	// Initialize the server
	server := grpc.NewServer(sopts...)

	// Health
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	// Register services
	pb.RegisterUserServer(server, users)
	pb.RegisterChapterServer(server, chapters)
	pb.RegisterGuildServer(server, guilds)
	pb.RegisterSquadServer(server, squads)
	pb.RegisterTribeServer(server, tribes)
	pb.RegisterGraphServer(server, graph)

	// Reflection
	reflection.Register(server)

	// Return no error
	return server, nil
}

// -----------------------------------------------------------------------------

func setupRethinkDB(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {

	wire.Build(
		rdbConfig,
		rethinkdb.RepositorySet,
		serviceSet,
		grpcServer,
	)

	return nil, nil
}

func setupMongoDB(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {

	wire.Build(
		mgoConfig,
		mongodb.RepositorySet,
		serviceSet,
		grpcServer,
	)

	return nil, nil
}

func setupPostgresDB(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {

	wire.Build(
		pgConfig,
		postgresql.RepositorySet,
		serviceSet,
		grpcServer,
	)

	return nil, nil
}
