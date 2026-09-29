package main

import (
	"time"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
	"github.com/JtorresBear/pokedex_cli/internal/pokecache"
)

func main() {
	initialPage := pokeapi.BaseURL + "?limit=20"

	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(cache)
	cfg := &config{
		commands: getCommands(),
		nextPage: &initialPage,
		client:   client,
		pokedex:  make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
