package main

import (
	"time"

	"github.com/stwaldo/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	config := &config{
		Client: client,
	}
	startRepl(config)
}
