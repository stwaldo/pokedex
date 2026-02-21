package main

import (
	"errors"
	"fmt"
)

type locationListResponse struct {
	Count	int `json:"count"`
	Next	string `json:"next"`
	Previous string `json:"previous"`
	Results []locationListItem `json:"results"`
}

type locationListItem struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func commandMap(config *config) error {
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

func commandMapB(config *config) error {
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
