package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResponseShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	locationResponse := ResponseShallowLocations{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	
	return locationResponse, nil
}