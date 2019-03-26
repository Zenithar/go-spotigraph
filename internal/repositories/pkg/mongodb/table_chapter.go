package mongodb

import (
	"context"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type mgoChapterRepository struct {
	adapter *db.Default
}

// NewChapterRepository returns an initialized MongoDB repository for chapters
func NewChapterRepository(cfg *db.Configuration, session *mongo.Client) repositories.Chapter {
	return &mgoChapterRepository{
		adapter: db.NewCRUDTable(session, cfg.DatabaseName, ChapterTableName),
	}
}

// ------------------------------------------------------------

func (r *mgoChapterRepository) Create(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *mgoChapterRepository) Get(ctx context.Context, id string) (*models.Chapter, error) {
	var entity models.Chapter

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoChapterRepository) Update(ctx context.Context, entity *models.Chapter) error {
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

func (r *mgoChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *mgoChapterRepository) Search(ctx context.Context, filter *repositories.ChapterSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Chapter, int, error) {
	panic("Not implemented")
}

func (r *mgoChapterRepository) FindByName(ctx context.Context, name string) (*models.Chapter, error) {
	var entity models.Chapter

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
