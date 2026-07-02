package main

import (
	"strings"

	"github.com/jnguye580/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of the location areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous location areas in the Pokemon world",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore",
			description: "See a list of all the Pokémon at the given location",
			callback:    commandExplore,
		},
	}
}
