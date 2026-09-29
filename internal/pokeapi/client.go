package pokeapi

import "github.com/JtorresBear/pokedex_cli/internal/pokecache"

const BaseURL = "https://pokeapi.co/api/v2/location-area/"
const PokeBase = "https://pokeapi.co/api/v2/pokemon/"

type Client struct {
	Cache *pokecache.Cache
}

func NewClient(cache *pokecache.Cache) *Client {
	return &Client{
		Cache: cache,
	}
}
