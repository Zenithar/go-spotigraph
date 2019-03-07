package services

import (
	"context"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// User defines user service contract
type User interface {
	Create(ctx context.Context, req *spotigraph.UserCreateReq) (*spotigraph.SingleUserRes, error)
	Get(ctx context.Context, req *spotigraph.UserGetReq) (*spotigraph.SingleUserRes, error)
	Update(ctx context.Context, req *spotigraph.UserUpdateReq) (*spotigraph.SingleUserRes, error)
	Delete(ctx context.Context, req *spotigraph.UserGetReq) (*spotigraph.EmptyRes, error)
}
