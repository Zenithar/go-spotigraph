package internal

import (
	"context"

	"github.com/sony/gobreaker"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type chapterWithBreaker struct {
	next services.Chapter

	createBreaker *gobreaker.CircuitBreaker
	deleteBreaker *gobreaker.CircuitBreaker
	getBreaker    *gobreaker.CircuitBreaker
	searchBreaker *gobreaker.CircuitBreaker
	updateBreaker *gobreaker.CircuitBreaker
}

// NewChapterWithBreaker instruments a chapter service to add circuit breaker
func NewChapterWithBreaker(next services.Chapter) services.Chapter {
	failureFunc := func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio <= 0.6
	}

	return &chapterWithBreaker{
		next: next,
		createBreaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "Chapter.Create",
			ReadyToTrip: failureFunc,
		}),
		deleteBreaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "Chapter.Delete",
			ReadyToTrip: failureFunc,
		}),
		getBreaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "Chapter.Get",
			ReadyToTrip: failureFunc,
		}),
		searchBreaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "Chapter.Search",
			ReadyToTrip: failureFunc,
		}),
		updateBreaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "Chapter.Update",
			ReadyToTrip: failureFunc,
		}),
	}
}

// -----------------------------------------------------------------------------

// Create implements services.Chapter
func (d chapterWithBreaker) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (*spotigraph.SingleChapterRes, error) {
	body, err := d.createBreaker.Execute(func() (interface{}, error) {
		return d.next.Create(ctx, req)
	})
	return body.(*spotigraph.SingleChapterRes), err
}

// Delete implements services.Chapter
func (d chapterWithBreaker) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.EmptyRes, err error) {
	body, err := d.createBreaker.Execute(func() (interface{}, error) {
		return d.next.Delete(ctx, req)
	})
	return body.(*spotigraph.EmptyRes), err
}

// Get implements services.Chapter
func (d chapterWithBreaker) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.SingleChapterRes, err error) {
	body, err := d.createBreaker.Execute(func() (interface{}, error) {
		return d.next.Get(ctx, req)
	})
	return body.(*spotigraph.SingleChapterRes), err
}

// Search implements services.Chapter
func (d chapterWithBreaker) Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (res *spotigraph.PaginatedChapterRes, err error) {
	body, err := d.createBreaker.Execute(func() (interface{}, error) {
		return d.next.Search(ctx, req)
	})
	return body.(*spotigraph.PaginatedChapterRes), err
}

// Update implements services.Chapter
func (d chapterWithBreaker) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (res *spotigraph.SingleChapterRes, err error) {
	body, err := d.createBreaker.Execute(func() (interface{}, error) {
		return d.next.Update(ctx, req)
	})
	return body.(*spotigraph.SingleChapterRes), err
}
