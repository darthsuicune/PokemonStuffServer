package raids

import "net/http"

func AddRoutes() {
	http.HandleFunc("/raids/view/", addRaid)
	http.HandleFunc("/raids/add/", viewRaid)

}

func addRaid(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("some stuff"))
}

func viewRaid(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("some other stuff"))
}