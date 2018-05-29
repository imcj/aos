package interfaces

import (
	// "github.com/aos-stack/env"
	"container/list"
)

var commands * list.List

// ApplicationCommand interface
type ApplicationCommand interface {
	Execute()
}

// Add command to commands list
func AddCommand(command ApplicationCommand) {
	if nil == commands {
		commands = list.New()
	}
	commands.PushFront(command)
}

func ExecuteCommands() {
	var command ApplicationCommand
	for e := commands.Front(); e != nil; e = e.Next() {
		
		command = e.Value.(ApplicationCommand)
		command.Execute()
	}
}