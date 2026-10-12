package main

import "fmt"

func commandHelp(cfg *config, _ []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.commands {
		fmt.Printf("%-10s %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
