package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location area name")
	}
	locationArea := args[0]

	fmt.Printf("Exploring %s...\n", locationArea)
	resp, err := cfg.pokeapiClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
