package monitor

import (
	"io"

	log "github.com/Sirupsen/logrus"

	api_property "github.com/wunderkraut/radi-api/property"
)

// A handler base that writes to an outputter
type MonitorBaseWriterOperation struct{}

// A utility function to write a message to the configured writer
func (op *MonitorBaseWriterOperation) WriteMessage(message string) bool {
	props := op.Properties()

	if writerConfig, exists := props.Get(OPERATION_PROPERTY_ID_MONITOR_WRITER); exists {
		confValue := writerConfig.Get()
		if writer, ok := confValue.(io.Writer); ok {
			writer.Write([]byte(message))
			return true
		} else {
			log.WithFields(log.Fields{"writer": writer}).Warning("Could not write status, as the output configuration contains an invalid writer.")
		}
	}
	return false
}

// Add a writer configuration
func (op *MonitorBaseWriterOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(&MonitorOutputProperty{}))

	return props.Properties()
}
