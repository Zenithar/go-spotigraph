package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/service.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Chapter -t ../../../../../tools/templates/services/grpc/service.txt -o chapter.go

import (
	"context"

	"github.com/pkg/errors"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/grpc/v1/spotigraph/pb"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type grpcChapterClient struct {
	factory ConnectionFactory
}

// NewChapterClient returns a service client wrapped for gRPC
func NewChapterClient(factory ConnectionFactory) services.Chapter {
	return &grpcChapterClient{
		factory: factory,
	}
}

// -----------------------------------------------------------------------------

func (c *grpcChapterClient) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	return cli.Create(ctx, req)
}

func (c *grpcChapterClient) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	return cli.Delete(ctx, req)
}

func (c *grpcChapterClient) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	return cli.Get(ctx, req)
}

func (c *grpcChapterClient) Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (pp1 *spotigraph.PaginatedChapterRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	return cli.Search(ctx, req)
}

func (c *grpcChapterClient) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	return cli.Update(ctx, req)
}
