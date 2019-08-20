package guild

import (
	"context"

	"go.zenithar.org/pkg/reactor"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/commands"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
)

type service struct {
	r reactor.Reactor
}

// New returns a service instance
func New(guilds repositories.Guild, persons repositories.Person, memberships repositories.Membership) services.Guild {
	// Initialize reactor
	r := reactor.New("spotigraph.guild.v1")

	// Register messages
	r.RegisterHandler(&guildv1.CreateRequest{}, commands.CreateHandler(guilds))
	r.RegisterHandler(&guildv1.GetRequest{}, commands.GetHandler(guilds))
	r.RegisterHandler(&guildv1.UpdateRequest{}, commands.UpdateHandler(guilds))
	r.RegisterHandler(&guildv1.DeleteRequest{}, commands.DeleteHandler(guilds))
	r.RegisterHandler(&guildv1.SearchRequest{}, commands.SearchHandler(guilds))
	r.RegisterHandler(&guildv1.JoinRequest{}, commands.JoinHandler(guilds, persons, memberships))
	r.RegisterHandler(&guildv1.LeaveRequest{}, commands.LeaveHandler(guilds, persons, memberships))

	// Service instance
	return &service{
		r: r,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *guildv1.CreateRequest) (*guildv1.CreateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *guildv1.GetRequest) (*guildv1.GetResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *guildv1.UpdateRequest) (*guildv1.UpdateResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *guildv1.DeleteRequest) (*guildv1.DeleteResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *guildv1.SearchRequest) (*guildv1.SearchResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.SearchResponse), err
}

func (s *service) Join(ctx context.Context, req *guildv1.JoinRequest) (*guildv1.JoinResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.JoinResponse), err
}

func (s *service) Leave(ctx context.Context, req *guildv1.LeaveRequest) (*guildv1.LeaveResponse, error) {
	res, err := s.r.Do(ctx, req)
	return res.(*guildv1.LeaveResponse), err
}
