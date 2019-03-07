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

// Squad defines squad service contract
type Squad interface {
	Create(ctx context.Context, req *spotigraph.SquadCreateReq) (*spotigraph.SingleSquadRes, error)
	Get(ctx context.Context, req *spotigraph.SquadGetReq) (*spotigraph.SingleSquadRes, error)
	Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (*spotigraph.SingleSquadRes, error)
	Delete(ctx context.Context, req *spotigraph.SquadGetReq) (*spotigraph.EmptyRes, error)
}
