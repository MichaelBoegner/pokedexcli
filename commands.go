package main

import (
	"errors"
	"fmt"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func commandHelp() error {
	fmt.Println("This is helpful")
	return errors.New("Test Error")
}

func Commands() map[string]cliCommand {
	commandsList := make(map[string]cliCommand)
	helpCommand := cliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp,
	}
	commandsList["help"] = helpCommand
	return commandsList
}
