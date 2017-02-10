package builder

import (
	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * A project is a top level struct used by
 * API consumers who want to implement the builder
 * approach.
 * If a projct is used, then you can just
 * add builders, activate them, and then ask for
 * operations.
 */

/**
 * An interface that marks struct that builds handlers
 * using Builders
 */

type Project interface {
	// Produc an API implementation from this builder
	API() api.API
	// Make a new Builder available in the Project (available to be activated)
	AddBuilder(builder Builder)
	// Activate some implementations for a handler on the builder. The builder should use these to decide what handlers to include
	ActivateBuilder(id string, implementations Implementations, settings SettingsProvider) error
	// Ask a BuilderAPI to validate itself, after it has been fully activated, before we ask for operations.
	Validate() operation.Result
	// Return a list of operations for the constructor from all of the activated Builders
	Operations() operation.Operations
}
