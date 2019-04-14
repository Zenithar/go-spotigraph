package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/client.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Guild -t ../../../../../tools/templates/services/grpc/client.txt -o guild.go

import (
	"context"

	"github.com/pkg/errors"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/grpc/v1/spotigraph/pb"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type grpcGuildClient struct {
	factory ConnectionFactory
}

// NewGuildClient returns a service client wrapped for gRPC
func NewGuildClient(factory ConnectionFactory) services.Guild {
	return &grpcGuildClient{
		factory: factory,
	}
}

// -----------------------------------------------------------------------------

func (c *grpcGuildClient) Create(ctx context.Context, req *spotigraph.GuildCreateReq) (res *spotigraph.SingleGuildRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.GuildClient.Create")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	res, err = cli.Create(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.GuildClient.Create"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcGuildClient) Delete(ctx context.Context, req *spotigraph.GuildGetReq) (res *spotigraph.EmptyRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.GuildClient.Delete")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	res, err = cli.Delete(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.GuildClient.Delete"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcGuildClient) Get(ctx context.Context, req *spotigraph.GuildGetReq) (res *spotigraph.SingleGuildRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.GuildClient.Get")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	res, err = cli.Get(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.GuildClient.Get"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcGuildClient) Search(ctx context.Context, req *spotigraph.GuildSearchReq) (res *spotigraph.PaginatedGuildRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.GuildClient.Search")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	res, err = cli.Search(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.GuildClient.Search"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcGuildClient) Update(ctx context.Context, req *spotigraph.GuildUpdateReq) (res *spotigraph.SingleGuildRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.GuildClient.Update")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	res, err = cli.Update(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.GuildClient.Update"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}
