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
	OPERATION_ID_MONITOR_LOG = "monitor.log"
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
	logTypeProp, _ := props.Get(OPERATION_PROPERTY_ID_MONITOR_LOG_TYPE)
	messageProp, _ := props.Get(OPERATION_PROPERTY_ID_MONITOR_LOG_MESSAGE)

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
