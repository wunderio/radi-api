package command

import (
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * Base Operation classes for command implementation
 */

// A Base command operation that provides a single command key string and a command object
type BaseCommandKeyCommandOperation struct{}

// Return a static keys list Property
func (base *BaseCommandKeyCommandOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&CommandKeyProperty{}))
	props.Add(operation.Property(&CommandCommandProperty{}))

	return props
}

// A Base command operation that returns a list of keys
type BaseCommandKeyKeysOperation struct{}

// Return a static keys list Property
func (base *BaseCommandKeyKeysOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&CommandKeyProperty{}))
	props.Add(operation.Property(&CommandKeysProperty{}))

	return props
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandContextOperation struct{}

// get static Properties
func (base *BaseCommandContextOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&CommandContextProperty{}))

	return props
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandInputOutputOperation struct{}

// get static Properties
func (base *BaseCommandInputOutputOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&CommandInputProperty{}))
	props.Add(operation.Property(&CommandOutputProperty{}))
	props.Add(operation.Property(&CommandErrorProperty{}))

	return props
}

// A base command operation that provides akey, flags
type BaseCommandFlagsOperation struct {
	properties *operation.Properties
}

// get static Properties
func (base *BaseCommandFlagsOperation) Properties() operation.Properties {
	props := operation.Properties{}

	props.Add(operation.Property(&CommandFlagsProperty{}))

	return props
}
