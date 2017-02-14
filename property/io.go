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
func (property *WriterProperty) Type() string {
	return "io.Writer"
}

func (property *WriterProperty) Get() interface{} {
	if property.value == nil {
		// writer := log.StandardLogger().Writer()
		// defer writer.Close()
		// property.value = io.Writer(writer)
		property.value = io.Writer(os.Stdout)
	}
	return interface{}(property.value)
}
func (property *WriterProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Writer); ok {
		property.value = converted
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
func (property *ReaderProperty) Type() string {
	return "io.Reader"
}

func (property *ReaderProperty) Get() interface{} {
	if property.value == nil {
		property.value = io.Reader(os.Stdin)
	}
	return interface{}(property.value)
}
func (property *ReaderProperty) Set(value interface{}) bool {
	if converted, ok := value.(io.Reader); ok {
		property.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected io.Reader")
		return false
	}
}
