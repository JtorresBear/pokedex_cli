package main

import (
	"fmt"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	fmt.Println("explore was called")
	if len(args) == 0 {
		fmt.Println("You need to add an area after \"explore\"")
		return nil
	}
	locationURL := pokeapi.BaseURL + args[0]

	pokemonEncounters, err := cfg.client.GetPokemonList(locationURL)
	if err != nil {
		return err
	}

	for _, encounter := range pokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
