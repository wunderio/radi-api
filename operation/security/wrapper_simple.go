package security

import (
	api_operation "github.com/james-nesbitt/kraut-api/operation"
)

/**
 * A simple blocking impementation of the security
 * wrapper.
 */

// Constructor for SimpleSecurityWrapper
func New_SimpleSecurityWrapper(operations *api_operation.Operations) *SimpleSecurityWrapper {
	return &SimpleSecurityWrapper{
		operations: operations,
	}
}

// Simple blocking security wrapper
type SimpleSecurityWrapper struct {
	operations *api_operation.Operations
}

// Convert this into a SecurityWrapper
func (simple *SimpleSecurityWrapper) SecurityWrapper() SecurityWrapper {
	return SecurityWrapper(simple)
}

// Authorize an operation
func (simple *SimpleSecurityWrapper) AuthorizeOperation(operation api_operation.Operation) RuleResult {
	authOp, exists := simple.operations.Get(OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION)

	if !exists {
		return nil
	}

	props := authOp.Properties()
	opProp, _ := props.Get(SECURITY_AUTHORIZATION_OPERATION_PROPERTY_KEY)
	opProp.Set(operation)

	if success, _ := authOp.Exec().Success(); !success {
		return nil
	}

	resultProp, _ := props.Get(SECURITY_AUTHORIZATION_RULERESULT_PROPERTY_KEY)
	return resultProp.Get().(RuleResult)

	return nil
}

// Return a representation of the current user
func (simple *SimpleSecurityWrapper) CurrentUser() SecurityUser {
	authOp, exists := simple.operations.Get(OPERATION_KEY_SECURITY_USER_OPERATION)

	if !exists {
		return nil
	}

	if success, _ := authOp.Exec().Success(); !success {
		props := authOp.Properties()
		opProp, _ := props.Get(SECURITY_USER_PROPERTY_KEY)
		user := opProp.Get().(SecurityUser)
		return user
	}

	return nil
}
