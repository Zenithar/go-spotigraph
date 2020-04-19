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

	"go.opencensus.io/trace"
	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type pgSquadRepository struct {
	adapter *db.Default
}

// NewSquadRepository returns an initialized PostgreSQL repository for squads
func NewSquadRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Squad {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label", "meta", "product_owner_id",
	}

	// Sortable columns
	sortableColumns := []string{
		"label", "product_owner_id",
	}

	return &pgSquadRepository{
		adapter: db.NewCRUDTable(session, "", SquadTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlSquad struct {
	ID             string `db:"id"`
	Label          string `db:"label"`
	Meta           string `db:"meta"`
	ProductOwnerID string `db:"product_owner_id"`
}

func toSquadSQL(entity *models.Squad) (*sqlSquad, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &sqlSquad{
		ID:             entity.ID,
		Label:          entity.Label,
		Meta:           string(meta),
		ProductOwnerID: entity.ProductOwnerID,
	}, nil
}

func (dto *sqlSquad) ToEntity() (*models.Squad, error) {
	entity := &models.Squad{
		ID:             dto.ID,
		Label:          dto.Label,
		ProductOwnerID: dto.ProductOwnerID,
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

func (r *pgSquadRepository) Create(ctx context.Context, entity *models.Squad) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.Create")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toSquadSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgSquadRepository) Get(ctx context.Context, id string) (*models.Squad, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.Get")
	defer span.End()

	var entity sqlSquad

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgSquadRepository) Update(ctx context.Context, entity *models.Squad) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.Update")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toSquadSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"label":            obj.Label,
		"meta":             obj.Meta,
		"product_owner_id": obj.ProductOwnerID,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgSquadRepository) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.Delete")
	defer span.End()

	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgSquadRepository) Search(ctx context.Context, filter *repositories.SquadSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Squad, int, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.Search")
	defer span.End()

	var results []sqlSquad

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Squad, len(results))
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

func (r *pgSquadRepository) FindByLabel(ctx context.Context, label string) (*models.Squad, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.FindByLabel")
	defer span.End()

	var entity sqlSquad

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"label": label,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgSquadRepository) buildFilter(filter *repositories.SquadSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.SquadID)) > 0 {
			clauses["id"] = filter.SquadID
		}
		if len(strings.TrimSpace(filter.Label)) > 0 {
			clauses["label"] = filter.Label
		}

		return clauses
	}

	return nil
}
