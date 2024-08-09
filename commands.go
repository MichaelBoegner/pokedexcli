package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type ResponseBody struct {
	Locations []Location `json:"results"`
	NextPage  string     `json:"next"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocalSession struct {
	NextPage string
}

var localSession = &LocalSession{
	NextPage: "",
}

func (c cliCommand) Commands() map[string]cliCommand {
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
		"map": {
			Name:        "map",
			Description: "Displays the names of the last 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations.",
			Callback:    commandMap,
		},
	}
}

func commandHelp() error {
	commands := cliCommand{}.Commands()
	fmt.Printf("Welcome to the Pokedex!\n Usage\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}

func commandExit() error {
	return errors.New("Exiting")
}

func commandMap() error {
	var (
		unmarshaledBody ResponseBody
		response        *http.Response
		err             error
	)

	if localSession.NextPage != "" {
		response, err = http.Get(localSession.NextPage)
		if err != nil {
			return err
		}
	} else {
		response, err = http.Get("https://pokeapi.co/api/v2/location")
		if err != nil {
			return err
		}
	}

	defer response.Body.Close() // Ensure the body is closed after reading

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &unmarshaledBody)

	// Store `next` pagination URL for next 20 locations
	localSession.NextPage = unmarshaledBody.NextPage

	// Print the response body
	for _, location := range unmarshaledBody.Locations {
		fmt.Println(location.Name)
	}
	return nil
}
