package property

import (
	log "github.com/Sirupsen/logrus"
)

// A base Property that provides a String value
type StringProperty struct {
	value string
}

// Give an idea of what type of value the property consumes
func (prop *StringProperty) Type() string {
	return "string"
}

// Property accessors
func (prop *StringProperty) Get() interface{} {
	return interface{}(prop.value)
}
func (prop *StringProperty) Set(value interface{}) bool {
	if converted, ok := value.(string); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected string")
		return false
	}
}

// A base Property that provides a slice of string values
type StringSliceProperty struct {
	value []string
}

// Give an idea of what type of value the property consumes
func (prop *StringSliceProperty) Type() string {
	return "[]string"
}

// Property accessors
func (prop *StringSliceProperty) Get() interface{} {
	return interface{}(prop.value)
}
func (prop *StringSliceProperty) Set(value interface{}) bool {
	if converted, ok := value.([]string); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected []string")
		return false
	}
}

// A base Property that provides a Bytes Array value
type BytesArrayProperty struct {
	value []byte
}

// Give an idea of what type of value the property consumes
func (prop *BytesArrayProperty) Type() string {
	return "[]byte"
}

// Property accessors
func (prop *BytesArrayProperty) Get() interface{} {
	return interface{}(prop.value)
}
func (prop *BytesArrayProperty) Set(value interface{}) bool {
	if converted, ok := value.([]byte); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected []byte")
		return false
	}
}
