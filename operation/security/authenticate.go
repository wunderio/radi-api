package security

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Operations for authenticating access to the API
 * which can come in a few forms.  This files holds
 * the Base operations, on which handlers should
 * build.
 */

// Base class for security authenticate Operation
type BaseSecurityAuthenticateOperation struct{}

// Id the operation
func (authenticate *BaseSecurityAuthenticateOperation) Id() string {
	return "security.authenticate"
}

// Label the operation
func (authenticate *BaseSecurityAuthenticateOperation) Label() string {
	return "Authenticate"
}

// Description for the operation
func (authenticate *BaseSecurityAuthenticateOperation) Description() string {
	return "Authenticate access to the app."
}

// Man page for the operation
func (authenticate *BaseSecurityAuthenticateOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (authenticate *BaseSecurityAuthenticateOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
