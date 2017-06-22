package main

import (
	"net/http"

	"github.com/willcruse/ComputingCoursework/dbOperations"
	"github.com/willcruse/ComputingCoursework/htmlOperations"
)

//Main Function
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/setsPage", setsPage)
	http.HandleFunc("/termsPage", termsPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/signup", signUpPage)
	http.ListenAndServe(":8080", nil)
}

//Page Functions
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func setsPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "sets.html")
		return
	}
	uID := 2
	if uID != -1 {
		htmlOperations.GenerateNewTable(uID)
	}
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
	if req.Method != "POST" {
		http.ServeFile(res, req, "loginPage.html")
		return
	}
	var data []string
	userName := req.FormValue("userName")
	pw := req.FormValue("pw")
	data = dbOperations.UserDataUname(userName)
	if len(data) == 0 {
		http.ServeFile(res, req, "signUpPage.html")
	} else if pw == data[1] {
		http.ServeFile(res, req, "sets.html")
	} else {
		http.ServeFile(res, req, "index.html")
	}
}

func signUpPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signUpPage.html")
		return
	}
	uName := req.FormValue("userName")
	pw := req.FormValue("pw")
	email := req.FormValue("email")
	dbOperations.NewUser(email, uName, pw)
	http.ServeFile(res, req, "loginPage.html")
}
