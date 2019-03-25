package rethinkdb

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	rdb "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type rdbTribeRepository struct {
	adapter *db.Default
}

// NewTribeRepository returns an initialized RethinkDB repository for tribes
func NewTribeRepository(cfg *db.Configuration, session *rdb.Session) repositories.Tribe {
	return &rdbTribeRepository{
		adapter: db.NewCRUDTable(session, cfg.Database, TribeTableName),
	}
}

// ------------------------------------------------------------

func (r *rdbTribeRepository) Create(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *rdbTribeRepository) Get(ctx context.Context, id string) (*models.Tribe, error) {
	var entity models.Tribe

	// Do the query
	err := r.adapter.FindOneBy(ctx, "id", id, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}

func (r *rdbTribeRepository) Update(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, entity.ID, map[string]interface{}{
		"name": entity.Name,
	})
}

func (r *rdbTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	panic("Not implemented")
}

func (r *rdbTribeRepository) FindByName(ctx context.Context, name string) (*models.Tribe, error) {
	var entity models.Tribe

	// Do the query
	err := r.adapter.FindOneBy(ctx, "name", name, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}
