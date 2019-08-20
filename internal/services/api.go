package services

import (
	"context"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
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
