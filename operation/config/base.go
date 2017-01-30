package config

import (
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * Base operations for configs, primarily giving some base Property structs
 */

// Config Base peration that has just a Key property
type BaseConfigKeyOperation struct{}

// Return operation properties
func (base *BaseConfigKeyOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&ConfigKeyProperty{}))

	return props
}

// Base Config operation that has a string key, and bytes array value property pair
type BaseConfigKeyValueOperation struct {
	properties *operation.Properties
}

// Return operation properties
func (base *BaseConfigKeyValueOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&ConfigKeyProperty{}))
	props.Add(operation.Property(&ConfigValueProperty{}))

	return props
}

// Base Config operation that has a string key, and io.Reader value property pair
type BaseConfigKeyReadersOperation struct{}

// Return operation properties
func (base *BaseConfigKeyReadersOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&ConfigKeyProperty{}))
	props.Add(operation.Property(&ConfigValueScopedReadersProperty{}))

	return props
}

// Base Config operation that has a string key, and io.Writer value property pair
type BaseConfigKeyWritersOperation struct{}

// Return operation properties
func (base *BaseConfigKeyWritersOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&ConfigKeyProperty{}))
	props.Add(operation.Property(&ConfigValueScopedWritersProperty{}))

	return props
}

// Base Config operation that has a parent key, and key slice property pair
type BaseConfigKeyKeysOperation struct{}

// Return Operation properties
func (base *BaseConfigKeyKeysOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&ConfigKeyProperty{}))
	props.Add(operation.Property(&ConfigKeysProperty{}))

	return props
}
