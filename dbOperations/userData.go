package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func userDataUID(uID int) []string {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	rows, err := db.Query("SELECT * FROM users WHERE uID=?", uID)
	checkError(err)
	var data []string
	for rows.Next() {
		var email string
		var uName string
		var pw string
		err = rows.Scan(&email, &uName, &pw)
		data = append(data, email)
		data = append(data, uName)
		data = append(data, pw)

	}
	return data
}

func UserDataUname(uName string) ([]string, int) {
	var db *sql.DB
	var data []string
	uID := -1
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	rows, err := db.Query("SELECT uID, pw FROM users WHERE uName=?", uName)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return data, uID
	}

	for rows.Next() {
		var uName string
		var pw string
		err = rows.Scan(&uName, &pw, &uID)
		data = append(data, uName)
		data = append(data, pw)
	}
	return data, uID
}
