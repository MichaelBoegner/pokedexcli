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
	Locations    []Location `json:"results"`
	NextPage     string     `json:"next"`
	PreviousPage string     `json:"previous"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocalSession struct {
	NextPage     string
	PreviousPage string
}

var (
	localSession = &LocalSession{
		NextPage:     "",
		PreviousPage: "",
	}
	response        *http.Response
	err             error
	unmarshaledBody ResponseBody
)

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
			Description: "Displays the names of the next 20 location areas in the Pokemon world.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			Callback:    commandMapB,
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
		url string
	)
	locations := make([]byte, 0)
	if localSession.NextPage != "" {
		url = localSession.NextPage
	} else {
		url = "https://pokeapi.co/api/v2/location"
	}

	response, err = http.Get(url)
	if err != nil {
		return err
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
	localSession.PreviousPage = unmarshaledBody.PreviousPage

	// Print the response body
	for _, location := range unmarshaledBody.Locations {
		fmt.Println(location.Name)
		locations = append(locations, []byte(location.Name)...)
	}

	// Add the latest data to the cache
	cache.Add(url, locations)
	return nil
}

func commandMapB() error {
	fmt.Printf("\nlocalSession.PreviousPage: %v", localSession.PreviousPage)
	if localSession.PreviousPage != "" {
		response, err = http.Get(localSession.PreviousPage)
		if err != nil {
			return err
		}
	} else {
		return errors.New("No previous page available.")
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
	localSession.PreviousPage = unmarshaledBody.PreviousPage

	// Print the response body
	for _, location := range unmarshaledBody.Locations {
		fmt.Println(location.Name)
	}
	return nil
}
