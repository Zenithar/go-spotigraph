package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../../tools/templates/services/grpc/client.txt -o squad.go

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

type grpcSquadClient struct {
	factory ConnectionFactory
}

// NewSquadClient returns a service client wrapped for gRPC
func NewSquadClient(factory ConnectionFactory) services.Squad {
	return &grpcSquadClient{
		factory: factory,
	}
}

// -----------------------------------------------------------------------------

func (c *grpcSquadClient) Create(ctx context.Context, req *spotigraph.SquadCreateReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.SquadClient.Create")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	res, err := cli.Create(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.SquadClient.Create"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcSquadClient) Delete(ctx context.Context, req *spotigraph.SquadGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.SquadClient.Delete")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	res, err := cli.Delete(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.SquadClient.Delete"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcSquadClient) Get(ctx context.Context, req *spotigraph.SquadGetReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.SquadClient.Get")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	res, err := cli.Get(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.SquadClient.Get"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcSquadClient) Search(ctx context.Context, req *spotigraph.SquadSearchReq) (pp1 *spotigraph.PaginatedSquadRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.SquadClient.Search")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	res, err := cli.Search(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.SquadClient.Search"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcSquadClient) Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.SquadClient.Update")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	res, err := cli.Update(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.SquadClient.Update"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}
