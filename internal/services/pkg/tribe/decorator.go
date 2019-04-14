package tribe

import (
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe/internal"
)

// Decorator represents service decorator builder
type Decorator func(s services.Tribe) services.Tribe

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(tribes repositories.Tribe, decorators ...Decorator) services.Tribe {
	// Initialize base
	s := New(tribes)

	// Add decorators
	for _, wrapper := range decorators {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// WithLogger initialize the tribe service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.Tribe) services.Tribe {
		// Initialize the decorator
		return internal.NewTribeWithLogger(
			s,
			factory.With(zap.String("service", "core.spotigraph.services.Tribe")),
		)
	}
}

// WithTracer initialize the tribe service tracer decorator
func WithTracer() Decorator {
	return func(s services.Tribe) services.Tribe {
		// Initialize the decorator
		return internal.NewTribeWithOpenCensus(
			s,
			"core.spotigraph",
		)
	}
}
