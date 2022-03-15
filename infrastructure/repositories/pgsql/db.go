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

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// -----------------------------------------------------------------------------

type DBConn interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) pgx.Row
}

// -----------------------------------------------------------------------------

type PreparedStatementName string

const (
	// DefaultPageSize defines the default page size for pagination.
	DefaultPageSize = 50

	// Person
	dbPersonGetByID        = "dbPersonGetByID"
	dbPersonGetByObjectID  = "dbPersonGetByObjectID"
	dbPersonGetByPrincipal = "dbPersonGetByPrincipal"
	dbPersonCreate         = "dbPersonCreate"
	dbPersonUpdate         = "dbPersonUpdate"
	dbPersonDelete         = "dbPersonDelete"
)

// MakePool create a postgres database pool per configuration.
func MakePool(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	poolCfg.AfterConnect = onConnect
	poolCfg.MaxConns = int32(4)

	level, err := pgx.LogLevelFromString("warn")
	if err != nil {
		level = pgx.LogLevelWarn
	}

	//poolCfg.ConnConfig.Logger = zerologadapter.NewLogger(log.Logger)
	poolCfg.ConnConfig.LogLevel = level

	return pgxpool.ConnectConfig(ctx, poolCfg)
}

// NewQueryBuilder returns a new query builder
func NewQueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type QueryBuilder interface {
	ToSql() (string, []any, error)
}

// onConnect is executed after connection pool creation.
func onConnect(ctx context.Context, conn *pgx.Conn) error {
	// Check arguments
	if conn == nil {
		return fmt.Errorf("unable to prepare state with a nil connection")
	}

	if err := prepareQueries(ctx, conn); err != nil {
		return fmt.Errorf("unable to prepare core queries: %w", err)
	}

	// No error
	return nil
}

func prepareQueries(ctx context.Context, conn *pgx.Conn) error {
	// Check argument
	if conn == nil {
		return errors.New("unable to prepare queries with nil connection")
	}

	// Query collection
	queries := map[string]string{
		// Person --------------------------------------------------------------
		dbPersonGetByID:        "SELECT person_id, person_oid, principal, locked, created_at FROM persons WHERE person_id = $1;",
		dbPersonGetByObjectID:  "SELECT person_id, person_oid, principal, locked, created_at FROM persons WHERE person_oid = $1;",
		dbPersonGetByPrincipal: "SELECT person_id, person_oid, principal, locked, created_at FROM persons WHERE principal = $1;",
		dbPersonCreate:         "INSERT INTO persons (person_oid, principal) VALUES ($1, $2);",
		dbPersonUpdate:         "SELECT principal FROM persons WHERE person_id = $1 FOR UPDATE; UPDATE persons SET principal = $2 WHERE person_id = $1;",
		dbPersonDelete:         "DELETE FROM persons WHERE person_id = $1;",
	}

	// Prepare all queries
	for name, stmt := range queries {
		if _, err := conn.Prepare(ctx, name, stmt); err != nil {
			return fmt.Errorf("unable to prepare '%s' query: %w", name, err)
		}
	}

	// No error
	return nil
}
