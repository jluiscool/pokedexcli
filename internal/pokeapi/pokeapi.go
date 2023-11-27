package pokeapi

import (
	"net/http"
	"time"

	pokeacache "github.com/jluiscool/pokedexcli/internal/pokecache"
)

// string from pokeapi url
const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokeacache.Cache
	httpClient http.Client
}

// always a good idea to make a function for the http client, kind of like a constructor in an object oriented language
// gives the client a timeout, the timeout is set to one minute
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokeacache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
