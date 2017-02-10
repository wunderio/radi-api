package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
)

/**
 * A standard Project struct
 */

// StandardProject provides an API that is picks builders based on configuration
type StandardProject struct {
	builders  Builders
	activated []string
}

// Constructor for StandardProject
func New_StandardProject() *StandardProject {
	return &StandardProject{
		builders: Builders{},
		activated: []string{},
	}
}

// Convert this to a Project interface
func (project *StandardProject) Project() Project {
	return Project(project)
}

// Make a new Builder available in the Project (available to be activated)
func (project *StandardProject) API() api.API {
	return api.API(project)
}

// Make a new Builder available in the Project (available to be activated)
func (project *StandardProject) AddBuilder(builder Builder) {
	project.builders.Add(builder.Id(), builder)
}

// Activate some implementations for a builder in the Project
func (project *StandardProject) ActivateBuilder(id string, implementations Implementations, settings SettingsProvider) error {
	if builder, err := project.builders.Get(id); err == nil {
		builder.SetAPI(api.API(project))

		if err := builder.Activate(implementations, settings); err == nil {
			found := false
			for _, activated := range project.activated {
				if activated == id {
					found = true
					break
				}
			}
			if !found {
				project.activated = append(project.activated, id)
			}
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("Builder not found")
	}
}

// Ask a StandardProject to validate itself, after it has been fully activated, before we ask for operations.
func (project *StandardProject) Validate() operation.Result {
	result := operation.New_StandardResult()
	result.MarkSuccess()
	result.MarkFinished()
	return result.Result()
}

// Return a list of operations for Project from all of the activated Builders
func (project *StandardProject) Operations() operation.Operations {
	ops := operation.Operations{}
	for _, id := range project.activated {
		builder, _ := project.builders.Get(id)
		ops.Merge(builder.Operations())
	}
	return ops
}
