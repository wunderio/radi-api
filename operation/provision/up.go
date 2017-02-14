package provision

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Bring up the machines related to the project
 */

const (
	OPERATION_KEY_PROVISION_UP_OPERATION = "provision.up"
)

// Up Operation
type BaseProvisionUpOperation struct{}

// Id the operation
func (up *BaseProvisionUpOperation) Id() string {
	return OPERATION_KEY_PROVISION_UP_OPERATION
}

// Label the operation
func (up *BaseProvisionUpOperation) Label() string {
	return "Provision"
}

// Description for the operation
func (up *BaseProvisionUpOperation) Description() string {
	return "Provision machine elements needed for the project."
}

// Man page for the operation
func (up *BaseProvisionUpOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (up *BaseProvisionUpOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
