package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("No Pokemon given")
	}

	key := args[0]

	pokemon, exists := cfg.CaughtPokemon[key]
	if !exists {
		return errors.New("Invalid Pokemon given")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf("  - %s\n", stat.Type.Name)
	}

	return nil
}
