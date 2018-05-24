package main

import (
	"net/http"
	"log"
	"github.com/darthsuicune/pogo-raids-web/raids"
	"github.com/darthsuicune/pogo-raids-web/dex"
)

func main() {
	raids.AddRoutes()
	dex.AddRoutes()
	println("Listening on port :8008")
	log.Fatal(http.ListenAndServe(":8008", nil))
}