package operation

import (
	"context"
)

/**
 * A result that is aware of a net context, and
 * times out either when it is set to finished,
 * or when the context ends
 */

type ContextualResult struct {
	StandardResult
}

// constructor for ContextualResult
func New_ContextualResult(ctx context.Context) *ContextualResult {
	result := &ContextualResult{
		StandardResult: *New_StandardResult(),
	}
	result.expireOnContext(ctx)

	return result
}

func (result *ContextualResult) expireOnContext(ctx context.Context) {
	go func() {

		for {
			select {
			case <-ctx.Done():
				result.MarkFailed()
				if err := ctx.Err(); err != nil {
					result.AddError(err)
				}
				result.MarkFinished()
				return
			}
		}

	}()
}
