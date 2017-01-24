package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * This provides a more advanced API handler for a project
 * which allows a stronger definition of which handlers are used
 * in a project, and which operations should come from which
 * handlers.
 */

/**
 * The BuildAPI
 */

// Constructor for BuilderAPI

// BuilderAPI provides an API that is picks builders based on configuration
type BuilderAPI struct {
	builders  Builders
	activated []string
}

// Make a new Builder available in the API (available to be activated)
func (builderApi *BuilderAPI) AddBuilder(builder Builder) {
	builderApi.builders.Add(builder.Id(), builder)
}
func (builderApi *BuilderAPI) ActivateBuilder(id string, implementations Implementations, settings SettingsProvider) error {
	if builder, err := builderApi.builders.Get(id); err == nil {
		builder.SetAPI(api.API(builderApi))

		if err := builder.Activate(implementations, settings); err == nil {
			found := false
			for _, activated := range builderApi.activated {
				if activated == id {
					found = true
					break
				}
			}
			if !found {
				builderApi.activated = append(builderApi.activated, id)
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
func (builderApi *BuilderAPI) Operations() operation.Operations {
	ops := operation.Operations{}
	for _, id := range builderApi.activated {
		builder, _ := builderApi.builders.Get(id)
		ops.Merge(builder.Operations())
	}
	return ops
}
