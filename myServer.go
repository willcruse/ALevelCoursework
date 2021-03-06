package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/willcruse/ComputingCoursework/dbOperations"
	"os"
	"bufio"
	"crypto"
	"encoding/hex"
)

//Main Function
func main() {
	mux := http.NewServeMux()
	server := http.Server{ //Create Custom HTTP server on port 8080 with a Read/Write timeout of 15s
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//Add path recognition match URLs and assign them to relevant functions
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/setsPage", setsPage)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/signUpPage", signUpPage)
	mux.HandleFunc("/teacherTools", teacherTools)
	mux.HandleFunc("/signUpPage/signUp", signUp)
	mux.HandleFunc("/loginPage/login", login)
	mux.HandleFunc("/setsPage/newSets", newSets)
	mux.HandleFunc("/setsPage/deleteSets", deleteSets)
	mux.HandleFunc("/setsPage/getTerms", getTermsFunc)
	mux.HandleFunc("/setsPage/deleteTerms", delTerms)
	mux.HandleFunc("/setsPage/addTerms", addTerms)
	mux.HandleFunc("/teachertools/timer", timer)
	mux.HandleFunc("/teachertools/stopwatch", stopWatch)
	mux.HandleFunc("/games/quizMove", quizMove)
	mux.HandleFunc("/games/getFirstTerm", getFirstTerm)
	mux.HandleFunc("/games/checkQuizRes", checkQuizRes)
	//Add path recognition to match URLs for static resources such as style sheets and js
	mux.Handle("/html/cache/", http.StripPrefix("/html/cache/", http.FileServer(http.Dir("html/cache"))))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	err := server.ListenAndServe() //Set the server to start listening
	checkErr(err)                  //Check that the server started up correctly
}


//Page Functions
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html") //Just serves the file
}

func setsPage(res http.ResponseWriter, req *http.Request) { //The function that is invoked when navigating to the sets page
	if req.Method != "POST" { //Checks to see if the req is a POST request and if is not serves the HTML
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
	res.Header().Set("Content-Type", "application/json") //Sets headers --> Writes data
	res.Write(dataJS)
	return
}


func loginPage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/loginPage.html") //Serves the client with the logingPage html file
}

func signUpPage(res http.ResponseWriter, req *http.Request) { //Delivers sign Page
	http.ServeFile(res, req, "html/signUpPage.html")
}

func login(res http.ResponseWriter, req *http.Request) {
	var loginSuccess int //Defines login success in the scope above other sub-scopes allowing it to be used later out
	type rec struct {    //struct for recieving from the json data
		UName string
		Pw    string
		pwHash string
	}
	type Success struct { //Struct to send to the client
		Succ int `json:"loginsuccess"` //Success code
		UID  int `json:"UID"`          //Uid to be placed as a cookie
	}
	decoder := json.NewDecoder(req.Body) //Creates a json decoder for the request body
	var recS rec
	err := decoder.Decode(&recS) //Decodes the json into an instance of the rec struct
	checkErr(err)
	recS.pwHash = getHash(recS.Pw)
	var pwR string
	var uID int
	pwR, uID = dbOperations.UserDataUname(recS.UName) //Fetches UserData linked to that username and returns the pw and uID linked to that user
	if pwR == "notFound" {                       //incorrect userName
		loginSuccess = 0
		fmt.Println("notFound")
	} else if recS.pwHash == pwR { //login as pw match
		loginSuccess = 1
	} else { //incorrect pw
		loginSuccess = 2
	}
	var success Success
	success = Success{ //Creates a new instance of the Success struct with the returned uID and the loginSuccess code
		UID:  uID,
		Succ: loginSuccess}
	js, err := json.Marshal(success) //Turns the struct into json to be sent
	if err != nil {
		fmt.Println("JError", err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js) //Sends the json to the client with the correct headers
	return
}

func signUp(res http.ResponseWriter, req *http.Request) {
	type rec struct { //Struct to recieve data into
		UName string
		PW    string
		pwHash string
	}
	var recS rec
	decoder := json.NewDecoder(req.Body) //Creates a new decoder then decoded the received data into it
	err := decoder.Decode(&recS)
	checkErr(err)
	recS.pwHash = getHash(recS.PW)
	resp := dbOperations.NewUser(recS.UName, recS.pwHash) //Inserts a new user into the db with the received data
	type resS struct {                                //Defines a new struct for the success code
		Succ int `json:"success"`
	}
	resSS := resS{ //Makes a new instance of this struct with the response code
		Succ: resp}
	jsonR, err := json.Marshal(resSS) //Converts to json
	checkErr(err)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonR) //Sends to client with appropriate headers
	return
}

