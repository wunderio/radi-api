package provision

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

// Is this operation meant to be yused internally only
func (stop *BaseProvisionStopOperation) Internal() bool {
	return false
}

// Run a validation check on the Operation
func (stop *BaseProvisionStopOperation) Validate() bool {
	return true
}
