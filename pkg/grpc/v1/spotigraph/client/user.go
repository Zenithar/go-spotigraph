package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i User -t ../../../../../tools/templates/services/grpc/service.txt -o user.go

import (
	"context"

	"github.com/pkg/errors"
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
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	return cli.Create(ctx, req)
}

func (c *grpcUserClient) Delete(ctx context.Context, req *spotigraph.UserGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	return cli.Delete(ctx, req)
}

func (c *grpcUserClient) Get(ctx context.Context, req *spotigraph.UserGetReq) (sp1 *spotigraph.SingleUserRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	return cli.Get(ctx, req)
}

func (c *grpcUserClient) Search(ctx context.Context, req *spotigraph.UserSearchReq) (pp1 *spotigraph.PaginatedUserRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	return cli.Search(ctx, req)
}

func (c *grpcUserClient) Update(ctx context.Context, req *spotigraph.UserUpdateReq) (sp1 *spotigraph.SingleUserRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewUserClient(conn)

	// Call remote service
	return cli.Update(ctx, req)
}
