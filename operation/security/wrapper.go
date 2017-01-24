package security

import (
	api_operation "github.com/james-nesbitt/radi-api/operation"
)

/**
 * Security wrapper definition
 */

// SecurityWrapper definition
type SecurityWrapper interface {
	AuthorizeOperation(api_operation.Operation) RuleResult
	CurrentUser() SecurityUser
}
