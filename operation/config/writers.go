package config

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_CONFIG_WRITERS = "config.writers"
)

/**
 * Retrieve writers for the Config handler
 */

// Base class for config retrieve writers Operation
type BaseConfigWritersOperation struct{}

// Id the operation
func (writers *BaseConfigWritersOperation) Id() string {
	return OPERATION_ID_CONFIG_WRITERS
}

// Label the operation
func (writers *BaseConfigWritersOperation) Label() string {
	return "Config Writers"
}

// Description for the operation
func (writers *BaseConfigWritersOperation) Description() string {
	return "Get a set of scoped writers for a configuration."
}

// Man Page for the operation
func (writers *BaseConfigWritersOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (writers *BaseConfigWritersOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}

// Return operation properties
func (writers *BaseConfigWritersOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigKeyProperty{}),
		api_property.Usage_Required(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigValueScopedWritersProperty{}),
		api_property.Usage_ReadOnly(),
	))

	return props.Properties()
}
