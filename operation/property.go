package operation

import (
	log "github.com/Sirupsen/logrus"
)

/**
 * A github.com/wunderkraut/radi-api/property.Property that holds
 * an operation.  This is used for decorators.
 */

// A base Property that provides an operation
type OperationProperty struct {
	value Operation
}

// Give an idea of what type of value the property consumes
func (op *OperationProperty) Type() string {
	return "github.com/wunderkraut/radi-api/operation.Operation"
}

// Retrieve the context, or retrieve a Background context by default
func (op *OperationProperty) Get() interface{} {
	return interface{}(op.value)
}
func (op *OperationProperty) Set(value interface{}) bool {
	if converted, ok := value.(Operation); ok {
		op.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected an github.com/wunderkraut/radi-api/operation.Operation")
		return false
	}
}
