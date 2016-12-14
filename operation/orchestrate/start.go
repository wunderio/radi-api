package orchestrate

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

// Is this an internal API operation
func (start *BaseOrchestrationStartOperation) Internal() bool {
	return false
}
