package command

import (
	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * A base Command definition, which defines the command
 * container property, but may receive overrides for
 * flags, input/error/output
 *
 * It turns out that a Command has a very similar need to
 * operations, and so it makes sense to write a command
 * interface that can be used as an operation, and can
 * give an operation
 */

// Command definition
type Command interface {
	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string

	// Run a validation check on the Operation
	Validate() api_result.Result

	// Is this operation meant to be used only inside the API
	Usage() api_usage.Usage

	// What settings does the Operation provide to an implemenentor
	Properties() api_property.Properties

	// Execute the Operation
	Exec(api_property.Properties) api_result.Result
}
