package monitor

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Status Operations are meant to return long human readable
 * messages, about the state of the app
 */

const (
	OPERATION_ID_MONITOR_INFO = "monitor.info"
)

// Base class for monitor info Operation
type BaseMonitorInfoOperation struct {
	MonitorBaseWriterOperation
}

// Id the operation
func (info *BaseMonitorInfoOperation) Id() string {
	return OPERATION_ID_MONITOR_INFO
}

// Label the operation
func (info *BaseMonitorInfoOperation) Label() string {
	return "Information"
}

// Description for the operation
func (info *BaseMonitorInfoOperation) Description() string {
	return "App information."
}

// Man page for the operation
func (info *BaseMonitorInfoOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (info *BaseMonitorInfoOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
