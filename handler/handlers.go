package handler

import (
	"errors"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * An ordered collection of Handler objects
 */

type Handlers interface {
	// Add a new handler to the list
	Add(hand Handler) error
	// Get a single handler
	Get(key string) (Handler, error)
	// Get the handler ordered keys
	Order() []string
	// Return a list of operations from all of the Handlers
	Operations() api_operation.Operations
}

// Ordered list of handlers
type SimpleHandlers struct {
	handlers map[string]Handler
	order    []string
}

// Constructor for SimpleHandlers
func New_SimpleHandlers() *SimpleHandlers {
	return &SimpleHandlers{}
}

// Convert this SimpleHandlers to a Handlers interface
func (handlers *SimpleHandlers) Handlers() Handlers {
	return Handlers(handlers)
}

// safe intitializer
func (handlers *SimpleHandlers) safe() {
	if handlers.order == nil {
		handlers.handlers = map[string]Handler{}
		handlers.order = []string{}
	}
}

// Add a handler
func (handlers *SimpleHandlers) Add(hand Handler) error {
	key := hand.Id()
	handlers.safe()
	if _, exists := handlers.handlers[key]; !exists {
		handlers.order = append(handlers.order, key)
	}
	handlers.handlers[key] = hand
	return nil
}

// Get a single handler
func (handlers *SimpleHandlers) Get(key string) (Handler, error) {
	handlers.safe()
	if hand, found := handlers.handlers[key]; found {
		return hand, nil
	} else {
		return hand, errors.New("No such handler found")
	}
}

// Get the handler ordered keys
func (handlers *SimpleHandlers) Order() []string {
	handlers.safe()
	return handlers.order
}

// Return a list of operations from all of the Handlers
func (handlers *SimpleHandlers) Operations() api_operation.Operations {
	ops := api_operation.New_SimpleOperations()
	for _, key := range handlers.Order() {
		hand, _ := handlers.Get(key)
		ops.Merge(hand.Operations())
	}
	return ops.Operations()
}
