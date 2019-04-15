package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/cache"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type chapterWithCache struct {
	next services.Chapter
	ttl  time.Duration

	mgr                    cache.Storage
	cacheHitChapterCount   *stats.Int64Measure
	cacheMissChapterCount  *stats.Int64Measure
	cacheErrorChapterCount *stats.Int64Measure
}

// NewChapterWithCache instruments a chapter service to add cache
func NewChapterWithCache(next services.Chapter, mgr cache.Storage, ttl time.Duration) services.Chapter {
	s := &chapterWithCache{
		next:                   next,
		mgr:                    mgr,
		ttl:                    ttl,
		cacheHitChapterCount:   stats.Int64("cache_hit_chapter_count", "Number of chapters cache hit", stats.UnitDimensionless),
		cacheMissChapterCount:  stats.Int64("cache_miss_chapter_count", "Number of chapters cache miss", stats.UnitDimensionless),
		cacheErrorChapterCount: stats.Int64("cache_error_chapter_count", "Number of chapters cache error", stats.UnitDimensionless),
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
func (d chapterWithCache) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (res *spotigraph.SingleChapterRes, err error) {
	return d.next.Create(ctx, req)
}

// Delete implements services.Chapter
func (d chapterWithCache) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.EmptyRes, err error) {
	defer func() {
		log.CheckErrCtx(ctx, "Unable to remove chapter from cache", d.mgr.Remove(ctx, d.key(req.Id)))
	}()
	return d.next.Delete(ctx, req)
}

// Get implements services.Chapter
func (d chapterWithCache) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (res *spotigraph.SingleChapterRes, err error) {
	var k = d.key(req.Id)

	// Check from cache
	payload, err := d.mgr.Get(ctx, k)
	if err != nil {
		if err == cache.ErrCacheMiss {
			log.For(ctx).Debug("Object not found in cache", zap.String("key", k))
			stats.Record(ctx, d.cacheMissChapterCount.M(1))
		} else {
			log.For(ctx).Error("Unable to retrieve object from cache", zap.Error(err), zap.String("key", k))
			stats.Record(ctx, d.cacheErrorChapterCount.M(1))
			return d.next.Get(ctx, req)
		}
	}

	// Object value
	if payload != nil {
		// Decode payload
		var cached spotigraph.SingleChapterRes
		if err := proto.Unmarshal(payload, &cached); err != nil {
			log.For(ctx).Error("Unable to decode object payload from cache", zap.String("key", k), zap.Error(err))
			stats.Record(ctx, d.cacheErrorChapterCount.M(1))
			return d.next.Get(ctx, req)
		}

		// Return cached result
		stats.Record(ctx, d.cacheHitChapterCount.M(1))
		return &cached, nil
	}

	// Do the service call
	res, err = d.next.Get(ctx, req)

	// No cache if error
	if err != nil {
		return res, err
	}
	if res != nil && res.Error != nil {
		return res, err
	}

	// Encode result
	payload, err = proto.Marshal(res)
	if err != nil {
		log.For(ctx).Error("Unable to encode object payload to cache", zap.String("key", k), zap.Error(err))
		stats.Record(ctx, d.cacheErrorChapterCount.M(1))
		return res, err
	}

	// Set in cache
	if err = d.mgr.Set(ctx, k, payload, d.ttl); err != nil {
		log.For(ctx).Error("Unable to encode object payload to cache", zap.String("key", k), zap.Error(err))
		stats.Record(ctx, d.cacheErrorChapterCount.M(1))
		return res, err
	}

	// Return result
	return res, err
}

// Search implements services.Chapter
func (d chapterWithCache) Search(ctx context.Context, req *spotigraph.ChapterSearchReq) (res *spotigraph.PaginatedChapterRes, err error) {
	return d.next.Search(ctx, req)
}

// Update implements services.Chapter
func (d chapterWithCache) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (res *spotigraph.SingleChapterRes, err error) {
	defer func() {
		log.CheckErrCtx(ctx, "Unable to remove chapter from cache", d.mgr.Remove(ctx, req.Id))
	}()
	return d.next.Update(ctx, req)
}

// -----------------------------------------------------------------------------

func (d *chapterWithCache) key(id string) string {
	return fmt.Sprintf("core:spotigraph:services:Chapter:%s", id)
}

func (d *chapterWithCache) enableViews() error {
	cacheHitChapterCountView := &view.View{
		Name:        "chapter_cache_hit_count",
		Description: "Count of chapters cache hits",
		Measure:     d.cacheHitChapterCount,
		Aggregation: view.Count(),
	}

	cacheMissChapterCountView := &view.View{
		Name:        "chapter_cache_miss_count",
		Description: "Count of chapters cache miss",
		Measure:     d.cacheMissChapterCount,
		Aggregation: view.Count(),
	}

	cacheErrorChapterCountView := &view.View{
		Name:        "chapter_cache_error_count",
		Description: "Count of chapters cache error",
		Measure:     d.cacheErrorChapterCount,
		Aggregation: view.Count(),
	}

	// Register all views
	return view.Register(
		cacheHitChapterCountView,
		cacheMissChapterCountView,
		cacheErrorChapterCountView,
	)
}
