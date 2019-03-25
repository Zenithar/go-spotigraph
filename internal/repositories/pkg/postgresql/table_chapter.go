package postgresql

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgChapterRepository struct {
	adapter *db.Default
}

// NewChapterRepository returns an initialized PostgreSQL repository for chapters
func NewChapterRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Chapter {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label",
	}

	// Sortable columns
	sortableColumns := []string{
		"name",
	}

	return &pgChapterRepository{
		adapter: db.NewCRUDTable(session, "", ChapterTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

func (r *pgChapterRepository) Create(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Create(ctx, entity)
}

func (r *pgChapterRepository) Get(ctx context.Context, id string) (*models.Chapter, error) {
	var entity models.Chapter

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *pgChapterRepository) Update(ctx context.Context, entity *models.Chapter) error {
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

func (r *pgChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgChapterRepository) Search(ctx context.Context, filter *repositories.ChapterSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Chapter, int, error) {
	panic("Not implemented")
}

func (r *pgChapterRepository) FindByName(ctx context.Context, name string) (*models.Chapter, error) {
	var entity models.Chapter

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
