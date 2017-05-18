package main

import (
	"net/http"

	"github.com/ComputingCoursework/dbOperations"
	_ "github.com/go-sql-driver/mysql"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

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

func termsPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "termsPage.html")
		return
	}
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	setID := db.
		dbOperations.NewTerm(termA, termB, setID)
	http.ServeFile(res, req, "success.html")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/setsPage", setsPage)
	http.HandleFunc("/termsPage", termsPage)
	http.ListenAndServe(":8080", nil)
}
