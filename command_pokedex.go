package main

import "fmt"

func commandPokedex(cfg *config, _ []string) error {
	caughtPokemon := cfg.pokedex
	if len(caughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty. you have to catch some pokemon")
		return nil
	}

	for pokemon := range caughtPokemon {
		fmt.Println(pokemon)
	}
	return nil
}
