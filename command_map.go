package main

import "fmt"

func commandMap(cfg *Config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
