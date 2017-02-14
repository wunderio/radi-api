package config

/**
 * The config wrapper provides an easier way of using config operations
 * without having to deal with properties and results.
 */
type ConfigWrapper interface {
	Get(key string) (ConfigScopedValues, error)
	Set(key string, values ConfigScopedValues) error
	List(parent string) ([]string, error)
}

// Some scoped config values kept in order of scope
type ConfigScopedValues struct {
	configMap map[string]ConfigScopedValue
	order     []string
}

// An individual scoped config value
type ConfigScopedValue []byte

// Save JIT initializer
func (values *ConfigScopedValues) safe() {
	if values.configMap == nil {
		values.configMap = map[string]ConfigScopedValue{}
		values.order = []string{}
	}
}

// Get a FileSource from the set
func (values *ConfigScopedValues) Get(key string) (ConfigScopedValue, bool) {
	values.safe()

	value, found := values.configMap[key]
	return value, found
}

// Add a FileSource to the set
func (values *ConfigScopedValues) Set(key string, source ConfigScopedValue) {
	values.safe()

	if _, found := values.configMap[key]; !found {
		values.order = append(values.order, key)
	}
	values.configMap[key] = source
}

// Get the key order for the set
func (values *ConfigScopedValues) Order() []string {
	values.safe()
	return values.order
}
