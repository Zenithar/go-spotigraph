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
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"

	"zntr.io/spotigraph/infrastructure/repositories"
	"zntr.io/spotigraph/pkg/cursor"
	"zntr.io/spotigraph/pkg/types"
)

type Entity interface{}

func queryExec[E Entity](ctx context.Context, conn DBConn, log *zerolog.Logger, entity E, builder func(tx DBConn) (pgconn.CommandTag, error)) error {
	// Check arguments
	if types.IsNil(entity) {
		return errors.New("unable to process with nil entity")
	}

	// Prepate query
	queryFunc := func(tx DBConn) error {
		// Do the query
		cmdTag, errRow := builder(tx)
		if errRow != nil {
			return fmt.Errorf("unable to execute query: %w", errRow)
		}
		if cmdTag.RowsAffected() == 0 {
			return repositories.ErrNoChanges
		}

		// No error
		return nil
	}

	// Execute in transaction
	if txErr := withTransaction(ctx, conn, log, queryFunc); txErr != nil {
		return fmt.Errorf("unable to execute the query: %w", txErr)
	}

	// No error
	return nil
}

func findOne[E Entity](ctx context.Context, conn DBConn, log *zerolog.Logger, builder func(tx DBConn) (entity E, err error)) (E, error) {
	var entity E

	// Prepare query
	queryFunc := func(conn DBConn) error {
		var errRow error

		// Do the query
		entity, errRow = builder(conn)
		if errRow != nil {
			// If query returned no result.
			if errors.Is(errRow, pgx.ErrNoRows) {
				return repositories.ErrNoResult
			}

			return fmt.Errorf("unable to unpack entity from database: %w", errRow)
		}

		// No error
		return nil
	}

	// Execute query
	if qErr := queryFunc(conn); qErr != nil {
		return entity, fmt.Errorf("unable to retrieve entity: %w", qErr)
	}

	// No error
	return entity, nil
}

type sqlBuilderFunc func() (string, []interface{}, error)

func cursorPaginate[E Entity, C any](
	ctx context.Context, conn DBConn, log *zerolog.Logger,
	entities []E,
	pageSize uint64,
	builder sqlBuilderFunc,
	unpack func(rows pgx.Rows) (E, error),
	cursorBuilder func(last E) C,
) ([]E, *cursor.PageInfo, error) {
	var pi *cursor.PageInfo

	// Prepare query
	queryFunc := func(conn DBConn) error {
		// Compile sql query
		sql, params, err := builder()
		if err != nil {
			return fmt.Errorf("unable to compile sql statement: %w", err)
		}

		// Do the query
		rows, err := conn.Query(ctx, sql, params...)
		if err != nil {
			return fmt.Errorf("unable to query the database: %w", err)
		}
		defer rows.Close()

		// Unpack rows
		var (
			count = uint64(0)
			last  E
		)

		for rows.Next() {
			// Extract from row
			entity, errRow := unpack(rows)
			if errRow != nil {
				return fmt.Errorf("unable to unpack entity from database: %w", errRow)
			}
			count++

			// Get relative references
			if count < pageSize {
				entities = append(entities, entity)
			} else {
				last = entity
			}
		}

		// No more page
		if count < pageSize {
			return nil
		}

		// Use the last entity as reference
		cursorClaims := cursorBuilder(last)

		// Encode the cursor
		nextToken, err := encodeCursor(cursorClaims)
		if err != nil {
			return fmt.Errorf("unable to encode next_page token: %w", err)
		}

		// Assign to response
		pi = &cursor.PageInfo{
			NextPageToken: types.StringRef(nextToken),
		}

		// No error
		return nil
	}

	// Execute query
	if qErr := queryFunc(conn); qErr != nil {
		// If query returned no result.
		if errors.Is(qErr, pgx.ErrNoRows) {
			return []E{}, nil, nil
		}

		return nil, nil, fmt.Errorf("unable to list entities: %w", qErr)
	}

	// No error
	return entities, pi, nil
}
