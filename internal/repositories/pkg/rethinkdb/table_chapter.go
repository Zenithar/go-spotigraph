package rethinkdb

import (
	"context"
	"strings"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	rdb "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type rdbChapterRepository struct {
	adapter *db.Default
}

// NewChapterRepository returns an initialized RethinkDB repository for chapters
func NewChapterRepository(cfg *db.Configuration, session *rdb.Session) repositories.Chapter {
	return &rdbChapterRepository{
		adapter: db.NewCRUDTable(session, cfg.Database, ChapterTableName),
	}
}

// ------------------------------------------------------------

func (r *rdbChapterRepository) Create(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *rdbChapterRepository) Get(ctx context.Context, id string) (*models.Chapter, error) {
	var entity models.Chapter

	// Do the query
	err := r.adapter.FindOneBy(ctx, "id", id, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}

func (r *rdbChapterRepository) Update(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, entity.ID, map[string]interface{}{
		"name":       entity.Name,
		"leader_id":  entity.LeaderID,
		"member_ids": entity.MemberIDs,
		"meta":       entity.Meta,
	})
}

func (r *rdbChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbChapterRepository) Search(ctx context.Context, filter *repositories.ChapterSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Chapter, int, error) {
	var results []*models.Chapter

	// Build filter
	filterFunc := func(row rdb.Term) rdb.Term {
		var term = rdb.Expr(true)

		// Chapter ID
		if len(strings.TrimSpace(filter.ChapterID)) > 0 {
			term = term.And(row.Field("id").Eq(filter.ChapterID))
		}

		// Name
		if len(strings.TrimSpace(filter.Name)) > 0 {
			term = term.And(row.Field("name").Eq(filter.Name))
		}

		return term
	}

	// Run the count
	count, err := r.adapter.WhereCount(ctx, filterFunc)
	if err != nil {
		return nil, 0, err
	}

	// Run the query
	err = r.adapter.Search(ctx, &results, filterFunc, sortParams, pagination)
	if err != nil {
		return nil, 0, err
	}

	if len(results) == 0 {
		err = api.ErrNoResult
	} else {
		err = nil
	}

	return results, count, err
}

func (r *rdbChapterRepository) FindByName(ctx context.Context, name string) (*models.Chapter, error) {
	var entity models.Chapter

	// Do the query
	err := r.adapter.FindOneBy(ctx, "name", name, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}
