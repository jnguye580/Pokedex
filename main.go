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
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
	}

	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		firstWord := cleanedInput[0]
		args := cleanedInput[1:]

		if cmd, ok := commands[firstWord]; ok {
			if err := cmd.callback(cfg, args); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}