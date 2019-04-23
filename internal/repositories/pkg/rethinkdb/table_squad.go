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
		"name":             entity.Name,
		"product_owner_id": entity.ProductOwnerID,
		"member_ids":       entity.MemberIDs,
		"meta":             entity.Meta,
	})
}

func (r *rdbSquadRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbSquadRepository) Search(ctx context.Context, filter *repositories.SquadSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Squad, int, error) {
	var results []*models.Squad

	// Build filter
	filterFunc := func(row rdb.Term) rdb.Term {
		var term = rdb.Expr(true)

		// Squad ID
		if len(strings.TrimSpace(filter.SquadID)) > 0 {
			term = term.And(row.Field("id").Eq(filter.SquadID))
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
