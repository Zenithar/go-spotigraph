package postgresql

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgGuildRepository struct {
	adapter *db.Default
}

// NewGuildRepository returns an initialized PostgreSQL repository for guilds
func NewGuildRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Guild {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label", "meta",
	}

	// Sortable columns
	sortableColumns := []string{
		"name",
	}

	return &pgGuildRepository{
		adapter: db.NewCRUDTable(session, "", GuildTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

func (r *pgGuildRepository) Create(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Create(ctx, entity)
}

func (r *pgGuildRepository) Get(ctx context.Context, id string) (*models.Guild, error) {
	var entity models.Guild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *pgGuildRepository) Update(ctx context.Context, entity *models.Guild) error {
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

func (r *pgGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	panic("Not implemented")
}

func (r *pgGuildRepository) FindByName(ctx context.Context, name string) (*models.Guild, error) {
	var entity models.Guild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
