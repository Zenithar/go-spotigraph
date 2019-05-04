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

type rdbUserRepository struct {
	adapter *db.Default
}

// NewUserRepository returns an initialized RethinkDB repository for users
func NewUserRepository(cfg *db.Configuration, session *rdb.Session) repositories.User {
	return &rdbUserRepository{
		adapter: db.NewCRUDTable(session, cfg.Database, UserTableName),
	}
}

// ------------------------------------------------------------

func (r *rdbUserRepository) Create(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *rdbUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var entity models.User

	// Do the query
	err := r.adapter.FindOneBy(ctx, "id", id, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}

func (r *rdbUserRepository) Update(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, entity.ID, map[string]interface{}{
		"principal": entity.Principal,
	})
}

func (r *rdbUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	var results []*models.User

	// Build filter
	filterFunc := func(row rdb.Term) rdb.Term {
		term := rdb.Expr(true)

		// User ID
		if len(strings.TrimSpace(filter.UserID)) > 0 {
			term = term.And(row.Field("id").Eq(filter.UserID))
		}

		// Name
		if len(strings.TrimSpace(filter.Principal)) > 0 {
			term = term.And(row.Field("principal").Eq(filter.Principal))
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

func (r *rdbUserRepository) FindByPrincipal(ctx context.Context, principal string) (*models.User, error) {
	var entity models.User

	// Do the query
	err := r.adapter.FindOneBy(ctx, "principal", principal, &entity)
	if err != nil {
		return nil, err
	}

	// Return result
	return &entity, nil
}
