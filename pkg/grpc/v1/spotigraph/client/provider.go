package client

import "github.com/google/wire"

// Services is a wire provider to wrap remote gRPC service as local interface
var Services = wire.NewSet(
	NewUserClient,
	NewSquadClient,
	NewChapterClient,
	NewGuildClient,
	NewTribeClient,
)
