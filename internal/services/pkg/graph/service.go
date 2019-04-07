package graph

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	users    repositories.User
	squads   repositories.Squad
	chapters repositories.Chapter
	guilds   repositories.Guild
	tribes   repositories.Tribe

	nodeResolverMap map[string]nodeResolverFunc
}

type nodeResolverFunc func(ctx context.Context, id string) (*spotigraph.Graph_Node, error)

const (
	urnPrefix      = "urn"
	urnApplication = "spfg"
	urnVersion     = "v1"
)

// New returns a service instance
func New(users repositories.User, squads repositories.Squad, chapters repositories.Chapter, guilds repositories.Guild, tribes repositories.Tribe) services.Graph {
	s := &service{
		users:    users,
		squads:   squads,
		chapters: chapters,
		guilds:   guilds,
		tribes:   tribes,
	}

	// Build resolvermap
	s.nodeResolverMap = map[string]nodeResolverFunc{
		"user": s.resolveUser,
	}

	// REtrun service
	return s
}

// -----------------------------------------------------------------------------

func (s *service) Expand(ctx context.Context, req *spotigraph.NodeInfoReq) (*spotigraph.GraphRes, error) {
	res := &spotigraph.GraphRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Resolve node first
	node, err := s.resolveNode(ctx, req.Urn)
	if err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "Unable to resolve requested node",
		}
		return res, err
	}

	// Expand edges
	res.Graph, err = s.expandNode(ctx, node)
	if err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Unable to retrieve node's edges",
		}
		return res, err
	}

	// Return result
	return res, nil
}

// -----------------------------------------------------------------------------

func (s *service) resolveNode(ctx context.Context, urn string) (*spotigraph.Graph_Node, error) {
	// Split urn as parts
	parts := strings.SplitN(urn, ":", 4)

	// Check part length
	if len(parts) < 4 {
		return nil, errors.New("invalid urn: wrong part length")
	}

	// Check static parts
	if parts[0] != urnPrefix {
		return nil, errors.New("invalid urn: invalid prefix")
	}
	if parts[1] != urnApplication {
		return nil, errors.New("invalid urn: invalid application")
	}
	if parts[2] != urnVersion {
		return nil, errors.New("invalid urn: invalid version")
	}

	// Rerieve node from repository
	if resolver, ok := s.nodeResolverMap[parts[3]]; ok {
		return resolver(ctx, parts[4])
	}

	// Error as default
	return nil, errors.New("invalid urn: invalid type")
}

func (s *service) resolveUser(ctx context.Context, id string) (*spotigraph.Graph_Node, error) {
	// Retrieve from repository
	entity, err := s.users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Transform as node
	return &spotigraph.Graph_Node{
		Type: spotigraph.Graph_Node_USER,
		Urn:  entity.URN(),
		Properties: map[string]string{
			"principal": entity.Principal,
		},
	}, nil
}

func (s *service) expandNode(ctx context.Context, n *spotigraph.Graph_Node) (*spotigraph.Graph, error) {
	return nil, nil
}
