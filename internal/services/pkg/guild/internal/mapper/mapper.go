package mapper

import (
	"fmt"

	"go.zenithar.org/spotigraph/internal/models"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Guild) *guildv1.Guild {
	return &guildv1.Guild{
		Id:    entity.ID,
		Label: entity.Label,
		Urn:   fmt.Sprintf("spfg:v1::guild:%s", entity.ID),
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.Guild) []*guildv1.Guild {
	res := make([]*guildv1.Guild, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
