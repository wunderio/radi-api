package property

/**
 * Tools to manage collections of properties
 */

// A set of Properties
type Properties interface {
	Get(string) (Property, bool)
	Order() []string
}

// A set of Properties
type SimpleProperties struct {
	propMap map[string]Property
	order   []string
}

func New_SimplePropertiesEmpty() *SimpleProperties {
	return &SimpleProperties{}
}

// Convert these properties into a Properties interface
func (properties *SimpleProperties) Properties() Properties {
	return Properties(properties)
}

// safe initialization of vars
func (properties *SimpleProperties) makeSafe() {
	if properties.propMap == nil {
		properties.propMap = map[string]Property{}
		properties.order = []string{}
	}
}

// Add a property
func (properties *SimpleProperties) Add(property Property) {
	properties.makeSafe()
	id := property.Id()
	if _, exists := properties.propMap[id]; !exists {
		properties.order = append(properties.order, id)
	}
	properties.propMap[id] = property
}

// Merge in one set of properties into this configurations
func (properties *SimpleProperties) Merge(merge Properties) {
	for _, id := range merge.Order() {
		property, _ := merge.Get(id)
		properties.Add(property)
	}
}

// Retrieve a single property based on key id
func (properties *SimpleProperties) Get(id string) (Property, bool) {
	properties.makeSafe()
	property, ok := properties.propMap[id]
	return property, ok
}

// Retrieve and ordered list of property keys
func (properties *SimpleProperties) Order() []string {
	properties.makeSafe()
	return properties.order
}
