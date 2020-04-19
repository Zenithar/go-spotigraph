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

	"github.com/gobuffalo/packr"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	// Load postgresql drivers
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/pkg/log"
)

// ----------------------------------------------------------

var (
	// PersonTableName represents persons collection name
	PersonTableName = "persons"
	// ChapterTableName represents chapters collection name
	ChapterTableName = "chapters"
	// GuildTableName represents guilds collection name
	GuildTableName = "guilds"
	// SquadTableName represents squads collection name
	SquadTableName = "squads"
	// TribeTableName represents tribes collection name
	TribeTableName = "tribes"
	// MembershipTableName represents membership collection name
	MembershipTableName = "memberships"
)

// ----------------------------------------------------------

// RepositorySet exposes Google Wire providers
var RepositorySet = wire.NewSet(
	AutoMigrate,
	NewPersonRepository,
	NewChapterRepository,
	NewGuildRepository,
	NewSquadRepository,
	NewTribeRepository,
	NewMembershipRepository,
)

// ----------------------------------------------------------

//go:generate packr

// migrations contains all schema migrations
var migrations = &migrate.PackrMigrationSource{
	Box: packr.NewBox("./migrations"),
}

// CreateSchemas create or updates the current database schema
func CreateSchemas(conn *sqlx.DB) (int, error) {
	// Migrate schema
	migrate.SetTable("spfg_schema_migration")

	n, err := migrate.Exec(conn.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d migrations", n)
	}

	return n, nil
}

// AutoMigrate provider for auto schema migration feature
func AutoMigrate(ctx context.Context, cfg *db.Configuration) (*sqlx.DB, error) {
	// Initialize database connection
	conn, err := db.Connection(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.AutoMigrate {
		log.For(ctx).Info("Migrating database schema ...")

		n, err := CreateSchemas(conn)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to migrate database schema")
		}

		log.For(ctx).Info("Schema migrated", zap.Int("migrations", n))
	}

	// No error
	return conn, nil
}
