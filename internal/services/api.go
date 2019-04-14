package services

import (
	"context"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

//go:generate mockgen -destination test/mock/user.gen.go -package mock go.zenithar.org/spotigraph/internal/services User

// User defines user service contract
type User interface {
	Create(ctx context.Context, req *spotigraph.UserCreateReq) (res *spotigraph.SingleUserRes, err error)
	Get(ctx context.Context, req *spotigraph.UserGetReq) (res *spotigraph.SingleUserRes, err error)
	Update(ctx context.Context, req *spotigraph.UserUpdateReq) (res *spotigraph.SingleUserRes, err error)
	Delete(ctx context.Context, req *spotigraph.UserGetReq) (res *spotigraph.EmptyRes, err error)
	Search(ctx context.Context, req *spotigraph.UserSearchReq) (res *spotigraph.PaginatedUserRes, err error)
}

//go:generate mockgen -destination test/mock/squad.gen.go -package mock go.zenithar.org/spotigraph/internal/services Squad

// Squad defines squad service contract
type Squad interface {
	Create(ctx context.Context, req *spotigraph.SquadCreateReq) (res *spotigraph.SingleSquadRes, err error)
	Get(ctx context.Context, req *spotigraph.SquadGetReq) (res *spotigraph.SingleSquadRes, err error)
	Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (res *spotigraph.SingleSquadRes, err error)
	Delete(ctx context.Context, req *spotigraph.SquadGetReq) (res *spotigraph.EmptyRes, err error)
	Search(ctx context.Context, req *spotigraph.SquadSearchReq) (res *spotigraph.PaginatedSquadRes, err error)
}

//go:generate mockgen -destination test/mock/chapter.gen.go -package mock go.zenithar.org/spotigraph/internal/services Chapter

// Chapter defines chapter service contract
type Chapter interface {
	Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (res *spotigraph.SingleChapterRes, err error)
	Get(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.SingleChapterRes, err error)
	Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (res *spotigraph.SingleChapterRes, err error)
	Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.EmptyRes, err error)
	Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (res *spotigraph.PaginatedChapterRes, err error)
}

//go:generate mockgen -destination test/mock/guild.gen.go -package mock go.zenithar.org/spotigraph/internal/services Guild

// Guild defines guild service contract
type Guild interface {
	Create(ctx context.Context, req *spotigraph.GuildCreateReq) (res *spotigraph.SingleGuildRes, err error)
	Get(ctx context.Context, req *spotigraph.GuildGetReq) (res *spotigraph.SingleGuildRes, err error)
	Update(ctx context.Context, req *spotigraph.GuildUpdateReq) (res *spotigraph.SingleGuildRes, err error)
	Delete(ctx context.Context, req *spotigraph.GuildGetReq) (res *spotigraph.EmptyRes, err error)
	Search(ctx context.Context, req *spotigraph.GuildSearchReq) (res *spotigraph.PaginatedGuildRes, err error)
}

//go:generate mockgen -destination test/mock/tribe.gen.go -package mock go.zenithar.org/spotigraph/internal/services Tribe

// Tribe defines tribe service contract
type Tribe interface {
	Create(ctx context.Context, req *spotigraph.TribeCreateReq) (res *spotigraph.SingleTribeRes, err error)
	Get(ctx context.Context, req *spotigraph.TribeGetReq) (res *spotigraph.SingleTribeRes, err error)
	Update(ctx context.Context, req *spotigraph.TribeUpdateReq) (res *spotigraph.SingleTribeRes, err error)
	Delete(ctx context.Context, req *spotigraph.TribeGetReq) (res *spotigraph.EmptyRes, err error)
	Search(ctx context.Context, req *spotigraph.TribeSearchReq) (res *spotigraph.PaginatedTribeRes, err error)
}

//go:generate mockgen -destination test/mock/graph.gen.go -package mock go.zenithar.org/spotigraph/internal/services Graph

// Graph defines graph service contract
type Graph interface {
	Expand(cx context.Context, req *spotigraph.NodeInfoReq) (res *spotigraph.GraphRes, err error)
}
