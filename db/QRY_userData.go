package dbOperations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func userData(uID int) []string {
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
