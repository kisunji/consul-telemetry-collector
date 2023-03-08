package confresolver

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
)

// Config is a helper type to create a new opentelemetry server configuration.
// It implements a map[string]interface{} representation of the opentelemetry-collector configuration.
// More information can be found here: https://opentelemetry.io/docs/collector/configuration/
type Config struct {
	Receivers  telemetryComponents `mapstructure:"receivers"`
	Exporters  telemetryComponents `mapstructure:"exporters"`
	Processors telemetryComponents `mapstructure:"processors"`
	Connectors telemetryComponents `mapstructure:"connectors"`
	Extensions telemetryComponents `mapstructure:"extensions"`
	Service    service.Config      `mapstructure:"service"`
}

type telemetryComponents map[component.ID]interface{}
type componentConfig map[string]interface{}

// ComponentConfig is an interface that lets us set key/value entries or child maps on the component
type ComponentConfig interface {
	Set(k, v string)
	Map(k string) ComponentConfig
}

func (t componentConfig) Set(k, v string) {
	t[k] = v
}

func (t componentConfig) Map(k string) ComponentConfig {
	tc := make(componentConfig)
	t[k] = tc
	return tc
}

// addComponent ensures that when a new otel component is created it is always added to the pipeline and the appropriate component for further configuration.
func addComponent(comp telemetryComponents, id component.ID,
	pipelineComponent []component.ID) ([]component.ID, componentConfig) {
	// create the new config
	ccfg := make(componentConfig)
	comp[id] = ccfg

	// add to the pipeline slice (and make sure that's not nil)
	if pipelineComponent == nil {
		pipelineComponent = make([]component.ID, 0, 1)
	}
	pipelineComponent = append(pipelineComponent, id)

	// return both the slice and config
	return pipelineComponent, ccfg
}
