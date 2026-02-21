package main

import (
	"fmt"
)

func commmandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("location area cannot be empty")
	}
	locationArea := args[0]
	locationAreaResponse, err := config.Client.ListPokemonAreas(locationArea)
	if err != nil {
		return err
	}

	for _, pokemon := range locationAreaResponse.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
