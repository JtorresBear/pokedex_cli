package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemonList(url string) ([]PokemonEncounter, error) {
	cache, ok := c.Cache.Get(url)
	if ok {
		var location location

		if err := json.Unmarshal(cache, &location); err != nil {
			return []PokemonEncounter{}, err
		}
		return location.PokemonEncounters, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []PokemonEncounter{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return []PokemonEncounter{}, fmt.Errorf("Returned a bad status code")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []PokemonEncounter{}, err
	}

	var location location

	if err := json.Unmarshal(data, &location); err != nil {
		return []PokemonEncounter{}, err
	}
	c.Cache.Add(url, data)

	return location.PokemonEncounters, nil
}
