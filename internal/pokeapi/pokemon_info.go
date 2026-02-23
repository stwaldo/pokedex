package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	if pokemonName == "" {
		return Pokemon{}, fmt.Errorf("pokemon name cannot be empty")
	}

	url := baseURL + "/pokemon/" + pokemonName + "/"
	
	if cached, exists := c.cache.Get(url); exists {
		pokemonInfo := Pokemon{}
		err := json.Unmarshal(cached, &pokemonInfo)
		if err == nil {
			return pokemonInfo, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	
	pokemonInfo := Pokemon{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonInfo, nil
}