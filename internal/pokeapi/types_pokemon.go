package pokeapi

type Pokemon struct {
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Name           string  `json:"name"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
	Weight         int     `json:"weight"`
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Types struct {
	Slot int `json:"name"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func (p *Pokemon) IsEmpty() bool {
	return p.BaseExperience == 0 && p.Height == 0 && p.Name == "" && len(p.Stats) == 0 && len(p.Types) == 0 && p.Weight == 0
}
