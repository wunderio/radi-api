package property

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

/**
 * Base properties for io package interfaces/structs
 */

// A base Property that provides an IO.Writer
type WriterProperty struct {
	value io.Writer
}

// Give an idea of what type of value the property consumes
func (prop *WriterProperty) Type() string {
	return "io.Writer"
}

// Property accessors
func (prop *WriterProperty) Get() interface{} {
	if prop.value == nil {
		prop.value = io.Writer(os.Stdout)
	}
	return interface{}(prop.value)
}
func (prop *WriterProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Writer); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected io.Writer")
		return false
	}
}

// A base Property that provides an IO.Reader
type ReaderProperty struct {
	value io.Reader
}

// Give an idea of what type of value the property consumes
func (prop *ReaderProperty) Type() string {
	return "io.Reader"
}

// Property accessors
func (prop *ReaderProperty) Get() interface{} {
	if prop.value == nil {
		prop.value = io.Reader(os.Stdin)
	}
	return interface{}(prop.value)
}
func (prop *ReaderProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Reader); ok {
		prop.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected io.Reader")
		return false
	}
}
