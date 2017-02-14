package config

import (
	log "github.com/Sirupsen/logrus"

	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Here are the commond shared propertys for the various
 * Config operations.
 */

const (
	// config for a single config key
	OPERATION_PROPERTY_CONFIG_KEY = "config.key"
	// config for an orerered list of keys
	OPERATION_PROPERTY_CONFIG_KEYS = "config.keys"
	// config for a single config value ([]byte])
	OPERATION_PROPERTY_CONFIG_VALUE = "config.value"
	// config for a single config value ([]byte])
	OPERATION_PROPERTY_CONFIG_SCOPE = "config.scope"
	// config for a single config value (as an io.readet)
	OPERATION_PROPERTY_CONFIG_VALUE_READERS = "config.value.reader"
	// config for a single config value (as an io.writer)
	OPERATION_PROPERTY_CONFIG_VALUE_WRITERS = "config.value.writer"
)

// property for a single config ket
type ConfigKeyProperty struct {
	api_property.StringProperty
}

// Id for the property
func (confKey *ConfigKeyProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEY
}

// Label for the property
func (confKey *ConfigKeyProperty) Label() string {
	return "property key."
}

// Description for the property
func (confKey *ConfigKeyProperty) Description() string {
	return "property key."
}

// Is the Property internal only
func (confKey *ConfigKeyProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// property for an ordered list of config keys
type ConfigKeysProperty struct {
	api_property.StringSliceProperty
}

// Id for the property
func (keyValue *ConfigKeysProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEYS
}

// Label for the property
func (keyValue *ConfigKeysProperty) Label() string {
	return "Configuration key list."
}

// Description for the property
func (keyValue *ConfigKeysProperty) Description() string {
	return "List of configuration keys."
}

// Is the Property internal only
func (keyValue *ConfigKeysProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// property for a single config value
type ConfigValueProperty struct {
	api_property.BytesArrayProperty
}

// Id for the property
func (property *ConfigValueProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE
}

// Label for the property
func (property *ConfigValueProperty) Label() string {
	return "Configuration content"
}

// Description for the property
func (property *ConfigValueProperty) Description() string {
	return "Content of the configuration"
}

// Is the Property internal only
func (property *ConfigValueProperty) Usage() api_usage.Usage {
	return api_property.Usage_Internal()
}

// property for a value as a set of io.Readers
type ConfigValueScopedReadersProperty struct {
	value ScopedReaders
}

// Id for the property
func (property *ConfigValueScopedReadersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_READERS
}

// Label for the property
func (property *ConfigValueScopedReadersProperty) Label() string {
	return "Config value readers."
}

// Description for the property
func (property *ConfigValueScopedReadersProperty) Description() string {
	return "Config value in the form of an ScopeReaders, which is an ordered map of io.Readers."
}

// Is the Property internal only
func (property *ConfigValueScopedReadersProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (property *ConfigValueScopedReadersProperty) Type() string {
	return "operation/config.ScopeReaders"
}

// Retreive the property value
func (property *ConfigValueScopedReadersProperty) Get() interface{} {
	return interface{}(property.value)
}

// Assign the property value
func (property *ConfigValueScopedReadersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedReaders); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedReaders")
		return false
	}
}

// property for a single value as an io.Writer
type ConfigValueScopedWritersProperty struct {
	value ScopedWriters
}

// Id for the property
func (property *ConfigValueScopedWritersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_WRITERS
}

// Label for the property
func (property *ConfigValueScopedWritersProperty) Label() string {
	return "Config value writers."
}

// Description for the property
func (property *ConfigValueScopedWritersProperty) Description() string {
	return "Config value in the form of an io.Writer."
}

// Is the Property internal only
func (property *ConfigValueScopedWritersProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (property *ConfigValueScopedWritersProperty) Type() string {
	return "operation/config.ScopeWriters"
}

// Retreive the property value
func (property *ConfigValueScopedWritersProperty) Get() interface{} {
	return interface{}(property.value)
}

// Assign the property value
func (property *ConfigValueScopedWritersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedWriters); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedWriters")
		return false
	}
}
