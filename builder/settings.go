package builder

import (
	"errors"
)

// An abstracted settings provider that can unmarshal to a target
type SettingsProvider interface {
	AssignSettings(target interface{}) error
}

/**
 * The following 2 structs are used to keep track of settings
 * as a string map, but where each value knows from what config
 * scope it was derived.
 */

// Constructor for BuildComponent
func New_BuildComponent(buildType string, implementations Implementations, settingsProvider SettingsProvider) *BuildComponent {
	return &BuildComponent{
		Type:             buildType,
		Implementations:  implementations,
		SettingsProvider: settingsProvider,
	}
}

// A single Builder instruction set.
type BuildComponent struct {
	Type             string           // which builder should be used
	Implementations  Implementations  // which implementations the builder should activate (strings that the builder will understand)
	SettingsProvider SettingsProvider // this will provide abstract settings that the builder should understand
}

// An ordered list of BuildComponents
type BuildComponents struct {
	buildComponents map[string]BuildComponent
	order           []string
}

// safe intitializer
func (buildComponents *BuildComponents) safe() {
	if buildComponents.order == nil {
		buildComponents.buildComponents = map[string]BuildComponent{}
		buildComponents.order = []string{}
	}
}

// Add a builder
func (buildComponents *BuildComponents) Set(key string, builderSetting BuildComponent) error {
	buildComponents.safe()
	if _, exists := buildComponents.buildComponents[key]; !exists {
		buildComponents.order = append(buildComponents.order, key)
	}
	buildComponents.buildComponents[key] = builderSetting
	return nil
}

// Get a single builder
func (buildComponents *BuildComponents) Get(key string) (BuildComponent, error) {
	buildComponents.safe()
	if builderSetting, found := buildComponents.buildComponents[key]; found {
		return builderSetting, nil
	} else {
		return builderSetting, errors.New("No such builderler found")
	}
}

// Get the builder ordered keys
func (buildComponents *BuildComponents) Order() []string {
	buildComponents.safe()
	return buildComponents.order
}

// Does this list have any items
func (buildComponents *BuildComponents) Empty() bool {
	return (buildComponents.buildComponents == nil) || (len(buildComponents.buildComponents) == 0)
}
