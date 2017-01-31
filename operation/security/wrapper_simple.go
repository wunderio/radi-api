package security

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
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

	result := authOp.Exec(&props)
	<-result.Finished()

	if result.Success() {
		resultProp, _ := props.Get(SECURITY_AUTHORIZATION_RULERESULT_PROPERTY_KEY)
		return resultProp.Get().(RuleResult)
	} else {
		return RuleResult(New_SimpleRuleResult(OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION, "Authorization operation failed to execute", -1))
	}
}

// Return a representation of the current user
func (simple *SimpleSecurityWrapper) CurrentUser() SecurityUser {
	authOp, exists := simple.operations.Get(OPERATION_KEY_SECURITY_USER_OPERATION)

	if !exists {
		return nil
	}

	props := authOp.Properties()
	result := authOp.Exec(&props)

	if result.Success() {
		opProp, _ := props.Get(SECURITY_USER_PROPERTY_KEY)
		user := opProp.Get().(SecurityUser)
		return user
	} else {
		return nil
	}
}
