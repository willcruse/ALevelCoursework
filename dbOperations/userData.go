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
	var uNameRes string
	var pw string
	uID := -1
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	err = db.QueryRow("SELECT uID, uName, pw FROM users WHERE uName=?", uName).Scan(&uID, &uNameRes, &pw)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		fmt.Println("No rows")
		return data, uID
	}
	data = append(data, uNameRes)
	data = append(data, pw)
	for i := range data {
		fmt.Println(data[i])
	}
	return data, uID
}
