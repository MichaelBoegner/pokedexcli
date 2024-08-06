package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\npokedex> ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Printf("\npokedex> %v", input)
		commandsList := Commands()

		switch input {
		case "exit":
			break
		case "help":
			output := commandsList["help"].Description
			fmt.Printf("\n%v", output)
		}
	}
}
