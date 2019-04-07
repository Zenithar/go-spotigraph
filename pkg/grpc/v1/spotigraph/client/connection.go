package client

import (
	"context"

	"google.golang.org/grpc"
)

// ConnectionFactory describes the gRPC connection factory function
type ConnectionFactory func(ctx context.Context) (*grpc.ClientConn, func(), error)
