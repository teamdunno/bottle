package bot

type CommandFunc func(ctx Context)

type Command struct {
	Run  CommandFunc
	Help string
}

type CommandRegistry struct {
	commands map[string]Command
}

func NewCommandRegistry() CommandRegistry {
	return CommandRegistry{
		commands: make(map[string]Command),
	}
}

func (r *CommandRegistry) AddCommand(name string, command CommandFunc) {
	r.commands[name] = Command{
		Run: command,
	}
}

func (r *CommandRegistry) SetHelp(name string, help string) {
	command, exists := r.GetCommand(name)
	if !exists {
		return
	}
	command.Help = help
	r.commands[name] = *command
}

func (r *CommandRegistry) RemoveCommand(name string) {
	delete(r.commands, name)
}

func (r *CommandRegistry) GetCommand(name string) (*Command, bool) {
	command, exists := r.commands[name]
	return &command, exists
}

func (r *CommandRegistry) GetCommands() map[string]Command {
	return r.commands
}

func (r *CommandRegistry) ExecuteCommand(name string, ctx Context) bool {
	command, exists := r.GetCommand(name)
	if !exists {
		return false
	}
	command.Run(ctx)
	return true
}
