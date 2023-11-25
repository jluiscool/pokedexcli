package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {

		type cliCommand struct {
			name        string
			description string
			callback    func() error
		}

		fmt.Println("pokedex >")

		// Create a new scanner to read from standard input (keyboard)
		scanner := bufio.NewScanner(os.Stdin)

		// Read the next line from the input
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println("That's not a valid command")
		}

		userInput := scanner.Text()

		if userInput == "help" {
			fmt.Println("This message describes how to use the podedex")
		}
		if userInput == "exit" {
			fmt.Println("Exiting program")
			return
		}
	}
}
