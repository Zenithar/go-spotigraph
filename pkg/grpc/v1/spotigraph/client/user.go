package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i User -t ../../../../../tools/templates/services/grpc/client.txt -o user.go

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

type grpcUserClient struct {
	factory ConnectionFactory
}

// NewUserClient returns a service client wrapped for gRPC
func NewUserClient(factory ConnectionFactory) services.User {
	return &grpcUserClient{
		factory: factory,
	}
}

// -----------------------------------------------------------------------------

func (c *grpcUserClient) Create(ctx context.Context, req *spotigraph.UserCreateReq) (sp1 *spotigraph.SingleUserRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.UserClient.Create")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	res, err := cli.Create(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.UserClient.Create"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcUserClient) Delete(ctx context.Context, req *spotigraph.UserGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.UserClient.Delete")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	res, err := cli.Delete(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.UserClient.Delete"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcUserClient) Get(ctx context.Context, req *spotigraph.UserGetReq) (sp1 *spotigraph.SingleUserRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.UserClient.Get")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	res, err := cli.Get(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.UserClient.Get"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcUserClient) Search(ctx context.Context, req *spotigraph.UserSearchReq) (pp1 *spotigraph.PaginatedUserRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.UserClient.Search")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	res, err := cli.Search(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.UserClient.Search"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcUserClient) Update(ctx context.Context, req *spotigraph.UserUpdateReq) (sp1 *spotigraph.SingleUserRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.UserClient.Update")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	res, err := cli.Update(ctx, req)
	if err != nil {
		log.For(ctx).Error("Remote call go error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.UserClient.Update"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}
