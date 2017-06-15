package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //JUSTIFIED
)

func NewTerm(term1, term2 string, setID int64) { //Function to make new terms in db
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

func TermsExisting(term1, term2, setName string) {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	var setID int
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)
	rows, err := db.Query("SELECT setID FROM cards WHERE setName = ?", setName)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&setID)
		checkError(err)
	}

}
