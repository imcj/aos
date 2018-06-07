package interfaces

import (
	"fmt"
	// "github.com/apex/log"
)

var commands map[string]ApplicationCommand

// ApplicationCommand interface
type ApplicationCommand interface {
	Execute()
}


type ClosureCommand struct {
	closure func()
	// Execute()
}

func (c ClosureCommand)Execute() {
	c.closure()
}

// Add command to commands list
func AddCommand(name string, command ApplicationCommand) {
	if nil == commands {
		commands = make(map[string]ApplicationCommand)
	}
	commands[name] = command
}

func AddClosureCommand(name string, closure func()) {
	if nil == commands {
		commands = make(map[string]ApplicationCommand)
	}
	commands[name] = ClosureCommand{closure}
}

func ExecuteCommands(aCommands []string) {
	for _, name := range aCommands {
		fmt.Println("Will load command " + name)
		// log.Debug("Will load command" + name)
		command := commands[name]
		if nil == command {
			fmt.Println(name + " not exists.")
			continue
		}
		command.Execute()
	}
}