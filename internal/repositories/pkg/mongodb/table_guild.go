package mongodb

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type mgoGuildRepository struct {
	adapter *db.Default
}

// NewGuildRepository returns an initialized MongoDB repository for guilds
func NewGuildRepository(cfg *db.Configuration, session *mongo.Client) repositories.Guild {
	return &mgoGuildRepository{
		adapter: db.NewCRUDTable(session, cfg.DatabaseName, GuildTableName),
	}
}

// ------------------------------------------------------------

func (r *mgoGuildRepository) Create(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *mgoGuildRepository) Get(ctx context.Context, id string) (*models.Guild, error) {
	var entity models.Guild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoGuildRepository) Update(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":       entity.Name,
		"meta":       entity.Meta,
		"member_ids": entity.Members,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *mgoGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *mgoGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	panic("Not implemented")
}

func (r *mgoGuildRepository) FindByName(ctx context.Context, name string) (*models.Guild, error) {
	var entity models.Guild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
