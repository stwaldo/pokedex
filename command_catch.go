package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("pokemon name cannot be empty")
	} else if len (args) > 1 {
		return fmt.Errorf("too many arguments provided. Usage: catch [pokemon name]")
	}
	pokemonName := args[0]
	pokemonInfo, err := config.Client.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)

	// simulate catch chance based on base experience - higher base experience means harder to catch
	catchChance := rand.Intn(pokemonInfo.BaseExperience)
	if catchChance <= 40 {
		config.Pokedex[pokemonInfo.Name] = pokemonInfo
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}

	return nil
}