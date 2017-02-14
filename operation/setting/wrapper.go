package setting

/**
 * An easy to use wrapper around the settings operations to provide
 * a more functional approach to retrieving settings
 */

type SettingWrapper interface {
	// Get a setting value from the wrapper
	Get(key string) (string, error)
	// Set a setting through the wrapper
	Set(key, value string) error
	// List all settings through the wrapper
	List(parent string) ([]string, error)
}
