package main

import (
	"net/http"

	"github.com/ComputingCoursework/dbOperations"
)

func setsPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "sets.html")
		return
	}
	uID := 2
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	setID := dbOperations.NewSet(setName, uID)
	dbOperations.NewTerm(termA, termB, setID)
	http.ServeFile(res, req, "success.html")
}

func main() {
	http.HandleFunc("/", setsPage)
	http.ListenAndServe(":8080", nil)
}
