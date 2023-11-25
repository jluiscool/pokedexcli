package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type pokemonCity struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type apiResponse struct {
	Results []pokemonCity `json:"results"`
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Fatal(err)
	}
	results := apiResp.Results
	for c, city := range results {
		fmt.Printf("%s \n", city.Name)
		fmt.Printf("%v \n", c)
	}

	return nil
}
