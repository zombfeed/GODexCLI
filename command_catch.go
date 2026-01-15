package main

import (
	"fmt"
	"math/rand"

	"github.com/zombfeed/godexcli/internal/pokeapi"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		return fmt.Errorf("invalid amount of arguments; expected call: 'catch <pokemon_name>'")
	}

	catchResp, err := config.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if !catchResp.IsEmpty() {
		if caught := attemptCatch(catchResp); caught {
			if _, ok := config.Pokedex[args[0]]; !ok {
				config.Pokedex[args[0]] = catchResp
			}
			fmt.Println("You may now inspect it with the inspect command.")
		}
	}

	return nil
}

func attemptCatch(pokemon pokeapi.Pokemon) bool {
	baseExperience := pokemon.BaseExperience
	catchValue := (256.0 / (float64(baseExperience + 1))) * (5.0 / 256.0) * 10.0
	if catchValue >= 1.0 || rand.Float64() <= catchValue {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return true
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return false
	}
}
