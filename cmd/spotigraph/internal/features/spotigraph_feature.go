package features

import (
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
)

const (
	// OpenTracingDecorator controls repository layer to add or remove
	// tracing feature
	OpenTracingDecorator feature.Feature = "OpenTracingDecorator"

	// PrometheusDecorator controls repositroy layer to add or remove
	// metric feature
	PrometheusDecorator feature.Feature = "PrometheusDecorator"

	// LoggerDecorator controls repositroy layer to add or remove
	// logging feature
	LoggerDecorator feature.Feature = "LoggerDecorator"

	// RESTv1 controls HTTP REST v1 API exposure from server
	RESTv1 feature.Feature = "RESTv1"

	// GraphQLv1 controls GraphQL v1 API exposure from server
	GraphQLv1 feature.Feature = "GraphQLv1"
)

func init() {
	log.CheckErr("Unable to register feature flags", feature.DefaultMutableGate.Add(defaultSpotigraphFeatureGates))
}

// defaultKubernetesFeatureGates consists of all known Spotigraph-specific feature keys.
var defaultSpotigraphFeatureGates = map[feature.Feature]feature.Spec{
	OpenTracingDecorator: {Default: false, PreRelease: feature.GA},
	PrometheusDecorator:  {Default: false, PreRelease: feature.GA},
	LoggerDecorator:      {Default: false, PreRelease: feature.GA},
	RESTv1:               {Default: true, PreRelease: feature.GA},
	GraphQLv1:            {Default: false, PreRelease: feature.Alpha},
}
