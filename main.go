package main // Required package declaration

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/jnguye580/pokedexcli/internal/pokeapi"
)

// The entry point of the executable
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]

		if cmd, ok := commands[firstWord]; ok {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}