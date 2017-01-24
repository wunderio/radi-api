package builder

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	"github.com/wunderkraut/radi-api/operation"
	"github.com/wunderkraut/radi-api/operation/security"
)

/**
 * SecureBuilderAPI
 *
 * A BuilderAPI that implements a security layer
 */
type SecureBuilderAPI struct {
	BuilderAPI
}

// Return operations where each operation is decorated with the authorize operation
// Return a list of operations for the API from all of the activated Builders
func (secureBuilderApi *SecureBuilderAPI) Operations() operation.Operations {
	ops := operation.Operations{}

	// Build all of the actual operations, which we will then decorate
	builderOps := secureBuilderApi.BuilderAPI.Operations()

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

	} else {
		log.WithError(errors.New("Secure Builder API desn't have access to any security Authorization operation.")).Error("Operations could not be authorized, so none are available")
	}

	return ops
}
