package config

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_CONFIG_READERS = "config.readers"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseConfigReadersOperation struct{}

// Id the operation
func (readers *BaseConfigReadersOperation) Id() string {
	return OPERATION_ID_CONFIG_READERS
}

// Label the operation
func (readers *BaseConfigReadersOperation) Label() string {
	return "Config Readers"
}

// Description for the operation
func (readers *BaseConfigReadersOperation) Description() string {
	return "Retrieve a keyed configuration scoped reader set."
}

// Man Page for the operation
func (readers *BaseConfigReadersOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (readers *BaseConfigReadersOperation) Usage() api_usage.Usage {
	return api_operation.Usage_Internal()
}

// Return operation properties
func (readers *BaseConfigReadersOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigKeyProperty{}),
		api_property.Usage_Required(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigValueScopedReadersProperty{}),
		api_property.Usage_ReadOnly(),
	))

	return props.Properties()
}
