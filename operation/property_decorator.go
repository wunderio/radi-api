package operation

/**
 * Various decorating properties
 */

// constructor for DecoratingInternalizerProperty
func New_DecoratingInternalizerProperty(prop Property) *DecoratingInternalizerProperty {
	return &DecoratingInternalizerProperty{
		decorated: prop,
	}
}

// A decorating property that internalizes any property
type DecoratingInternalizerProperty struct {
	decorated Property
}

// Convert this into a Property
func (decProp *DecoratingInternalizerProperty) Property() Property {
	return Property(decProp)
}

// ID returns string unique property Identifier
func (decProp *DecoratingInternalizerProperty) Id() string {
	return decProp.decorated.Id()
}

// Label returns a short user readable label for the property
func (decProp *DecoratingInternalizerProperty) Label() string {
	return decProp.decorated.Label()
}

// Description provides a longer multi-line string description of what the property does
func (decProp *DecoratingInternalizerProperty) Description() string {
	return decProp.decorated.Description()
}

// Mark a property as being for internal use only (no shown to users)
func (decProp *DecoratingInternalizerProperty) Internal() bool {
	return true
}

// Give an idea of what type of value the property consumes
func (decProp *DecoratingInternalizerProperty) Type() string {
	return decProp.decorated.Type()
}

// Value allows the retrieval and setting of unknown Typed values for the property.
func (decProp *DecoratingInternalizerProperty) Get() interface{} {
	return decProp.decorated.Get()
}
func (decProp *DecoratingInternalizerProperty) Set(value interface{}) bool {
	return decProp.decorated.Set(value)
}
