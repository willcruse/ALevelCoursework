package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewTerm(term1, term2 string, setID int) { //Function to make new terms in db
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(setID, term1, term2)
	checkError(err)
	fmt.Println(res2)
}
