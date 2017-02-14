package operation

import (
	"github.com/wunderkraut/radi-api/property"
	"github.com/wunderkraut/radi-api/result"
	"github.com/wunderkraut/radi-api/usage"
)

/**
 * This file holds the definition for an API Operation, and
 * also defines a usefull Operation list struct, as well as
 * a utility BaseOperation struct, which can be used for
 * Operation inheritance.
 */

// A single operation
type Operation interface {

	// METADATA

	// Return the string machinename/id of the Operation
	Id() string
	// Return a user readable string label for the Operation
	Label() string
	// return a multiline string description for the Operation
	Description() string
	// return a multiline string description for the Operation
	Help() string

	// In what ways is this operation meant to be used
	Usage() usage.Usage

	// FUNCTIONAL

	// Run a validation check on the Operation
	Validate() result.Result

	// What settings/values does the Operation provide to an implemenentor
	Properties() property.Properties

	// Execute the Operation
	Exec(property.Properties) result.Result
}
