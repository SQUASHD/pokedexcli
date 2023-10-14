package pokeapi

import (
	"net/http"
	"time"

	"github.com/squashd/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient
// timeout - timeout for http requests
// cacheInterval - interval to cache http requests
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
