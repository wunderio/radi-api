package api

/**
 * The API interface defines the simplest top level radi-api components that
 * can be included in an application.  It's sole purpose is to provide a list
 * of operations.
 *
 * The API can be manually build, or you can look a the builder options.
 */

import (
	"github.com/wunderkraut/radi-api/operation"
	"github.com/wunderkraut/radi-api/result"
)

// API is an instance of the API library that can return Operations
type API interface {
	// Validate that the API is ready to give operations.
	Validate() result.Result
	// Operations returns a list of executable operations
	Operations() operation.Operations
}
