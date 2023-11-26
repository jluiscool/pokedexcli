package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//actually executes the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	//closes body after done reading it
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	//reads the body of the response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	//unmarshal takes the data and assigns it to a variable with a pointer
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
