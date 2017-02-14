package setting

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_SETTING_GET = "setting.get"
)

/**
 * Retrieve keyed configurations for the config handler
 */

// Base class for config get Operation
type BaseSettingGetOperation struct{}

// Id the operation
func (get *BaseSettingGetOperation) Id() string {
	return OPERATION_ID_SETTING_GET
}

// Label the operation
func (get *BaseSettingGetOperation) Label() string {
	return "Setting Get"
}

// Description for the operation
func (get *BaseSettingGetOperation) Description() string {
	return "Retrieve a keyed setting."
}

// Man page for the operation
func (get *BaseSettingGetOperation) Help() string {
	return ""
}

// Return External usage
func (get *BaseSettingGetOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}

// Return properties for the operation
func (base *BaseSettingGetOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingKeyProperty{}),
		api_property.Usage_Required(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingScopeProperty{}),
		api_property.Usage_Optional(),
	))
	props.Add(api_property.New_UsageDecoratedProperty(
		api_property.Property(&SettingValueProperty{}),
		api_property.Usage_ReadOnly(),
	))

	return props.Properties()
}
