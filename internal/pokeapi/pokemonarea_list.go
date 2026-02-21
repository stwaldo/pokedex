package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) ListPokemonAreas(locationArea string) (ResponseLocationArea, error) {

	if locationArea == "" {
		return ResponseLocationArea{}, fmt.Errorf("location area cannot be empty")
	}

	url := baseURL + "/location-area/" + locationArea + "/"
	
	if cached, exists := c.cache.Get(url); exists {
		locationAreaResponse := ResponseLocationArea{}
		err := json.Unmarshal(cached, &locationAreaResponse)
		if err == nil {
			return locationAreaResponse, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	locationAreaResponse := ResponseLocationArea{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	fmt.Printf("Parsed response: %+v\n", locationAreaResponse)

	c.cache.Add(url, data)

	return locationAreaResponse, nil
}