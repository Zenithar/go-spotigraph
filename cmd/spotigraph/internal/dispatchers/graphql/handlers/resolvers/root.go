package resolvers

import (
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/graphql/handlers/generated"
	"go.zenithar.org/spotigraph/internal/services"
)

// Resolver represents GraphQL RootResolver implementation
type resolver struct {
	users    services.User
	squads   services.Squad
	chapters services.Chapter
	guilds   services.Guild
	tribes   services.Tribe
}

// NewResolver returns a GraphQL root resolver implementation
func NewResolver(users services.User, squads services.Squad, chapters services.Chapter, guilds services.Guild, tribes services.Tribe) generated.Config {
	return generated.Config{
		Resolvers: &resolver{
			users:    users,
			squads:   squads,
			chapters: chapters,
			guilds:   guilds,
			tribes:   tribes,
		},
	}
}

// -----------------------------------------------------------------------------

func (r *resolver) Chapter() generated.ChapterResolver {
	return &chapterResolver{r}
}

func (r *resolver) Error() generated.ErrorResolver {
	return &errorResolver{r}
}

func (r *resolver) Guild() generated.GuildResolver {
	return &guildResolver{r}
}

func (r *resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{
		root:     r,
		users:    r.users,
		squads:   r.squads,
		chapters: r.chapters,
		guilds:   r.guilds,
		tribes:   r.tribes,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return &queryResolver{
		root:     r,
		users:    r.users,
		squads:   r.squads,
		chapters: r.chapters,
		guilds:   r.guilds,
		tribes:   r.tribes,
	}
}

func (r *resolver) Squad() generated.SquadResolver {
	return &squadResolver{r}
}

func (r *resolver) Tribe() generated.TribeResolver {
	return &tribeResolver{r}
}

func (r *resolver) User() generated.UserResolver {
	return &userResolver{r}
}

// -----------------------------------------------------------------------------
