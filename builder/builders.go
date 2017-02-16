package builder

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * An ordered list of builders
 */

type Builders interface {
	// Add a builder
	Add(key string, builder Builder) error
	// Get a single builder
	Get(key string) (Builder, error)
	// Get the builder ordered keys
	Order() []string
	// Is the builder list empty
	Empty() bool
	// Return a list of operations from all of the Builders
	Operations() api_operation.Operations
}

// A simple ordered collection of Builder objects
type SimpleBuilders struct {
	builders map[string]Builder
	order    []string
}

// Constructor for SimpleBuilders
func New_SimpleBuilders() *SimpleBuilders {
	return &SimpleBuilders{}
}

// convert this to a Builders interface
func (builders *SimpleBuilders) Builders() Builders {
	return Builders(builders)
}

// safe intitializer
func (builders *SimpleBuilders) safe() {
	if builders.order == nil {
		builders.builders = map[string]Builder{}
		builders.order = []string{}
	}
}

// Add a builder
func (builders *SimpleBuilders) Add(key string, builder Builder) error {
	builders.safe()
	if _, exists := builders.builders[key]; !exists {
		builders.order = append(builders.order, key)
	}
	builders.builders[key] = builder
	return nil
}

// Get a single builder
func (builders *SimpleBuilders) Get(key string) (Builder, error) {
	builders.safe()
	if builder, found := builders.builders[key]; found {
		return builder, nil
	} else {
		return builder, errors.New("No such builder found")
	}
}

// Get the builder ordered keys
func (builders *SimpleBuilders) Order() []string {
	builders.safe()
	return builders.order
}

// Is the builder list empty
func (builders *SimpleBuilders) Empty() bool {
	return builders.builders == nil
}

// Return a list of operations from all of the Builders
func (builders *SimpleBuilders) Operations() api_operation.Operations {
	ops := api_operation.New_SimpleOperations()
	for _, key := range builders.Order() {
		builder, _ := builders.Get(key)
		ops.Merge(builder.Operations())
	}
	return ops.Operations()
}
