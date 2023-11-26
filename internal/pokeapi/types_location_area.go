package pokeapi

// the best way to represent a json field that is a string but sometimes null is to use a *string type
// when the value is null, the pointer will be nill

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
