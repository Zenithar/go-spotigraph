package mapper

import (
	"fmt"

	"go.zenithar.org/spotigraph/internal/models"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Person) *personv1.Person {
	return &personv1.Person{
		Id:  entity.ID,
		Urn: fmt.Sprintf("spfg:v1::person:%s", entity.ID),
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.Person) []*personv1.Person {
	res := make([]*personv1.Person, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
