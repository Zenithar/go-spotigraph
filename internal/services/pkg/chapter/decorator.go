package chapter

import (
	"time"

	"go.uber.org/zap"

	"go.zenithar.org/pkg/cache"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal"
)

// Decorator represents service decorator builder
type Decorator func(s services.Chapter) services.Chapter

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(chapters repositories.Chapter, decorators ...Decorator) services.Chapter {
	// Initialize base
	s := New(chapters)

	// Add decorators
	for _, wrapper := range decorators {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}

// WithLogger initialize the chapter service logger decorator
func WithLogger(factory log.LoggerFactory) Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return internal.NewChapterWithLogger(
			s,
			factory.With(zap.String("service", "core.spotigraph.services.Chapter")),
		)
	}
}

// WithTracer initialize the chapter service tracer decorator
func WithTracer() Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return internal.NewChapterWithOpenCensus(
			s,
			"core.spotigraph",
		)
	}
}

// WithMetric initialize the chapter service metric decorator
func WithMetric() Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return internal.NewChapterWithMetrics(s)
	}
}

// WithCache initialize the chapter service cache decorator
func WithCache(storage cache.Storage, ttl time.Duration) Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return internal.NewChapterWithCache(s, storage, ttl)
	}
}

// WithBreaker initialize the chapter service circuit breaker decorator
func WithBreaker() Decorator {
	return func(s services.Chapter) services.Chapter {
		// Initialize the decorator
		return internal.NewChapterWithBreaker(s)
	}
}
