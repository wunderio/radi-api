package security

import (
	api_operation "github.com/james-nesbitt/kraut-api/operation"
)

/**
 * Security specifc properties
 */

const (
	SECURITY_AUTHENTICATION_SUCCEEDED_PROPERTY_KEY = "security.authentication.success"
	SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY  = "security.authoriation.success"
)

// A boolean property for if an authentication succeeded
type SecurityAuthenticationSucceededProperty struct {
	api_operation.BooleanProperty
}

// ID returns string unique property Identifier
func (authSuccessOp *SecurityAuthenticationSucceededProperty) Id() string {
	return SECURITY_AUTHENTICATION_SUCCEEDED_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (authSuccessOp *SecurityAuthenticationSucceededProperty) Label() string {
	return "Authentication success"
}

// Description provides a longer multi-line string description of what the property does
func (authSuccessOp *SecurityAuthenticationSucceededProperty) Description() string {
	return "Indicated whether or not an authentication operation has succeeded"
}

// Mark a property as being for internal use only (no shown to users)
func (authSuccessOp *SecurityAuthenticationSucceededProperty) Internal() bool {
	return false
}

// A boolean property for if an authorization succeeded (was granted)
type SecurityAuthorizationSucceededProperty struct {
	api_operation.BooleanProperty
}

// ID returns string unique property Identifier
func (authSuccessOp *SecurityAuthorizationSucceededProperty) Id() string {
	return SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (authSuccessOp *SecurityAuthorizationSucceededProperty) Label() string {
	return "Authorization success"
}

// Description provides a longer multi-line string description of what the property does
func (authSuccessOp *SecurityAuthorizationSucceededProperty) Description() string {
	return "Indicated whether or not an autheorization operation has succeeded"
}

// Mark a property as being for internal use only (no shown to users)
func (authSuccessOp *SecurityAuthorizationSucceededProperty) Internal() bool {
	return false
}
