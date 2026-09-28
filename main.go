package main

import (
	"time"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
	"github.com/JtorresBear/pokedex_cli/internal/pokecache"
)

func main() {
	initialPage := "https://pokeapi.co/api/v2/location-area/?limit=20"
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(cache)
	cfg := &config{
		commands: getCommands(),
		nextPage: &initialPage,
		client:   client,
	}
	startRepl(cfg)
}
