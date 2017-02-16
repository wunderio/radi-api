package orchestrate

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Orchestration START - like docker-compose start
 *
 * Start all app containers, volumes and networks.
 */

const (
	OPERATION_ID_ORCHESTRATE_START = "orchestrate.start"
)

// Base class for orchestration Up Operation
type BaseOrchestrationStartOperation struct{}

// Id the operation
func (start *BaseOrchestrationStartOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_START
}

// Label the operation
func (start *BaseOrchestrationStartOperation) Label() string {
	return "Start"
}

// Description for the operation
func (start *BaseOrchestrationStartOperation) Description() string {
	return "This operation will start all containers, volumes and networks related to an application."
}

// Man page for the operation
func (start *BaseOrchestrationStartOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (start *BaseOrchestrationStartOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
