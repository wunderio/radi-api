package provision

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Stop any running servers
 */

const (
	OPERATION_KEY_PROVISION_STOP_OPERATION = "provision.stop"
)

// Down Operation
type BaseProvisionStopOperation struct{}

// Return the string machinename/id of the Operation
func (stop *BaseProvisionStopOperation) Id() string {
	return OPERATION_KEY_PROVISION_STOP_OPERATION
}

// Return a user readable string label for the Operation
func (stop *BaseProvisionStopOperation) Label() string {
	return "Stop servers"
}

// return a multiline string description for the Operation
func (stop *BaseProvisionStopOperation) Description() string {
	return "Stop the servers for this project."
}

// return a multiline string man page for the Operation
func (stop *BaseProvisionStopOperation) Help() string {
	return ""
}

// Run a validation check on the Operation
func (stop *BaseProvisionStopOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
