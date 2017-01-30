package operation

// ChainResult is a Result that aggregates multiple results
type ChainResult struct {
	StandardResult
}

// Add A result to the chain
// func (chain *ChainResult) AddResult(add Result) {
// 	if !add.Success() {
// 		chain.MarkFailed()
// 	}
// 	chain.addErrors(add.Errors())
// }
