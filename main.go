package main

func main() {
	initialPage := "https://pokeapi.co/api/v2/location-area/?limit=20"
	cfg := &config{
		commands: getCommands(),
		nextPage: &initialPage,
	}
	startRepl(cfg)
}
