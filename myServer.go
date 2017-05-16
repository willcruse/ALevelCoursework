package main

import (
	"net/http"
)

func setsPage(res http.ResponseWriter, req *http.Request) {
	if req != "POST" {
		http.ServeFile(res, req, "sets.html")
		return
	}

	setName := 

}

func main() {
	http.HandleFunc("/", setsPage)
	http.ListenAndServe(":8080", nil)
}
