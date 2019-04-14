package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../../tools/templates/services/grpc/client.txt -o tribe.go

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

type grpcTribeClient struct {
	factory ConnectionFactory
}

// NewTribeClient returns a service client wrapped for gRPC
func NewTribeClient(factory ConnectionFactory) services.Tribe {
	return &grpcTribeClient{
		factory: factory,
	}
}

// -----------------------------------------------------------------------------

func (c *grpcTribeClient) Create(ctx context.Context, req *spotigraph.TribeCreateReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.TribeClient.Create")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	res, err := cli.Create(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.TribeClient.Create"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcTribeClient) Delete(ctx context.Context, req *spotigraph.TribeGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.TribeClient.Delete")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	res, err := cli.Delete(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.TribeClient.Delete"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcTribeClient) Get(ctx context.Context, req *spotigraph.TribeGetReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.TribeClient.Get")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	res, err := cli.Get(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.TribeClient.Get"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcTribeClient) Search(ctx context.Context, req *spotigraph.TribeSearchReq) (pp1 *spotigraph.PaginatedTribeRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.TribeClient.Search")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	res, err := cli.Search(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.TribeClient.Search"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcTribeClient) Update(ctx context.Context, req *spotigraph.TribeUpdateReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.TribeClient.Update")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	res, err := cli.Update(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.TribeClient.Update"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}
