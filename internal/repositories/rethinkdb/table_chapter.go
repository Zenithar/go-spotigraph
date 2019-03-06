package rethinkdb

import (
	"context"

	"go.zenithar.org/spotimap/internal/models"
	"go.zenithar.org/spotimap/internal/repositories"

	"go.uber.org/zap"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	"go.zenithar.org/pkg/log"
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
		log.For(ctx).Error("Unable to query database", zap.Error(err))
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

	return r.adapter.UpdateID(ctx, entity.ID, entity)
}

func (r *rdbChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}
