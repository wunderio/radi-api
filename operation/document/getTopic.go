package document

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Retrieve the documentation for a single Documentation topic,
 * using the Documentation handler
 */

// Base class for documentation topic get Operation
type BaseDocumentTopicGetOperation struct{}

// Id the operation
func (get *BaseDocumentTopicGetOperation) Id() string {
	return "document.get"
}

// Label the operation
func (get *BaseDocumentTopicGetOperation) Label() string {
	return "Documentation topic get"
}

// Description for the operation
func (get *BaseDocumentTopicGetOperation) Description() string {
	return "List document topics."
}

// Man page for the operation
func (get *BaseDocumentTopicGetOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (get *BaseDocumentTopicGetOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
