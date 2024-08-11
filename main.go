package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/michaelboegner/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Second)

func main() {
	// Initialize commands and reader
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

		splitInput := strings.Split(input, " ")
		if command, exists := command[splitInput[0]]; exists {
			if len(splitInput) == 2 {
				err = command.Callback(splitInput[1])
			} else {
				err = command.Callback("")
			}
			if err != nil {
				log.Printf("\nCallback failed. Err: %s\n", err)
			}
		}
	}
}
