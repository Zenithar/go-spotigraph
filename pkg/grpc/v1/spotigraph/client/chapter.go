package client

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../../tools/templates/services/grpc/client.txt template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Chapter -t ../../../../../tools/templates/services/grpc/client.txt -o chapter.go

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

func (c *grpcChapterClient) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (res *spotigraph.SingleChapterRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.ChapterClient.Create")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	res, err = cli.Create(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.ChapterClient.Create"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcChapterClient) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.EmptyRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.ChapterClient.Delete")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	res, err = cli.Delete(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.ChapterClient.Delete"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcChapterClient) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.SingleChapterRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.ChapterClient.Get")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	res, err = cli.Get(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.ChapterClient.Get"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcChapterClient) Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (res *spotigraph.PaginatedChapterRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.ChapterClient.Search")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	res, err = cli.Search(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.ChapterClient.Search"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}

func (c *grpcChapterClient) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (res *spotigraph.SingleChapterRes, err error) {
	ctx, span := trace.StartSpan(ctx, "grpc.spotigraph.v1.ChapterClient.Update")
	defer span.End()

	// Retrieve a connection from factory
	conn, releaser, err := c.factory(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.Wrap(err, "unable to initialize gRPC connection")
	}
	defer releaser()

	// Wrap the connection
	cli := pb.NewChapterClient(conn)

	// Call remote service
	res, err = cli.Update(ctx, req)
	if err != nil {
		log.For(ctx).Error("gRPC remote call raised an error", zap.Error(err), zap.String("service", "grpc.spotigraph.v1.ChapterClient.Update"))
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	// Return result
	return res, err
}
