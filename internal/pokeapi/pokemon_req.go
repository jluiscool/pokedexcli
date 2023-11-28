package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// get a single location area, dont' need a page url, but a name
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	//key should be the url, url should return the same []byte
	data, ok := c.cache.Get(fullURL)
	if ok {
		//if ok is true it means cache was hit
		fmt.Println("cache hit")
		//declare empty variable
		pokemon := LocationAreasResp{}
		//unmarshal takes the data and assigns it to a variable with a pointer
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
	} else {
		//if cache was not hit:
		fmt.Println("cache missed")
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	//actually executes the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	//closes body after the function returns
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v on url %v", res.StatusCode, fullURL)
	}

	//reads the body of the response
	//no need to initialize data, just assign
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//declare empty variable
	Pokemon := Pokemon{}
	//unmarshal takes the data and assigns it to a variable with a pointer
	err = json.Unmarshal(data, &Pokemon)
	if err != nil {
		return Pokemon, err
	}

	//before returning, save cache data
	c.cache.Add(fullURL, data)

	return Pokemon, nil
}
