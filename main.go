package main // Required package declaration

import (
	"bufio"
	"fmt"
	"os"
)

// The entry point of the executable
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()  
		input := scanner.Text()
		
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		
		if cmd, ok := commands[firstWord]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}