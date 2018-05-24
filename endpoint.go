package main

import (
	"net/http"
	"log"
	"github.com/darthsuicune/PokemonStuffServer/dex"
	"github.com/darthsuicune/PokemonStuffServer/raids"
	"github.com/darthsuicune/PokemonStuffServer/db"
)

func main() {
	db.CreateDbIfNeeded()
	raids.AddRoutes()
	dex.AddRoutes()
	println("Listening on port :8008")
	log.Fatal(http.ListenAndServe(":8008", nil))
}