package interfaces

import (
	"fmt"
)

var commands map[string]ApplicationCommand

// ApplicationCommand interface
type ApplicationCommand interface {
	Execute()
}

// Add command to commands list
func AddCommand(name string, command ApplicationCommand) {
	if nil == commands {
		commands = make(map[string]ApplicationCommand)
	}
	commands[name] = command
}

func ExecuteCommands(aCommands []string) {
	for _, name := range aCommands {
		fmt.Println(commands[name])
		commands[name].Execute()
	}
}