package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stwaldo/pokedexcli/internal/pokeapi"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Print("Unknown command\n")
			continue
		}
		err := command.callback(config)
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "Displays the previous page of locations in the Pokemon world",
			callback:	commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type config struct {
	Next     *string
	Previous *string
	Client pokeapi.Client
}
