package mongodb

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type mgoTribeRepository struct {
	adapter *db.Default
}

// NewTribeRepository returns an initialized MongoDB repository for tribes
func NewTribeRepository(cfg *db.Configuration, session *mongo.Client) repositories.Tribe {
	return &mgoTribeRepository{
		adapter: db.NewCRUDTable(session, cfg.DatabaseName, TribeTableName),
	}
}

// ------------------------------------------------------------

func (r *mgoTribeRepository) Create(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *mgoTribeRepository) Get(ctx context.Context, id string) (*models.Tribe, error) {
	var entity models.Tribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoTribeRepository) Update(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":      entity.Name,
		"meta":      entity.Meta,
		"squad_ids": entity.Squads,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *mgoTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *mgoTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	panic("Not implemented")
}

func (r *mgoTribeRepository) FindByName(ctx context.Context, name string) (*models.Tribe, error) {
	var entity models.Tribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
