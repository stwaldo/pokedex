package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many arguments provided. Usage: pokedex")
	}
	if len(config.Pokedex) == 0 {
		fmt.Println("your pokedex is empty. Try catching some pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range config.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
