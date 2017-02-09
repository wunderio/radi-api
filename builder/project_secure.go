package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/operation"
	"github.com/wunderkraut/radi-api/operation/security"
)

/**
 * SecureProject
 *
 * An API Project that implements a security layer
 * on top of the standard project.
 *
 * The security layer is implemented by adding authorization
 * on top of most operations, using authorize and user
 * operations to implement that.
 */
type SecureProject struct {
	StandardProject
}

// Ask a SecureProject to validate itself, after it has been fully activated, before we ask for operations.
func (project *SecureProject) Validate() operation.Result {
	result := operation.New_StandardResult()

	// Build all of the actual operations, which we will then decorate
	builderOps := project.StandardProject.Operations()

	/**
	 * @NOTE that at this point we should also have the Authorize
	 * operation that we need to use as a decorator
	 *
	 * Use some default values
	 */

	if _, found := builderOps.Get(security.OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION); !found {
		result.AddError(errors.New("Secure Builder API desn't have access to any security Authorization operation."))
		result.MarkFailed()
	} else {
		result.MarkSuccess()
	}
	result.MarkFinished()

	return result.Result()
}

// Return operations where each operation is decorated with the authorize operation
// Return a list of operations for the Project from all of the activated Builders
func (project *SecureProject) Operations() operation.Operations {
	ops := operation.Operations{}

	// Build all of the actual operations, which we will then decorate
	builderOps := project.StandardProject.Operations()

	/**
	 * @NOTE that at this point we should also have the Authorize
	 * operation that we need to use as a decorator
	 *
	 * Use some default values
	 */

	if authOp, found := builderOps.Get(security.OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION); found {

		for _, id := range builderOps.Order() {
			op, _ := builderOps.Get(id)
			switch op.Id() {
			case security.OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION:
				ops.Add(op)
			default:
				ops.Add(operation.Operation(security.New_SecureBuilderAuthorizationDecoratorOperation(authOp, op, security.SECURITY_AUTHORIZATION_OPERATION_PROPERTY_KEY, security.SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY)))
			}

		}

	}

	return ops
}