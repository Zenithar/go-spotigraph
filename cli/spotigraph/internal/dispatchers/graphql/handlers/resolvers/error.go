package resolvers

import (
	"context"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type errorResolver struct{ *resolver }

func (r *errorResolver) Code(ctx context.Context, obj *spotigraph.Error) (int, error) {
	panic("not implemented")
}
