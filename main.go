package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
		if description, exists := command[input]; exists {
			fmt.Printf("\n%v: %v", description.Name, description.Description)
			err := command[input].Callback
			if err != nil {
				log.Printf("Callback failed.")
			}
		}

	}
}
