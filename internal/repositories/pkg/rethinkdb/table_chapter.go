package rethinkdb

import (
	"context"

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
		"name": entity.Name,
	})
}

func (r *rdbChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbChapterRepository) Search(ctx context.Context, filter *repositories.ChapterSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Chapter, int, error) {
	panic("Not implemented")
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
