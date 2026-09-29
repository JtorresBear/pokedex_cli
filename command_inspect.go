package main

import (
	"fmt"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
)

func commandInspect(cfg *config, args []string) error {
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	printStats(pokemon.Stats)
	fmt.Println("Types")
	printTypes(pokemon)

	return nil
}

func printStats(stats []pokeapi.Stats) {

	for _, stat := range stats {
		statName := stat.Stat.Name
		baseStat := stat.BaseStat
		fmt.Printf(" - %v: %v\n", statName, baseStat)
	}
}
func printTypes(pokemon pokeapi.Pokemon) {
	types := pokemon.Types

	for _, x := range types {
		fmt.Printf(" - %v\n", x.Type.Name)
	}
}
