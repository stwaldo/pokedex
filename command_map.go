package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config, args ...string) error {
	locationResponse, err := config.Client.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(config *config, args ...string) error {
	if config.Previous == nil {
		return errors.New("you're on the first page.")
	}

	locationResponse, err := config.Client.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous
	
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
