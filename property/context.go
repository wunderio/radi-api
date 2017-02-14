package property

import (
	"context"

	log "github.com/Sirupsen/logrus"
)

/**
 * Base properties for providing core context interfaces and structs
 */

// A base Property that provides an core context
type ContextProperty struct {
	value context.Context
}

// Give an idea of what type of value the property consumes
func (property *ContextProperty) Type() string {
	return "context.Context"
}

// Retrieve the context, or retrieve a Background context by default
func (property *ContextProperty) Get() interface{} {
	if property.value == nil {
		property.value = context.Background()
	}

	return interface{}(property.value)
}
func (property *ContextProperty) Set(value interface{}) bool {
	if converted, ok := value.(context.Context); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected context/Context")
		return false
	}
}
