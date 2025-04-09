package pokeapi

type Pokemon struct {
	Next           string
	Prev           string
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Data           struct {
		Height    int `json:"height"`
		Weight    int `json:"weight"`
		Abilities []struct {
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
			Ability  []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"ability"`
		} `json:"abilities"`
		Forms []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"forms"`
		HeldItems []struct {
			Item []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"item"`
		}
	}
}
