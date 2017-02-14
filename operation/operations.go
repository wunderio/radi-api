package operation

// An ordered list of operations
type Operations interface {
	// Add an operation
	Add(Operation) bool
	// Merge an Operations set into this one
	Merge(Operations)
	// Retrieve an operation by it's id
	Get(string) (Operation, bool)
	// List operation ids in order
	Order() []string
}

// SimpleOperations are a keyed map of individual Operations
type SimpleOperations struct {
	opMap map[string]Operation
	order []string
}

func New_SimpleOperations() *SimpleOperations {
	return &SimpleOperations{}
}

func (ops *SimpleOperations) Operations() Operations {
	return Operations(ops)
}

// Save internal initializer
func (ops *SimpleOperations) safe() {
	if ops.opMap == nil {
		ops.opMap = map[string]Operation{}
		ops.order = []string{}
	}
}

// Add a new Operation to the map
func (ops *SimpleOperations) Add(add Operation) bool {
	ops.safe()

	addId := add.Id()
	ops.opMap[addId] = add
	ops.order = append(ops.order, addId)
	return true
}

// Merge one Operations set into the current set
func (ops *SimpleOperations) Merge(merge Operations) {
	ops.safe()

	for _, operation := range merge.Order() {
		mergeOperation, _ := merge.Get(operation)
		ops.Add(mergeOperation)
	}
}

// Operation accessor by id
func (ops *SimpleOperations) Get(id string) (Operation, bool) {
	ops.safe()

	operation, ok := ops.opMap[id]
	return operation, ok
}

// Order returns a slice of operation ids, used in iterators to maintain an operation order
func (ops *SimpleOperations) Order() []string {
	ops.safe()

	return ops.order
}
