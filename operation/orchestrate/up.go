package orchestrate

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Orchestration UP - like docker-compose up
 *
 * Bring up all app containers, volumes and networks.
 */

const (
	OPERATION_ID_ORCHESTRATE_UP = "orchestrate.up"
)

// Base class for orchestration Up Operation
type BaseOrchestrationUpOperation struct{}

// Id the operation
func (up *BaseOrchestrationUpOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_UP
}

// Label the operation
func (up *BaseOrchestrationUpOperation) Label() string {
	return "Up"
}

// Description for the operation
func (up *BaseOrchestrationUpOperation) Description() string {
	return "This operation will bring up all containers, volumes and networks related to an application."
}

// Man page for the operation
func (up *BaseOrchestrationUpOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (up *BaseOrchestrationUpOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
