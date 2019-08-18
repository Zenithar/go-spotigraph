package chapter

import (
	"context"

	"go.zenithar.org/pkg/reactor"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/commands"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

type service struct {
	r reactor.Reactor
}

// New returns a service instance
func New(chapters repositories.Chapter) services.Chapter {
	// Initialize reactor
	r := reactor.New("spotigraph.chapter.v1")

	// Register messages
	r.RegisterHandler(&chapterv1.CreateRequest{}, commands.CreateHandler(chapters))
	r.RegisterHandler(&chapterv1.GetRequest{}, commands.GetHandler(chapters))
	r.RegisterHandler(&chapterv1.UpdateRequest{}, commands.UpdateHandler(chapters))
	r.RegisterHandler(&chapterv1.DeleteRequest{}, commands.DeleteHandler(chapters))
	r.RegisterHandler(&chapterv1.SearchRequest{}, commands.SearchHandler(chapters))

	// Service instance
	return &service{
		r: r,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *chapterv1.CreateRequest) (*chapterv1.CreateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*chapterv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *chapterv1.GetRequest) (*chapterv1.GetResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*chapterv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *chapterv1.UpdateRequest) (*chapterv1.UpdateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*chapterv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *chapterv1.DeleteRequest) (*chapterv1.DeleteResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*chapterv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *chapterv1.SearchRequest) (*chapterv1.SearchResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*chapterv1.SearchResponse), err
}
