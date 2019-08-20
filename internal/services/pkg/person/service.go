package person

import (
	"context"

	"go.zenithar.org/pkg/reactor"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/person/internal/commands"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
)

type service struct {
	r reactor.Reactor
}

// New returns a service instance
func New(persons repositories.Person) services.Person {
	// Initialize reactor
	r := reactor.New("spotigraph.person.v1")

	// Register messages
	r.RegisterHandler(&personv1.CreateRequest{}, commands.CreateHandler(persons))
	r.RegisterHandler(&personv1.GetRequest{}, commands.GetHandler(persons))
	r.RegisterHandler(&personv1.UpdateRequest{}, commands.UpdateHandler(persons))
	r.RegisterHandler(&personv1.DeleteRequest{}, commands.DeleteHandler(persons))
	r.RegisterHandler(&personv1.SearchRequest{}, commands.SearchHandler(persons))

	// Service instance
	return &service{
		r: r,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *personv1.CreateRequest) (*personv1.CreateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*personv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *personv1.GetRequest) (*personv1.GetResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*personv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *personv1.UpdateRequest) (*personv1.UpdateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*personv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *personv1.DeleteRequest) (*personv1.DeleteResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*personv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *personv1.SearchRequest) (*personv1.SearchResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*personv1.SearchResponse), err
}
