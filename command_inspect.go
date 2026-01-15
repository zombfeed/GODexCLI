package main

import "fmt"

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		return fmt.Errorf("invalid amount of arguments; expected call: 'inspect <pokemon_name>'")
	}

	if pokemon, ok := config.Pokedex[args[0]]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, ptype := range pokemon.Types {
			fmt.Printf("- %s\n", ptype.Type.Name)
		}
	} else {
		fmt.Printf("No entry for %s found\n", args[0])
	}
	return nil
}
