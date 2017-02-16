package security

import (
	"os/user"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * User retrieval related functionality
 */

const (
	OPERATION_KEY_SECURITY_USER_OPERATION = "security.user"
)

// An interface for how a user is defined
type SecurityUser interface {
	// Return a machine id for the user for things like authorization checks
	Id() string
	// Human readable display id for the user
	Label() string

	/**
	 * INCOMING
	 */

	/**
	 * Respond to an authentication challenge
	 *
	 * [not sure how the challenge can be abstracted yet]
	 *
	 * Responding to an authentication challenge will likely
	 * involve some UI element, at least for something
	 * like RSA passphrase or password entry.
	 */
	// Authenticate(AuthenticationChallenge) AuthenticationResponse
}

/**
 * A SecurityUser implementation that wraps the
 * core os/user.User object
 */
type CoreUserSecurityUser struct {
	user *user.User
}

// Constructor for CoreUserSecurityUser
func New_CoreUserSecurityUser(coreUser *user.User) *CoreUserSecurityUser {
	return &CoreUserSecurityUser{
		user: coreUser,
	}
}

// Return this object as a SecurityUser interface
func (coreUser *CoreUserSecurityUser) SecurityUser() SecurityUser {
	return SecurityUser(coreUser)
}

// Return a machine id for the user for things like authorization checks
func (coreUser *CoreUserSecurityUser) Id() string {
	return coreUser.user.Username
}

// Human readable display id for the user
func (coreUser *CoreUserSecurityUser) Label() string {
	return coreUser.user.Name
}

/**
 * User operations return information about the currently
 * authenticated user
 */

// Base class for security authenticate Operation
type BaseSecurityUserOperation struct{}

// Id the operation
func (authenticate *BaseSecurityUserOperation) Id() string {
	return OPERATION_KEY_SECURITY_USER_OPERATION
}

// Label the operation
func (authenticate *BaseSecurityUserOperation) Label() string {
	return "Get User"
}

// Description for the operation
func (authenticate *BaseSecurityUserOperation) Description() string {
	return "Retrieve information about the current app user."
}

// Man page for the operation
func (authenticate *BaseSecurityUserOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (authenticate *BaseSecurityUserOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
