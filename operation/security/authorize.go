package security

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION = "security.authorize"
)

/**
 * Authorize operations should be used to question whether or not
 * the authorized user should be given access to other operations
 *
 * @QUESTION should authentication be tied to Operation.Id() values,
 * or should we also allow user|action|key style options
 */

// Base class for security authorize Operation
type BaseSecurityAuthorizeOperation struct{}

// Id the operation
func (authorize *BaseSecurityAuthorizeOperation) Id() string {
	return OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION
}

// Label the operation
func (authorize *BaseSecurityAuthorizeOperation) Label() string {
	return "Authorize"
}

// Description for the operation
func (authorize *BaseSecurityAuthorizeOperation) Description() string {
	return "Authorize access to a part of the app."
}

// Man page for the operation
func (authorize *BaseSecurityAuthorizeOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (authorize *BaseSecurityAuthorizeOperation) Usage() api_usage.Usage {
	return api_operation.Usage_Internal()
}
