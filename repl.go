package main

import (
	"strings"
)

func cleanInput(text string) []string{
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}