func newSets(res http.ResponseWriter, req *http.Request) {
	type rec struct { //Struct to receive data into
		SetName  string
		UID      string
		uIDTrans int
	}
	var recS rec
	decoder := json.NewDecoder(req.Body) //Creates a new decoder then decodes the data recieved into the struct
	err := decoder.Decode(&recS)
	checkErr(err)
	recS.uIDTrans, err = strconv.Atoi(recS.UID) //Converts the uID received into an int
	checkErr(err)
	setID := dbOperations.NewSet(recS.SetName, recS.uIDTrans) //Makes a new set with the appropriate set name belonging to the right user
	var suc int
	if setID == -5 { //Generates the correct success code
		log.Println("Set already exists")
		suc = 1
	} else {
		suc = 0
	}
	type send struct {
		Succ int `json:"success"`
	}
	sendS := &send{suc} //Creates an instance of the send struct that contains the success code
	json, err := json.Marshal(sendS)
	checkErr(err)
	res.Header().Set("Content-Type", "application/json")
	res.Write(json) //Writes json to the client with the right headers
	return
}

func deleteSets(res http.ResponseWriter, req *http.Request) {
	type rec struct { //Struct to receive data into
		SetID    int
		UID      string
		UIDTrans int
	}
	var recS rec
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recS) //Decodes the recieved json into the rec stuct
	checkErr(err)
	recS.UIDTrans, err = strconv.Atoi(recS.UID)        //converts the received uID into an int
	dbOperations.DeleteSets(recS.SetID, recS.UIDTrans) //Deletes the set
	res.Write([]byte(":)"))                            //Writes to the client so it knows to update
}

func getTermsFunc(res http.ResponseWriter, req *http.Request) {
	type rec struct { //Struct to recieve data into
		SetID int
	}
	var recS rec
	decoder := json.NewDecoder(req.Body) //Creates a new decoder then decodes the data into it
	err := decoder.Decode(&recS)
	checkErr(err)
	terms, err := dbOperations.GetTerms(recS.SetID) //Fetches all the terms with the relevent setID
	checkErr(err)
	type send struct { //Struct to send the terms
		Terms [][]string `json:"terms"` //2D array each sub array contains a pair of terms ie. [[1, 2], [3, 4]]
	}
	sendS := send{terms}             //Creates an instance of this struct with the relevent terms
	json, err := json.Marshal(sendS) //Converts this to kson
	res.Header().Set("Content-Type", "application/json")
	res.Write(json) //Writes to client with the relevent headers
	return
}

func delTerms(res http.ResponseWriter, req *http.Request) {
	type rec struct {
		SetID int
		Term  []string
	}
	var recS rec
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recS)
	checkErr(err)
	log.Println("Terms (delTerms Func) ", recS.Term)
	log.Println("ID (delTerms Func) ", recS.SetID)
	dbOperations.DeleteTerms(recS.SetID, recS.Term)
	res.Write([]byte(":)")) //So the page knows to update the sets table
	return
}

func addTerms(res http.ResponseWriter, req *http.Request) {
	type rec struct { //struct to receive data into
		SetID    int
		TermA    string
		TermB    string
		UID      string
		uIDTrans int
	}
	var recS rec
	decoder := json.NewDecoder(req.Body) //Creates a new decoder which decodes the recieved data into the rec struct
	err := decoder.Decode(&recS)
	checkErr(err)
	setID64 := int64(recS.SetID)                //Converts the recieved setID to an int64
	recS.uIDTrans, err = strconv.Atoi(recS.UID) //Converts the received uID to an int
	checkErr(err)
	err = dbOperations.NewTerm(recS.TermA, recS.TermB, setID64, recS.uIDTrans) //Adds the terms to the db
	checkErr(err)
	res.Write([]byte(":)")) //writes so the page knows to update the sets table
	return
}

