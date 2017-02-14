package handler

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * An ordered collection of Handler objects
 *
 * @TODO should this be an interface?
 */

// Ordered list of handlers
type Handlers struct {
	handlers map[string]Handler
	order    []string
}

// safe intitializer
func (handlers *Handlers) safe() {
	if handlers.order == nil {
		handlers.handlers = map[string]Handler{}
		handlers.order = []string{}
	}
}

// Add a handler
func (handlers *Handlers) Add(hand Handler) error {
	key := hand.Id()
	handlers.safe()
	if _, exists := handlers.handlers[key]; !exists {
		handlers.order = append(handlers.order, key)
	}
	handlers.handlers[key] = hand
	return nil
}

// Get a single handler
func (handlers *Handlers) Get(key string) (Handler, error) {
	handlers.safe()
	if hand, found := handlers.handlers[key]; found {
		return hand, nil
	} else {
		return hand, errors.New("No such handler found")
	}
}

// Get the handler ordered keys
func (handlers *Handlers) Order() []string {
	handlers.safe()
	return handlers.order
}

// Return a list of operations from all of the Handlers
func (handlers *Handlers) Operations() api_operation.Operations {
	ops := api_operation.New_SimpleOperations()
	for _, key := range handlers.Order() {
		hand, _ := handlers.Get(key)
		ops.Merge(hand.Operations())
	}
	return ops.Operations()
}
