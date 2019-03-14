package server

import (
	"context"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/grpc/v1/spotigraph/pb"
)

// Setup returns a gRPC server
func Setup(ctx context.Context, users services.User, squads services.Squad, guilds services.Guild, chapters services.Chapter, tribes services.Tribe) *grpc.Server {
	s := grpc.NewServer(
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	// Health services
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthSrv)

	// Register all services
	pb.RegisterUserServiceServer(s, users)
	healthSrv.SetServingStatus("go.zenithar.org/spotigraph/srv/users", grpc_health_v1.HealthCheckResponse_SERVING)

	pb.RegisterSquadServiceServer(s, squads)
	healthSrv.SetServingStatus("go.zenithar.org/spotigraph/srv/squads", grpc_health_v1.HealthCheckResponse_SERVING)

	pb.RegisterGuildServiceServer(s, guilds)
	healthSrv.SetServingStatus("go.zenithar.org/spotigraph/srv/users", grpc_health_v1.HealthCheckResponse_SERVING)

	pb.RegisterChapterServiceServer(s, chapters)
	healthSrv.SetServingStatus("go.zenithar.org/spotigraph/srv/chapters", grpc_health_v1.HealthCheckResponse_SERVING)

	pb.RegisterTribeServiceServer(s, tribes)
	healthSrv.SetServingStatus("go.zenithar.org/spotigraph/srv/tribes", grpc_health_v1.HealthCheckResponse_SERVING)

	// Returns the server instance
	return s
}
