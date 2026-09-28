package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JtorresBear/pokedex_cli/internal/pokecache"
)

type Client struct {
	Cache *pokecache.Cache
}

type locationsArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(cache *pokecache.Cache) *Client {
	return &Client{
		Cache: cache,
	}
}

func (c Client) GetLocations(url string) (locationsArea, error) {
	cache, ok := c.Cache.Get(url)
	if ok {
		var locations locationsArea

		if err := json.Unmarshal(cache, &locations); err != nil {
			return locationsArea{}, err
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locationsArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return locationsArea{}, fmt.Errorf("Returned a bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationsArea{}, err
	}

	var locations locationsArea

	if err := json.Unmarshal(data, &locations); err != nil {
		return locationsArea{}, err
	}

	c.Cache.Add(url, data)

	return locations, nil
}
