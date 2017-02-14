package handler

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * This file defines the core handle functionality, which
 * actual handlers will implement.
 */

// A single API handler for providing operations
type Handler interface {
	// Return a string identifier for the Handler
	Id() string
	// Validate that the Handler has everything it needs to give operations
	Validate() api_result.Result
	// Return a list of Operations from the Handler
	Operations() api_operation.Operations
}
