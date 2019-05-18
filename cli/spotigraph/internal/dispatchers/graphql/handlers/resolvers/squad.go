package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type squadResolver struct{ *resolver }

func (r *squadResolver) ProductOwner(ctx context.Context, obj *spotigraph.Domain_Squad) (*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *squadResolver) Members(ctx context.Context, obj *spotigraph.Domain_Squad, paging *generated.PagingRequest) (*generated.UserPagingConnection, error) {
	panic("not implemented")
}
