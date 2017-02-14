package provision

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Bring up the machines related to the project
 */

const (
	OPERATION_KEY_PROVISION_DOWN_OPERATION = "provision.down"
)

// Down Operation
type BaseProvisionDownOperation struct{}

// Return the string machinename/id of the Operation
func (down *BaseProvisionDownOperation) Id() string {
	return OPERATION_KEY_PROVISION_DOWN_OPERATION
}

// Return a user readable string label for the Operation
func (down *BaseProvisionDownOperation) Label() string {
	return "Remove servers"
}

// return a multiline string description for the Operation
func (down *BaseProvisionDownOperation) Description() string {
	return "Remove the servers for this project."
}

// return a multiline string man page for the Operation
func (down *BaseProvisionDownOperation) Help() string {
	return ""
}

// Run a validation check on the Operation
func (down *BaseProvisionDownOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
