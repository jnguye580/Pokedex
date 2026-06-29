package main // Required package declaration

import (
	"bufio"
	"fmt"
	"os"
)

// The entry point of the executable
func main() {
    scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()  
		input := scanner.Text()
		
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}