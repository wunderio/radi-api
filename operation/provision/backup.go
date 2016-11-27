package provision

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

// Is this an internal API operation
func (backup *BaseProvisionBackupOperation) Internal() bool {
	return false
}
