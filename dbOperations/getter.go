package dbOperations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetSets(uID int) []string { //Gets a list of setnames and returns as a string array
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)                                                      //Makig sure the db is connected and responding
	rows, err := db.Query("SELECT setName FROM cards WHERE userOwn=?", uID) //Querying db for setNames where the user that owns is uID
	checkError(err)
	var data []string
	for rows.Next() { //Adds each setname from rows to a slice called data
		var setName string
		err = rows.Scan(&setName)
		data = append(data, setName)
	}
	return data
}

func GetTerms(setName string, uID int) [][]string { //Gets a list of sets and returns as a slice of slices each containing a pair of terms
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)                                                                           //Connect to db and ensre it is responding to requests
	rows, err := db.Query("SELECT setID FROM cards WHERE setName=? AND userOwn=?", setName, uID) //Query for setID from setName and uID to ensure no other users terms are displayed
	checkError(err)
	var setID int
	err = rows.Scan(&setID)                                                     //as all data will contain the same setID no for loop is needed and setID is set
	rows, err = db.Query("SELECT term1, term2 FROM terms WHERE setID=?", setID) //Query for each pair of terms within which have a setID
	checkError(err)
	var finalData [][]string
	for rows.Next() { //Loops through the rows and adds each pair of terms to the set and appends them as aslice to finalData
		var term1, term2 string
		var data []string
		err = rows.Scan(&term1)
		err = rows.Scan(&term2)
		data = []string{term1, term2}
		finalData = append(finalData, data)
	}
	return finalData
}
