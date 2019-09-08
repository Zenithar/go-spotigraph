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

package events

import (
	"github.com/dchest/uniuri"

	eventsv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/events/v1"
)

const (
	chapterAggregateType = "chapter"
)

// ChapterCreated is raised when a chapter entity has been saved
func ChapterCreated(urn, label, leaderID string) *eventsv1.Event {
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_CHAPTER_CREATED,
		EventId:       uniuri.NewLen(64),
		AggregateType: chapterAggregateType,
		AggregateId:   urn,
		Payload: &eventsv1.Event_ChapterCreated{
			ChapterCreated: &eventsv1.ChapterCreated{
				Urn:      urn,
				Label:    label,
				LeaderId: leaderID,
			},
		},
	}
}

// ChapterDeleted is raised when a chapter entity has been deleted
func ChapterDeleted(urn string) *eventsv1.Event {
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_CHAPTER_DELETED,
		EventId:       uniuri.NewLen(64),
		AggregateType: chapterAggregateType,
		AggregateId:   urn,
		Payload: &eventsv1.Event_ChapterDeleted{
			ChapterDeleted: &eventsv1.ChapterDeleted{
				Urn: urn,
			},
		},
	}
}

// ChapterLabelUpdated is raised when a chapter entity label has been updated
func ChapterLabelUpdated(urn, oldValue, newValue string) *eventsv1.Event {
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_CHAPTER_LABEL_UPDATED,
		EventId:       uniuri.NewLen(64),
		AggregateType: chapterAggregateType,
		AggregateId:   urn,
		Payload: &eventsv1.Event_ChapterLabelUpdated{
			ChapterLabelUpdated: &eventsv1.ChapterLabelUpdated{
				Urn: urn,
				Old: oldValue,
				New: newValue,
			},
		},
	}
}

// ChapterLeaderUpdated is raised when a chapter entity leader has been updated
func ChapterLeaderUpdated(urn, oldValue, newValue string) *eventsv1.Event {
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_CHAPTER_LEADER_UPDATED,
		EventId:       uniuri.NewLen(64),
		AggregateType: chapterAggregateType,
		AggregateId:   urn,
		Payload: &eventsv1.Event_ChapterLeaderUpdated{
			ChapterLeaderUpdated: &eventsv1.ChapterLeaderUpdated{
				Urn: urn,
				Old: oldValue,
				New: newValue,
			},
		},
	}
}
