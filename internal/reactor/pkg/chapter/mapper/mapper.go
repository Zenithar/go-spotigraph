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

package mapper

import (
	"fmt"

	"go.zenithar.org/spotigraph/internal/models"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Chapter) *chapterv1.Chapter {
	return &chapterv1.Chapter{
		Id:       entity.ID,
		Label:    entity.Label,
		LeaderId: entity.LeaderID,
		Urn:      fmt.Sprintf("spfg:v1::chapter:%s", entity.ID),
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.Chapter) []*chapterv1.Chapter {
	res := make([]*chapterv1.Chapter, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
