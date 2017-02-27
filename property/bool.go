package property

import (
	log "github.com/Sirupsen/logrus"
)

/**
 * Base properties for providing booleans
 */

// A base Property that provides a Boolean value
type BooleanProperty struct {
	value bool
}

// Give an idea of what type of value the property consumes
func (prop *BooleanProperty) Type() string {
	return "bool"
}

// Property accessors
func (prop *BooleanProperty) Get() interface{} {
	return interface{}(prop.value)
}
func (prop *BooleanProperty) Set(value interface{}) bool {
	if converted, ok := value.(bool); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected bool")
		return false
	}
}
