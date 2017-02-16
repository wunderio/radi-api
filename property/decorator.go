package property

import (
	"github.com/wunderkraut/radi-api/usage"
)

/**
 * Various decorating properties
 */

// A decorating property that alters the Usage for the decorated property
type UsageDecoratedProperty struct {
	decorated Property
	usage     usage.Usage
}

// Constructer for a DecoratingInternalizerProperty
func New_UsageDecoratedProperty(decorated Property, usageOverride usage.Usage) Property {
	// Build a UsageDecoratedProperty reference, and return it as a Property interface
	return (&UsageDecoratedProperty{
		decorated: decorated,
		usage:     usageOverride,
	}).Property()
}

// Convert this into a Property
func (usageDecProp *UsageDecoratedProperty) Property() Property {
	return Property(usageDecProp)
}

// ID returns string unique property Identifier
func (usageDecProp *UsageDecoratedProperty) Id() string {
	return usageDecProp.decorated.Id()
}

// Pull Usage information from the decorated property
func (usageDecProp *UsageDecoratedProperty) Usage() usage.Usage {
	return usageDecProp.usage
}

// Give an idea of what type of value the property consumes
func (usageDecProp *UsageDecoratedProperty) Type() string {
	return usageDecProp.decorated.Type()
}

// Label returns a short user readable label for the property
func (usageDecProp *UsageDecoratedProperty) Label() string {
	return usageDecProp.decorated.Label()
}

// Description provides a longer multi-line string description of what the property does
func (usageDecProp *UsageDecoratedProperty) Description() string {
	return usageDecProp.decorated.Description()
}

// Value allows the retrieval and setting of unknown Typed values for the property.
func (usageDecProp *UsageDecoratedProperty) Get() interface{} {
	return usageDecProp.decorated.Get()
}
func (usageDecProp *UsageDecoratedProperty) Set(value interface{}) bool {
	return usageDecProp.decorated.Set(value)
}
