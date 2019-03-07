package services

import (
	"context"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// User defines user service contract
type User interface {
	Create(ctx context.Context, req *spotigraph.UserCreateReq) spotigraph.SingleUserRes
	Get(ctx context.Context, req *spotigraph.UserGetReq) spotigraph.SingleUserRes
	Update(ctx context.Context, req *spotigraph.UserCreateReq) spotigraph.SingleUserRes
	Delete(ctx context.Context, req *spotigraph.UserGetReq) spotigraph.EmptyRes
}
