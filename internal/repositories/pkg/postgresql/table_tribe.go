package postgresql

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgTribeRepository struct {
	adapter *db.Default
}

// NewTribeRepository returns an initialized PostgreSQL repository for tribes
func NewTribeRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Tribe {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label",
	}

	// Sortable columns
	sortableColumns := []string{
		"name",
	}

	return &pgTribeRepository{
		adapter: db.NewCRUDTable(session, "", TribeTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

func (r *pgTribeRepository) Create(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Create(ctx, entity)
}

func (r *pgTribeRepository) Get(ctx context.Context, id string) (*models.Tribe, error) {
	var entity models.Tribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *pgTribeRepository) Update(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name": entity.Name,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	panic("Not implemented")
}

func (r *pgTribeRepository) FindByName(ctx context.Context, name string) (*models.Tribe, error) {
	var entity models.Tribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
