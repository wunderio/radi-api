package builder

import (
	"errors"

	"github.com/wunderkraut/radi-api/api"
	"github.com/wunderkraut/radi-api/operation"
	"github.com/wunderkraut/radi-api/result"
)

/**
 * A standard Project struct
 */

// StandardProject provides an API that is picks builders based on configuration
type StandardProject struct {
	builders Builders
}

// Constructor for StandardProject
func New_StandardProject() *StandardProject {
	return &StandardProject{
		builders: Builders{},
	}
}

// Convert this to a Project interface
func (project *StandardProject) Project() Project {
	return Project(project)
}

// Convert this to an API interface
func (project *StandardProject) API() api.API {
	return api.API(project)
}

// Make a new Builder available in the Project (available to be activated)
func (project *StandardProject) AddBuilder(builder Builder) {
	project.builders.Add(builder.Id(), builder)
}

// Activate some implementations for a builder in the Project
func (project *StandardProject) ActivateBuilder(id string, implementations Implementations, settings SettingsProvider) result.Result {
	if builder, err := project.builders.Get(id); err == nil {
		builder.SetAPI(api.API(project))
		return builder.Activate(implementations, settings)
	} else {
		res := result.New_StandardResult()
		res.AddError(errors.New("Builder not found"))
		res.MarkFailed()
		res.MarkFinished()
		return res.Result()
	}
}

// Ask a StandardProject to validate itself, after it has been fully activated, before we ask for operations.
func (project *StandardProject) Validate() result.Result {
	/**
	 * @TODO this needs to properly validate all of the builders
	 *  in a concurrent manner.  This would probably benefit from
	 *  having a threadsafe result object.
	 */

	// just return success for now.
	return result.MakeSuccessfulResult()
}

// Return a list of operations for Project from all of the activated Builders
func (project *StandardProject) Operations() operation.Operations {
	ops := operation.New_SimpleOperations()
	for _, id := range project.builders.Order() {
		builder, _ := project.builders.Get(id)
		ops.Merge(builder.Operations())
	}
	return ops.Operations()
}
