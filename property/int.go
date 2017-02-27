package property

import (
	log "github.com/Sirupsen/logrus"
)

// A base Property that provides an Integer value
type IntProperty struct {
	value int
}

// Give an idea of what type of value the property consumes
func (property *IntProperty) Type() string {
	return "int"
}

// Property accessors
func (property *IntProperty) Get() interface{} {
	return interface{}(property.value)
}
func (property *IntProperty) Set(value interface{}) bool {
	if converted, ok := value.(int); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected int")
		return false
	}
}
