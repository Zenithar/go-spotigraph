package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type chapterResolver struct{ *resolver }

func (r *chapterResolver) Leader(ctx context.Context, obj *spotigraph.Domain_Chapter) (*spotigraph.Domain_User, error) {
	panic("not implemented")
}

func (r *chapterResolver) Members(ctx context.Context, obj *spotigraph.Domain_Chapter, paging *generated.PagingRequest) (*generated.UserPagingConnection, error) {
	panic("not implemented")
}
