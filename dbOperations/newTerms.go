package dbOperations

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //JUSTIFIED
)

func NewTerm(term1, term2 string, setID int64) int { //Function to make new terms in db from a setID
	var db *sql.DB
	db, err := sql.Open("mysql", "root:somePass@/educationWebsite")
	checkError(err)
	setIDInt := int(setID)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err) //connects to db and checks for errors
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(setIDInt, term1, term2) //inserts the values into the db
	if err != nil {
		return 1
	}
	rows, err := res2.RowsAffected()
	checkError(err)
	log.Println("Success, rows Affected : ", rows)
	return 0
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
