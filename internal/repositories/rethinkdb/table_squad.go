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
		log.For(ctx).Error("Unable to query database", zap.Error(err))
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

	return r.adapter.UpdateID(ctx, entity.ID, entity)
}

func (r *rdbSquadRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}
