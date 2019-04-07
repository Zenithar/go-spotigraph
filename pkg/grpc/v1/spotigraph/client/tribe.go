package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../../tools/templates/services/grpc/service.txt -o tribe.go

import (
	"context"

	"github.com/pkg/errors"
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
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	return cli.Create(ctx, req)
}

func (c *grpcTribeClient) Delete(ctx context.Context, req *spotigraph.TribeGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	return cli.Delete(ctx, req)
}

func (c *grpcTribeClient) Get(ctx context.Context, req *spotigraph.TribeGetReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	return cli.Get(ctx, req)
}

func (c *grpcTribeClient) Search(ctx context.Context, req *spotigraph.TribeSearchReq) (pp1 *spotigraph.PaginatedTribeRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	return cli.Search(ctx, req)
}

func (c *grpcTribeClient) Update(ctx context.Context, req *spotigraph.TribeUpdateReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewTribeClient(conn)

	// Call remote service
	return cli.Update(ctx, req)
}
