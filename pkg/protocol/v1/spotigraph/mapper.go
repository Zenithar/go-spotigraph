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

// ---------------------------------------------------------------------------

// FromSquad returns a DTO instance from entity one
func FromSquad(entity *models.Squad) *Domain_Squad {
	return &Domain_Squad{
		Id:   entity.ID,
		Name: entity.Name,
		Urn:  entity.URN(),
	}
}

// FromSquads returns a DTO collection from model collection
func FromSquads(entities []*models.Squad) []*Domain_Squad {
	collection := make([]*Domain_Squad, len(entities))
	for k, v := range entities {
		collection[k] = FromSquad(v)
	}
	return collection
}

// ---------------------------------------------------------------------------

// FromChapter returns a DTO instance from entity one
func FromChapter(entity *models.Chapter) *Domain_Chapter {
	return &Domain_Chapter{
		Id:   entity.ID,
		Name: entity.Name,
		Urn:  entity.URN(),
	}
}

// FromChapters returns a DTO collection from model collection
func FromChapters(entities []*models.Chapter) []*Domain_Chapter {
	collection := make([]*Domain_Chapter, len(entities))
	for k, v := range entities {
		collection[k] = FromChapter(v)
	}
	return collection
}
