package builder

import (
	"errors"
)

/**
 * The following 2 structs are used to keep track of settings
 * as a string map, but where each value knows from what config
 * scope it was derived.
 */

// A single Builder instruction set.
type ProjectComponent struct {
	Type             string           // which builder should be used
	Implementations  Implementations  // which implementations the builder should activate (strings that the builder will understand)
	SettingsProvider SettingsProvider // this will provide abstract settings that the builder should understand
}

// Constructor for ProjectComponent
func New_ProjectComponent(buildType string, implementations Implementations, settingsProvider SettingsProvider) *ProjectComponent {
	return &ProjectComponent{
		Type:             buildType,
		Implementations:  implementations,
		SettingsProvider: settingsProvider,
	}
}

// An ordered list of ProjectComponents
type ProjectComponents struct {
	components map[string]ProjectComponent
	order           []string
}

// safe intitializer
func (projectComponents *ProjectComponents) safe() {
	if projectComponents.order == nil {
		projectComponents.components = map[string]ProjectComponent{}
		projectComponents.order = []string{}
	}
}

// Add a builder
func (projectComponents *ProjectComponents) Set(key string, builderSetting ProjectComponent) error {
	projectComponents.safe()
	if _, exists := projectComponents.components[key]; !exists {
		projectComponents.order = append(projectComponents.order, key)
	}
	projectComponents.components[key] = builderSetting
	return nil
}

// Get a single builder
func (projectComponents *ProjectComponents) Get(key string) (ProjectComponent, error) {
	projectComponents.safe()
	if builderSetting, found := projectComponents.components[key]; found {
		return builderSetting, nil
	} else {
		return builderSetting, errors.New("No such builderler found")
	}
}

// Get the builder ordered keys
func (projectComponents *ProjectComponents) Order() []string {
	projectComponents.safe()
	return projectComponents.order
}

// Does this list have any items
func (projectComponents *ProjectComponents) Empty() bool {
	return (projectComponents.components == nil) || (len(projectComponents.components) == 0)
}
