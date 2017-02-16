package config

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_CONFIG_LIST = "config.list"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config list Operation
type BaseConfigListOperation struct{}

// Id the operation
func (list *BaseConfigListOperation) Id() string {
	return OPERATION_ID_CONFIG_LIST
}

// Label the operation
func (list *BaseConfigListOperation) Label() string {
	return "Config List"
}

// Description for the operation
func (list *BaseConfigListOperation) Description() string {
	return "Retrieve a list of available configuration keys."
}

// Man page for the operation
func (list *BaseConfigListOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (list *BaseConfigListOperation) Usage() api_usage.Usage {
	return api_operation.Usage_Internal()
}

// Return Operation properties
func (list *BaseConfigListOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigKeyProperty{}),
		api_property.Usage_ReadOnly(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&ConfigKeysProperty{}),
		api_property.Usage_ReadOnly(),
	))

	return props.Properties()
}
