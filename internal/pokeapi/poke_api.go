package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationsArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (locationsArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationsArea{}, err
	}
	if res.StatusCode > 299 {
		return locationsArea{}, fmt.Errorf("Returned a bad status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	var locations locationsArea
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locations); err != nil {
		return locationsArea{}, err
	}

	return locations, nil
}
