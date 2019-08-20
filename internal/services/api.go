package services

import (
	"context"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

//go:generate mockgen -destination test/mock/chapter.gen.go -package mock go.zenithar.org/spotigraph/internal/services Chapter

// ChapterRetriever defines read-only service methods. (ISP)
type ChapterRetriever interface {
	Get(ctx context.Context, req *chapterv1.GetRequest) (res *chapterv1.GetResponse, err error)
	Search(ctx context.Context, req *chapterv1.SearchRequest) (res *chapterv1.SearchResponse, err error)
}

// ChapterModifier defines read-write service methods. (ISP)
type ChapterModifier interface {
	Create(ctx context.Context, req *chapterv1.CreateRequest) (res *chapterv1.CreateResponse, err error)
	Update(ctx context.Context, req *chapterv1.UpdateRequest) (res *chapterv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *chapterv1.DeleteRequest) (res *chapterv1.DeleteResponse, err error)
	Join(ctx context.Context, req *chapterv1.JoinRequest) (res *chapterv1.JoinResponse, err error)
	Leave(ctx context.Context, req *chapterv1.LeaveRequest) (res *chapterv1.LeaveResponse, err error)
}

// Chapter defines chapter service contract
type Chapter interface {
	ChapterRetriever
	ChapterModifier
}

//go:generate mockgen -destination test/mock/squad.gen.go -package mock go.zenithar.org/spotigraph/internal/services Squad

// SquadRetriever defines read-only service methods. (ISP)
type SquadRetriever interface {
	Get(ctx context.Context, req *squadv1.GetRequest) (res *squadv1.GetResponse, err error)
	Search(ctx context.Context, req *squadv1.SearchRequest) (res *squadv1.SearchResponse, err error)
}

// SquadModifier defines read-write service methods. (ISP)
type SquadModifier interface {
	Create(ctx context.Context, req *squadv1.CreateRequest) (res *squadv1.CreateResponse, err error)
	Update(ctx context.Context, req *squadv1.UpdateRequest) (res *squadv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *squadv1.DeleteRequest) (res *squadv1.DeleteResponse, err error)
	Join(ctx context.Context, req *squadv1.JoinRequest) (res *squadv1.JoinResponse, err error)
	Leave(ctx context.Context, req *squadv1.LeaveRequest) (res *squadv1.LeaveResponse, err error)
}

// Squad defines squad service contract
type Squad interface {
	SquadRetriever
	SquadModifier
}

//go:generate mockgen -destination test/mock/guild.gen.go -package mock go.zenithar.org/spotigraph/internal/services Guild

// GuildRetriever defines read-only service methods. (ISP)
type GuildRetriever interface {
	Get(ctx context.Context, req *guildv1.GetRequest) (res *guildv1.GetResponse, err error)
	Search(ctx context.Context, req *guildv1.SearchRequest) (res *guildv1.SearchResponse, err error)
}

// GuildModifier defines read-write service methods. (ISP)
type GuildModifier interface {
	Create(ctx context.Context, req *guildv1.CreateRequest) (res *guildv1.CreateResponse, err error)
	Update(ctx context.Context, req *guildv1.UpdateRequest) (res *guildv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *guildv1.DeleteRequest) (res *guildv1.DeleteResponse, err error)
	Join(ctx context.Context, req *guildv1.JoinRequest) (res *guildv1.JoinResponse, err error)
	Leave(ctx context.Context, req *guildv1.LeaveRequest) (res *guildv1.LeaveResponse, err error)
}

// Guild defines guild service contract
type Guild interface {
	GuildRetriever
	GuildModifier
}

//go:generate mockgen -destination test/mock/person.gen.go -package mock go.zenithar.org/spotigraph/internal/services Person

// PersonRetriever defines read-only service methods. (ISP)
type PersonRetriever interface {
	Get(ctx context.Context, req *personv1.GetRequest) (res *personv1.GetResponse, err error)
	Search(ctx context.Context, req *personv1.SearchRequest) (res *personv1.SearchResponse, err error)
}

// PersonModifier defines read-write service methods. (ISP)
type PersonModifier interface {
	Create(ctx context.Context, req *personv1.CreateRequest) (res *personv1.CreateResponse, err error)
	Update(ctx context.Context, req *personv1.UpdateRequest) (res *personv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *personv1.DeleteRequest) (res *personv1.DeleteResponse, err error)
}

// Person defines person service contract
type Person interface {
	PersonModifier
	PersonRetriever
}
