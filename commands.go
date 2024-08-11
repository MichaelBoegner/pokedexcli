package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(string) error
}

type ResponseBody struct {
	Locations         []Location   `json:"results"`
	NextPage          string       `json:"next"`
	PreviousPage      string       `json:"previous"`
	PokemonEncounters []Encounters `json:"pokemon_encounters"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Encounters struct {
	Pokemon map[string]string `json:"pokemon"`
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
		"explore": {
			Name:        "explore",
			Description: "explore <area_name> to explore that given area",
			Callback:    commandExplore,
		},
	}
}

func commandHelp(command string) error {
	commands := cliCommand{}.Commands()
	fmt.Printf("Welcome to the Pokedex!\n Usage\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}

func commandExit(command string) error {
	return errors.New("Exiting")
}

func commandMap(command string) error {
	var url string
	locations := make([]byte, 0)

	if localSession.NextPage != "" {
		url = localSession.NextPage
	} else {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	// Check for cached key and val. Return if succesful.
	if cachedEntry, ok := cache.Get(url); ok {
		locationsStr := string(cachedEntry.Data)
		words := strings.Split(locationsStr, " ")
		for _, word := range words {
			fmt.Println(word)
		}
		localSession.NextPage = cachedEntry.NextPage
		localSession.PreviousPage = cachedEntry.PreviousPage
		return nil
	}

	// Get locations
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
		locations = append(locations, ' ')
	}

	// Add the latest data to the cache
	cache.Add(url, localSession.NextPage, localSession.PreviousPage, locations)

	return nil
}

func commandMapB(command string) error {
	var url string

	if localSession.PreviousPage != "" {
		url = localSession.PreviousPage
	} else {
		return errors.New("No previous page available.")
	}

	// Check for cached key and val. Return cached locations if successful.
	if cachedEntry, ok := cache.Get(url); ok {
		locationsStr := string(cachedEntry.Data)
		words := strings.Split(locationsStr, " ")
		for _, word := range words {
			fmt.Println(word)
		}
		localSession.NextPage = cachedEntry.NextPage
		localSession.PreviousPage = cachedEntry.PreviousPage
		return nil
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
	}
	return nil
}

func commandExplore(command string) error {
	pokemons := make([]byte, 0)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", command)

	// Check for cached key and val. Return if succesful.
	if cachedEntry, ok := cache.Get(url); ok {
		pokemonsStr := string(cachedEntry.Data)
		pokemons := strings.Split(pokemonsStr, " ")
		fmt.Printf("\nExploring pastoria-city-area...\nFound Pokemon:\n")
		for _, pokemon := range pokemons {
			fmt.Printf(" - %v\n", pokemon)
		}
		return nil
	}

	// Get locations
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
	fmt.Printf("\nExploring pastoria-city-area...\nFound Pokemon:\n")
	for i, encounter := range unmarshaledBody.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon["name"])
		pokemons = append(pokemons, []byte(encounter.Pokemon["name"])...)
		if i < len(unmarshaledBody.PokemonEncounters)-1 {
			pokemons = append(pokemons, ' ')
		}
	}

	cache.Add(url, "", "", pokemons)

	return nil
}
