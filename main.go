package main

import "github.com/jluiscool/pokedexcli/internal/pokeapi"

type config struct {
	//allows us to keep reusing the client
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
