package postgresql

import (
	"context"
	"encoding/json"
	"strings"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type pgChapterRepository struct {
	adapter *db.Default
}

// NewChapterRepository returns an initialized PostgreSQL repository for chapters
func NewChapterRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Chapter {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "name", "meta", "leader_id", "member_ids",
	}

	// Sortable columns
	sortableColumns := []string{
		"name", "leader_id",
	}

	return &pgChapterRepository{
		adapter: db.NewCRUDTable(session, "", ChapterTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlChapter struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Meta    string `db:"meta"`
	Leader  string `db:"leader_id"`
	Members string `db:"member_ids"`
}

func toChapterSQL(entity *models.Chapter) (*sqlChapter, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := json.Marshal(entity.Members)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlChapter{
		ID:      entity.ID,
		Name:    entity.Name,
		Meta:    string(meta),
		Leader:  entity.Leader,
		Members: string(members),
	}, nil
}

func (dto *sqlChapter) ToEntity() (*models.Chapter, error) {
	entity := &models.Chapter{
		ID:     dto.ID,
		Name:   dto.Name,
		Leader: dto.Leader,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Membership
	err = json.Unmarshal([]byte(dto.Members), &entity.Members)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgChapterRepository) Create(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toChapterSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgChapterRepository) Get(ctx context.Context, id string) (*models.Chapter, error) {
	var entity sqlChapter

	if err := r.adapter.WhereAndFetchOne(ctx, sq.Eq{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgChapterRepository) Update(ctx context.Context, entity *models.Chapter) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toChapterSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":       obj.Name,
		"meta":       obj.Meta,
		"leader_id":  obj.Leader,
		"member_ids": obj.Members,
	}, sq.Eq{
		"id": entity.ID,
	})
}

func (r *pgChapterRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, sq.Eq{
		"id": id,
	})
}

func (r *pgChapterRepository) Search(ctx context.Context, filter *repositories.ChapterSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Chapter, int, error) {
	var results []*models.Chapter

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	if len(results) == 0 {
		return results, count, api.ErrNoResult
	}

	// Return results and total count
	return results, count, nil
}

func (r *pgChapterRepository) FindByName(ctx context.Context, name string) (*models.Chapter, error) {
	var entity sqlChapter

	if err := r.adapter.WhereAndFetchOne(ctx, sq.Eq{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgChapterRepository) buildFilter(filter *repositories.ChapterSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.ChapterID)) > 0 {
			clauses["chapter_id"] = filter.ChapterID
		}
		if len(strings.TrimSpace(filter.Name)) > 0 {
			clauses["name"] = filter.Name
		}

		return clauses
	}

	return nil
}
