package cache

import (
	"context"
	"errors"
	"time"
)

// ErrCacheMiss is raised when item is not found in cache
var ErrCacheMiss = errors.New("cache: item not found")

//go:generate mockgen -destination mock/storage.gen.go -package mock go.zenithar.org/spotigraph/pkg/cache Storage

// Storage describes cache storage contract
type Storage interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, duration time.Duration) error
	Remove(ctx context.Context, key string) error
}
