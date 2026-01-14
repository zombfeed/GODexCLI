package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
