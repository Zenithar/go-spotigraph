package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type userResolver struct{ *resolver }

func (r *userResolver) Squads(ctx context.Context, obj *spotigraph.Domain_User, paging *generated.PagingRequest) (*generated.SquadPagingConnection, error) {
	panic("not implemented")
}

func (r *userResolver) Chapter(ctx context.Context, obj *spotigraph.Domain_User) (*spotigraph.Domain_Chapter, error) {
	panic("not implemented")
}

func (r *userResolver) Guilds(ctx context.Context, obj *spotigraph.Domain_User, paging *generated.PagingRequest) (*generated.GuildPagingConnection, error) {
	panic("not implemented")
}
