package graphql

import (
	"context"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Error() ErrorResolver {
	return &errorResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type errorResolver struct{ *Resolver }

func (r *errorResolver) Code(ctx context.Context, obj *spotigraph.Error) (int, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input spotigraph.UserCreateReq) (*spotigraph.SingleUserRes, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]spotigraph.Domain_User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*spotigraph.SingleUserRes, error) {
	panic("not implemented")
}
