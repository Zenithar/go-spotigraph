package postgresql

import (
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
)

// ----------------------------------------------------------

// RepositorySet exposes Google Wire providers
var RepositorySet = wire.NewSet(
	db.Connection,
	NewUserRepository,
	NewChapterRepository,
	NewGuildRepository,
	NewSquadRepository,
	NewTribeRepository,
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
	migrate.SetTable("schema_migration")

	n, err := migrate.Exec(conn.DB, conn.DriverName(), migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d migrations", n)
	}

	return n, nil
}
