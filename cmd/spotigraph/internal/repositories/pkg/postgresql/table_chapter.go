// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgresql

import (
	"context"
	"encoding/json"
	"strings"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories"

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
		"id", "label", "meta", "leader_id",
	}

	// Sortable columns
	sortableColumns := []string{
		"label", "leader_id",
	}

	return &pgChapterRepository{
		adapter: db.NewCRUDTable(session, "", ChapterTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlChapter struct {
	ID       string `db:"id"`
	Label    string `db:"label"`
	Meta     string `db:"meta"`
	LeaderID string `db:"leader_id"`
}

func toChapterSQL(entity *models.Chapter) (*sqlChapter, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlChapter{
		ID:       entity.ID,
		Label:    entity.Label,
		Meta:     string(meta),
		LeaderID: entity.LeaderID,
	}, nil
}

func (dto *sqlChapter) ToEntity() (*models.Chapter, error) {
	entity := &models.Chapter{
		ID:       dto.ID,
		Label:    dto.Label,
		LeaderID: dto.LeaderID,
	}

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
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
		"label":     obj.Label,
		"meta":      obj.Meta,
		"leader_id": obj.LeaderID,
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
	var results []sqlChapter

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Chapter, len(results))
	if len(results) == 0 {
		return entities, count, api.ErrNoResult
	}

	for i, entity := range results {
		e, err := entity.ToEntity()
		if err != nil {
			continue
		}
		entities[i] = e
	}

	// Return results and total count
	return entities, count, nil
}

func (r *pgChapterRepository) FindByLabel(ctx context.Context, label string) (*models.Chapter, error) {
	var entity sqlChapter

	if err := r.adapter.WhereAndFetchOne(ctx, sq.Eq{
		"label": label,
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
		if len(strings.TrimSpace(filter.Label)) > 0 {
			clauses["label"] = filter.Label
		}

		return clauses
	}

	return nil
}
