package command

/**
 * An ordered set of commands used internall
 * in the commands handlers.
 *
 * @TODO should this be an interface?
 */

type Commands struct {
	commands map[string]Command
	order    []string
}

// Safe lazy constructor
func (commands *Commands) safe() {
	if &commands.commands == nil {
		commands.commands = map[string]Command{}
		commands.order = []string{}
	}
}

// Get a command
func (commands *Commands) Get(key string) (Command, bool) {
	commands.safe()
	comm, found := commands.commands[key]
	return comm, found
}

// Add a command
func (commands *Commands) Set(key string, comm Command) error {
	commands.safe()
	if _, exists := commands.commands[key]; !exists {
		commands.order = append(commands.order, key)
	}
	commands.commands[key] = comm
	return nil
}

// Order of commands
func (commands *Commands) Order() []string {
	commands.safe()
	return commands.order
}
