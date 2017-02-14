package builder

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * An ordered list of builders
 *
 * @TODO should this be an interface
 */

// An ordered collection of Builder objects
type Builders struct {
	builders map[string]Builder
	order    []string
}

// safe intitializer
func (builders *Builders) safe() {
	if builders.order == nil {
		builders.builders = map[string]Builder{}
		builders.order = []string{}
	}
}

// Add a builder
func (builders *Builders) Add(key string, builder Builder) error {
	builders.safe()
	if _, exists := builders.builders[key]; !exists {
		builders.order = append(builders.order, key)
	}
	builders.builders[key] = builder
	return nil
}

// Get a single builder
func (builders *Builders) Get(key string) (Builder, error) {
	builders.safe()
	if builder, found := builders.builders[key]; found {
		return builder, nil
	} else {
		return builder, errors.New("No such builder found")
	}
}

// Get the builder ordered keys
func (builders *Builders) Order() []string {
	builders.safe()
	return builders.order
}

// Is the builder list empty
func (builders *Builders) Empty() bool {
	return builders.builders == nil
}

// Return a list of operations from all of the Builders
func (builders *Builders) Operations() api_operation.Operations {
	ops := api_operation.New_SimpleOperations()
	for _, key := range builders.Order() {
		builder, _ := builders.Get(key)
		ops.Merge(builder.Operations())
	}
	return ops.Operations()
}
