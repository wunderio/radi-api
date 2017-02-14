package builder

/**
 * Implementations are string identifiers or gorups of operations, which will
 * likely always relate to the api/operations (but not strictly necessary)
 */

// A list of which implementations should be included.
type Implementations struct {
	implementations []string
}

// Constructor for Implementations
func New_Implementations(implementations []string) *Implementations {
	return &Implementations{
		implementations: implementations,
	}
}

// Provides the implementations as an ordered string list
func (implementations *Implementations) Order() []string {
	return implementations.implementations
}

// Provides the implementations as an ordered string list
func (implementations *Implementations) Merge(merge Implementations) {
	implementations.implementations = append(implementations.implementations, merge.Order()...)
}
