package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// needs user input
func commandCatch(cfg *config, args ...string) error {
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

	const threshold = 50
	//returns a random value from 0 to variable
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	fmt.Printf("%v was caught! \n", pokemon.Name)
	return nil
}
