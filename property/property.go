package property

import (
	"github.com/wunderkraut/radi-api/usage"
)

/**
 * Properties are abstract Operations values which are meant to be accessed
 * from outside of an operation to either retrieve or assign. This file
 * provides a property collection struct, and the interface for a single
 * Property.
 *
 * A Property consumer should either recognize the Operation by it's keys, and
 * then handle it's Properties as "knowns", or it should iterate through the
 * Properties and use some user-interface to allow interaction.
 *
 * A Property is typically used, by overwriting it's value before running the
 * operation, or by retreiving it's value after the operation has executed.
 */

// A single Property for an operation
type Property interface {

	/**
	 * MetaData Properties
	 */

	// ID returns string unique property Identifier
	Id() string
	// Label returns a short user readable label for the property
	Label() string
	// Description provides a longer multi-line string description of what the property does
	Description() string

	// Give an idea of what type of value the property consumes
	Type() string

	/**
	 * Mark what usage consumers should expect
	 */

	Usage() usage.Usage

	/**
	 * Value Accessors
	 */

	// Value allows the retrieval and setting of unknown Typed values for the property.
	Get() interface{}
	Set(interface{}) bool
}
