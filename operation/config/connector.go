package config

/**
 * The config connector is an interface that allows any implementation
 * to provide full config operations, by providing the interface, and
 * then using the connector operations.
 *
 * The connector is not a wrapper
 *   - the wrapper wraps around multiple  operations to provide easy to use
 *     functions
 *   - the connector provbides easy to use methods from which to create multiple
 *     operations
 */

// Provide a Connector to Config sources
type ConfigConnector interface {
	// Get scoped key readers for a particular config key
	Readers(key string) ScopedReaders
	// Get scoped key writers for a particular config key
	Writers(key string) ScopedWriters
	// List all possible config keys, across all scopes
	List() []string
}

/**
 * Base operations for configs, primarily giving accesors for the connector
 */

// Constructor for BaseConfigConnectorOperation
func New_BaseConfigConnectorOperation(connector ConfigConnector) *BaseConfigConnectorOperation {
	return &BaseConfigConnectorOperation{
		connector: connector,
	}
}

// A Base config operation that provides a config connector
type BaseConfigConnectorOperation struct {
	connector ConfigConnector
}

// set the operation config connect
func (base *BaseConfigConnectorOperation) SetConnector(connector ConfigConnector) {
	base.connector = connector
}

// retrieve the operations config connnector
func (base *BaseConfigConnectorOperation) Connector() ConfigConnector {
	return base.connector
}
