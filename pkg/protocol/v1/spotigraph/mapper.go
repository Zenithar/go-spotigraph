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
