package provision

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

// Is this operation meant to be yused internally only
func (down *BaseProvisionDownOperation) Internal() bool {
	return false
}

// Run a validation check on the Operation
func (down *BaseProvisionDownOperation) Validate() bool {
	return true
}
