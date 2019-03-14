package server

import (
	"github.com/google/wire"

	"go.zenithar.org/spotigraph/internal/repositories/pkg/rethinkdb"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
)

var serviceSet = wire.Set(
	user.New,
	chapter.New,
	squad.New,
	guild.New,
	tribe.New,
)

// InitializeServerContext is used to prepare server context
func InitializeServerContext() error {
	wire.Build(
		rethinkdb.RepositorySet,
		serviceSet,
	)

	// Error will be set by injector
	return nil
}
