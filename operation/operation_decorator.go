package operation

import (
	"errors"
)

/**
 * This file provides a complete operation implementation that
 * uses one existing operation to decorate another.  The result
 * is an operation with metadata from the decorated operation
 * and properties from both decorated and decorator.  When
 * executing, the decorator is first executed, and if sucessful
 * then the decorated is executed.
 *
 * This operation blocks while performing the decorating operation
 * until it is finished.
 *
 * If more custom sequencing is desired, extend this operation
 * and override the Exec() method
 */

// A wrapping operation that decorates one operation with another
type DecoratedOperation struct {
	// Operation that wraps or decorates the other operation
	decorating Operation
	// Operation being decorated
	decorated Operation
}

// Get decorted operation id
func (operation *DecoratedOperation) Id() string {
	return operation.decorated.Id()
}

// Get decorted operation label
func (operation *DecoratedOperation) Label() string {
	return operation.decorated.Label() + " [" + operation.decorating.Label() + "]"
}

// return a multiline string description for the Operation
func (operation *DecoratedOperation) Description() string {
	return operation.decorated.Description() + " [" + operation.decorating.Description() + "]"
}

// Is this operation meant to be used only inside the API
func (operation *DecoratedOperation) Internal() bool {
	return operation.decorated.Internal()
}

// Run a validation check on the Operation
func (operation *DecoratedOperation) Validate() bool {
	return operation.decorating.Validate() && operation.decorated.Validate()
}

// Get Operation Properties from both operations
func (operation *DecoratedOperation) Properties() Properties {
	Properties := operation.decorated.Properties()
	Properties.Merge(operation.decorating.Properties())
	return Properties
}

// Execute the decorating operation, and then execute the decorated operation if the decorating was successful
func (operation *DecoratedOperation) Exec(props *Properties) Result {
	result := operation.decorating.Exec(props)
	<-result.Finished()

	if !result.Success() {
		return result
	}

	return operation.decorated.Exec(props)
}

/**
 * This extension to the decorator evaluates the decorated operation
 * only if a targeted bool property on the decorating operation is
 * true after execution.
 */

// A constructor for DecoratedBooleanPropertyBasedOperation
func New_DecoratedBooleanPropertyBasedOperation(decorating Operation, decorated Operation, property string) *DecoratedBooleanPropertyBasedOperation {
	return &DecoratedBooleanPropertyBasedOperation{
		DecoratedOperation: DecoratedOperation{
			decorating: decorating,
			decorated:  decorated,
		},
		property: property,
	}
}

// A decorator operation that uses a property to decide on success
type DecoratedBooleanPropertyBasedOperation struct {
	DecoratedOperation

	property string
}

// Execute the decorating operation, and then execute the decorated operation if the decorator property is true
func (operation *DecoratedBooleanPropertyBasedOperation) Exec(props *Properties) Result {
	boolValue := false
	result := New_StandardResult()

	decoratingResult := operation.decorating.Exec(props)
	<-decoratingResult.Finished()

	if decoratingResult.Success() {
		props := operation.decorating.Properties()
		if boolProp, found := props.Get(operation.property); !found {
			result.AddError(errors.New("Decorator operation did not have the targeted property"))
			result.MarkFailed()
		} else if boolProp.Type() != "bool" {
			result.AddError(errors.New("Decorator operation targeted property was not a boolean [" + operation.property + ":" + boolProp.Type() + "]"))
			result.MarkFailed()
		} else {
			boolValue = boolProp.Get().(bool)
			if !boolValue {
				result.MarkFailed()
			}
		}
	}

	if result.Success() && boolValue {
		return operation.decorated.Exec(props)
	} else {
		return Result(result)
	}
}
