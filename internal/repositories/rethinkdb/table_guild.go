package rethinkdb

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.uber.org/zap"
	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	"go.zenithar.org/pkg/log"
	rdb "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type rdbGuildRepository struct {
	adapter *db.Default
}

// NewGuildRepository returns an initialized RethinkDB repository for guilds
func NewGuildRepository(cfg *db.Configuration, session *rdb.Session) repositories.Guild {
	return &rdbGuildRepository{
		adapter: db.NewCRUDTable(session, cfg.Database, GuildTableName),
	}
}

// ------------------------------------------------------------

func (r *rdbGuildRepository) Create(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *rdbGuildRepository) Get(ctx context.Context, id string) (*models.Guild, error) {
	var entity models.Guild

	// Do the query
	err := r.adapter.FindOneBy(ctx, "id", id, &entity)
	if err != nil {
		log.For(ctx).Error("Unable to query database", zap.Error(err))
		return nil, err
	}

	// Return result
	return &entity, nil
}

func (r *rdbGuildRepository) Update(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.UpdateID(ctx, entity.ID, entity)
}

func (r *rdbGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	panic("Not implemented")
}
