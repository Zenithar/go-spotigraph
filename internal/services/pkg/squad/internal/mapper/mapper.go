package mapper

import (
	"fmt"

	"go.zenithar.org/spotigraph/internal/models"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Squad) *squadv1.Squad {
	return &squadv1.Squad{
		Id:    entity.ID,
		Label: entity.Label,
		Urn:   fmt.Sprintf("spfg:v1::squad:%s", entity.ID),
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.Squad) []*squadv1.Squad {
	res := make([]*squadv1.Squad, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
