package result

// Make a successfully finished result with no message and no properties
func MakeSuccessfulResult() Result {
	res := New_StandardResult()
	res.MarkSuccess()
	res.MarkFinished()
	return res.Result()
}

// StandardResult is a base class for results which keep success boolean and errors slice as variables
type StandardResult struct {
	finished chan bool
	success  bool
	errors   []error
}

// Constructor for StandardResult
func New_StandardResult() *StandardResult {
	return &StandardResult{
		finished: make(chan bool),
		success:  true, // results default to success, to prevent silly issues.
		errors:   []error{},
	}
}

// Return this struct as a Result interfacew
func (base *StandardResult) Result() Result {
	return Result(base)
}

// Mark the result operation as finihsed [non blocking]
func (base *StandardResult) MarkFinished() {
	go func(finished chan bool) { finished <- true }(base.finished)
}

// Mark the result as succeeded
func (base *StandardResult) SetSuccess(success bool) {
	base.success = success
}

// Mark the result as succeeded
func (base *StandardResult) MarkSuccess() {
	base.success = true
}

// Mark the result as failed
func (base *StandardResult) MarkFailed() {
	base.success = false
}

// Add some errors to the result
func (base *StandardResult) AddErrors(errs []error) {
	base.errors = append(base.errors, errs...)
}

// Add some errors to the result
func (base *StandardResult) AddError(err error) {
	base.errors = append(base.errors, err)
}

// Did the operation succeed
func (base *StandardResult) Finished() chan bool {
	return base.finished
}

// Did the operation succeed
func (base *StandardResult) Success() bool {
	return base.success
}

// return any errors that occured
func (base *StandardResult) Errors() []error {
	return base.errors
}

// Merge a result into this result
func (base *StandardResult) Merge(merge Result) {
	if !merge.Success() {
		base.MarkFailed()
	}
	base.AddErrors(merge.Errors())
}
