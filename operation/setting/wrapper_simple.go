package setting

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
)

/**
 * A simplified Settings wrapper, that performs blocking operations
 * using the operation.Operations.  This can be used to simplify
 * using the settings operations as a single struct.
 */

// A simple, blocking, Settings operations wrapper
type SimpleSettingWrapper struct {
	operations api_operation.Operations
}

// Constructor for SimpleSettingWrapper
func New_SimpleSettingWrapper(operations api_operation.Operations) *SimpleSettingWrapper {
	return &SimpleSettingWrapper{
		operations: operations,
	}
}

// Set a setting through the wrapped operations
func (wrapper *SimpleSettingWrapper) Get(key string) (string, error) {
	var found bool
	var getOp api_operation.Operation
	var keyProp, valueProp api_property.Property

	value := ""

	if getOp, found = wrapper.operations.Get(OPERATION_ID_SETTING_GET); !found {
		return value, errors.New("No get operation available in Setting Wrapper")
	}

	props := getOp.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return value, errors.New("No key property available in Get operation in Setting Wrapper")
	}

	if !keyProp.Set(key) {
		return value, errors.New("Key property value failed to set in Setting Wrapper")
	}

	if valueProp, found = props.Get(OPERATION_PROPERTY_SETTING_VALUE); !found {
		return value, errors.New("No value property available in Get operation in Setting Wrapper")
	}

	res := getOp.Exec(props)
	<-res.Finished()

	if !res.Success() {
		if errs := res.Errors(); len(errs) == 0 {
			return value, errors.New("Set operation failed to execute in Setting Wrapper")
		} else {
			return value, errs[0]
		}
	}

	return string(valueProp.Get().([]byte)), nil
}

// Retrieve a setting through the wrapped operations
func (wrapper *SimpleSettingWrapper) Set(key, value string) error {
	var found bool
	var setOp api_operation.Operation
	var keyProp, valueProp api_property.Property

	if setOp, found = wrapper.operations.Get(OPERATION_ID_SETTING_SET); !found {
		return errors.New("No set operation available in Setting Wrapper")
	}

	props := setOp.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return errors.New("No key property available in Set operation in Setting Wrapper")
	}

	if valueProp, found = props.Get(OPERATION_PROPERTY_SETTING_VALUE); !found {
		return errors.New("No value property available in Set operation in Setting Wrapper")
	}

	if !keyProp.Set(key) {
		return errors.New("Key property failed to set in Setting Wrapper")
	}

	if !valueProp.Set([]byte(value)) {
		return errors.New("Value property failed to set in Setting Wrapper")
	}

	result := setOp.Exec(props)
	<-result.Finished()

	if !result.Success() {
		if errs := result.Errors(); len(errs) == 0 {
			return errors.New("Get operation failed to execute in Setting Wrapper")
		} else {
			return errs[0]
		}
	}

	return nil
}

// List settins through the wrapped opertions
func (wrapper *SimpleSettingWrapper) List(parent string) ([]string, error) {
	var found bool
	var listOp api_operation.Operation
	var keyProp, keysProp api_property.Property

	list := []string{}

	if listOp, found = wrapper.operations.Get(OPERATION_ID_SETTING_LIST); !found {
		return list, errors.New("No list operation available in Setting Wrapper")
	}

	props := listOp.Properties()

	if keyProp, found = props.Get(OPERATION_PROPERTY_SETTING_KEY); !found {
		return list, errors.New("No Parent key property available in Get operation in Setting Wrapper")
	}

	if !keyProp.Set(parent) {
		return list, errors.New("Parent key property value failed to set in Setting Wrapper")
	}

	if keysProp, found = props.Get(OPERATION_PROPERTY_SETTING_KEYS); !found {
		return list, errors.New("No keys value property available in list operation in Setting Wrapper")
	}

	result := listOp.Exec(props)
	<-result.Finished()

	if !result.Success() {
		if errs := result.Errors(); len(errs) == 0 {
			return list, errors.New("List operation failed to execute in Setting Wrapper")
		} else {
			return list, errs[0]
		}
	}

	return keysProp.Get().([]string), nil
}
