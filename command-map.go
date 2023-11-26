package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {

	res, err := cfg.pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	//sets the config's next location url based on the response from API
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous
	return nil

	// type pokemonCity struct {
	// 	Name string `json:"name"`
	// 	URL  string `json:"url"`
	// }

	// type apiResponse struct {
	// 	Results []pokemonCity `json:"results"`
	// }

	// res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }

	// var apiResp apiResponse
	// if err := json.Unmarshal(body, &apiResp); err != nil {
	// 	log.Fatal(err)
	// }
	// results := apiResp.Results
	// fmt.Printf("%v \n", results)
	// for c, city := range results {
	// 	fmt.Printf("%s \n", city.Name)
	// 	fmt.Printf("%v \n", c+1)
	// }

	// return nil
}
