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

type pgPersonRepository struct {
	adapter *db.Default
}

// NewPersonRepository returns an initialized PostgreSQL repository for persons
func NewPersonRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Person {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "principal", "meta",
	}

	// Sortable columns
	sortableColumns := []string{
		"principal",
	}

	return &pgPersonRepository{
		adapter: db.NewCRUDTable(session, "", PersonTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlPerson struct {
	ID        string `db:"id"`
	Principal string `db:"principal"`
	Meta      string `db:"meta"`
}

func toPersonSQL(entity *models.Person) (*sqlPerson, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlPerson{
		ID:        entity.ID,
		Principal: entity.Principal,
		Meta:      string(meta),
	}, nil
}

func (dto *sqlPerson) ToEntity() (*models.Person, error) {
	entity := &models.Person{
		ID:        dto.ID,
		Principal: dto.Principal,
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

func (r *pgPersonRepository) Create(ctx context.Context, entity *models.Person) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.Create")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toPersonSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgPersonRepository) Get(ctx context.Context, id string) (*models.Person, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.Get")
	defer span.End()

	var entity sqlPerson

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgPersonRepository) Update(ctx context.Context, entity *models.Person) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.Update")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toPersonSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"meta": obj.Meta,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgPersonRepository) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.Delete")
	defer span.End()

	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgPersonRepository) Search(ctx context.Context, filter *repositories.PersonSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Person, int, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.Search")
	defer span.End()

	var results []sqlPerson

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Person, len(results))
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

func (r *pgPersonRepository) FindByPrincipal(ctx context.Context, principal string) (*models.Person, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.person.FindByPrincipal")
	defer span.End()

	var entity sqlPerson

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"principal": principal,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgPersonRepository) buildFilter(filter *repositories.PersonSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.PersonID)) > 0 {
			clauses["id"] = filter.PersonID
		}
		if len(strings.TrimSpace(filter.Principal)) > 0 {
			clauses["principal"] = filter.Principal
		}

		return clauses
	}

	return nil
}
