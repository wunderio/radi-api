package orchestrate

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Orchestration DOWN - like docker-compose down
 *
 * Bring down all app containers, and remove them an all
 * related volumes and networks.
 */

const (
	OPERATION_ID_ORCHESTRATE_DOWN = "orchestrate.down"
)

// Base class for orchestration Down Operation
type BaseOrchestrationDownOperation struct{}

// Id the operation
func (down *BaseOrchestrationDownOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_DOWN
}

// Label the operation
func (down *BaseOrchestrationDownOperation) Label() string {
	return "Down"
}

// Description for the operation
func (down *BaseOrchestrationDownOperation) Description() string {
	return "This operation will bring down all containers, volumes and networks related to an application."
}

// Man page for the operation
func (down *BaseOrchestrationDownOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (down *BaseOrchestrationDownOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
