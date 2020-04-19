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

type pgGuildRepository struct {
	adapter *db.Default
}

// NewGuildRepository returns an initialized PostgreSQL repository for guilds
func NewGuildRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Guild {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label", "meta", "leader_id",
	}

	// Sortable columns
	sortableColumns := []string{
		"label",
	}

	return &pgGuildRepository{
		adapter: db.NewCRUDTable(session, "", GuildTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlGuild struct {
	ID       string `db:"id"`
	Label    string `db:"label"`
	Meta     string `db:"meta"`
	LeaderID string `db:"leader_id"`
}

func toGuildSQL(entity *models.Guild) (*sqlGuild, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlGuild{
		ID:       entity.ID,
		Label:    entity.Label,
		Meta:     string(meta),
		LeaderID: entity.LeaderID,
	}, nil
}

func (dto *sqlGuild) ToEntity() (*models.Guild, error) {
	entity := &models.Guild{
		ID:       dto.ID,
		Label:    dto.Label,
		LeaderID: dto.LeaderID,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgGuildRepository) Create(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}
	// Convert to DTO
	data, err := toGuildSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgGuildRepository) Get(ctx context.Context, id string) (*models.Guild, error) {
	var entity sqlGuild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgGuildRepository) Update(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toGuildSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"label":     obj.Label,
		"meta":      obj.Meta,
		"leader_id": obj.LeaderID,
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
	var results []sqlGuild

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Guild, len(results))
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

func (r *pgGuildRepository) FindByLabel(ctx context.Context, label string) (*models.Guild, error) {
	var entity sqlGuild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"label": label,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgGuildRepository) buildFilter(filter *repositories.GuildSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.GuildID)) > 0 {
			clauses["guild_id"] = filter.GuildID
		}
		if len(strings.TrimSpace(filter.Label)) > 0 {
			clauses["label"] = filter.Label
		}

		return clauses
	}

	return nil
}
