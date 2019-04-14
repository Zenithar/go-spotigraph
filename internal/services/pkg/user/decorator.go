package user

import (
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/user/internal"
)

// Decorator represents service decorator builder
type Decorator func(s services.User) services.User

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(users repositories.User, decorators ...Decorator) services.User {
	// Initialize base
	s := New(users)

	// Add decorators
	for _, wrapper := range decorators {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// WithLogger initialize the user service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.User) services.User {
		// Initialize the decorator
		return internal.NewUserWithLogger(
			s,
			factory.With(zap.String("service", "core.spotigraph.services.User")),
		)
	}
}

// WithTracer initialize the user service tracer decorator
func WithTracer() Decorator {
	return func(s services.User) services.User {
		// Initialize the decorator
		return internal.NewUserWithOpenCensus(
			s,
			"core.spotigraph",
		)
	}
}
