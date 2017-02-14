package provision

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Bring up the machines related to the project
 */

const (
	OPERATION_KEY_PROVISION_BACKUP_OPERATION = "provision.backup"
)

// Up Operation
type BaseProvisionBackupOperation struct{}

// Id the operation
func (backup *BaseProvisionBackupOperation) Id() string {
	return OPERATION_KEY_PROVISION_BACKUP_OPERATION
}

// Label the operation
func (backup *BaseProvisionBackupOperation) Label() string {
	return "Provision"
}

// Description for the operation
func (backup *BaseProvisionBackupOperation) Description() string {
	return "Provision machine elements needed for the project."
}

// Man page for the operation
func (backup *BaseProvisionBackupOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (backup *BaseProvisionBackupOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
