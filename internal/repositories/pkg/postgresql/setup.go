package postgresql

import (
	"github.com/google/wire"

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
