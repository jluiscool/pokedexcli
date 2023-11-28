package main

import (
	"time"

	"github.com/jluiscool/pokedexcli/internal/pokeapi"
)

type config struct {
	//allows us to keep reusing the client
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
