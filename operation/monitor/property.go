package monitor

import (
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_PROPERTY_ID_MONITOR_WRITER      = "monitor.output.writer"
	OPERATION_PROPERTY_ID_MONITOR_LOG_TYPE    = "monitor.log.type"
	OPERATION_PROPERTY_ID_MONITOR_LOG_MESSAGE = "monitor.log.message"
)

// Configuration for an outputter for monitoring
type MonitorOutputProperty struct {
	api_property.WriterProperty
}

// Id for the configuration
func (output *MonitorOutputProperty) Id() string {
	return OPERATION_PROPERTY_ID_MONITOR_WRITER
}

// Label for the configuration
func (output *MonitorOutputProperty) Label() string {
	return "Output handler for the monitor"
}

// Description for the operation
func (output *MonitorOutputProperty) Description() string {
	return "Attach an io.Writer to the operation, and it will be used to capture the output.  By default, the output will go to log."
}

// Is the Property internal only
func (output *MonitorOutputProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (output *MonitorOutputProperty) Copy() api_property.Property {
	prop := &MonitorOutputProperty{}
	prop.Set(output.Get())
	return api_property.Property(prop)
}

// property for a monitoring log status : error|info
type MonitorLogTypeProperty struct {
	api_property.StringProperty
}

// Constructor for MonitorLogTypeProperty
func NewMonitorLogTypeProperty(logType string) *MonitorLogTypeProperty {
	conf := MonitorLogTypeProperty{}
	conf.Set(logType)
	return &conf
}

// Id for the property
func (logType *MonitorLogTypeProperty) Property() api_property.Property {
	return api_property.Property(logType)
}

// Id for the property
func (logType *MonitorLogTypeProperty) Id() string {
	return OPERATION_PROPERTY_ID_MONITOR_LOG_TYPE
}

// Label for the property
func (logType *MonitorLogTypeProperty) Label() string {
	return "Message type."
}

// Description for the property
func (logType *MonitorLogTypeProperty) Description() string {
	return "Message type, which can be either info or error."
}

// Man page for the property
func (logType *MonitorLogTypeProperty) Help() string {
	return ""
}

// Is the Property internal only
func (logType *MonitorLogTypeProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// Copy the property
func (logType *MonitorLogTypeProperty) Copy() api_property.Property {
	prop := &MonitorLogTypeProperty{}
	prop.Set(logType.Get())
	return api_property.Property(prop)
}

// property for a monitoring log message
type MonitorLogMessageProperty struct {
	api_property.StringProperty
}

// Id for the property
func (message *MonitorLogMessageProperty) Id() string {
	return OPERATION_PROPERTY_ID_MONITOR_LOG_MESSAGE
}

// Label for the property
func (message *MonitorLogMessageProperty) Label() string {
	return "Message to be logged."
}

// Description for the property
func (message *MonitorLogMessageProperty) Description() string {
	return "Message which will be sent to the standard logger."
}

// Is the Property internal only
func (message *MonitorLogMessageProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// Copy the property
func (message *MonitorLogMessageProperty) Copy() api_property.Property {
	prop := &MonitorLogMessageProperty{}
	prop.Set(message.Get())
	return api_property.Property(prop)
}
