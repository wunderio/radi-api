package config

import (
	"errors"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
)

/**
 * A wrapper for config operations
 *
 * @NOTE this is currently a blocking inline process, which would stall
 *   if the backend operations timeout.  A thread-safe implementation should
 *   be written, but we should see this one in operation before we do that.
 *
 * @TODO Make this much more intelligent, right now it is just a quick operator
 */

// Constructor for SimpleConfigWrapper
func New_SimpleConfigWrapper(operations api_operation.Operations) *SimpleConfigWrapper {
	return &SimpleConfigWrapper{operations: operations}
}

// Simple config wrapper
type SimpleConfigWrapper struct {
	operations api_operation.Operations
}

// Perform the Get Operation
func (wrapper *SimpleConfigWrapper) Get(key string) (ConfigScopedValues, error) {
	var found bool
	var op api_operation.Operation
	var keyProp, readersProp api_property.Property

	values := ConfigScopedValues{}

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_READERS); !found {
		return values, errors.New("No get operation available in Config Simple Wrapper")
	}

	props := op.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return values, errors.New("No key configuraiton available in Config Simple Wrapper")
	}

	if !keyProp.Set(key) {
		return values, errors.New("Key property value failed to set in Config Simple Wrapper")
	}

	if readersProp, found = props.Get(OPERATION_PROPERTY_CONFIG_VALUE_READERS); !found {
		return values, errors.New("No value property available in Config Simple Wrapper")
	}

	// Execute the wrapped operation, and wait for it to finish
	result := op.Exec(props)
	<-result.Finished()

	if !result.Success() {
		if errs := result.Errors(); len(errs) > 0 {
			return values, errs[0]
		} else {
			return values, errors.New("Unknown error prevented Get execution in Config Simple Wrapper")
		}
	}

	readers := readersProp.Get().(ScopedReaders)
	for _, scope := range readers.Order() {
		reader, _ := readers.Get(scope)
		if contents, err := ioutil.ReadAll(reader); err == nil {
			values.Set(scope, ConfigScopedValue(contents))
		}
	}
	return values, nil
}

// Perform the Set Operation
func (wrapper *SimpleConfigWrapper) Set(key string, values ConfigScopedValues) error {
	var found bool
	var op api_operation.Operation
	var keyProp, writersProp api_property.Property

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_WRITERS); !found {
		return errors.New("No get operation available in Config Simple Wrapper")
	}

	props := op.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return errors.New("No key configuraiton available in Config Simple Wrapper")
	}

	if !keyProp.Set(key) {
		return errors.New("Key property value failed to set in Config Simple Wrapper")
	}

	if writersProp, found = props.Get(OPERATION_PROPERTY_CONFIG_VALUE_WRITERS); !found {
		return errors.New("No writers property available in Config Simple Wrapper")
	}

	// Execute the wrapped operation, and wait for it to finish
	result := op.Exec(props)
	<-result.Finished()

	if !result.Success() {
		if errs := result.Errors(); len(errs) > 0 {
			return errs[0]
		} else {
			return errors.New("Unknown error prevented Set execution in Config Simple Wrapper")
		}
	}

	writers := writersProp.Get().(ScopedWriters)

	var returnError error
	for _, scope := range values.Order() {
		content, _ := values.Get(scope)
		if writer, found := writers.Get(scope); found {
			log.WithFields(log.Fields{"scope": scope, "content": string(content)}).Debug("ConfigWrapper: writing config to writer")
			byteContent := []byte(content)
			if _, err := writer.Write(byteContent); err != nil {
				returnError = err
			}
		} else {
			log.WithFields(log.Fields{"scope": scope, "scopes": writers.Order()}).Error("ConfigWrapper could not find wrapper for targeted Config Set")
		}

		/**
		 * @TODO should we allow an attempt to create a new writer?
		 */
	}

	return returnError
}

// Performe the List Operation
func (wrapper *SimpleConfigWrapper) List(parent string) ([]string, error) {
	var found bool
	var op api_operation.Operation
	var keyProp, keysProp api_property.Property

	list := []string{}

	if op, found = wrapper.operations.Get(OPERATION_ID_CONFIG_LIST); !found {
		return list, errors.New("No list operation available in Config Wrapper")
	}

	props := op.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_CONFIG_KEY); !found {
		return list, errors.New("No key property available in Config Wrapper")
	}

	if !keyProp.Set(parent) {
		return list, errors.New("Key property value failed to set in Config Wrapper")
	}

	if keysProp, found = props.Get(OPERATION_PROPERTY_CONFIG_KEYS); !found {
		return list, errors.New("No keys property available in Config Wrapper")
	}

	result := op.Exec(props)
	<-result.Finished()

	if !result.Success() {
		if errs := result.Errors(); len(errs) > 0 {
			return list, errs[0]
		} else {
			return list, errors.New("Unknown error prevented List execution in Config Simple Wrapper")
		}
	}

	list = keysProp.Get().([]string)
	return list, nil
}
