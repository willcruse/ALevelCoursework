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
var loginSuccess = 0

//Main Function
func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/setsPage", setsPage)
	mux.HandleFunc("/termsPage", termsPage)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/signUpPage", signUpPage)
	mux.HandleFunc("uIDRequest", uIDPost)
	mux.HandleFunc("login", login)
	server.ListenAndServe()
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
	fmt.Println(data)
	if len(data) == 0 { //incorrect userName
		loginSuccess = 0
	} else if pw == data[1] { //login
		client.uID = uID
		client.uName = userName
		loginSuccess = 1
		http.ServeFile(res, req, "setsPage.html")
	} else { //incorrect pw
		loginSuccess = 2
	}
	http.ServeFile(res, req, "loginPage.html")
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

func uIDPost(res http.ResponseWriter, req *http.Request) {
	uIDbyte := byte(client.uID)
	var uIDbyteArr []byte
	uIDbyteArr[0] = uIDbyte
	res.Write(uIDbyteArr)
}

func login(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(string(loginSuccess)))
}
