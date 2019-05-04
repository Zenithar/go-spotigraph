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
		"name":      entity.Name,
		"squad_ids": entity.SquadIDs,
		"meta":      entity.Meta,
		"leader_id": entity.LeaderID,
	})
}

func (r *rdbTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	var results []*models.Tribe

	// Build filter
	filterFunc := func(row rdb.Term) rdb.Term {
		term := rdb.Expr(true)

		// Tribe ID
		if len(strings.TrimSpace(filter.TribeID)) > 0 {
			term = term.And(row.Field("id").Eq(filter.TribeID))
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
