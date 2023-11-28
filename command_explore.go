package main

import (
	"errors"
	"fmt"
)

// needs user input
func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location area provided")
	}
	if len(args) > 1 {
		return errors.New("too many areas provided")
	}
	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s: \n", locationArea.Name)
	for _, pokemonDetails := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonDetails.Pokemon.Name)
	}
	return nil
}
