package chapter

import (
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/decorators"
)

// Decorator represents service decorator builder
type Decorator func(s services.Chapter) services.Chapter

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(chapters repositories.Chapter, users repositories.User, memberships repositories.Membership, dcrs ...Decorator) services.Chapter {
	// Initialize base
	s := New(chapters, users, memberships)

	// Add decorators
	for _, wrapper := range dcrs {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// -----------------------------------------------------------------------------

// WithLogger initialize the chapter service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return decorators.NewChapterWithLogger(
			s,
			factory.With(zap.String("service", "spotigraph.services.Chapter")),
		)
	}
}

// WithTracer initialize the chapter service tracer decorator
func WithTracer() Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return decorators.NewChapterWithOpenCensus(
			s,
			"spotigraph",
		)
	}
}
