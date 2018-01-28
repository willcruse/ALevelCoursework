package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewUser(uName, pw string) int {

	if checkTaken(uName) {
		return 0 //Users Name already taken
	}
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping() //Setup db connection
	checkError(errCon)
	stmt, err := db.Prepare("INSERT INTO users (uName, pw) VALUES(?, ?)") //Insert uname and pw
	checkError(err)
	res, err := stmt.Exec(uName, pw) //Exec and check err and resp
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Rows:", affect) //print affected rows
	return 1                     //Success
}

func checkTaken(uName string) bool {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping() //standard db setup
	checkError(errCon)
	rows, err := db.Query("SELECT * FROM users WHERE uName=?", uName) //Query for uname
	defer rows.Close()
	checkError(err)
	var count int
	for rows.Next() {
		count++
	}
	return count > 0
}
