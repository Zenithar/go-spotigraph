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
)

func init() {
	log.CheckErr("Unable to register feature flags", feature.DefaultMutableGate.Add(defaultSpotigraphFeatureGates))
}

// defaultKubernetesFeatureGates consists of all known Spotigraph-specific feature keys.
var defaultSpotigraphFeatureGates = map[feature.Feature]feature.Spec{
	OpenTracingDecorator: {Default: false, PreRelease: feature.GA},
	PrometheusDecorator:  {Default: false, PreRelease: feature.GA},
	LoggerDecorator:      {Default: false, PreRelease: feature.GA},
}
