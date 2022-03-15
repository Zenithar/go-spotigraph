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

	// Used to embed schema.sql into the sql variable.
	"context"
	"database/sql"
	"embed"

	// Support natively embedded files
	"errors"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Imported for migration loader features
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	// Database driver
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

//go:embed migrations/*.sql
var fs embed.FS

//go:embed data/seed.sql
var seedDoc string

// Migrate database schema
func Migrate(connString string, log *zerolog.Logger) error {
	// Check arguments
	if log == nil {
		return errors.New("unable to process with nil logger")
	}

	log.Info().Msg("==== Setting up connection for Migrations ====")
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Error().Str("uri", connString).Err(err).Msg("Fail open db")
		return fmt.Errorf("unable to connecto to database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Fail open migration driver")
		return fmt.Errorf("unable to prepare migration driver: %w", err)
	}
	defer driver.Close()

	log.Info().Msg("==== Running Migrations ====")

	log.Info().Msg("Using embedded migrations ...")

	// Use embedded migrations
	d, errFs := iofs.New(fs, "migrations")
	if errFs != nil {
		log.Error().Err(err).Msg("unable to initialize embedded filesystem")
		return fmt.Errorf("unable to initialize embedded filesystem: %w", errFs)
	}

	// Use embedded migrations
	m, errMigrate := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if errMigrate != nil {
		log.Error().Err(errMigrate).Msg("Fail create migration instance")
		return fmt.Errorf("unable to initialize migration engine: %w", errMigrate)
	}
	defer m.Close()

	// Apply migrations
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msg("==== No Migration Changes ====")
			// Return no error
			return nil
		}

		return fmt.Errorf("unable to apply migration(s): %w", err)
	}

	log.Info().Msg("==== Finished Migrations ====")

	// No error
	return nil
}

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(ctx context.Context, conn DBConn) error {
	// Start a transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	// Execute seed file
	if _, errExec := tx.Exec(ctx, seedDoc); errExec != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return errRollback
		}
		return errExec
	}

	// Commit transaction
	return tx.Commit(ctx)
}
