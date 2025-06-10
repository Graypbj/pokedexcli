package main

import (
	"fmt"
	"github/Graypbj/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	fmt.Println("                                  ,'\\\n" +
		"    _.----.        ____         ,'  _\\   ___    ___     ____\n" +
		"_,-'       `.     |    |  /`.   \\,-'    |   \\  /   |   |    \\  |`.\n" +
		"\\      __    \\    '-.  | /   `.  ___    |    \\/    |   '-.   \\ |  |\n" +
		" \\.    \\ \\   |  __  |  |/    ,','_  `.  |          | __  |    \\|  |\n" +
		"   \\    \\/   /,' _`.|      ,' / / / /   |          ,' _`.|     |  |\n" +
		"    \\     ,-'/  /   \\    ,'   | \\/ / ,`.|         /  /   \\  |     |\n" +
		"     \\    \\ |   \\_/  |   `-.  \\    `'  /|  |    ||   \\_/  | |\\    |\n" +
		"      \\    \\ \\      /       `-.`.___,-' |  |\\  /| \\      /  | |   |\n" +
		"       \\    \\ `.__,'|  |`-._    `|      |__| \\/ |  `.__,'|  | |   |\n" +
		"        \\_.-'       |__|    `-._ |              '-.|     '-.| |   |\n" +
		"                                `'                            '-._|")
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
