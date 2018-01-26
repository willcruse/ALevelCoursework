package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func userDataUID(uID int) []string {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon) //Connects to db and pings to ensure connection
	rows, err := db.Query("SELECT uName, pw FROM users WHERE uID=?", uID) //Queries for the user data
	checkError(err)
	var data []string
	for rows.Next() { //Extracts the user data
		var uName string
		var email string
		var pw string
		err = rows.Scan(&email, &uName, &pw)
		data = append(data, email)
		data = append(data, uName)
		data = append(data, pw)

	}
	return data
}

func UserDataUname(uName string) (string, int) {
	var db *sql.DB
	uID := -1
	pw := "notFound"
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon) //Connects to db and pings to ensure connection
	rows, err := db.Query("SELECT uID, pw FROM users WHERE uName=?", uName) //Queries for the user Data
	if err == sql.ErrNoRows { //Checks to see if rows were returned
		fmt.Println("No rows")
		return pw, uID
	}
	checkError(err) //Checks for error
	defer rows.Close()
	for rows.Next() { //Extracts data from rows
		var uIDl int
		var pwl string
		err = rows.Scan(&uIDl, &pwl)
		checkError(err)
		return pwl, uIDl
	}
	return pw, uID
}
