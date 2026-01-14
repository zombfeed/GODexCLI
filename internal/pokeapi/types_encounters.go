package pokeapi

type RespShallowEncounters struct {
	ID                int     `json:"id"`
	Name              *string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"URL"`
		} `json:"pokemon"`
	} `json:"pokemon_ecounters"`
}
