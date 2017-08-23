package main

import (
	"fmt"
	"net/http"

	"github.com/willcruse/ComputingCoursework/dbOperations"
)

type clientInfo struct {
	uID      int
	uName    string
	setIDs   []int
	setNames []string
}

var uID = -1
var client clientInfo

//Main Function
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/setsPage", setsPage)
	http.HandleFunc("/termsPage", termsPage)
	http.HandleFunc("/loginPage", loginPage)
	http.HandleFunc("/signUpPage", signUpPage)
	//http.HandleFunc("/getSetsData", getSetsData)
	http.ListenAndServe(":8080", nil)
}

//Page Functions
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func setsPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "setsPage.html")
		return
	}
	if uID == -1 {
		fmt.Println("Not logged in")
		http.ServeFile(res, req, "signUpPage.html")
	}
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	setID := dbOperations.NewSet(setName, uID)
	dbOperations.NewTerm(termA, termB, setID)
	http.ServeFile(res, req, "setsPage.html")
}

func termsPage(res http.ResponseWriter, req *http.Request) { //makes new terms
	if req.Method != "POST" {
		http.ServeFile(res, req, "termsPage.html")
		return
	}
	if uID == -1 {
		fmt.Println("Not logged in")
		http.ServeFile(res, req, "signUpPage.html")
	}
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	dbOperations.TermsExisting(termA, termB, setName, uID)
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
	data, uID = dbOperations.UserDataUname(userName)
	if len(data) == 0 {
		http.ServeFile(res, req, "signUpPage.html")
	} else if pw == data[1] {
		client.uID = uID
		client.uName = userName
		http.ServeFile(res, req, "setsPage.html")
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
	fmt.Println("ADDED USER")
	http.ServeFile(res, req, "loginPage.html")
}
