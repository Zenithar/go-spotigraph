package postgresql

import (
	"context"
	"encoding/json"
	"strings"

	"go.opencensus.io/trace"
	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

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
		"id", "name", "meta", "product_owner_id", "member_ids",
	}

	// Sortable columns
	sortableColumns := []string{
		"name", "product_owner_id",
	}

	return &pgSquadRepository{
		adapter: db.NewCRUDTable(session, "", SquadTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlSquad struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	Meta           string `db:"meta"`
	ProductOwnerID string `db:"product_owner_id"`
	MemberIDs      string `db:"member_ids"`
}

func toSquadSQL(entity *models.Squad) (*sqlSquad, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := json.Marshal(entity.MemberIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlSquad{
		ID:             entity.ID,
		Name:           entity.Name,
		Meta:           string(meta),
		MemberIDs:      string(members),
		ProductOwnerID: entity.ProductOwnerID,
	}, nil
}

func (dto *sqlSquad) ToEntity() (*models.Squad, error) {
	entity := &models.Squad{
		ID:             dto.ID,
		Name:           dto.Name,
		ProductOwnerID: dto.ProductOwnerID,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Membership
	err = json.Unmarshal([]byte(dto.MemberIDs), &entity.MemberIDs)
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
		"name":             obj.Name,
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

func (r *pgSquadRepository) FindByName(ctx context.Context, name string) (*models.Squad, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.squad.FindByName")
	defer span.End()

	var entity sqlSquad

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
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
		if len(strings.TrimSpace(filter.Name)) > 0 {
			clauses["name"] = filter.Name
		}

		return clauses
	}

	return nil
}
