package operation

import (
	"errors"

	"github.com/wunderkraut/radi-api/property"
	"github.com/wunderkraut/radi-api/result"
	"github.com/wunderkraut/radi-api/usage"
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
func (decOp *DecoratedOperation) Id() string {
	return decOp.decorated.Id()
}

// Get decorted operation label
func (decOp *DecoratedOperation) Label() string {
	return decOp.decorated.Label() + " [" + decOp.decorating.Label() + "]"
}

// return a multiline string description for the Operation
func (decOp *DecoratedOperation) Description() string {
	return decOp.decorated.Description() + " [" + decOp.decorating.Description() + "]"
}

// return a man type help for the Operation
func (decOp *DecoratedOperation) Help() string {
	return decOp.decorated.Help()
}

// Is this operation meant to be used only inside the API
func (decOp *DecoratedOperation) Usage() usage.Usage {
	combUsage := usage.New_SimpleMapUsageEmpty()

	combUsage.Merge(decOp.decorated.Usage())
	combUsage.Merge(decOp.decorating.Usage())

	return combUsage.Usage()
}

// Run a validation check on the Operation
func (decOp *DecoratedOperation) Validate() result.Result {
	/**
	 * @TODO developer a result type for merging, and use it here
	 */
	return decOp.decorating.Validate()
}

// Get Operation Properties from both operations
func (decOp *DecoratedOperation) Properties() property.Properties {
	props := property.New_SimplePropertiesEmpty()

	props.Merge(decOp.decorated.Properties())
	props.Merge(decOp.decorating.Properties())

	return props.Properties()
}

// Execute the decorating operation, and then execute the decorated operation if the decorating was successful
func (decOp *DecoratedOperation) Exec(props property.Properties) result.Result {
	res := result.New_StandardResult()

	decRes := decOp.decorating.Exec(props)
	<-decRes.Finished()

	res.Merge(decRes)

	if !res.Success() {
		return res.Result()
	}

	execResult := decOp.decorated.Exec(props)
	res.Merge(execResult)

	return res.Result()
}

/**
 * This extension to the decorator evaluates the decorated operation
 * only if a targeted bool property on the decorating operation is
 * true after that op is executed.
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
func (decOp *DecoratedBooleanPropertyBasedOperation) Exec(props property.Properties) result.Result {
	boolValue := false
	res := result.New_StandardResult()

	decoratingResult := decOp.decorating.Exec(props)
	<-decoratingResult.Finished()

	if decoratingResult.Success() {
		props := decOp.decorating.Properties()
		if boolProp, found := props.Get(decOp.property); !found {
			res.AddError(errors.New("Decorator operation did not have the targeted property"))
			res.MarkFailed()
		} else if boolProp.Type() != "bool" {
			res.AddError(errors.New("Decorator operation targeted property was not a boolean [" + decOp.property + ":" + boolProp.Type() + "]"))
			res.MarkFailed()
		} else {
			boolValue = boolProp.Get().(bool)
			if !boolValue {
				res.MarkFailed()
			}
		}
	}

	if res.Success() && boolValue {
		res.Merge(decOp.decorated.Exec(props))
	}

	return res.Result()
}
