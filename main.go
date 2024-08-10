package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/michaelboegner/pokedexcli/internal/pokecache"
)

func main() {
	// Initialize cache, commands, and reader
	cache := pokecache.NewCache()
	fmt.Printf("\ncache: %v", cache)

	command := cliCommand{}.Commands()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\npokedex> ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if command, exists := command[input]; exists {
			err := command.Callback()
			if err != nil {
				log.Printf("\nCallback failed. Err: %s\n", err)
			}
		}
	}
}
