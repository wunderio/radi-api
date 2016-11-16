package operation

/**
 * This file provides an Operation base class, that chains together
 * set of other operations to be executed in sequence.  The chain
 * provides properties as a union of the properties of all of the
 * chained operetaionts, and executes in sequence, optionally stopping
 * on the first to suceed.
 *
 * @NOTE this operation does not complete the interface, as it needs
 * metadata methods for Id(), Label() etc.
 */

// ChainOperation runs multiple operations in sequence.  Extend this and add ID/Label/Properties handling
type ChainOperation struct {
	// Tell this operation to stop processing the chained operations on the first TRUE result
	stopOnSuccess bool
	// The ordered chain of Operations to process when this operation is called
	operations Operations
}

// Add an operation to the chain
func (chain *ChainOperation) AddOperation(operation Operation) {
	if &chain.operations == nil {
		chain.operations = Operations{}
	}
	chain.operations.Add(operation)
}

// Get Operation Configuration from all operations
func (chain *ChainOperation) Properties() *Properties {
	Properties := Properties{}
	for _, key := range chain.operations.Order() {
		op, _ := chain.operations.Get(key)
		Properties.Merge(*op.Properties())
	}
	return &Properties
}

// Exec the chain operation by running Exec on each child
func (chain *ChainOperation) Exec() Result {
	chainResult := ChainResult{
		BaseResult{
			success: true,
			errors:  []error{},
		},
	}

	for _, id := range chain.operations.Order() {
		operation, _ := chain.operations.Get(id)
		result := operation.Exec()
		chainResult.AddResult(result)

		if resultSuccess, _ := result.Success(); chain.stopOnSuccess && resultSuccess {
			break
		}
	}

	return Result(&chainResult)
}
