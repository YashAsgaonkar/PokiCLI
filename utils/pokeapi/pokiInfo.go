package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(url); ok {
		var pokemonInfo Pokemon
		if err := json.Unmarshal(val, &pokemonInfo); err != nil {
			return Pokemon{}, err
		}
		return pokemonInfo, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	var pokemonInfo Pokemon
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pokemonInfo); err != nil {
		return Pokemon{}, err
	}
	return pokemonInfo, nil
}