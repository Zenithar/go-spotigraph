package features

import (
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
)

const (
	// TracingRepositoryDecorator controls repository layer to add or remove
	// tracing feature
	TracingRepositoryDecorator feature.Feature = "TracingRepositoryDecorator"

	// MetricRepositoryDecorator controls repositroy layer to add or remove
	// metric feature
	MetricRepositoryDecorator feature.Feature = "MetricRepositoryDecorator"
)

func init() {
	log.CheckErr("Unable to register feature flags", feature.DefaultMutableGate.Add(defaultSpotigraphFeatureGates))
}

// defaultKubernetesFeatureGates consists of all known Spotigraph-specific feature keys.
var defaultSpotigraphFeatureGates = map[feature.Feature]feature.Spec{
	TracingRepositoryDecorator: {Default: true, PreRelease: feature.GA},
	MetricRepositoryDecorator:  {Default: true, PreRelease: feature.GA},
}
