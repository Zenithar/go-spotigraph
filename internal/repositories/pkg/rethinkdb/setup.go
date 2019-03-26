package rethinkdb

import (
	"github.com/google/wire"

	db "go.zenithar.org/pkg/db/adapter/rethinkdb"

	rdb "gopkg.in/rethinkdb/rethinkdb-go.v5"
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

// Setup configures the RethinkDB server
func Setup(opts rdb.ConnectOpts) error {

	// Initialize a new setup connection
	ss, err := rdb.Connect(opts)
	if err != nil {
		return err
	}

	// Create the database
	rdb.DBCreate(opts.Database).Exec(ss)

	// Create schema

	rdb.DB(opts.Database).TableCreate(UserTableName).Exec(ss)
	rdb.DB(opts.Database).Table(UserTableName).IndexCreate("prn").Exec(ss)

	rdb.DB(opts.Database).TableCreate(TribeTableName).Exec(ss)
	rdb.DB(opts.Database).Table(TribeTableName).IndexCreate("name").Exec(ss)

	rdb.DB(opts.Database).TableCreate(SquadTableName).Exec(ss)
	rdb.DB(opts.Database).Table(SquadTableName).IndexCreate("name").Exec(ss)

	rdb.DB(opts.Database).TableCreate(GuildTableName).Exec(ss)
	rdb.DB(opts.Database).Table(GuildTableName).IndexCreate("name").Exec(ss)

	rdb.DB(opts.Database).TableCreate(ChapterTableName).Exec(ss)
	rdb.DB(opts.Database).Table(ChapterTableName).IndexCreate("name").Exec(ss)

	// No Error
	return nil
}
