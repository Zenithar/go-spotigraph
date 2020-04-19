// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package publisher

import (
	"context"

	eventsv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/events/v1"
)

//go:generate mockgen -destination mock/publisher.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/internal/publisher Publisher

// Publisher decribes event publisher contract.
type Publisher interface {
	Publish(ctx context.Context, event *eventsv1.Event) error
}
