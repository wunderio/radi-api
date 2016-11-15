package builder

import (
	"errors"

	"github.com/james-nesbitt/kraut-api/api"
	"github.com/james-nesbitt/kraut-api/operation"	
)

/**
 * This provides a more advanced API handler for a project
 * which allows a stronger definition of which handlers are used
 * in a project, and which operations should come from which
 * handlers.
 */

/**
 * Implementations are string identifiers or gorups of
 * operations, which will likely always relate to 
 * the api/operations (but not strictly necessary)
 */

// Constructor for Implementations
func New_Implementations(implementations []string) *Implementations {
	return &Implementations{
		implementations: implementations,
	}
}

// A list of which implementations should be included.
type Implementations struct {
	implementations []string
}

// Provides the implementations as an ordered string list
func (implementations *Implementations) Order() []string {
	return implementations.implementations
}
// Provides the implementations as an ordered string list
func (implementations *Implementations) Merge(merge Implementations) {
	for _, add := range merge.Order() {
		implementations.implementations = append(implementations.implementations, add)
	}
}

/**
 * The BuildAPI
 */

// BuilderAPI provides an API that is picks handlers and operations based on configuration
type BuilderAPI struct {
	availableBuilders Builders
	activated []string
}

// Make a new handler available in the API (available to be activated)
func (builder *BuilderAPI) AddBuilder(hand Builder) {
	builder.availableBuilders.Add(hand.Id(), hand)
}
func (builder *BuilderAPI) ActivateBuilder(id string, implementations Implementations, settings interface{}) error {
	if hand, err := builder.availableBuilders.Get(id); err == nil {
		hand.SetAPI(api.API(builder))

		if err := hand.Activate(implementations, settings); err == nil {
			found := false
			for _, activated := range builder.activated {
				if activated == id {
					found = true
					break
				}
			}
			if !found {
				builder.activated = append(builder.activated, id)
			}
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("Builder not found")
	}
}

// Return a list of operations for the API from all of the activated Builders
func (builder *BuilderAPI) Operations() operation.Operations {
	ops := operation.Operations{}
	for _, id := range builder.activated {
		hand, _ := builder.availableBuilders.Get(id)
		ops.Merge(hand.Operations())
	}
	return ops
}
