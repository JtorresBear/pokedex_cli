package pokeapi

type Pokemon struct {
	Name       string  `json:"name"`
	Experience int     `json:"base_experience"`
	Height     int     `json:"height"`
	Weight     int     `json:"weight"`
	Stats      []Stats `json:"stats"`
	Types      []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}
