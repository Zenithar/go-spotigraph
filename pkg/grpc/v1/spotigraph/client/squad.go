package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../../tools/templates/services/grpc/service.txt -o squad.go

import (
	"context"

	"github.com/pkg/errors"
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
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	return cli.Create(ctx, req)
}

func (c *grpcSquadClient) Delete(ctx context.Context, req *spotigraph.SquadGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	return cli.Delete(ctx, req)
}

func (c *grpcSquadClient) Get(ctx context.Context, req *spotigraph.SquadGetReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	return cli.Get(ctx, req)
}

func (c *grpcSquadClient) Search(ctx context.Context, req *spotigraph.SquadSearchReq) (pp1 *spotigraph.PaginatedSquadRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	return cli.Search(ctx, req)
}

func (c *grpcSquadClient) Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewSquadClient(conn)

	// Call remote service
	return cli.Update(ctx, req)
}
