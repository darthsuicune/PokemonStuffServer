package dex

type Pokedex struct {
	Pokemon_stats []Pokemon `json:"mon_stats"`
}

type Pokemon struct {
	Id int `json:"id"`
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	Stamina int `json:"stamina"`
}

type Version int

func findVersion() string {
	return "0"
}