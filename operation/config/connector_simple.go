package config

import (
	"errors"

	"github.com/wunderkraut/radi-api/operation"
)

/**
 * The Simple config connector operations are config operations that use
 * a connector is the simplest way, by directly calling the operation and
 * waiting for a response.
 *
 * @NOTE these operations can result in blocking of the connector stalls
 *   or fails.  This means that either we rely on an advanced connector
 *   which can itself time out, or we write more advanced operations.
 */

// Config Get operation that relies on a ConfigConnector for an io.Reader
type ConfigSimpleConnectorReadersOperation struct {
	BaseConfigReadersOperation
	BaseConfigKeyReadersOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (readers ConfigSimpleConnectorReadersOperation) Validate() bool {
	return true
}

// Execute the operation
func (readers ConfigSimpleConnectorReadersOperation) Exec(props *operation.Properties) operation.Result {
	result := operation.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	readersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_READERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		readersValue := readers.Connector().Readers(key)
		readersProp.Set(readersValue)
		result.MarkSuccess()
		if len(readersValue.Order()) == 0 {
			result.AddError(errors.New("No config found for requested key"))
		}
	} else {
		result.AddError(errors.New("Invalid config key requested"))
		result.MarkFailed()
	}

	result.MarkFinished()

	return operation.Result(result)
}

// Config Set operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorWritersOperation struct {
	BaseConfigWritersOperation
	BaseConfigKeyWritersOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (writers ConfigSimpleConnectorWritersOperation) Validate() bool {
	return true
}

// Execute the operation
func (writers ConfigSimpleConnectorWritersOperation) Exec(props *operation.Properties) operation.Result {
	result := operation.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	writersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_WRITERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		if writersValue := writers.Connector().Writers(key); len(writersValue.Order()) > 0 {
			writersProp.Set(writersValue)
			result.MarkSuccess()
		} else {
			result.MarkFailed()
			result.AddError(errors.New("Unknown config key requested"))
		}
	} else {
		result.MarkFailed()
		result.AddError(errors.New("Invalid config key for config writers"))
	}

	result.MarkFinished()

	return operation.Result(result)
}

// Config List operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorListOperation struct {
	BaseConfigListOperation
	BaseConfigKeyKeysOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (list ConfigSimpleConnectorListOperation) Validate() bool {
	return true
}

// Execute the operation
func (list ConfigSimpleConnectorListOperation) Exec(props *operation.Properties) operation.Result {
	result := operation.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	keysProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEYS)

	if key, ok := keyProp.Get().(string); ok || key == "" {
		if list := list.Connector().List(); len(list) > 0 {
			keysProp.Set(list)
			result.MarkSuccess()
		} else {
			result.MarkFailed()
			result.AddError(errors.New("Config has no keys"))
		}
	} else {
		result.MarkFailed()
		result.AddError(errors.New("Invalid config parent key provided for config list"))
	}

	result.MarkFinished()

	return operation.Result(result)
}
