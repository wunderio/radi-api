package command

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * A command operation to retrieve a command object
 */

const (
	OPERATION_ID_COMMAND_GET = "command.get"
)

/**
 * Execute a command using the command handler
 */

// Base class for command get Operation
type BaseCommandGetOperation struct{}

// Id the operation
func (get *BaseCommandGetOperation) Id() string {
	return OPERATION_ID_COMMAND_GET
}

// Label the operation
func (get *BaseCommandGetOperation) Label() string {
	return "Command Get"
}

// Description for the operation
func (get *BaseCommandGetOperation) Description() string {
	return "Retrieve a specified command.  Retrieve a command object."
}

// Man page for the operation
func (get *BaseCommandGetOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (get *BaseCommandGetOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
