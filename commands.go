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

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("This is helpful")
	return errors.New("Test Error")
}

func commandExit() error {
	fmt.Println("This exits the program.")
	return errors.New("couldn't exit")
}
