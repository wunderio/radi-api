package command

/**
 * An operations wrapper that provides an easy to use tool
 * for accessing commands without having to manage operations
 * an properties.
 */

type CommandWrapper interface {
	// Retrieve a command that matches a key
	Get(key string) (Command, error)
	// List all of the command keys available
	List(parent string) ([]string, error)
}