func quizMove(res http.ResponseWriter, req *http.Request) {
	type rec struct {
		ID int
	} //struct to be received into
	var recS rec
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recS)
	checkErr(err) //Decodes json into recS
	htmlTemplate := template.New("quiz") //makes a new template called quiz
	htmlTemplate, err = template.ParseFiles("html/quiz.html") //parses the quiz template
	checkErr(err)
	fileName := "html/cache/"
	fileName += strconv.Itoa(recS.ID)
	fileName += ".html" //Creates new HTML fileName for the template to be saved into with the setID as the fileName
	file, err := os.Create(fileName) //Creates the file
	checkErr(err)
	writer := bufio.NewWriter(file) //New buffered writer to write into the file
	err = htmlTemplate.Execute(writer, recS.ID) //Executes the template writing the result to the file
	checkErr(err)
	writer.Flush() //Flushes the writer to make sure it clear
	file.Close() //Closes file
	res.Write([]byte(":)")) //writes so the page knows to redirect
}

func getFirstTerm(res http.ResponseWriter, req *http.Request) {
	type rec struct {
		ID int
	} //Struct to receive into
	type send struct {
		Term []string `json:"terms"`
	} //Struct to send out of
	var recS rec
	decoder := json.NewDecoder(req.Body) //new json decoder which decodes into an instance of rec struct
	err := decoder.Decode(&recS)
	checkErr(err)
	terms, err := dbOperations.GetTerms(recS.ID) //Gets terms from appropriate set
	var termA []string
	for _, value := range terms {
		termA = append(termA, value[0]) //Gets first term value
	}
	sendS := send{termA} //Puts array into new instance of send Struct
	json, err := json.Marshal(sendS) //Converts to json
	res.Header().Set("Content-Type", "application/json")
	res.Write(json) //Writes to client with the relevant headers
	return
}

func checkQuizRes(res http.ResponseWriter, req *http.Request){
	type rec struct { //Struct to be received into
		ID int
		Ans []string
	}
	type send struct { //Struct to send from
		Ans []string 	`json:"ansArr"`
		Cor []bool		`json:"corArr"`
		Score int 		`json:"score"`
	}
	var recS rec
	var sendS send
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recS) //Creates new json decoder which decodes into an instance of rec Struct
	checkErr(err)
	terms, err := dbOperations.GetTerms(recS.ID) //Fetches terms with appropriate id
	var termB []string
	for _, val := range terms {
		termB = append(termB, val[1]) //Appends teh second terms to an array
	}
	for index, val := range termB {
			if index < len(recS.Ans) { //Checks that index is available in received array
				sendS.Cor = append(sendS.Cor, val == recS.Ans[index]) //Appends whether the users anwser and the stored answer match
				sendS.Ans = append(sendS.Ans, val) //Appends answer to ans struct
				if val == recS.Ans[index] {
					sendS.Score++ //Increments score if appropriate
				}
			}
	}
	json, err := json.Marshal(sendS) //Creates json from sendS struct
	res.Header().Set("Content-Type", "application/json")
	res.Write(json) //Writes to client with the relevant headers
	return
}

func teacherTools(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/teachertools.html") //Serves the html
}

func timer(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/timer.html") //Serves the html
}

func stopWatch(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/stopWatch.html") //Serves the html
}

func checkErr(e error) {
	if e != nil {
		log.Println(e) //Checks for error and if there is an error it prints with a timestamp
	}
}

func getHash(text string) string {
	h := crypto.SHA256.New() //Creates new SHA256 hasher
	h.Write([]byte(text)) //Hashes the pw
	return hex.EncodeToString(h.Sum(nil)) //Returns the string version of the result
}
