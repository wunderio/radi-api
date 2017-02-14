package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
	"github.com/wunderkraut/radi-api/operation/security"
	"github.com/wunderkraut/radi-api/result"
)

/**
 * SecureProject
 *
 * An API Project that implements a security layer on top of the standard
 * project.
 *
 * The security layer is implemented by adding authorization on top of most
 * operations, using authorize and user operations to implement that.
 */
type SecureProject struct {
	StandardProject
}

// Constructor for SecureProject
func New_SecureProject() *SecureProject {
	return &SecureProject{
		StandardProject: *New_StandardProject(),
	}
}

// Convert this to a Project interface
func (project *SecureProject) Project() Project {
	return Project(project)
}

// Convert his project to an API interface
func (project *SecureProject) API() api.API {
	return api.API(project)
}

// Ask a SecureProject to validate itself, after it has been fully activated, before we ask for operations.
func (project *SecureProject) Validate() result.Result {
	res := result.New_StandardResult()

	parentRes := project.StandardProject.Validate()
	<-parentRes.Finished()

	if !parentRes.Success() {
		return parentRes
	}
	res.Merge(parentRes)

	// Build all of the actual operations, which we will then decorate
	builderOps := project.StandardProject.Operations()

	/**
	 * @NOTE that at this point we should also have the Authorize
	 * operation that we need to use as a decorator
	 *
	 * Use some default values
	 */

	if _, found := builderOps.Get(security.OPERATION_KEY_SECURITY_AUTHORIZE_OPERATION); !found {
		res.AddError(errors.New("Secure Builder API desn't have access to any security Authorization operation."))
		res.MarkFailed()
	} else {
		res.MarkSuccess()
	}
	res.MarkFinished()

	return res.Result()
}

// Return operations where each operation is decorated with the authorize operation
// Return a list of operations for the Project from all of the activated Builders
func (project *SecureProject) Operations() operation.Operations {
	ops := operation.New_SimpleOperations()

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

	return ops.Operations()
}
