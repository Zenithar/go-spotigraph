package graph

import (
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func toNode(entity interface{}) *spotigraph.Graph_Node {
	switch e := entity.(type) {

	case *models.User:
		return &spotigraph.Graph_Node{
			Type: spotigraph.Graph_Node_USER,
			Urn:  e.URN(),
			Properties: map[string]string{
				"principal": e.Principal,
			},
		}

	case *models.Squad:
		return &spotigraph.Graph_Node{
			Type: spotigraph.Graph_Node_SQUAD,
			Urn:  e.URN(),
			Properties: map[string]string{
				"name": e.Name,
			},
		}

	case *models.Chapter:
		return &spotigraph.Graph_Node{
			Type: spotigraph.Graph_Node_CHAPTER,
			Urn:  e.URN(),
			Properties: map[string]string{
				"name": e.Name,
			},
		}

	case *models.Guild:
		return &spotigraph.Graph_Node{
			Type: spotigraph.Graph_Node_GUILD,
			Urn:  e.URN(),
			Properties: map[string]string{
				"name": e.Name,
			},
		}

	case *models.Tribe:
		return &spotigraph.Graph_Node{
			Type: spotigraph.Graph_Node_TRIBE,
			Urn:  e.URN(),
			Properties: map[string]string{
				"name": e.Name,
			},
		}

	}

	return &spotigraph.Graph_Node{
		Type: spotigraph.Graph_Node_UNDEFINED,
	}
}
