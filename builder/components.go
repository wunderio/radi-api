package builder

import (
	"errors"
)

// An ordered list of ProjectComponents
type ProjectComponents struct {
	components map[string]ProjectComponent
	order      []string
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
