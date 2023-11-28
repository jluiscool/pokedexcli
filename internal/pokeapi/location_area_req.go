package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	//key should be the url, url should return the same []byte
	data, ok := c.cache.Get(fullURL)
	if ok {
		//if ok is true it means cache was hit
		fmt.Println("cache hit")
		//declare empty variable
		locationAreasResp := LocationAreasResp{}
		//unmarshal takes the data and assigns it to a variable with a pointer
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
	} else {
		//if cache was not hit:
		fmt.Println("cache missed")
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//actually executes the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	//closes body after the function returns
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	//reads the body of the response
	//no need to initialize data, just assign
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//declare empty variable
	locationAreasResp := LocationAreasResp{}
	//unmarshal takes the data and assigns it to a variable with a pointer
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//before returning, save cache data
	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

// get a single location area, dont' need a page url, but a name
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//key should be the url, url should return the same []byte
	data, ok := c.cache.Get(fullURL)
	if ok {
		//if ok is true it means cache was hit
		fmt.Println("cache hit")
		//declare empty variable
		locationAreasResp := LocationAreasResp{}
		//unmarshal takes the data and assigns it to a variable with a pointer
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationArea{}, err
		}
	} else {
		//if cache was not hit:
		fmt.Println("cache missed")
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	//actually executes the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	//closes body after the function returns
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v on url %v", res.StatusCode, fullURL)
	}

	//reads the body of the response
	//no need to initialize data, just assign
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	//declare empty variable
	LocationArea := LocationArea{}
	//unmarshal takes the data and assigns it to a variable with a pointer
	err = json.Unmarshal(data, &LocationArea)
	if err != nil {
		return LocationArea, err
	}

	//before returning, save cache data
	c.cache.Add(fullURL, data)

	return LocationArea, nil
}
