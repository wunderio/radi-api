package builder

/**
 * This defines what a BuilderConfigWrapper must provide.
 *
 * This way different wrappers could be used to interpret
 * JSON or YML or whatever.
 */
type ProjectConfigWrapper interface {
	DefaultScope() string
	Get(key string) (ProjectComponent, bool)
	Set(key string, values ProjectComponent) bool
	List() []string
}
