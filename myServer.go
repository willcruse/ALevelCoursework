package main

import (
	"net/http"

	"github.com/ComputingCoursework/dbOperations"
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
	http.ServeFile(res, req, "sets.html")
}

func termsPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "termsPage.html")
		return
	}
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	dbOperations.TermsExisting(termA, termB, setName)
	http.ServeFile(res, req, "termsPage.html")
}
func loginPage(res http.ResponseWriter, req *http.Request) {
	var data []string
	if req.Method != "POST" {
		http.ServeFile(res, req, "loginPage.html")
		return
	}
	uName := req.FormValue("userName")
	pw := req.FormValue("pw")
	data = dbOperations.UserDataUname(uName)
	if len(data) == 0 {
		http.ServeFile(res, req, "index.html")
	}
	if pw == data[1] {
		http.ServeFile(res, req, "sets.html")
	} else {
		http.ServeFile(res, req, "index.html")
	}
}

func signUpPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
	}
	uName := req.FormValue("uName")
	pw := req.FormValue("pw")
	email := req.FormValue("email")
	dbOperations.NewUser(email, uName, pw)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/setsPage", setsPage)
	http.HandleFunc("/termsPage", termsPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/signup", signUpPage)
	http.ListenAndServe(":8080", nil)
}
