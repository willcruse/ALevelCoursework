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
	checkError(errCon)
	rows, err := db.Query("SELECT * FROM users WHERE uID=?", uID)
	checkError(err)
	var data []string
	for rows.Next() {
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
	checkError(errCon)
	rows, err := db.Query("SELECT uID, pw FROM users WHERE uName=?", uName)
	fmt.Println("Error", err)
	if err == sql.ErrNoRows {
		fmt.Println("No rows")
		return pw, uID
	}
	if err != nil {
		checkError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var uIDl = -1
		var pwl = "notFound"
		err = rows.Scan(&uID, &pw)
		checkError(err)
		fmt.Println("PW", pw)
		return pwl, uIDl
	}

	return pw, uID
}
