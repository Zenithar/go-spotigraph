package resolvers

import (
	"context"
	"encoding/base64"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type queryResolver struct {
	root     *resolver
	users    services.User
	squads   services.Squad
	chapters services.Chapter
	guilds   services.Guild
	tribes   services.Tribe
}

func (r *queryResolver) Me(ctx context.Context) (*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchForUsers(ctx context.Context, paging *generated.PagingRequest) (*generated.UserPagingConnection, error) {
	// Do service request
	res, err := r.users.Search(ctx, &spotigraph.UserSearchReq{})
	if err != nil {
		return nil, err
	}

	// Retrieve last element
	nextToken := base64.URLEncoding.EncodeToString([]byte(res.Members[len(res.Members)-1].GetId()))

	// Format result
	return &generated.UserPagingConnection{
		Items:     res.Members,
		NextToken: &nextToken,
	}, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*spotigraph.Domain_User, error) {
	// Do service request
	res, err := r.users.Get(ctx, &spotigraph.UserGetReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	// No Error
	return res.Entity, nil
}

func (r *queryResolver) GetUsers(ctx context.Context, ids []string) ([]*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchForSquads(ctx context.Context, paging *generated.PagingRequest) (*generated.SquadPagingConnection, error) {
	// Do service request
	res, err := r.squads.Search(ctx, &spotigraph.SquadSearchReq{})
	if err != nil {
		return nil, err
	}

	// Retrieve last element
	nextToken := base64.URLEncoding.EncodeToString([]byte(res.Members[len(res.Members)-1].GetId()))

	// Format result
	return &generated.SquadPagingConnection{
		Items:     res.Members,
		NextToken: &nextToken,
	}, nil
}

func (r *queryResolver) GetSquad(ctx context.Context, id string) (*spotigraph.Domain_Squad, error) {
	// Do service request
	res, err := r.squads.Get(ctx, &spotigraph.SquadGetReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	// No Error
	return res.Entity, nil
}

func (r *queryResolver) GetSquads(ctx context.Context, ids []string) ([]*spotigraph.Domain_Squad, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchForTribes(ctx context.Context, paging *generated.PagingRequest) (*generated.TribePagingConnection, error) {
	// Do service request
	res, err := r.tribes.Search(ctx, &spotigraph.TribeSearchReq{})
	if err != nil {
		return nil, err
	}

	// Retrieve last element
	nextToken := base64.URLEncoding.EncodeToString([]byte(res.Members[len(res.Members)-1].GetId()))

	// Format result
	return &generated.TribePagingConnection{
		Items:     res.Members,
		NextToken: &nextToken,
	}, nil
}

func (r *queryResolver) GetTribe(ctx context.Context, id string) (*spotigraph.Domain_Tribe, error) {
	// Do service request
	res, err := r.tribes.Get(ctx, &spotigraph.TribeGetReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	// No Error
	return res.Entity, nil
}

func (r *queryResolver) GetTribes(ctx context.Context, ids []string) ([]*spotigraph.Domain_Tribe, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchForChapters(ctx context.Context, paging *generated.PagingRequest) (*generated.ChapterPagingConnection, error) {
	// Do service request
	res, err := r.chapters.Search(ctx, &spotigraph.ChapterSearchReq{})
	if err != nil {
		return nil, err
	}

	// Retrieve last element
	nextToken := base64.URLEncoding.EncodeToString([]byte(res.Members[len(res.Members)-1].GetId()))

	// Format result
	return &generated.ChapterPagingConnection{
		Items:     res.Members,
		NextToken: &nextToken,
	}, nil
}

func (r *queryResolver) GetChapter(ctx context.Context, id string) (*spotigraph.Domain_Chapter, error) {
	// Do service request
	res, err := r.chapters.Get(ctx, &spotigraph.ChapterGetReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	// No Error
	return res.Entity, nil
}

func (r *queryResolver) GetChapters(ctx context.Context, ids []string) ([]*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchForGuilds(ctx context.Context, paging *generated.PagingRequest) (*generated.GuildPagingConnection, error) {
	// Do service request
	res, err := r.guilds.Search(ctx, &spotigraph.GuildSearchReq{})
	if err != nil {
		return nil, err
	}

	// Retrieve last element
	nextToken := base64.URLEncoding.EncodeToString([]byte(res.Members[len(res.Members)-1].GetId()))

	// Format result
	return &generated.GuildPagingConnection{
		Items:     res.Members,
		NextToken: &nextToken,
	}, nil
}

func (r *queryResolver) GetGuild(ctx context.Context, id string) (*spotigraph.Domain_Guild, error) {
	// Do service request
	res, err := r.guilds.Get(ctx, &spotigraph.GuildGetReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	// No Error
	return res.Entity, nil
}

func (r *queryResolver) GetGuilds(ctx context.Context, ids []string) ([]*spotigraph.Domain_Guild, error) {
	panic("not implemented")
}
