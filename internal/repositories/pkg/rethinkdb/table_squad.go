package rethinkdb

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	rdb "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type rdbSquadRepository struct {
	adapter *db.Default
}

// NewSquadRepository returns an initialized RethinkDB repository for squads
func NewSquadRepository(cfg *db.Configuration, session *rdb.Session) repositories.Squad {
	return &rdbSquadRepository{
		adapter: db.NewCRUDTable(session, cfg.Database, SquadTableName),
	}
}

// ------------------------------------------------------------

func (r *rdbSquadRepository) Create(ctx context.Context, entity *models.Squad) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *rdbSquadRepository) Get(ctx context.Context, id string) (*models.Squad, error) {
	var entity models.Squad

	// Do the query
	err := r.adapter.FindOneBy(ctx, "id", id, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}

func (r *rdbSquadRepository) Update(ctx context.Context, entity *models.Squad) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, entity.ID, map[string]interface{}{
		"name": entity.Name,
	})
}

func (r *rdbSquadRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbSquadRepository) Search(ctx context.Context, filter *repositories.SquadSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Squad, int, error) {
	panic("Not implemented")
}

func (r *rdbSquadRepository) FindByName(ctx context.Context, name string) (*models.Squad, error) {
	var entity models.Squad

	// Do the query
	err := r.adapter.FindOneBy(ctx, "name", name, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}
