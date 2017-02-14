package command

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
)

// A simple implementation of a CommandWrapper
type SimpleCommandWrapper struct {
	operations api_operation.Operations
}

// Constructor for SimpleCommandWrapper
func New_SimpleCommandWrapper(operations api_operation.Operations) *SimpleCommandWrapper {
	return &SimpleCommandWrapper{
		operations: operations,
	}
}

// Retrieve a command that matches a key
func (wrapper *SimpleCommandWrapper) Get(key string) (Command, error) {
	var found bool
	var op api_operation.Operation
	var keyProp, commandProp api_property.Property

	var comm Command

	if op, found = wrapper.operations.Get(OPERATION_ID_COMMAND_GET); !found {
		return comm, errors.New("No get operation available in Command Wrapper")
	}

	props := op.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_COMMAND_KEY); !found {
		return comm, errors.New("No key property available in Get operation in Command Wrapper")
	}

	if !keyProp.Set(key) {
		return comm, errors.New("Key property value failed to set in Command Wrapper")
	}

	if commandProp, found = props.Get(OPERATION_PROPERTY_COMMAND_COMMAND); !found {
		return comm, errors.New("No command property available in Get operation in Command Wrapper")
	}

	result := op.Exec(props)
	<-result.Finished()

	if !result.Success() {
		errs := result.Errors()
		if len(errs) == 0 {
			errors.New("Operation get failed to execute in Command Wrapper")
		} else {
			return comm, errs[0]
		}
	}

	comm = commandProp.Get().(Command)

	return comm, nil
}

// List all of the command keys available
func (wrapper *SimpleCommandWrapper) List(parent string) ([]string, error) {
	var found bool
	var op api_operation.Operation
	var keyProp, keysProp api_property.Property

	list := []string{}

	if op, found = wrapper.operations.Get(OPERATION_ID_COMMAND_LIST); !found {
		return list, errors.New("No list operation available in Command Wrapper")
	}

	props := op.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_COMMAND_KEY); !found {
		return list, errors.New("No key property available in Command Wrapper")
	}

	if !keyProp.Set(parent) {
		return list, errors.New("Key property value failed to set in Command Wrapper")
	}

	if keysProp, found = props.Get(OPERATION_PROPERTY_COMMAND_KEYS); !found {
		return list, errors.New("No keys property available in Command Wrapper")
	}

	result := op.Exec(props)
	<-result.Finished()

	if !result.Success() {
		errs := result.Errors()
		if len(errs) == 0 {
			return list, errors.New("Unknown error occured listing commands")
		} else {
			return list, errs[0]
		}
	}

	list = keysProp.Get().([]string)
	return list, nil
}
