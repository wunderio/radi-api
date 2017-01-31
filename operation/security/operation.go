package security

import (
	"errors"

	operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Security operations
 */

// A constructor for DecoratedBooleanPropertyBasedOperation
func New_SecureBuilderAuthorizationDecoratorOperation(authorizing operation.Operation, authorized operation.Operation, operationProperty string, successProperty string) *SecureBuilderAuthorizationDecoratorOperation {
	return &SecureBuilderAuthorizationDecoratorOperation{
		authorizing:       authorizing,
		authorized:        authorized,
		operationProperty: operationProperty,
		successProperty:   successProperty,
	}
}

/**
 * An operation decorator that uses one operation, with a boolean property
 * to authorized the execution of a second operation
 *
 * This operation blocks the second operation until the first operation has
 * marked itself as ->Finished()
 */

// Custom Authorization Decorator operation
type SecureBuilderAuthorizationDecoratorOperation struct {
	authorizing       operation.Operation
	authorized        operation.Operation
	operationProperty string
	successProperty   string
}

// Get decorted operation id
func (op *SecureBuilderAuthorizationDecoratorOperation) Id() string {
	return op.authorized.Id()
}

// Return a short description of the operation
func (op *SecureBuilderAuthorizationDecoratorOperation) Label() string {
	return op.authorized.Label() + " [authorization required]"
}

// Return a long description of the operation
func (op *SecureBuilderAuthorizationDecoratorOperation) Description() string {
	return op.authorized.Description() + " [authorization required]"
}

// Is this operation meant to be used only inside the API
func (op *SecureBuilderAuthorizationDecoratorOperation) Internal() bool {
	return op.authorized.Internal()
}

// Run a validation check on the Operation
func (op *SecureBuilderAuthorizationDecoratorOperation) Validate() bool {
	return op.authorizing.Validate() && op.authorized.Validate()
}

// Get Operation Properties from both operations
func (op *SecureBuilderAuthorizationDecoratorOperation) Properties() operation.Properties {
	props := operation.Properties{}

	// add auth operation props as internal only, by decorating them
	authProps := op.authorizing.Properties()
	for _, index := range authProps.Order() {
		authProp, _ := authProps.Get(index)

		switch authProp.Id() {
		case SECURITY_AUTHENTICATION_SUCCEEDED_PROPERTY_KEY:
			fallthrough
		case SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY:
			props.Add(authProp)
		default:
			props.Add(operation.New_DecoratingInternalizerProperty(authProp).Property())
		}
	}

	props.Merge(op.authorized.Properties())
	return props
}

// Execute the authorizing operation, and then execute the authorized operation if the authorizing property is true
func (op *SecureBuilderAuthorizationDecoratorOperation) Exec(props *operation.Properties) operation.Result {
	result := operation.New_StandardResult()

	successProp, successPropFound := props.Get(op.successProperty)
	operationProp, operationPropFound := props.Get(op.operationProperty)

	if !(successPropFound && operationPropFound) || successProp.Type() != "bool" {
		// this authorization op is not valid, it is either missing its op or success property

		if !operationPropFound {
			result.MarkFailed()
			result.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the operation authorization property."))
		} else if !successPropFound {
			result.MarkFailed()
			result.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the authorization success property."))
		} else if successProp.Type() != "bool" {
			result.MarkFailed()
			result.AddError(errors.New("Secure Builder API Authorize operation is invalid.  The authorization success property is not a bool."))
		}

		result.MarkFailed()
		result.AddError(errors.New("Secure Builder API could not execute authorized Operation."))
	} else {

		operationProp.Set(op.authorized)

		authResult := op.authorizing.Exec(props)
		<-authResult.Finished()

		result.Merge(authResult)

		if !result.Success() {
			result.MarkFailed()
			result.AddErrors(append(result.Errors(), errors.New("Operation authorization failed to execute.")))
		} else {

			if successProp.Get().(bool) {
				// The Auth op returned a TRUE success value

				execResult := op.authorized.Exec(props)
				<-execResult.Finished()

				result.Merge(execResult)
				// we should return quicklly after this if the execResult had .finished already at TRUE
			} else {
				// The Auth op returned a FALSE success value
				result.MarkFailed()
				result.AddError(errors.New("Authorization failed.  You are not permitted to execute the requested operation: " + op.authorized.Id()))
			}

		}

	}

	result.MarkFinished()

	return operation.Result(result)
}
