package spotigraph

import (
	"go.zenithar.org/spotigraph/internal/models"
)

// FromUser returns a DTO instance from entity one
func FromUser(entity *models.User) *Domain_User {
	return &Domain_User{
		Id:        entity.ID,
		Principal: entity.Principal,
	}
}

// FromUsers returns a DTO collection from model collection
func FromUsers(entities []*models.User) []*Domain_User {
	collection := make([]*Domain_User, len(entities))
	for k, v := range entities {
		collection[k] = FromUser(v)
	}
	return collection
}
