package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type guildResolver struct{ *resolver }

func (r *guildResolver) Members(ctx context.Context, obj *spotigraph.Domain_Guild, paging *generated.PagingRequest) (*generated.UserPagingConnection, error) {
	panic("not implemented")
}
