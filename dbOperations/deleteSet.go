package dbOperations

import (
	"database/sql"
	"log"
)

func DeleteSets(setID, uID int) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	stmt, err := db.Prepare("DELETE FROM terms WHERE (setID=?)")
	checkError(err)
	res, err := stmt.Exec(setID)
	checkError(err)
	log.Println(res.RowsAffected())
	stmt2, err := db.Prepare("DELETE FROM cards WHERE (setID=? AND userOwn=?)")
	checkError(err)
	res, err = stmt2.Exec(setID, uID)
	checkError(err)
	log.Println(res.RowsAffected())
	return
}
