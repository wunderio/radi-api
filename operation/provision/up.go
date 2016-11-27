package provision

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

// Is this an internal API operation
func (up *BaseProvisionUpOperation) Internal() bool {
	return false
}
