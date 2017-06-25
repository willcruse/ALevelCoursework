package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //JUSTIFIED
)

func NewTerm(term1, term2 string, setID int64) { //Function to make new terms in db from a setID
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err) //connects to db and checks for errors
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(setID, term1, term2) //inserts the values into the db
	checkError(err)
	fmt.Println(res2)
}

func TermsExisting(term1, term2, setName string, uID int) { //Func to make new terms in db from setName and uID
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
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
