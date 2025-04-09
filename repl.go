package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github/Graypbj/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	CaughtPokemon    map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays all locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays all previous locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon in a location. explore [name]",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a named Pokemon. catch [name]",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists all information for a Pokemon if it has been caught. inspect [name]",
			callback:    commandInspect,
		},
	}
}
