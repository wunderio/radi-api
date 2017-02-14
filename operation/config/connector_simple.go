package config

import (
	"errors"

	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
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
	BaseConfigConnectorOperation
}

// Validate the operation
func (readers ConfigSimpleConnectorReadersOperation) Validate() api_result.Result {
	// @TODO write real validation
	return api_result.MakeSuccessfulResult()
}

// Execute the operation
func (readers ConfigSimpleConnectorReadersOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	readersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_READERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		readersValue := readers.Connector().Readers(key)
		readersProp.Set(readersValue)
		res.MarkSuccess()
		if len(readersValue.Order()) == 0 {
			res.AddError(errors.New("No config found for requested key"))
		}
	} else {
		res.AddError(errors.New("Invalid config key requested"))
		res.MarkFailed()
	}

	res.MarkFinished()

	return res.Result()
}

// Config Set operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorWritersOperation struct {
	BaseConfigWritersOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (writers ConfigSimpleConnectorWritersOperation) Validate() api_result.Result {
	// @TODO write real validation
	return api_result.MakeSuccessfulResult()
}

// Execute the operation
func (writers ConfigSimpleConnectorWritersOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	writersProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_VALUE_WRITERS)

	if key, ok := keyProp.Get().(string); ok && key != "" {
		if writersValue := writers.Connector().Writers(key); len(writersValue.Order()) > 0 {
			writersProp.Set(writersValue)
			res.MarkSuccess()
		} else {
			res.MarkFailed()
			res.AddError(errors.New("Unknown config key requested"))
		}
	} else {
		res.MarkFailed()
		res.AddError(errors.New("Invalid config key for config writers"))
	}

	res.MarkFinished()

	return res.Result()
}

// Config List operation that relies on a ConfigConnector for an io.Writer
type ConfigSimpleConnectorListOperation struct {
	BaseConfigListOperation
	BaseConfigConnectorOperation
}

// Validate the operation
func (list ConfigSimpleConnectorListOperation) Validate() api_result.Result {
	// @TODO write real validation
	return api_result.MakeSuccessfulResult()
}

// Execute the operation
func (list ConfigSimpleConnectorListOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	keyProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEY)
	keysProp, _ := props.Get(OPERATION_PROPERTY_CONFIG_KEYS)

	if key, ok := keyProp.Get().(string); ok || key == "" {
		// @TODO Currently the key proerty doesn't do anything, shouldn't it?

		if list := list.Connector().List(); len(list) > 0 {
			keysProp.Set(list)
			res.MarkSuccess()
		} else {
			res.MarkFailed()
			res.AddError(errors.New("Config has no keys"))
		}
	} else {
		res.MarkFailed()
		res.AddError(errors.New("Invalid config parent key provided for config list"))
	}

	res.MarkFinished()

	return res.Result()
}
