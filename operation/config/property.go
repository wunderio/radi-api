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
func (key *ConfigKeyProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEY
}

// Label for the property
func (key *ConfigKeyProperty) Label() string {
	return "property key."
}

// Description for the property
func (key *ConfigKeyProperty) Description() string {
	return "property key."
}

// Is the Property internal only
func (key *ConfigKeyProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// Copy the property
func (key *ConfigKeyProperty) Copy() api_property.Property {
	prop := &ConfigKeyProperty{}
	prop.Set(key.Get())
	return api_property.Property(prop)
}

// property for an ordered list of config keys
type ConfigKeysProperty struct {
	api_property.StringSliceProperty
}

// Id for the property
func (keys *ConfigKeysProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_KEYS
}

// Label for the property
func (keys *ConfigKeysProperty) Label() string {
	return "Configuration key list."
}

// Description for the property
func (keys *ConfigKeysProperty) Description() string {
	return "List of configuration keys."
}

// Is the Property internal only
func (keys *ConfigKeysProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Copy the property
func (keys *ConfigKeysProperty) Copy() api_property.Property {
	prop := &ConfigKeysProperty{}
	prop.Set(keys.Get())
	return api_property.Property(prop)
}

// property for a single config value
type ConfigValueProperty struct {
	api_property.BytesArrayProperty
}

// Id for the property
func (val *ConfigValueProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE
}

// Label for the property
func (val *ConfigValueProperty) Label() string {
	return "Configuration content"
}

// Description for the property
func (val *ConfigValueProperty) Description() string {
	return "Content of the configuration"
}

// Is the Property internal only
func (val *ConfigValueProperty) Usage() api_usage.Usage {
	return api_property.Usage_Internal()
}

// Copy the property
func (val *ConfigValueProperty) Copy() api_property.Property {
	prop := &ConfigValueProperty{}
	prop.Set(val.Get())
	return api_property.Property(prop)
}

// property for a value as a set of io.Readers
type ConfigValueScopedReadersProperty struct {
	value ScopedReaders
}

// Id for the property
func (scopedReaders *ConfigValueScopedReadersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_READERS
}

// Label for the property
func (scopedReaders *ConfigValueScopedReadersProperty) Label() string {
	return "Config value readers."
}

// Description for the property
func (scopedReaders *ConfigValueScopedReadersProperty) Description() string {
	return "Config value in the form of an ScopeReaders, which is an ordered map of io.Readers."
}

// Is the Property internal only
func (scopedReaders *ConfigValueScopedReadersProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (scopedReaders *ConfigValueScopedReadersProperty) Type() string {
	return "operation/config.ScopeReaders"
}

// Retreive the property value
func (scopedReaders *ConfigValueScopedReadersProperty) Get() interface{} {
	return interface{}(scopedReaders.value)
}

// Assign the property value
func (scopedReaders *ConfigValueScopedReadersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedReaders); ok {
		scopedReaders.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedReaders")
		return false
	}
}

// Copy the property
func (scopedReaders *ConfigValueScopedReadersProperty) Copy() api_property.Property {
	prop := &ConfigValueScopedReadersProperty{}
	prop.Set(scopedReaders.Get())
	return api_property.Property(prop)
}

// property for a single value as an io.Writer
type ConfigValueScopedWritersProperty struct {
	value ScopedWriters
}

// Id for the property
func (scopedWriters *ConfigValueScopedWritersProperty) Id() string {
	return OPERATION_PROPERTY_CONFIG_VALUE_WRITERS
}

// Label for the property
func (scopedWriters *ConfigValueScopedWritersProperty) Label() string {
	return "Config value writers."
}

// Description for the property
func (scopedWriters *ConfigValueScopedWritersProperty) Description() string {
	return "Config value in the form of an io.Writer."
}

// Is the Property internal only
func (scopedWriters *ConfigValueScopedWritersProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (scopedWriters *ConfigValueScopedWritersProperty) Type() string {
	return "operation/config.ScopeWriters"
}

// Retreive the property value
func (scopedWriters *ConfigValueScopedWritersProperty) Get() interface{} {
	return interface{}(scopedWriters.value)
}

// Assign the property value
func (scopedWriters *ConfigValueScopedWritersProperty) Set(value interface{}) bool {
	if converted, ok := value.(ScopedWriters); ok {
		scopedWriters.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected ScopedWriters")
		return false
	}
}

// Copy the property
func (scopedWriters *ConfigValueScopedWritersProperty) Copy() api_property.Property {
	prop := &ConfigValueScopedWritersProperty{}
	prop.Set(scopedWriters.Get())
	return api_property.Property(prop)
}
