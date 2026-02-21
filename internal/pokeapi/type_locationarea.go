package pokeapi

type ResponseLocationArea struct {
	ID int `json:"id"`
	Name string `json:"name"`
	GameIndex int `json:"game_index"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		}
	} `json:"pokemon_encounters"`
}
