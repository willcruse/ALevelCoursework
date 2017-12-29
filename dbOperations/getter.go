package dbOperations

import (
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetSets(uID int) ([][]string, error) { //Gets a list of setnames and their respective ids returned as arrays
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)                                                             //Making sure the db is connected and responding
	rows, err := db.Query("SELECT setName, setID FROM cards WHERE userOwn=?", uID) //Querying db for setNames where the user that owns is uID
	if err == sql.ErrNoRows {
		return nil, errors.New("GetSets: No rows returned")
	}
	checkError(err)
	var data [][]string
	for rows.Next() { //Adds each setname from rows to a slice called data
		var setName string
		var setID int
		err = rows.Scan(&setName, &setID)
		checkError(err)
		tempStr := strconv.Itoa(setID)
		tempArr := []string{tempStr, setName}
		data = append(data, tempArr)
	}
	return data, nil
}

func GetTerms(setID int) ([][]string, error) { //Gets a list of sets and returns as a slice of slices each containing a pair of terms
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)                                                           //Connect to db and ensre it is responding to requests                                                  //as all data will contain the same setID no for loop is needed and setID is set
	rows, err := db.Query("SELECT term1, term2 FROM terms WHERE setID=?", setID) //Query for each pair of terms within which have a setID
	checkError(err)
	var finalData [][]string
	for rows.Next() { //Loops through the rows and adds each pair of terms to the set and appends them as aslice to finalData
		var term1, term2 string
		var data []string
		err = rows.Scan(&term1, &term2)
		data = []string{term1, term2}
		finalData = append(finalData, data)
	}
	if len(finalData) == 0 {
		return finalData, errors.New("GetTerms: No rows returned")
	}
	return finalData, nil
}
