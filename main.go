package main

import (
	"time"

	"github.com/stwaldo/pokedexcli/internal/pokeapi"
	"github.com/stwaldo/pokedexcli/internal/pokecache"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	config := &config{
		Client: client,
		Cache: cache,
	}
	startRepl(config)
}
