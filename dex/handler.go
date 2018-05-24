package dex

import (
	"net/http"
	"encoding/json"
)

func AddRoutes() {
	http.HandleFunc("/dex/go/version/", getVersion)
	http.HandleFunc("/dex/go/stats/", getStats)
}

func getVersion(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte (findVersion()))
}

func getStats(w http.ResponseWriter, r *http.Request) {
	var mon1 = Pokemon{0,0,0,0}
	var mon2 = Pokemon{1,1,1,1}
	var mons = &Pokedex{
		[]Pokemon{mon1, mon2},
	}
	enc := json.NewEncoder(w)
	enc.Encode(mons)
}