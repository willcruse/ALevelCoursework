package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewUser(email, uName, pw string) int {
	i, b := checkTaken(email, uName)
	if i {
		return 0
	} else if b {
		return 1
	}
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)
	stmt, err := db.Prepare("INSERT INTO users (email, uName, pw) VALUES(?, ?, ?)")
	checkError(err)
	res, err := stmt.Exec(email, uName, pw)
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Rows:", affect)
	return 2
}

func checkTaken(email, uName string) (bool, bool) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)
	rows, err := db.Query("SELECT * FROM users WHERE uName=? OR email=?", uName, email)
	if err != sql.ErrNoRows {
		return false, false
	}
	var uNameRes = "!###!"
	var emailRes = "!##~##!"
	rows.Scan(&uNameRes, &emailRes)
	if uNameRes != "!###!" {
		return true, false
	} else if emailRes != "!##~##!" {
		return false, true
	}
	return false, false
}
