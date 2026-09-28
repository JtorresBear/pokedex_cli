package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JtorresBear/pokedex_cli/internal/pokeapi"
)

type config struct {
	commands     map[string]cliCommand
	client       *pokeapi.Client
	nextPage     *string
	previousPage *string
}
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("there was an error: %v", err)
				break
			}
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	split_input := strings.Fields(input)

	return split_input
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays last 20 map locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows the pokemon encounters for that location. use full location name",
			callback:    commandExplore,
		},
	}
}
