package dbOperations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getSets(uID int) []string {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	rows, err := db.Query("SELECT setName FROM cards WHERE userOwn=?", uID)
	checkError(err)
	var data []string
	for rows.Next() {
		var setName string
		err = rows.Scan(&setName)
		data = append(data, setName)
	}
	return data
}

func getTerms(setName string) [][]string {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
}
