package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Invalid Pokemon given")
	}
	pokemon := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	caughtChance := rand.IntN(pokemonResp.BaseExperience)
	if caughtChance > pokemonResp.BaseExperience/2 {
		cfg.CaughtPokemon[pokemon] = pokemonResp
		fmt.Printf("%s was caught!\n", pokemon)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
}
