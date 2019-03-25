package rethinkdb

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
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

	return r.adapter.Update(ctx, entity.ID, map[string]interface{}{
		"name": entity.Name,
	})
}

func (r *rdbGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	panic("Not implemented")
}

func (r *rdbGuildRepository) FindByName(ctx context.Context, name string) (*models.Guild, error) {
	var entity models.Guild

	// Do the query
	err := r.adapter.FindOneBy(ctx, "name", name, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}
