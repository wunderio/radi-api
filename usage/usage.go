package usage

// Usage nao fir wgeb a property should be considered usageable
type Usage interface {
	// >Retrieve usage boolean for a key
	Get(string) bool
	// Answer a boolean usage question for a usage name
	Has(string) bool
	// List usages
	List() []string
}

/**
 * Usage Implementations
 */

// A simple Usage struct that runs of an internally managed map
type SimpleMapUsage struct {
	usageMap map[string]bool
}

// A constructor of a Usage Interface from a
func MakeUsageFromMap(source map[string]bool) Usage {
	return (&SimpleMapUsage{usageMap: source}).Usage()
}

// A constructor of a Usage Interface from a
func New_SimpleMapUsageEmpty() *SimpleMapUsage {
	return &SimpleMapUsage{}
}

// Convert this struct to a Usage interface
func (simple *SimpleMapUsage) Usage() Usage {
	return Usage(simple)
}

// Answer a boolean usage question for a usage name
func (simple *SimpleMapUsage) Get(key string) bool {
	if has, found := simple.usageMap[key]; has {
		return found
	}
	return false
}

// Answer a boolean usage question for a usage name
func (simple *SimpleMapUsage) Has(key string) bool {
	return simple.Get(key)
}

// Set a specific usage flag
func (simple *SimpleMapUsage) Set(key string, val bool) {
	simple.usageMap[key] = val
}

// List registered usage keys
func (simple *SimpleMapUsage) List() []string {
	list := []string{}
	for key, _ := range simple.usageMap {
		list = append(list, key)
	}
	return list
}

// List registered usage keys
func (simple *SimpleMapUsage) Merge(merge Usage) {
	for _, key := range merge.List() {
		simple.Set(key, merge.Has(key))
	}
}
