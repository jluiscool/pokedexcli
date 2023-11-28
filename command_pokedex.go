package main

import (
	"fmt"
)

// needs user input
func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("You have caught:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("%s \n", pokemon.Name)
	}
	return nil
}
