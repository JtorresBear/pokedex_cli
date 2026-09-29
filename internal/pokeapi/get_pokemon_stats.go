package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c Client) GetPokemonStats(url string) (Pokemon, error) {
	cache, ok := c.Cache.Get(url)
	if ok {
		var pokemon Pokemon

		if err := json.Unmarshal(cache, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, errors.New("The status code is bad.")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	c.Cache.Add(url, data)
	return pokemon, nil
}
