package guild

import (
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/decorators"
)

// Decorator represents service decorator builder
type Decorator func(s services.Guild) services.Guild

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(guilds repositories.Guild, persons repositories.Person, memberships repositories.Membership, dcrs ...Decorator) services.Guild {
	// Initialize base
	s := New(guilds, persons, memberships)

	// Add decorators
	for _, wrapper := range dcrs {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// -----------------------------------------------------------------------------

// WithLogger initialize the guild service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.Guild) services.Guild {
		// Initialize the decorator
		return decorators.NewGuildWithLogger(
			s,
			factory.With(zap.String("service", "spotigraph.services.Guild")),
		)
	}
}

// WithTracer initialize the guild service tracer decorator
func WithTracer() Decorator {
	return func(s services.Guild) services.Guild {
		// Initialize the decorator
		return decorators.NewGuildWithOpenCensus(
			s,
			"spotigraph",
		)
	}
}
