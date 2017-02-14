package command

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * An operation for listing commands that are available
 * in the app
 */

const (
	OPERATION_ID_COMMAND_LIST = "command.list"
)

// Base class for command list Operation
type BaseCommandListOperation struct{}

// Id the operation
func (list *BaseCommandListOperation) Id() string {
	return "command.list"
}

// Label the operation
func (list *BaseCommandListOperation) Label() string {
	return "Command List"
}

// Description for the operation
func (list *BaseCommandListOperation) Description() string {
	return "List all available commands."
}

// Man page for the operation
func (list *BaseCommandListOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (list *BaseCommandListOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
