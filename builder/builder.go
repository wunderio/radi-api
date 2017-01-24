package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * Builders are essentially meta handlers, which can
 * produce handlers from simpler settings.
 *
 * Builders often need a back reference to the api
 * as they may need other operations, for example
 * to build Config based operations
 */

// A single API handler for providing operations
type Builder interface {
	// Set a API for this Handler
	SetAPI(parent api.API)
	// Initialize and activate the Handler
	Activate(Implementations, SettingsProvider) error
	// Rturn a string identifier for the Handler (not functionally needed yet)
	Id() string
	// Return a list of Operations from the Handler
	Operations() *operation.Operations
}

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
func (builders *Builders) Operations() operation.Operations {
	ops := operation.Operations{}
	for _, key := range builders.Order() {
		builder, _ := builders.Get(key)
		ops.Merge(builder.Operations())
	}
	return ops
}
