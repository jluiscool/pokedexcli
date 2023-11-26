package pokeapi

import (
	"net/http"
	"time"
)

// string from pokeapi url
const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

// always a good idea to make a function for the http client, kind of like a constructor in an object oriented language
// gives the client a timeout, the timeout is set to one minute
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
