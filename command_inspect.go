package main

import (
	"errors"
	"fmt"
)

// needs user input
func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon provided")
	}
	if len(args) > 1 {
		return errors.New("too many pokemon provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	//make a new key with that pokemon
	pokemonDetails, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	pokemonDetailList := pokemonDetails.Abilities

	fmt.Printf("%v has these different abilities:\n", pokemon.Name)
	for _, pokemonDetail := range pokemonDetailList {
		fmt.Printf("%v \n", pokemonDetail.Ability.Name)
	}
	return nil
}
