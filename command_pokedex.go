package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	if len(config.Pokedex) == 0 {
		return fmt.Errorf("There are no entries in the Pokedex")
	}
	fmt.Println("Your Pokedex:")
	for k := range config.Pokedex {
		fmt.Printf("- %s\n", k)
	}

	return nil
}
