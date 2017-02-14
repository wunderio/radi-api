package project

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_PROJECT_CREATE = "project.create"
)

/**
 * Generate a new project
 */

// Generate a new project operation
type ProjectCreateOperation struct{}

// Id the operation
func (create *ProjectCreateOperation) Id() string {
	return OPERATION_ID_PROJECT_CREATE
}

// Label the operation
func (create *ProjectCreateOperation) Label() string {
	return "Create new project"
}

// Description for the operation
func (create *ProjectCreateOperation) Description() string {
	return "Create a new project from a templating source."
}

// Man page for the operation
func (create *ProjectCreateOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (create *ProjectCreateOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
