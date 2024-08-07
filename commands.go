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

func (c cliCommand) Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
	}
}

func commandHelp() error {
	commands := cliCommand{}.Commands()
	fmt.Printf("Welcome to the Pokedex!\n Usage\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}

	return errors.New("Command not found")
}
