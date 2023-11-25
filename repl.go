package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//initiate a reader
	reader := bufio.NewScanner(os.Stdin)
	//welcome message
	fmt.Print("Welcome to the Pokedex, type help to get started: \n")
	for {
		fmt.Print("Pokedex > ")
		// gets user input
		reader.Scan()

		//filter words to get a list of words typed
		words := cleanInput(reader.Text())
		//if nothing was typed, keep the loop going.
		if len(words) == 0 {
			continue
		}

		//command looks for the first word of the list
		commandName := words[0]

		//checks to see if command exists from getCommands func, uses commandName as the key
		command, exists := getCommands()[commandName]
		//if it exists
		if exists {
			//err message is the command.Callback()'s return, executes func
			err := command.callback()
			//prints error if any
			if err != nil {
				fmt.Println(err)
			}
			continue
			//if it doesn't exists
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleans input from CLI
func cleanInput(text string) []string {
	//make words to lowercase
	output := strings.ToLower(text)
	// splits the lowercase string into a slice of words separated by space
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// list of commands
func getCommands() map[string]cliCommand {
	//returns a map with a key string, and a value that is cliCommand structured
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
