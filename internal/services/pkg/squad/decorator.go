package squad

import (
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad/internal/decorators"
)

// Decorator represents service decorator builder
type Decorator func(s services.Squad) services.Squad

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(squads repositories.Squad, persons repositories.Person, memberships repositories.Membership, dcrs ...Decorator) services.Squad {
	// Initialize base
	s := New(squads, persons, memberships)

	// Add decorators
	for _, wrapper := range dcrs {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// -----------------------------------------------------------------------------

// WithLogger initialize the squad service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.Squad) services.Squad {
		// Initialize the decorator
		return decorators.NewSquadWithLogger(
			s,
			factory.With(zap.String("service", "spotigraph.services.Squad")),
		)
	}
}

// WithTracer initialize the squad service tracer decorator
func WithTracer() Decorator {
	return func(s services.Squad) services.Squad {
		// Initialize the decorator
		return decorators.NewSquadWithOpenCensus(
			s,
			"spotigraph",
		)
	}
}
