package project

import (
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
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
	api_property.StringProperty
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
func (createSource *ProjectCreateSourceProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// Property for the create source
type ProjectInitDemoModeProperty struct {
	api_property.BooleanProperty
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
func (initDemo *ProjectInitDemoModeProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}
