package dex

import "fmt"

type Pokedex struct {
	Version Version `json:"version"`
	Stats []Pokemon `json:"mon_stats"`
}

type Pokemon struct {
	Id int `json:"id"`
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	Stamina int `json:"stamina"`
}

type Version int

func findVersion() string {
	var dex = getDex()
	return fmt.Sprintf("%v", dex.Version)
}

func getDex() Pokedex {
	var mon1 = Pokemon{0,0,0,0}
	var mon2 = Pokemon{1,1,1,1}
	return Pokedex{
		Version: 0,
		Stats: []Pokemon{mon1, mon2},
	}

}