package setting

import (
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * @QUESTION can't we just use a global definition for these, as they
 * are used for config and for settings (and properly others)
 */

// A setting operation base that provides a key string
type BaseSettingKeyScopeOperation struct{}

// Provides properties
func (base *BaseSettingKeyScopeOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&SettingKeyProperty{}))
	props.Add(operation.Property(&SettingScopeProperty{}))

	return props
}

// A setting operation base that provides a string key/value property pair
type BaseSettingKeyScopeValueOperation struct{}

// Provide properties
func (base *BaseSettingKeyScopeValueOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&SettingKeyProperty{}))
	props.Add(operation.Property(&SettingScopeProperty{}))
	props.Add(operation.Property(&SettingValueProperty{}))

	return props
}

// A setting operation base that provides a key list string array
type BaseSettingKeyScopeKeysOperation struct{}

// Provide properties
func (base *BaseSettingKeyScopeKeysOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&SettingKeyProperty{}))
	props.Add(operation.Property(&SettingScopeProperty{}))
	props.Add(operation.Property(&SettingKeysProperty{}))

	return props
}
