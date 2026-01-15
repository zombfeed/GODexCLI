package main

import "fmt"

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		return fmt.Errorf("invalid amount of arguments; expected call: 'explore <location_area_name>'")
	}
	encounterResp, err := config.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, enc := range encounterResp.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}

	return nil
}
