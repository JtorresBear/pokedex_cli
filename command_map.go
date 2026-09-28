package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.nextPage == nil {
		fmt.Println("You're on the last page")
		return nil
	}
	locations, err := cfg.client.GetLocations(*cfg.nextPage)
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
		fmt.Println("You're on the first page or you haven't started yet.")
		return nil
	}
	locations, err := cfg.client.GetLocations(*cfg.previousPage)
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
