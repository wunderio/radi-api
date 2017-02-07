package security

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

// Is this an internal API operation
func (authenticate *BaseSecurityUserOperation) Internal() bool {
	return true
}
