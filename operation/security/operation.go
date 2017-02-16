package security

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Security operations
 */

/**
 * An operation decorator that uses one operation, with a boolean property
 * to authorized the execution of a second operation
 *
 * This operation blocks the second operation until the first operation has
 * marked itself as ->Finished()
 */

// Custom Authorization Decorator operation
type SecureBuilderAuthorizationDecoratorOperation struct {
	authorizing       api_operation.Operation
	authorized        api_operation.Operation
	operationProperty string
	successProperty   string
}

// A constructor for DecoratedBooleanPropertyBasedOperation
func New_SecureBuilderAuthorizationDecoratorOperation(authorizing api_operation.Operation, authorized api_operation.Operation, operationProperty string, successProperty string) *SecureBuilderAuthorizationDecoratorOperation {
	return &SecureBuilderAuthorizationDecoratorOperation{
		authorizing:       authorizing,
		authorized:        authorized,
		operationProperty: operationProperty,
		successProperty:   successProperty,
	}
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

// Return a long man-page for the operation
func (op *SecureBuilderAuthorizationDecoratorOperation) Help() string {
	return op.authorized.Help()
}

// Is this operation meant to be used only inside the API
func (op *SecureBuilderAuthorizationDecoratorOperation) Usage() api_usage.Usage {
	return op.authorized.Usage()
}

// Run a validation check on the Operation
func (op *SecureBuilderAuthorizationDecoratorOperation) Validate() api_result.Result {
	/**
	 * @TODO developer a result type for merging, and use it here
	 */
	return op.authorized.Validate()
}

// Get Operation Properties from both operations
func (op *SecureBuilderAuthorizationDecoratorOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

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
			props.Add(api_property.New_UsageDecoratedProperty(
				authProp,
				api_property.Usage_Internal(),
			))

		}
	}

	props.Merge(op.authorized.Properties())
	return props.Properties()
}

// Execute the authorizing operation, and then execute the authorized operation if the authorizing property is true
func (op *SecureBuilderAuthorizationDecoratorOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	successProp, successPropFound := props.Get(op.successProperty)
	operationProp, operationPropFound := props.Get(op.operationProperty)

	if !(successPropFound && operationPropFound) || successProp.Type() != "bool" {
		// this authorization op is not valid, it is either missing its op or success property

		if !operationPropFound {
			res.MarkFailed()
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the operation authorization property."))
		} else if !successPropFound {
			res.MarkFailed()
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the authorization success property."))
		} else if successProp.Type() != "bool" {
			res.MarkFailed()
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  The authorization success property is not a bool."))
		}

		res.MarkFailed()
		res.AddError(errors.New("Secure Builder API could not execute authorized Operation."))
	} else {

		operationProp.Set(op.authorized)

		authResult := op.authorizing.Exec(props)
		<-authResult.Finished()

		res.Merge(authResult)

		if !res.Success() {
			res.MarkFailed()
			res.AddErrors(append(res.Errors(), errors.New("Operation authorization failed to execute.")))
		} else {

			if successProp.Get().(bool) {
				// The Auth op returned a TRUE success value

				execResult := op.authorized.Exec(props)
				<-execResult.Finished()

				res.Merge(execResult)
				// we should return quicklly after this if the execResult had .finished already at TRUE
			} else {
				// The Auth op returned a FALSE success value
				res.MarkFailed()
				res.AddError(errors.New("Authorization failed.  You are not permitted to execute the requested operation: " + op.authorized.Id()))
			}

		}

	}

	res.MarkFinished()

	return res.Result()
}
