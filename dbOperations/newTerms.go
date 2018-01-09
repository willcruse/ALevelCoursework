package dbOperations

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql" //JUSTIFIED
)

func NewTerm(term1, term2 string, setID int64, uID int) error { //Function to make new terms in db from a setID
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	setIDInt := int(setID)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err) //connects to db and checks for errors
	rows, err := db.Query("SELECT userOwn FROM cards WHERE setID=?", setID)
	var uIDRes int
	for rows.Next() {
		err = rows.Scan(&uIDRes)
		checkError(err)
	}
	if uIDRes != uID {
		return errors.New("New Terms: Error you do not own the sets")
	}
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(setIDInt, term1, term2) //inserts the values into the db
	if err != nil {
		return err
	}
	rowsAff, err := res2.RowsAffected()
	checkError(err)
	log.Println("Success, rows Affected : ", rowsAff)
	return nil
}

func TermsExisting(term1, term2, setName string, uID int) { //Func to make new terms in db from setName and uID
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)                                                                                  //Connects to db and checks for error
	rows, err := db.Query("SELECT setID FROM cards WHERE setName = ? AND userOwn = ?", setName, uID) //
	checkError(err)
	var setID int
	for rows.Next() {
		err := rows.Scan(&setID)
		checkError(err)
	}

}
