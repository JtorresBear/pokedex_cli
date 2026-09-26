package main

import (
	"fmt"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.nextPage == nil {
		fmt.Println("You're on the last page")
		return nil
	}
	locations, err := pokeapi.GetLocations(*cfg.nextPage)
	if err != nil {
		return err
	}

	cfg.nextPage = locations.Next

	cfg.previousPage = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousPage == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	locations, err := pokeapi.GetLocations(*cfg.previousPage)
	if err != nil {
		return err
	}
	cfg.nextPage = locations.Next

	cfg.previousPage = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
