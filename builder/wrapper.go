package builder

/**
 * This defines what a BuilderConfigWrapper must provide.
 *
 * This way different wrappers could be used to interpret
 * JSON or YML or whatever.
 */
type ProjectConfigWrapper interface {
	// Determine the default project builder config scope
	DefaultScope() string
	// Get the project components for a particular key
	Get(key string) (ProjectComponent, bool)
	Set(key string, values ProjectComponent) bool
	List() []string
}
