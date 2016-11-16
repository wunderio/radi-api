package builder

/**
 * This defines what a BuilderConfigWrapper must provide.
 *
 * This way different wrappers could be used to interpret
 * JSON or YML or whatever.
 */
type BuilderConfigWrapper interface {
	DefaultScope() string
	Get(key string) (BuildSetting, bool)
	Set(key string, values BuildSetting) bool
	List() []string
}
