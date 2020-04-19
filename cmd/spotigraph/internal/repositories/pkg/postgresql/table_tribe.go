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

type pgTribeRepository struct {
	adapter *db.Default
}

// NewTribeRepository returns an initialized PostgreSQL repository for tribes
func NewTribeRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Tribe {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label", "meta", "squad_ids", "leader_id",
	}

	// Sortable columns
	sortableColumns := []string{
		"label",
	}

	return &pgTribeRepository{
		adapter: db.NewCRUDTable(session, "", TribeTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlTribe struct {
	ID       string `db:"id"`
	Label    string `db:"label"`
	Meta     string `db:"meta"`
	SquadIDs string `db:"squad_ids"`
	LeaderID string `db:"leader_id"`
}

func toTribeSQL(entity *models.Tribe) (*sqlTribe, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	squads, err := json.Marshal(entity.SquadIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlTribe{
		ID:       entity.ID,
		Label:    entity.Label,
		Meta:     string(meta),
		SquadIDs: string(squads),
	}, nil
}

func (dto *sqlTribe) ToEntity() (*models.Tribe, error) {
	entity := &models.Tribe{
		ID:    dto.ID,
		Label: dto.Label,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Membership
	err = json.Unmarshal([]byte(dto.SquadIDs), &entity.SquadIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgTribeRepository) Create(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toTribeSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgTribeRepository) Get(ctx context.Context, id string) (*models.Tribe, error) {
	var entity sqlTribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgTribeRepository) Update(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toTribeSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"label":     obj.Label,
		"meta":      obj.Meta,
		"squad_ids": obj.SquadIDs,
		"leader_id": obj.LeaderID,
	}, map[string]interface{}{
		"id": obj.ID,
	})
}

func (r *pgTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	var results []sqlTribe

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Tribe, len(results))
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

func (r *pgTribeRepository) FindByLabel(ctx context.Context, label string) (*models.Tribe, error) {
	var entity sqlTribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"label": label,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgTribeRepository) buildFilter(filter *repositories.TribeSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.TribeID)) > 0 {
			clauses["id"] = filter.TribeID
		}
		if len(strings.TrimSpace(filter.Label)) > 0 {
			clauses["label"] = filter.Label
		}

		return clauses
	}

	return nil
}
