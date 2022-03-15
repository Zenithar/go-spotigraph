// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pgsql

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/rs/xid"
	"github.com/rs/zerolog"

	"zntr.io/spotigraph/domain/person"
	"zntr.io/spotigraph/pkg/cursor"
	"zntr.io/spotigraph/pkg/types"
)

func Persons(conn DBConn, log *zerolog.Logger) person.Repository {
	return &pgPersonRepository{
		conn: conn,
		log:  log,
	}
}

// -----------------------------------------------------------------------------

type sqlPerson struct {
	ID        string
	Principal string
	Locked    bool
	CreatedAt time.Time
}

var _ person.Person = (*sqlPerson)(nil)

func (s *sqlPerson) GetID() person.ID     { return person.ID(s.ID) }
func (s *sqlPerson) GetPrincipal() string { return s.Principal }
func (s *sqlPerson) IsLocked() bool       { return s.Locked }

// persionListCursor -----------------------------------------------------------

type personListCursor struct {
	_         struct{}  `cbor:",toarray"`
	ID        person.ID `cbor:"1,keyasint,omitempty"`
	Principal *string   `cbor:"2,keyasint,omitempty"`
}

// -----------------------------------------------------------------------------

type pgPersonRepository struct {
	conn DBConn
	log  *zerolog.Logger
}

// -----------------------------------------------------------------------------

func (r *pgPersonRepository) NextID(ctx context.Context) (person.ID, error) {
	return person.ID(xid.New().String()), nil
}

// -----------------------------------------------------------------------------

func (r *pgPersonRepository) List(ctx context.Context, filter person.SearchFilter) ([]person.Person, *cursor.PageInfo, error) {
	collection := []person.Person{}

	// Apply limit boundary
	pageSize := uint64(DefaultPageSize + 1)
	if filter.Limit != nil {
		if *filter.Limit > 0 && *filter.Limit <= DefaultPageSize {
			// Limit N+1 for cursor computation.
			pageSize = *filter.Limit + 1
		}
	}

	// Prepare query
	queryFunc := func() (string, []interface{}, error) {
		// Prepare query
		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		q := psql.Select(
			"person_id",
			"principal",
			"locked", "created_at",
		).From("persons").Limit(pageSize)

		// Filters -------------------------------------------------------------

		// Filter by objectIDs
		if len(filter.ObjectIDs) > 0 {
			q = q.Where("person_id = ANY(?)", pq.Array(filter.ObjectIDs))
		}

		// Filter by principal
		if filter.Principal != nil {
			q = q.Where(sq.Like{"principal": fmt.Sprint("%", *filter.Principal, "%")})
		}

		// Controls ------------------------------------------------------------

		// Process cursor
		var cursor personListCursor
		if filter.Cursor != nil {
			if err := decodeCursor(*filter.Cursor, &cursor); err != nil {
				r.log.Debug().Msgf("skipping invalid cursor processing '%s'", *filter.Cursor)
			}
		}

		// Default relative offset
		if cursor.ID != "" {
			q = q.Where("person_id <= ?", cursor.ID)
		}

		// Apply default order for pagination
		q = q.OrderBy("person_id DESC")

		// Override limit
		if pageSize > 0 {
			q = q.Limit(pageSize)
		}

		// Compile query
		return q.ToSql()
	}

	// Row unpacker
	rowUnpacker := func(rows pgx.Rows) (person.Person, error) {
		var entity sqlPerson

		// Extract from row
		if errRow := rows.Scan(
			&entity.ID,
			&entity.Principal,
			&entity.Locked, &entity.CreatedAt,
		); errRow != nil {
			return nil, fmt.Errorf("unable to unpack entity from database: %w", errRow)
		}

		// No error
		return &entity, nil
	}

	// Cursor builder
	cursorBuilder := func(last person.Person) personListCursor {
		return personListCursor{
			ID:        last.GetID(),
			Principal: types.StringRef(last.GetPrincipal()),
		}
	}

	// Deleagte to generic behavior implementation.
	return cursorPaginate(ctx, r.conn, r.log, collection, pageSize, queryFunc, rowUnpacker, cursorBuilder)
}

func (r *pgPersonRepository) GetByID(ctx context.Context, id person.ID) (person.Person, error) {
	// Check arguments
	if err := validation.Validate(id, validation.Required, is.PrintableASCII); err != nil {
		return nil, fmt.Errorf("invalid personID value: %w", err)
	}

	// Delegate to finder
	return findOne(ctx, r.conn, r.log, func(tx DBConn) (person.Person, error) {
		var entity sqlPerson

		if errRow := tx.QueryRow(ctx, dbPersonGetByID, id).Scan(
			&entity.ID,
			&entity.Principal,
			&entity.Locked, &entity.CreatedAt,
		); errRow != nil {
			return nil, fmt.Errorf("unable to unpack entity: %w", errRow)
		}

		// No error
		return &entity, nil
	})
}

func (r *pgPersonRepository) GetByPrincipal(ctx context.Context, principal string) (person.Person, error) {
	// Check arguments
	if err := validation.Validate(principal, validation.Required, is.PrintableASCII); err != nil {
		return nil, fmt.Errorf("invalid principal value: %w", err)
	}

	// Delegate to finder
	return findOne(ctx, r.conn, r.log, func(tx DBConn) (person.Person, error) {
		var entity sqlPerson

		if errRow := tx.QueryRow(ctx, dbPersonGetByPrincipal, principal).Scan(
			&entity.ID,
			&entity.Principal,
			&entity.Locked, &entity.CreatedAt,
		); errRow != nil {
			return nil, fmt.Errorf("unable to unpack entity: %w", errRow)
		}

		// No error
		return &entity, nil
	})
}

// -----------------------------------------------------------------------------

func (r *pgPersonRepository) Save(ctx context.Context, do person.Person) error {
	// Check arguments
	if types.IsNil(do) {
		return fmt.Errorf("unable to create with nil domain object")
	}

	// Convert models to entity
	persistable := &sqlPerson{
		ID:        string(do.GetID()),
		Principal: do.GetPrincipal(),
		Locked:    false,
		CreatedAt: time.Now().UTC(),
	}

	// Delegate to executor
	return queryExec(ctx, r.conn, r.log, persistable, func(tx DBConn) (pgconn.CommandTag, error) {
		return tx.Exec(ctx, dbPersonCreate,
			persistable.ID, persistable.Principal, persistable.Locked, persistable.CreatedAt,
		)
	})
}

func (r *pgPersonRepository) Remove(ctx context.Context, do person.Person) error {
	// Check arguments
	if types.IsNil(do) {
		return fmt.Errorf("unable to remove a nil domain object")
	}

	// Convert models to entity
	persistable := &sqlPerson{
		ID: string(do.GetID()),
	}

	// Delegate to executor
	return queryExec(ctx, r.conn, r.log, persistable, func(tx DBConn) (pgconn.CommandTag, error) {
		return tx.Exec(ctx, dbPersonDelete, persistable.ID)
	})
}
