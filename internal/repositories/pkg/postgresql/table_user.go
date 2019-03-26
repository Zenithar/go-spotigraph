package postgresql

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgUserRepository struct {
	adapter *db.Default
}

// NewUserRepository returns an initialized PostgreSQL repository for users
func NewUserRepository(cfg *db.Configuration, session *sqlx.DB) repositories.User {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "prn", "meta",
	}

	// Sortable columns
	sortableColumns := []string{
		"prn",
	}

	return &pgUserRepository{
		adapter: db.NewCRUDTable(session, "", UserTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

func (r *pgUserRepository) Create(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Create(ctx, entity)
}

func (r *pgUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *pgUserRepository) Update(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"prn": entity.Principal,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	panic("Not implemented")
}

func (r *pgUserRepository) FindByPrincipal(ctx context.Context, principal string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"prn": principal,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
