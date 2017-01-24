package project

import (
	"github.com/james-nesbitt/radi-api/operation"
)

/**
 * Properties used for various project operations
 */

const (
	// property id for create source
	OPERATION_PROPERTY_PROJECT_CREATE_SOURCE = "project.create.source"
	// property id for init mode
	OPERATION_PROPERTY_PROJECT_INIT_DEMOMODE = "project.init.demomode"
)

// Property for the create source
type ProjectCreateSourceProperty struct {
	operation.StringProperty
}

// Id for the Property
func (createSource *ProjectCreateSourceProperty) Id() string {
	return OPERATION_PROPERTY_PROJECT_CREATE_SOURCE
}

// Label for the Property
func (createSource *ProjectCreateSourceProperty) Label() string {
	return "Create template source."
}

// Description for the Property
func (createSource *ProjectCreateSourceProperty) Description() string {
	return "Template source used to create the project."
}

// Is the Property internal only
func (createSource *ProjectCreateSourceProperty) Internal() bool {
	return false
}

// Property for the create source
type ProjectInitDemoModeProperty struct {
	operation.BooleanProperty
}

// Id for the Property
func (initDemo *ProjectInitDemoModeProperty) Id() string {
	return OPERATION_PROPERTY_PROJECT_INIT_DEMOMODE
}

// Label for the Property
func (initDemo *ProjectInitDemoModeProperty) Label() string {
	return "Init in demo mode."
}

// Description for the Property
func (initDemo *ProjectInitDemoModeProperty) Description() string {
	return "Initialization as a demo by including much more content."
}

// Is the Property internal only
func (initDemo *ProjectInitDemoModeProperty) Internal() bool {
	return false
}
