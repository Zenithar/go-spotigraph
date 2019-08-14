package postgresql

import (
	"context"

	"github.com/gobuffalo/packr"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

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
	// UserTableName represents users collection name
	UserTableName = "users"
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
	NewUserRepository,
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

		_, err := CreateSchemas(conn)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to migrate database schema")
		}
	}

	// No error
	return conn, nil
}
