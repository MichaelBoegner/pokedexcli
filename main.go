package main

import (
	"bufio"
	"fmt"
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
		// fmt.Printf("\npokedex> %v", input)
		fmt.Printf("\ncommand[input]: %v", command[input])
	}
}
