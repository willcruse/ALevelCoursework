package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewSet(setName string, uID int) int64 { //function to create a new set
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)
	stmtC, err := db.Prepare("INSERT INTO cards (userOwn, setName) VALUES(3, ?)")
	checkError(err)
	res, err := stmtC.Exec(setName)
	checkError(err)
	fmt.Println(res)
	lastID, err := res.LastInsertId()
	checkError(err)
	return lastID
}
