package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/willcruse/ComputingCoursework/dbOperations"
)

type clientInfo struct {
	uID   int
	uName string
	pw    string
}

var uID = -1
var client clientInfo
var loginSuccess = 0

//Main Function
func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/setsPage", setsPage)
	mux.HandleFunc("/termsPage", termsPage)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/signUpPage", signUpPage)
	mux.HandleFunc("/loginPage/uIDRequest", uIDPost)
	mux.HandleFunc("/loginPage/login", login)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
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
	/**client.uName = req.FormValue("userName")
	client.pw = req.FormValue("pw")**/
	http.ServeFile(res, req, "loginPage.html")
}**/

func signUpPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signUpPage.html")
		return
	}
	uName := req.FormValue("userName")
	pw := req.FormValue("pw")
	email := req.FormValue("email")
	succ := dbOperations.NewUser(email, uName, pw)
	var result string
	switch succ {
	case 0:
		result = "Username Taken"
	case 1:
		result = "Email already used"
	case 2:
		result = "User Added"
		fmt.Println("ADDED USER")
		http.ServeFile(res, req, "loginPage.html")
	}
	fmt.Println(result)
}

func uIDPost(res http.ResponseWriter, req *http.Request) {
	fmt.Println("UID TRIG")
	uIDbyte := []byte(string(client.uID))
	res.Write(uIDbyte)
	return
}

func login(res http.ResponseWriter, req *http.Request) {
	type Success struct {
		Succ int `json:"loginsuccess"`
		UID  int `json:"UID"`
	}
	uName := req.FormValue("uName")
	pw := req.FormValue("pw")
	var data []string
	var uID int
	data, uID = dbOperations.UserDataUname(uName)
	fmt.Println(data)
	if len(data) == 0 { //incorrect userName
		loginSuccess = 0
	} else if pw == data[1] { //login
		loginSuccess = 1
	} else { //incorrect pw
		loginSuccess = 2
	}
	var success Success
	success = Success{
		UID:  uID,
		Succ: loginSuccess}
	js, err := json.Marshal(success)
	if err != nil {
		fmt.Println("JError", err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
	fmt.Println(string(js))
	return
}
