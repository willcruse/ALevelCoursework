package dbOperations

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteTerms(setID int, terms []string) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	stmt, err := db.Prepare("DELETE FROM terms WHERE (setID=? AND term1=? AND term2=?)")
	checkError(err)
	log.Println("Terms Array: ", terms)
	res, err := stmt.Exec(setID, terms[0], terms[1])
	checkError(err)
	log.Println(res)
}
