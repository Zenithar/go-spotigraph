package chapter

import (
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
)

// Decorator represents service decorator builder
type Decorator func(s services.Chapter) services.Chapter

// NewWithDecorators returns a service instance with decorators
func NewWithDecorators(chapters repositories.Chapter, dcrs ...Decorator) services.Chapter {
	// Initialize base
	s := New(chapters)

	// Add decorators
	for _, wrapper := range dcrs {
		s = wrapper(s)
	}

	// Return decorated service
	return s
}
