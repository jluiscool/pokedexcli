package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	//initiate a reader
	reader := bufio.NewScanner(os.Stdin)
	//welcome message
	fmt.Print("Welcome to the Pokedex, type help to get started: \n")
	for {
		fmt.Print("Pokedex > ")
		// gets user input
		reader.Scan()

		//filter words to get a list of words typed
		cleaned := cleanInput(reader.Text())
		//if nothing was typed, keep the loop going.
		if len(cleaned) == 0 {
			continue
		}

		//command looks for the first word of the list
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) == 0 {
			continue
		}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		//checks to see if command exists from getCommands func, uses commandName as the key
		command, exists := getCommands()[commandName]
		//if it exists
		if exists {
			//err message is the command.Callback()'s return, executes func
			err := command.callback(cfg, args...)
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
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "Explore the pokemon world by displaying the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists pokemon in a a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Try catching a pokemon",
			callback:    commandCatch,
		},
	}
}
