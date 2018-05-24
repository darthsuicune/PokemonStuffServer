package dex

import (
	"fmt"
	"github.com/darthsuicune/PokemonStuffServer/db"
)

func findVersion() string {
	var dex = db.GetDex()
	return fmt.Sprintf("%v", dex.Version)
}

func getDex() db.Pokedex {
	return db.GetDex()
}