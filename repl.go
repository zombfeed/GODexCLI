package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zombfeed/godexcli/internal/pokeapi"
	"github.com/zombfeed/godexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	Pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandArg := ""
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			if len(words) == 2 {
				commandArg = words[1]
			}
			err := command.callback(config, commandArg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	cleaned := strings.Fields(output)
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show all the pokemon located in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball at a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokedex entry",
			callback:    commandInspect,
		},
	}
}
