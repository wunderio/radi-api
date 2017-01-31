package orchestrate

import (
	"errors"

	"github.com/wunderkraut/radi-api/operation"
)

/**
 * A simple, potentially blocking orchestrate wrapper implementation
 */

// Constructor for SimpleOrchestrateWrapper
func New_SimpleOrchestrateWrapper(operations *operation.Operations) *SimpleOrchestrateWrapper {
	return &SimpleOrchestrateWrapper{
		operations: operations,
	}
}

// A simple orchestration operation wrapper
type SimpleOrchestrateWrapper struct {
	operations *operation.Operations
}

// Orchestrate Up method
func (wrapper *SimpleOrchestrateWrapper) Up() error {
	var found bool
	var op operation.Operation

	if op, found = wrapper.operations.Get(OPERATION_ID_ORCHESTRATE_DOWN); !found {
		return errors.New("No up operation available in Orchestrate Wrapper")
	}

	props := op.Properties()
	result := op.Exec(&props)
	<-result.Finished()

	if !result.Success() {
		errs := result.Errors()
		if len(errs) == 0 {
			return errors.New("Operation orchestrate UP failed to execute in Setting Wrapper")
		} else {
			return errs[0]
		}
	}

	return nil
}

// Orchestrate Down method
func (wrapper *SimpleOrchestrateWrapper) Down() error {
	var found bool
	var op operation.Operation

	if op, found = wrapper.operations.Get(OPERATION_ID_ORCHESTRATE_DOWN); !found {
		return errors.New("No down operation available in Orchestrate Wrapper")
	}

	props := op.Properties()
	result := op.Exec(&props)
	<-result.Finished()

	if !result.Success() {
		errs := result.Errors()
		if len(errs) == 0 {
			return errors.New("Operation orchestrate DOWN failed to execute in Setting Wrapper")
		} else {
			return errs[0]
		}
	}

	return nil
}
