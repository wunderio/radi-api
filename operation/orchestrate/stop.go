package orchestrate

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

// Is this an internal API operation
func (stop *BaseOrchestrationStopOperation) Internal() bool {
	return false
}
