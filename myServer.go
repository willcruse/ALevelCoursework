package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/willcruse/ComputingCoursework/dbOperations"
)

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
	mux.HandleFunc("/teacherTools", teacherTools)
	mux.HandleFunc("/loginPage/login", login)
	mux.HandleFunc("/setsPage/getSets", getSets)
	mux.HandleFunc("/setsPage/newSets", newSets)
	mux.HandleFunc("/setsPage/newSetPage", newSetsPage)
	mux.HandleFunc("/teachertools/timer", timer)
	mux.HandleFunc("/teachertools/stopwatch", stopWatch)
	mux.Handle("/teacherScripts/", http.StripPrefix("/teacherScripts", http.FileServer(http.Dir("teacherScripts"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	err := server.ListenAndServe()
	checkErr(err)
}

//Page Functions
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html") //Just serves the file
}

func setsPage(res http.ResponseWriter, req *http.Request) { //The function that is invoked when navigating to the sets page
	if req.Method != "POST" { //Checks to see if the req is a POST request
		http.ServeFile(res, req, "html/setsPage.html")
		return
	}
	type rec struct { //Struct to receive the UID
		UID string
	}
	type send struct { //Struct to export for json
		Sets [][]string `json:"sets"`
	}
	decoder := json.NewDecoder(req.Body) //Makes new decoder --> Decodes JSON into receive struct --> converts the UDI string into an int
	var recS rec
	err := decoder.Decode(&recS)
	checkErr(err)
	uID, err := strconv.Atoi(recS.UID)
	checkErr(err)
	if uID == -1 { //Checks to see if the user is logged in if not redirects to the sign up page --> I know this is annoying for user but its convinent to show why sets arent showing
		fmt.Println("Not logged in")
		http.ServeFile(res, req, "html/signUpPage.html")
		return
	}
	sets, err := dbOperations.GetSets(uID) //Fetches sets using the uID
	checkErr(err)
	sendS := send{sets}                //Makes a send struct containing the returned struct
	dataJS, err := json.Marshal(sendS) //Turns into json
	checkErr(err)
	fmt.Println(string(dataJS))
	fmt.Println(sets)
	res.Header().Set("Content-Type", "application/json") //Sets headers --> Writes data
	res.Write(dataJS)
	return
}

func termsPage(res http.ResponseWriter, req *http.Request) { //makes new terms
	if req.Method != "POST" { //Checks for POST
		http.ServeFile(res, req, "html/termsPage.html")
		return
	}
	type rec struct { //Makes new struct to recieve into
		UID string
	}
	decoder := json.NewDecoder(req.Body) //New Decoder on the body of the request --> Decodes into rec struct --> Converts uID into int
	var recS rec
	err := decoder.Decode(&recS)
	checkErr(err)
	uID, err := strconv.Atoi(recS.UID)
	checkErr(err)
	if uID == -1 {
		fmt.Println("Not logged in")
		http.ServeFile(res, req, "html/signUpPage.html")
	}
	setName := req.FormValue("setName")
	termA := req.FormValue("termA")
	termB := req.FormValue("termB")
	dbOperations.TermsExisting(termA, termB, setName, uID)
	http.ServeFile(res, req, "html/termsPage.html")
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" { //Serves file to client
		http.ServeFile(res, req, "html/loginPage.html")
		return
	}
	http.ServeFile(res, req, "html/loginPage.html")
}

func signUpPage(res http.ResponseWriter, req *http.Request) { //Gets users info and enters into db
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/signUpPage.html")
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
		http.ServeFile(res, req, "html/loginPage.html")
	}
	fmt.Println(result)
}

func login(res http.ResponseWriter, req *http.Request) {
	var loginSuccess int
	type rec struct {
		UName string
		Pw    string
	}
	type Success struct {
		Succ int `json:"loginsuccess"`
		UID  int `json:"UID"`
	}
	decoder := json.NewDecoder(req.Body)
	var recS rec
	err := decoder.Decode(&recS)
	checkErr(err)
	uName := recS.UName
	pw := recS.Pw
	var pwR string
	var uID int
	pwR, uID = dbOperations.UserDataUname(uName)
	fmt.Println("PWR", pwR)
	if pwR == "notFound" { //incorrect userName
		loginSuccess = 0
		fmt.Println("notFound")
	} else if pw == pwR { //login as pw match
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

func getSets(res http.ResponseWriter, req *http.Request) {
	type uidStuct struct {
		UID string
	}
	var uIDs uidStuct
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&uIDs)
	checkErr(err)
	fmt.Println(uIDs.UID)
	uIDInt, err := strconv.Atoi(uIDs.UID)
	checkErr(err)
	data, err := dbOperations.GetSets(uIDInt)
	checkErr(err)
	fmt.Println(data)

}

func newSets(res http.ResponseWriter, req *http.Request) {
	type rec struct {
		setName string
		termA   string
		termB   string
		uID     int
	}
	log.Println("New sets TRIGGEREDDDD")
	var recS rec
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recS)
	checkErr(err)
	setID := dbOperations.NewSet(recS.setName, recS.uID)
	suc := dbOperations.NewTerm(recS.termA, recS.termB, setID)
	type send struct {
		Succ int `json:"success"`
	}
	sendS := &send{suc}
	if sendS.Succ == -10 {
		log.Println("We have reached unreachable code")
	}
	json, err := json.Marshal(sendS)
	checkErr(err)
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
	return
}

func newSetsPage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/newSet.html")
}

func teacherTools(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/teachertools.html")
}

func timer(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/timer.html")
}

func stopWatch(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/stopWatch.html")
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
