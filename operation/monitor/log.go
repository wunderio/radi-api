package monitor

import (
	log "github.com/Sirupsen/logrus"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_result "github.com/wunderkraut/radi-api/result"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Write messages to the log
 */

const (
	OPERATION_ID_MONITOR_LOG                    = "monitor.log"
	OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE    = "monitor.log.type"
	OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE = "monitor.log.message"
)

// Base class for monitor log Operation
type BaseMonitorLogOperation struct {
	MonitorBaseWriterOperation
}

// Id the operation
func (logger *BaseMonitorLogOperation) Id() string {
	return OPERATION_ID_MONITOR_LOG
}

// Label the operation
func (logger *BaseMonitorLogOperation) Label() string {
	return "Log a message"
}

// Description for the operation
func (logger *BaseMonitorLogOperation) Description() string {
	return "Log a message."
}

// Man page for the operation
func (logger *BaseMonitorLogOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (logger *BaseMonitorLogOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}

// Is the log request valid
func (logger *BaseMonitorLogOperation) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// Standard output logger
type MonitorStandardLogOperation struct{}

// Id the operation
func (logger *MonitorStandardLogOperation) Id() string {
	return OPERATION_ID_MONITOR_LOG
}

// Label the operation
func (logger *MonitorStandardLogOperation) Label() string {
	return "Log a message"
}

// Description for the operation
func (logger *MonitorStandardLogOperation) Description() string {
	return "Log a message."
}

// Man page for the operation
func (logger *MonitorStandardLogOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (logger *MonitorStandardLogOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}

// Is the log request valid
func (logger *MonitorStandardLogOperation) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// Add a Message propery
func (logger *MonitorStandardLogOperation) Properties() api_property.Properties {
	props := api_property.New_SimplePropertiesEmpty()

	props.Add(api_property.Property(NewMonitorLogTypeProperty("info")))
	props.Add(api_property.Property(&MonitorLogMessageProperty{}))

	return props.Properties()
}

// Exec the log output
func (logger *MonitorStandardLogOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.New_StandardResult()

	// we ignore the conf tests, as we ensured that the conf would exist in the property() method
	logTypeProp, _ := props.Get(OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE)
	messageProp, _ := props.Get(OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE)

	var logType, message string
	var ok bool

	if message, ok = messageProp.Get().(string); !ok {
		log.Error("MonitorStandardLogOperation has no message assigned")
	} else {
		if logType, ok = logTypeProp.Get().(string); !ok {
			logType = "info"
		}
		switch logType {
		case "error":
			log.Error(message)

		case "info":
			fallthrough
		default:
			log.Info(message)

		}
	}

	res.MarkSuccess()
	res.MarkFinished()

	return res.Result()
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
	return OPERATION_PROPERTY_CONF_MONITOR_LOG_TYPE
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

// property for a monitoring log message
type MonitorLogMessageProperty struct {
	api_property.StringProperty
}

// Id for the property
func (message *MonitorLogMessageProperty) Id() string {
	return OPERATION_PROPERTY_CONF_MONITOR_LOG_MESSAGE
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
