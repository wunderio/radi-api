package command

import (
	api_property "github.com/wunderkraut/radi-api/property"
)

/**
 * Base Operation classes for command implementation
 */

// A Base command operation that provides a single command key string and a command object
type BaseCommandKeyCommandOperation struct{}

// Return a static keys list Property
func (base *BaseCommandKeyCommandOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&CommandKeyProperty{}))
	props.Add(api_property.Property(&CommandCommandProperty{}))

	return props.Properties()
}

// A Base command operation that returns a list of keys
type BaseCommandKeyKeysOperation struct{}

// Return a static keys list Property
func (base *BaseCommandKeyKeysOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&CommandKeyProperty{}))
	props.Add(api_property.Property(&CommandKeysProperty{}))

	return props.Properties()
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandContextOperation struct{}

// get static Properties
func (base *BaseCommandContextOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&CommandContextProperty{}))

	return props.Properties()
}

// A base command operation that provides a a key, flags strings, and various input and output Properties
type BaseCommandInputOutputOperation struct{}

// get static Properties
func (base *BaseCommandInputOutputOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&CommandInputProperty{}))
	props.Add(api_property.Property(&CommandOutputProperty{}))
	props.Add(api_property.Property(&CommandErrorProperty{}))

	return props.Properties()
}

// A base command operation that provides akey, flags
type BaseCommandFlagsOperation struct{}

// get static Properties
func (base *BaseCommandFlagsOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&CommandFlagsProperty{}))

	return props.Properties()
}
