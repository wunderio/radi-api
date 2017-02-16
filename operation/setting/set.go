package setting

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_SETTING_SET = "setting.set"
)

/**
 * Set keyed values into the Config handler
 */

// Base class for setting set Operation
type BaseSettingSetOperation struct{}

// Id the operation
func (set *BaseSettingSetOperation) Id() string {
	return OPERATION_ID_SETTING_SET
}

// Label the operation
func (set *BaseSettingSetOperation) Label() string {
	return "Config Set"
}

// Description for the operation
func (set *BaseSettingSetOperation) Description() string {
	return "Set a keyed configuration."
}

// Man page for the operation
func (set *BaseSettingSetOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (set *BaseSettingSetOperation) Usage() api_usage.Usage {
	return api_operation.Usage_Internal()
}

func (set *BaseSettingSetOperation) Properties() api_property.Properties {
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
		api_property.Usage_Required(),
	))

	return props.Properties()
}
