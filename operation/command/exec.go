package command

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_COMMAND_EXEC = "command.exec"
)

/**
 * Execute a command using the command handler
 */

// Base class for command exec Operation
type BaseCommandExecOperation struct{}

// Id the operation
func (exec *BaseCommandExecOperation) Id() string {
	return "command.exec"
}

// Label the operation
func (exec *BaseCommandExecOperation) Label() string {
	return "Command Exec"
}

// Description for the operation
func (exec *BaseCommandExecOperation) Description() string {
	return "Execute a specified command.  This is an abstract command executor, but commands should probably add their own operations (@TODO)."
}

// Map page for the operation
func (exec *BaseCommandExecOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (exec *BaseCommandExecOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
