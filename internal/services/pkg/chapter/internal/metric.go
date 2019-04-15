package internal

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type chapterWithMetrics struct {
	next services.Chapter

	createdChapterCount *stats.Int64Measure
	deletedChapterCount *stats.Int64Measure
}

// NewChapterWithMetrics instruments a chapter service to add metrics
func NewChapterWithMetrics(next services.Chapter) services.Chapter {
	s := &chapterWithMetrics{
		next:                next,
		createdChapterCount: stats.Int64("created_chapter_count", "Number of chapters created", stats.UnitDimensionless),
		deletedChapterCount: stats.Int64("deleted_chapter_count", "Number of chapters deleted", stats.UnitDimensionless),
	}

	// Register metric views
	if err := s.enableViews(); err != nil {
		log.Bg().Fatal("Unable to register views of Chapter service.")
		return nil
	}

	return s
}

// -----------------------------------------------------------------------------

// Create implements services.Chapter
func (d chapterWithMetrics) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (res *spotigraph.SingleChapterRes, err error) {
	res, err = d.next.Create(ctx, req)

	// Increment the measure
	stats.Record(ctx, d.createdChapterCount.M(1))

	return res, err
}

// Delete implements services.Chapter
func (d chapterWithMetrics) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.EmptyRes, err error) {
	res, err = d.next.Delete(ctx, req)

	// Increment the measure
	stats.Record(ctx, d.deletedChapterCount.M(1))

	return res, err
}

// Get implements services.Chapter
func (d chapterWithMetrics) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.SingleChapterRes, err error) {
	return d.next.Get(ctx, req)
}

// Search implements services.Chapter
func (d chapterWithMetrics) Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (res *spotigraph.PaginatedChapterRes, err error) {
	return d.next.Search(ctx, req)
}

// Update implements services.Chapter
func (d chapterWithMetrics) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (res *spotigraph.SingleChapterRes, err error) {
	return d.next.Update(ctx, req)
}

// -----------------------------------------------------------------------------

func (d chapterWithMetrics) enableViews() error {
	createdChapterCountView := &view.View{
		Name:        "chapter_created_count",
		Description: "Count of chapters created",
		Measure:     d.createdChapterCount,
		Aggregation: view.Count(),
	}

	deletedChapterCountView := &view.View{
		Name:        "chapter_deleted_count",
		Description: "Count of chapters deleted",
		Measure:     d.deletedChapterCount,
		Aggregation: view.Count(),
	}

	// Register all views
	return view.Register(
		createdChapterCountView,
		deletedChapterCountView,
	)
}
