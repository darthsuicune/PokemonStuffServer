package db

type Version int

type Pokedex struct {
	Version Version `json:"version"`
	Stats []Pokemon `json:"mon_stats"`
}

type Pokemon struct {
	Id int `json:"id"`
	Form int`json:"form"`
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	Stamina int `json:"stamina"`
}

const version = 0
const dexQuery = "SELECT id, form, attack, defense, stamina FROM pokemon"

func GetDex() Pokedex {
	pokedex := Pokedex {
		Version: 0,
	}
	database := OpenDb()

	rows, err := database.Query(dexQuery)
	checkError(err)
	if rows != nil {
		for rows.Next() {
			var mon = Pokemon{}
			err := rows.Scan(&mon.Id, &mon.Form, &mon.Attack, &mon.Defense, &mon.Stamina)
			checkError(err)
			pokedex.Stats = append(pokedex.Stats, mon)
		}
	}
	return pokedex
}