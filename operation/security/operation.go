package security

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Security operations
 */

// A constructor for DecoratedBooleanPropertyBasedOperation
func New_SecureBuilderAuthorizationDecoratorOperation(authorizing api_operation.Operation, authorized api_operation.Operation, operationProperty string, successProperty string) *SecureBuilderAuthorizationDecoratorOperation {
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
 */

// Custom Authorization Decorator operation
type SecureBuilderAuthorizationDecoratorOperation struct {
	authorizing       api_operation.Operation
	authorized        api_operation.Operation
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
func (op *SecureBuilderAuthorizationDecoratorOperation) Properties() *api_operation.Properties {
	props := api_operation.Properties{}

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
			props.Add(api_operation.New_DecoratingInternalizerProperty(authProp).Property())
		}
	}

	props.Merge(*op.authorized.Properties())
	return &props
}

// Execute the authorizing operation, and then execute the authorized operation if the authorizing property is true
func (op *SecureBuilderAuthorizationDecoratorOperation) Exec() api_operation.Result {
	result := api_operation.BaseResult{}
	result.Set(true, []error{})

	authOpProps := op.authorizing.Properties()
	successProp, successPropFound := authOpProps.Get(op.successProperty)
	operationProp, operationPropFound := authOpProps.Get(op.operationProperty)

	if !(successPropFound && operationPropFound) || successProp.Type() != "bool" {
		// this authorization op is not valid, it is either missing its op or success property

		if !operationPropFound {
			result.Set(false, []error{errors.New("Secure Builder API Authorize operation is invalid.  It is missing the operation authorization property.")})
		} else if !successPropFound {
			result.Set(false, []error{errors.New("Secure Builder API Authorize operation is invalid.  It is missing the authorization success property.")})
		} else if successProp.Type() != "bool" {
			result.Set(false, []error{errors.New("Secure Builder API Authorize operation is invalid.  The authorization success property is not a bool.")})
		}

		result.Set(false, []error{errors.New("Secure Builder API could not execute authorized Operation.")})
	} else {

		operationProp.Set(op.authorized)
		result.Merge(op.authorizing.Exec())

		if success, errs := result.Success(); !success {
			errs = append(errs, errors.New("Operation authorization failed to execute."))
			result.Set(false, errs)
		} else {

			if successProp.Get().(bool) {
				// The Auth op returned a TRUE success value
				result.Merge(op.authorized.Exec())
			} else {
				// The Auth op returned a FALSE success value
				result.Set(false, []error{errors.New("Authorization failed.  You are not permitted to execute the requested operation: " + op.authorized.Id())})
			}

		}

	}

	return api_operation.Result(&result)
}
