package publisher

import (
	"context"

	eventsv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/events/v1"
)

//go:generate mockgen -destination mock/publisher.gen.go -package mock go.zenithar.org/spotigraph/internal/reactor/internal/publisher Publisher

// Publisher decribes event publisher contract.
type Publisher interface {
	Publish(ctx context.Context, event *eventsv1.Event) error
}
