package main

import (
	"net/http"
	"log"
	"github.com/darthsuicune/PokemonStuffServer/dex"
	"github.com/darthsuicune/PokemonStuffServer/raids"
)

func main() {
	raids.AddRoutes()
	dex.AddRoutes()
	println("Listening on port :8008")
	log.Fatal(http.ListenAndServe(":8008", nil))
}