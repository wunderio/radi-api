package builder

/**
 * This defines what a BuilderConfigWrapper must provide.
 *
 * This way different wrappers could be used to interpret
 * JSON or YML or whatever.
 */
type BuilderConfigWrapper interface {
	DefaultScope() string
	Get(key string) (BuildComponent, bool)
	Set(key string, values BuildComponent) bool
	List() []string
}
