package command

import (
	log "github.com/Sirupsen/logrus"

	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Typical command package operation Properies.
 */

const (
	// Key string for a single operation
	OPERATION_PROPERTY_COMMAND_KEY = "command.key"
	// Command object for a single operation
	OPERATION_PROPERTY_COMMAND_COMMAND = "command.command"
	// List of keys
	OPERATION_PROPERTY_COMMAND_KEYS = "command.keys"

	// list of string flags passed to the command container
	OPERATION_PROPERTY_COMMAND_FLAGS = "command.flags"

	// Input/Output objects
	OPERATION_PROPERTY_COMMAND_OUTPUT = "command.output"
	OPERATION_PROPERTY_COMMAND_ERR    = "command.err"
	OPERATION_PROPERTY_COMMAND_INPUT  = "command.input"

	// Use a context when running, to allow remote control of execution
	OPERATION_PROPERTY_COMMAND_CONTEXT = "command.context"
)

// Command for a single command key
type CommandKeyProperty struct {
	api_property.StringProperty
}

// Id for the Property
func (key *CommandKeyProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_KEY
}

// Label for the Property
func (key *CommandKeyProperty) Label() string {
	return "Command key."
}

// Description for the Property
func (key *CommandKeyProperty) Description() string {
	return "Command key."
}

// Is the Property internal only
func (key *CommandKeyProperty) Usage() api_usage.Usage {
	return api_property.Usage_Required()
}

// Copy the property
func (key *CommandKeyProperty) Copy() api_property.Property {
	prop := &CommandKeyProperty{}
	prop.Set(key.Get())
	return api_property.Property(prop)
}

// Command for a single command object
type CommandCommandProperty struct {
	command Command
}

// Id for the Property
func (com *CommandCommandProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_COMMAND
}

// Label for the Property
func (com *CommandCommandProperty) Label() string {
	return "Command object."
}

// Description for the Property
func (com *CommandCommandProperty) Description() string {
	return "Command object."
}

// Is the Property internal only
func (com *CommandCommandProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (com *CommandCommandProperty) Type() string {
	return "operation/command.Command"
}

// Get the Command value
func (com *CommandCommandProperty) Get() interface{} {
	return interface{}(com.command)
}

// Set the Command value
func (com *CommandCommandProperty) Set(value interface{}) bool {
	if converted, ok := value.(Command); ok {
		com.command = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected Command")
		return false
	}
}

// Copy the property
func (com *CommandCommandProperty) Copy() api_property.Property {
	prop := &CommandCommandProperty{}
	prop.Set(com.Get())
	return api_property.Property(prop)
}

// Command for an ordered list of command keys
type CommandKeysProperty struct {
	api_property.StringSliceProperty
}

// Id for the Property
func (keys *CommandKeysProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_KEYS
}

// Label for the Property
func (keys *CommandKeysProperty) Label() string {
	return "Command key list."
}

// Description for the Property
func (keys *CommandKeysProperty) Description() string {
	return "Command key list."
}

// Is the Property internal only
func (keys *CommandKeysProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Copy the property
func (keys *CommandKeysProperty) Copy() api_property.Property {
	prop := &CommandKeysProperty{}
	prop.Set(keys.Get())
	return api_property.Property(prop)
}

// Command for an ordered list of command keys
type CommandFlagsProperty struct {
	api_property.StringSliceProperty
}

// Id for the Property
func (flags *CommandFlagsProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_FLAGS
}

// Label for the Property
func (flags *CommandFlagsProperty) Label() string {
	return "Command flags list."
}

// Description for the Property
func (flags *CommandFlagsProperty) Description() string {
	return "An ordered list of string flags to send to a command."
}

// Is the Property internal only
func (flags *CommandFlagsProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (flags *CommandFlagsProperty) Copy() api_property.Property {
	prop := &CommandFlagsProperty{}
	prop.Set(flags.Get())
	return api_property.Property(prop)
}

// A command Property for command output
type CommandOutputProperty struct {
	api_property.WriterProperty
}

// Id for the Property
func (out *CommandOutputProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_OUTPUT
}

// Label for the Property
func (out *CommandOutputProperty) Label() string {
	return "Command output io.Writer."
}

// Description for the Property
func (out *CommandOutputProperty) Description() string {
	return "An io.Writer, which will receive the command execution output.  Any io.writer can be used, the default here will be os.Stdout."
}

// Is the Property internal only
func (out *CommandOutputProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (out *CommandOutputProperty) Copy() api_property.Property {
	prop := &CommandOutputProperty{}
	prop.Set(out.Get())
	return api_property.Property(prop)
}

// A command Property for command error output
type CommandErrorProperty struct {
	api_property.WriterProperty
}

// Id for the Property
func (err *CommandErrorProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_ERR
}

// Label for the Property
func (err *CommandErrorProperty) Label() string {
	return "Command error io.Writer."
}

// Description for the Property
func (err *CommandErrorProperty) Description() string {
	return "An io.Writer, which will receive the command execution error output.  Any io.writer can be used, the default here will be os.Stdout."
}

// Is the Property internal only
func (err *CommandErrorProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (err *CommandErrorProperty) Copy() api_property.Property {
	prop := &CommandErrorProperty{}
	prop.Set(err.Get())
	return api_property.Property(prop)
}

// A command Property for command execution input
type CommandInputProperty struct {
	api_property.ReaderProperty
}

// Id for the Property
func (in *CommandInputProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_INPUT
}

// Label for the Property
func (in *CommandInputProperty) Label() string {
	return "Command input io.Reader."
}

// Description for the Property
func (in *CommandInputProperty) Description() string {
	return "An io.Reader, which will provide command execution input.  Any io.reader can be used, the default here will be os.Stdin"
}

// Is the Property internal only
func (in *CommandInputProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (in *CommandInputProperty) Copy() api_property.Property {
	prop := &CommandInputProperty{}
	prop.Set(in.Get())
	return api_property.Property(prop)
}

// A command Property for command execution net context
type CommandContextProperty struct {
	api_property.ContextProperty
}

// Id for the Property
func (cont *CommandContextProperty) Id() string {
	return OPERATION_PROPERTY_COMMAND_CONTEXT
}

// Label for the Property
func (cont *CommandContextProperty) Label() string {
	return "Command context limiter"
}

// Description for the Property
func (cont *CommandContextProperty) Description() string {
	return "A context for controling execution."
}

// Is the Property internal only
func (cont *CommandContextProperty) Usage() api_usage.Usage {
	return api_property.Usage_Optional()
}

// Copy the property
func (cont *CommandContextProperty) Copy() api_property.Property {
	prop := &CommandContextProperty{}
	prop.Set(cont.Get())
	return api_property.Property(prop)
}
