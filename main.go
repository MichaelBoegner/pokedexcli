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

		fmt.Print("Enter something: ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Printf("\nInput: %v", input)
		if input == "exit()" {
			break
		}
		fmt.Println("You entered:", input)
	}
}
