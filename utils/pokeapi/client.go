package pokeapi

import (
	"net/http"
	"time"

	"github.com/yashasgaonkar/pokedex/utils/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout,cachetime time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cachetime),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}