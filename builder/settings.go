package builder

/**
 * This file containers an interface designed to abstract how settings from
 * a config source could be applied to an unknown struct.  This follows the
 * model of the golang yaml model where a function is provided that can apply
 * settings to a passed target interface.
 *
 * This allows the building process to pass unpredictable settings on to a
 * handlers builder.
 */

// An abstracted settings provider that can unmarshal to a target
type SettingsProvider interface {
	AssignSettings(target interface{}) error
}
