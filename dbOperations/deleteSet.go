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
	checkError(errCon) //Connects to db and pings to ensure connected
	stmt, err := db.Prepare("DELETE FROM terms WHERE (setID=?)")
	checkError(err)
	res, err := stmt.Exec(setID) //Deletes all terms with setID supplied
	checkError(err)
	log.Println(res.RowsAffected())
	stmt2, err := db.Prepare("DELETE FROM cards WHERE (setID=? AND userOwn=?)")
	checkError(err)
	res, err = stmt2.Exec(setID, uID) //Deletes all sets with setID and userOwn as supplied
	checkError(err)
	log.Println(res.RowsAffected())
	return
}
