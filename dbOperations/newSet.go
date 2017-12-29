package dbOperations

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewSet(setName string, uID int) int64 { //function to create a new set
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)                                                                              //Checks for errors in the dbconnection
	rows, errN := db.Query("SELECT setID FROM cards WHERE (setName=? AND userOwn=?)", setName, uID) //Query for sets which share the setName and user
	log.Println("ErrN: ", errN)
	defer rows.Close()
	checkError(err)
	for rows.Next() {
		log.Println("We did rows exists")
		return -5 //err code if set already exists
	}
	stmtC, err := db.Prepare("INSERT INTO cards (userOwn, setName) VALUES(?, ?)") //Prepares statement that creates new set
	checkError(err)
	defer stmtC.Close()
	res, err := stmtC.Exec(uID, setName)
	checkError(err)
	log.Println("Result ", res)
	lastID, err := res.LastInsertId()
	checkError(err)
	log.Println("We did insert")
	return lastID //Returns id of the set
}
