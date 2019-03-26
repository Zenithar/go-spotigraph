package mongodb

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type mgoUserRepository struct {
	adapter *db.Default
}

// NewUserRepository returns an initialized MongoDB repository for users
func NewUserRepository(cfg *db.Configuration, session *mongo.Client) repositories.User {
	return &mgoUserRepository{
		adapter: db.NewCRUDTable(session, cfg.DatabaseName, UserTableName),
	}
}

// ------------------------------------------------------------

func (r *mgoUserRepository) Create(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *mgoUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoUserRepository) Update(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name": entity.Principal,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *mgoUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *mgoUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	panic("Not implemented")
}

func (r *mgoUserRepository) FindByPrincipal(ctx context.Context, prn string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"prn": prn,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
