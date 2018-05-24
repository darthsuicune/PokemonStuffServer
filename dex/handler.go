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
	writeResponse(w, findVersion())

}
func writeResponse(writer http.ResponseWriter, toWrite ...interface{}) {
	enc := json.NewEncoder(writer)
	enc.Encode(toWrite)
}

func getStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	writeResponse(w, getDex())
}