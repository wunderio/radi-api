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

// Constructor for BuildSetting
func New_BuildSetting(buildType string, implementations Implementations, settingsProvider SettingsProvider) *BuildSetting {
	return &BuildSetting{
		Type:             buildType,
		Implementations:  implementations,
		SettingsProvider: settingsProvider,
	}
}

// A single Builder instruction set.
type BuildSetting struct {
	Type             string           // which builder should be used
	Implementations  Implementations  // which implementations the builder should activate (strings that the builder will understand)
	SettingsProvider SettingsProvider // this will provide abstract settings that the builder should understand
}

// An ordered list of BuildSettings
type BuildSettings struct {
	builderSettings map[string]BuildSetting
	order           []string
}

// safe intitializer
func (builderSettings *BuildSettings) safe() {
	if builderSettings.order == nil {
		builderSettings.builderSettings = map[string]BuildSetting{}
		builderSettings.order = []string{}
	}
}

// Add a builder
func (builderSettings *BuildSettings) Set(key string, builderSetting BuildSetting) error {
	builderSettings.safe()
	if _, exists := builderSettings.builderSettings[key]; !exists {
		builderSettings.order = append(builderSettings.order, key)
	}
	builderSettings.builderSettings[key] = builderSetting
	return nil
}

// Get a single builder
func (builderSettings *BuildSettings) Get(key string) (BuildSetting, error) {
	builderSettings.safe()
	if builderSetting, found := builderSettings.builderSettings[key]; found {
		return builderSetting, nil
	} else {
		return builderSetting, errors.New("No such builderler found")
	}
}

// Get the builder ordered keys
func (builderSettings *BuildSettings) Order() []string {
	builderSettings.safe()
	return builderSettings.order
}

// Does this list have any items
func (builderSettings *BuildSettings) Empty() bool {
	return (builderSettings.builderSettings == nil) || (len(builderSettings.builderSettings) == 0)
}
