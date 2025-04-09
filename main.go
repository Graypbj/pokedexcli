package main

import (
	"github/Graypbj/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
