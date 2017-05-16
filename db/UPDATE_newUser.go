package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func newUser(email, uName, pw string) {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite") //will:somePass   root:roottoor
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	stmt, err := db.Prepare("INSERT INTO users (email, uName, pw) VALUES(?, ?, ?)")
	checkError(err)
	res, err := stmt.Exec(email, uName, pw)
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Rows:", affect)
}


