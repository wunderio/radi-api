package setting

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_SETTING_LIST = "setting.list"
)

/**
 * Retrieve keyed Properties for the config handler
 */

// Base class for config list Operation
type BaseSettingListOperation struct{}

// Id the operation
func (list *BaseSettingListOperation) Id() string {
	return OPERATION_ID_SETTING_LIST
}

// Label the operation
func (list *BaseSettingListOperation) Label() string {
	return "Config List"
}

// Description for the operation
func (list *BaseSettingListOperation) Description() string {
	return "Retrieve a list of available configuration keys."
}

// Man page for the operation
func (list *BaseSettingListOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (list *BaseSettingListOperation) usage() api_usage.Usage {
	return api_operation.Usage_External()
}

// Provide properties
func (list *BaseSettingListOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingKeyProperty{}),
		api_property.Usage_Optional(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingScopeProperty{}),
		api_property.Usage_Optional(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingKeysProperty{}),
		api_property.Usage_ReadOnly(),
	))

	return props.Properties()
}
