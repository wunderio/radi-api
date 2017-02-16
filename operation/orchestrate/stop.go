package orchestrate

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Orchestration START - like docker-compose stop
 *
 * Stop all app containers, volumes and networks.
 */

const (
	OPERATION_ID_ORCHESTRATE_STOP = "orchestrate.stop"
)

// Base class for orchestration Up Operation
type BaseOrchestrationStopOperation struct{}

// Id the operation
func (stop *BaseOrchestrationStopOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_STOP
}

// Label the operation
func (stop *BaseOrchestrationStopOperation) Label() string {
	return "Stop"
}

// Description for the operation
func (stop *BaseOrchestrationStopOperation) Description() string {
	return "This operation will stop all containers, volumes and networks related to an application."
}

// Man page for the operation
func (stop *BaseOrchestrationStopOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (stop *BaseOrchestrationStopOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
