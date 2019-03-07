package rethinkdb

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.uber.org/zap"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
	"go.zenithar.org/pkg/log"
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
		log.For(ctx).Error("Unable to query database", zap.Error(err))
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

	return r.adapter.UpdateID(ctx, entity.ID, entity)
}

func (r *rdbUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}
