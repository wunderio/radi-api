package builder

import (
	api_api "github.com/wunderkraut/radi-api/api"
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Builders are essentially meta handlers, which can produce handlers from
 * simpler settings.
 *
 * Builders often need a back reference to the api as they may need other
 * operations, for example to build Config based operations
 */

// A single API handler builder
type Builder interface {
	// Return a string identifier for the Handler (not functionally needed yet)
	Id() string
	// Set a API parent for this Builder
	// This makes some builders capable of depending on others' operations.
	SetAPI(parent api_api.API)
	// Initialize and activate the Handler
	Activate(Implementations, SettingsProvider) api_result.Result

	/*
	 * API interface
	 *
	 * We intentionally duplicate the api.API interface here
	 */

	// Validate that the API is ready to give operations.
	Validate() api_result.Result
	// Return a list of Operations from the Handler
	Operations() api_operation.Operations
}
