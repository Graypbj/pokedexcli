package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex")
	pokemon := cfg.CaughtPokemon
	for _, singlePokemon := range pokemon {
		fmt.Printf(" - %s\n", singlePokemon.Name)
	}
	return nil
}
