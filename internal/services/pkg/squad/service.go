package squad

import (
	"context"

	"go.zenithar.org/pkg/reactor"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad/internal/commands"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

type service struct {
	r reactor.Reactor
}

// New returns a service instance
func New(squads repositories.Squad, persons repositories.Person, memberships repositories.Membership) services.Squad {
	// Initialize reactor
	r := reactor.New("spotigraph.squad.v1")

	// Register messages
	r.RegisterHandler(&squadv1.CreateRequest{}, commands.CreateHandler(squads))
	r.RegisterHandler(&squadv1.GetRequest{}, commands.GetHandler(squads))
	r.RegisterHandler(&squadv1.UpdateRequest{}, commands.UpdateHandler(squads))
	r.RegisterHandler(&squadv1.DeleteRequest{}, commands.DeleteHandler(squads))
	r.RegisterHandler(&squadv1.SearchRequest{}, commands.SearchHandler(squads))
	r.RegisterHandler(&squadv1.JoinRequest{}, commands.JoinHandler(squads, persons, memberships))
	r.RegisterHandler(&squadv1.LeaveRequest{}, commands.LeaveHandler(squads, persons, memberships))

	// Service instance
	return &service{
		r: r,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *squadv1.CreateRequest) (*squadv1.CreateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *squadv1.GetRequest) (*squadv1.GetResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *squadv1.UpdateRequest) (*squadv1.UpdateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *squadv1.DeleteRequest) (*squadv1.DeleteResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *squadv1.SearchRequest) (*squadv1.SearchResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.SearchResponse), err
}

func (s *service) Join(ctx context.Context, req *squadv1.JoinRequest) (*squadv1.JoinResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.JoinResponse), err
}

func (s *service) Leave(ctx context.Context, req *squadv1.LeaveRequest) (*squadv1.LeaveResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*squadv1.LeaveResponse), err
}
