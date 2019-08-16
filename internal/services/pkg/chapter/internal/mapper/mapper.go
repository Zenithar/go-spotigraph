package mapper

import (
	"go.zenithar.org/spotigraph/internal/models"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Chapter) *chapterv1.Chapter {
	return &chapterv1.Chapter{
		Id:    entity.ID,
		Label: entity.Label,
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
