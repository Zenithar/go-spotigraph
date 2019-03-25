package rethinkdb

import (
	"context"

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
		"prn": entity.Principal,
	})
}

func (r *rdbUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *rdbUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	panic("Not implemented")
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
