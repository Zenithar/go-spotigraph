package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Guild -t ../../../../../tools/templates/services/grpc/service.txt -o guild.go

import (
	"context"

	"github.com/pkg/errors"
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

func (c *grpcGuildClient) Create(ctx context.Context, req *spotigraph.GuildCreateReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	return cli.Create(ctx, req)
}

func (c *grpcGuildClient) Delete(ctx context.Context, req *spotigraph.GuildGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	return cli.Delete(ctx, req)
}

func (c *grpcGuildClient) Get(ctx context.Context, req *spotigraph.GuildGetReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	return cli.Get(ctx, req)
}

func (c *grpcGuildClient) Search(ctx context.Context, req *spotigraph.GuildSearchReq) (pp1 *spotigraph.PaginatedGuildRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	return cli.Search(ctx, req)
}

func (c *grpcGuildClient) Update(ctx context.Context, req *spotigraph.GuildUpdateReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewGuildClient(conn)

	// Call remote service
	return cli.Update(ctx, req)
}
