package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	//adds a new line
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	//prints commands and their description
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	//returns no error
	return nil
}
