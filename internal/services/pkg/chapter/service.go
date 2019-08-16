package chapter

import (
	"context"

	"go.zenithar.org/pkg/errors"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/commands"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

type service struct {
	createCmd commands.Handler
	getCmd    commands.Handler
	updateCmd commands.Handler
	deleteCmd commands.Handler
	searchCmd commands.Handler
}

// New returns a service instance
func New(chapters repositories.Chapter) services.Chapter {
	return &service{
		createCmd: commands.CreateHandler(chapters),
		getCmd:    commands.GetHandler(chapters),
		updateCmd: commands.UpdateHandler(chapters),
		deleteCmd: commands.DeleteHandler(chapters),
		searchCmd: commands.SearchHandler(chapters),
	}
}

// -----------------------------------------------------------------------------

func (s *service) Do(ctx context.Context, req interface{}) (interface{}, error) {
	switch req.(type) {
	case *chapterv1.CreateRequest:
		return s.createCmd.Handle(ctx, req)
	case *chapterv1.GetRequest:
		return s.getCmd.Handle(ctx, req)
	case *chapterv1.UpdateRequest:
		return s.updateCmd.Handle(ctx, req)
	case *chapterv1.DeleteRequest:
		return s.deleteCmd.Handle(ctx, req)
	case *chapterv1.SearchRequest:
		return s.searchCmd.Handle(ctx, req)
	default:
		return nil, errors.Newf(errors.Internal, nil, "unhandled command received (%T)", req)
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *chapterv1.CreateRequest) (*chapterv1.CreateResponse, error) {
	res, err := s.createCmd.Handle(ctx, req)
	return res.(*chapterv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *chapterv1.GetRequest) (*chapterv1.GetResponse, error) {
	res, err := s.getCmd.Handle(ctx, req)
	return res.(*chapterv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *chapterv1.UpdateRequest) (*chapterv1.UpdateResponse, error) {
	res, err := s.updateCmd.Handle(ctx, req)
	return res.(*chapterv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *chapterv1.DeleteRequest) (*chapterv1.DeleteResponse, error) {
	res, err := s.deleteCmd.Handle(ctx, req)
	return res.(*chapterv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *chapterv1.SearchRequest) (*chapterv1.SearchResponse, error) {
	res, err := s.searchCmd.Handle(ctx, req)
	return res.(*chapterv1.SearchResponse), err
}
