package document

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * List all topics or subtopics using the Documentation Handler
 */

// Base class for documentation topic list Operation
type BaseDocumentTopicListOperation struct{}

// Id the operation
func (list *BaseDocumentTopicListOperation) Id() string {
	return "document.list"
}

// Label the operation
func (list *BaseDocumentTopicListOperation) Label() string {
	return "Documentation topic list"
}

// Description for the operation
func (list *BaseDocumentTopicListOperation) Description() string {
	return "List document topics."
}

// Man page for the operation
func (list *BaseDocumentTopicListOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (list *BaseDocumentTopicListOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
