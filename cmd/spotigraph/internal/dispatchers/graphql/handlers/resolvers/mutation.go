package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type mutationResolver struct {
	root     *resolver
	users    services.User
	squads   services.Squad
	chapters services.Chapter
	guilds   services.Guild
	tribes   services.Tribe
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *spotigraph.UserCreateReq) (*spotigraph.Domain_User, error) {
	res, err := r.users.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.Entity, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *generated.UserUpdateInput) (*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateSquad(ctx context.Context, input *spotigraph.SquadCreateReq) (*spotigraph.Domain_Squad, error) {
	res, err := r.squads.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.Entity, nil
}

func (r *mutationResolver) UpdateSquad(ctx context.Context, input *generated.SquadUpdateInput) (*spotigraph.Domain_Squad, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteSquad(ctx context.Context, id *string) (*spotigraph.Domain_Squad, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddSquadMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Squad, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveSquadMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Squad, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTribe(ctx context.Context, input *spotigraph.TribeCreateReq) (*spotigraph.Domain_Tribe, error) {
	res, err := r.tribes.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.Entity, nil
}

func (r *mutationResolver) UpdateTribe(ctx context.Context, input *generated.TribeUpdateInput) (*spotigraph.Domain_Tribe, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTribe(ctx context.Context, id *string) (*spotigraph.Domain_Tribe, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddTribeSquads(ctx context.Context, id string, squads []*string) (*spotigraph.Domain_Tribe, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveTribeSquads(ctx context.Context, id string, squads []*string) (*spotigraph.Domain_Tribe, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input *spotigraph.ChapterCreateReq) (*spotigraph.Domain_Chapter, error) {
	res, err := r.chapters.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.Entity, nil
}

func (r *mutationResolver) UpdateChapter(ctx context.Context, input *generated.ChapterUpdateInput) (*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteChapter(ctx context.Context, id *string) (*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddChapterMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveChapterMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateGuild(ctx context.Context, input *spotigraph.GuildCreateReq) (*spotigraph.Domain_Guild, error) {
	res, err := r.guilds.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.Entity, nil
}

func (r *mutationResolver) UpdateGuild(ctx context.Context, input *generated.GuildUpdateInput) (*spotigraph.Domain_Guild, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteGuild(ctx context.Context, id *string) (*spotigraph.Domain_Guild, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddGuildMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Guild, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveGuildMembers(ctx context.Context, id string, users []*string) (*spotigraph.Domain_Guild, error) {
	panic("not implemented")
}
