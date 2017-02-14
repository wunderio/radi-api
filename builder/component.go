package builder

/**
 * The following 2 structs are used to keep track of settings as a string map,
 * but where each value knows from what config scope it was derived.
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
