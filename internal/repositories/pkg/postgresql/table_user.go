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

type pgUserRepository struct {
	adapter *db.Default
}

// NewUserRepository returns an initialized PostgreSQL repository for users
func NewUserRepository(cfg *db.Configuration, session *sqlx.DB) repositories.User {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "principal", "meta",
	}

	// Sortable columns
	sortableColumns := []string{
		"principal",
	}

	return &pgUserRepository{
		adapter: db.NewCRUDTable(session, "", UserTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlUser struct {
	ID        string `db:"id"`
	Principal string `db:"principal"`
	Meta      string `db:"meta"`
}

func toUserSQL(entity *models.User) (*sqlUser, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlUser{
		ID:        entity.ID,
		Principal: entity.Principal,
		Meta:      string(meta),
	}, nil
}

func (dto *sqlUser) ToEntity() (*models.User, error) {
	entity := &models.User{
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

func (r *pgUserRepository) Create(ctx context.Context, entity *models.User) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.Create")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toUserSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.Get")
	defer span.End()

	var entity sqlUser

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgUserRepository) Update(ctx context.Context, entity *models.User) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.Update")
	defer span.End()

	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toUserSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"meta": obj.Meta,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgUserRepository) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.Delete")
	defer span.End()

	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.Search")
	defer span.End()

	var results []sqlUser

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.User, len(results))
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

func (r *pgUserRepository) FindByPrincipal(ctx context.Context, principal string) (*models.User, error) {
	ctx, span := trace.StartSpan(ctx, "postgresql.user.FindByPrincipal")
	defer span.End()

	var entity sqlUser

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"principal": principal,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgUserRepository) buildFilter(filter *repositories.UserSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.UserID)) > 0 {
			clauses["id"] = filter.UserID
		}
		if len(strings.TrimSpace(filter.Principal)) > 0 {
			clauses["principal"] = filter.Principal
		}

		return clauses
	}

	return nil
}
