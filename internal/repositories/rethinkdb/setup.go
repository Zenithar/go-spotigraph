package rethinkdb

import (
	"github.com/google/wire"
	db "go.zenithar.org/pkg/db/adapter/rethinkdb"
)

// ----------------------------------------------------------

var (
	// UserTableName represents users collection name
	UserTableName = "users"
)

// ----------------------------------------------------------

// RepositorySet exposes Google Wire providers
var RepositorySet = wire.NewSet(
	db.Connection,
	NewUserRepository,
)
