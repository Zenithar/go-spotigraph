package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type tribeResolver struct{ *resolver }

func (r *tribeResolver) Squads(ctx context.Context, obj *spotigraph.Domain_Tribe, paging *generated.PagingRequest) (*generated.SquadPagingConnection, error) {
	panic("not implemented")
}
