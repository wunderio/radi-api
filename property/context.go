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
func (prop *ContextProperty) Type() string {
	return "context.Context"
}

// Property accessors: Retrieve the context, or retrieve a Background context by default
func (prop *ContextProperty) Get() interface{} {
	if prop.value == nil {
		prop.value = context.Background()
	}

	return interface{}(prop.value)
}
func (prop *ContextProperty) Set(value interface{}) bool {
	if converted, ok := value.(context.Context); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected context/Context")
		return false
	}
}
