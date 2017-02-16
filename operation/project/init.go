package project

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	OPERATION_ID_PROJECT_INIT = "project.init"
)

/**
 * Generate a new project
 */

// Generate a new project operation
type ProjectInitOperation struct{}

// Id the operation
func (create *ProjectInitOperation) Id() string {
	return OPERATION_ID_PROJECT_INIT
}

// Label the operation
func (create *ProjectInitOperation) Label() string {
	return "Initialize current project"
}

// Description for the operation
func (create *ProjectInitOperation) Description() string {
	return "Initialize current project as a radi project, by adding radi configuration files."
}

// Man page for the operation
func (create *ProjectInitOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (create *ProjectInitOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
