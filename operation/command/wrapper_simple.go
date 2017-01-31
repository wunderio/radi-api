package command

import (
	"errors"

	// log "github.com/Sirupsen/logrus"

	"github.com/wunderkraut/radi-api/operation"
)

func New_SimpleCommandWrapper(operations *operation.Operations) *SimpleCommandWrapper {
	return &SimpleCommandWrapper{
		operations: operations,
	}
}

type SimpleCommandWrapper struct {
	operations *operation.Operations
}

func (wrapper *SimpleCommandWrapper) Get(key string) (Command, error) {
	var found bool
	var op operation.Operation
	var keyProp, commandProp operation.Property

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

	result := op.Exec(&props)
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

func (wrapper *SimpleCommandWrapper) List(parent string) ([]string, error) {
	var found bool
	var op operation.Operation
	var keyProp, keysProp operation.Property

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

	result := op.Exec(&props)
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
