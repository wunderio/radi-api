package project

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
	return OPERATION_ID_PROJECT_CREATE
}

// Label the operation
func (create *ProjectInitOperation) Label() string {
	return "Initialize current project"
}

// Description for the operation
func (create *ProjectInitOperation) Description() string {
	return "Initialize current project as a kraut project, by adding kraut configuration files."
}

// Is this an internal API operation
func (create *ProjectInitOperation) Internal() bool {
	return false
}
