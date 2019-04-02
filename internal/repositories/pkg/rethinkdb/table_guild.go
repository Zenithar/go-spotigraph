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
		"name":       entity.Name,
		"member_ids": entity.Members,
		"meta":       entity.Meta,
	})
}

func (r *rdbGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	var results []*models.Guild

	// Build filter
	filterFunc := func(row rdb.Term) rdb.Term {
		var term = rdb.Expr(true)

		// Guild ID
		if len(strings.TrimSpace(filter.GuildID)) > 0 {
			term = term.And(row.Field("id").Eq(filter.GuildID))
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
