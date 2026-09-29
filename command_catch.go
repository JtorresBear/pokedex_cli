package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("You need to add a pokemon with the command catch")
		return nil
	}
	pokemonUrl := pokeapi.PokeBase + args[0]

	pokemon, err := cfg.client.GetPokemonStats(pokemonUrl)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	time.Sleep(1 * time.Second)
	caught := pokeCatch(pokemon.Experience)
	if caught {
		fmt.Printf("%v was caught\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}

func pokeCatch(experience int) bool {
	chance := 60 - experience/8

	if chance < 5 {
		chance = 5
	}

	if chance > 60 {
		chance = 60
	}

	caught := rand.Intn(100) < chance
	return caught
}
