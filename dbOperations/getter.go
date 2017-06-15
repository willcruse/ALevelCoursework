package dbOperations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetSets(uID int) []string {
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

func GetTerms(setName string) [][]string {
	var db *sql.DB
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	rows, err := db.Query("SELECT setID FROM cards WHERE setName=?", setName)
	checkError(err)
	var setID int
	err = rows.Scan(&setID)
	rows, err = db.Query("SELECT term1, term2 FROM terms WHERE setID=?", setID)
	checkError(err)
	var finalData [][]string
	for rows.Next() {
		var term1, term2 string
		var data []string
		err = rows.Scan(&term1)
		err = rows.Scan(&term2)
		data = []string{term1, term2}
		finalData = append(finalData, data)
	}
	return finalData
}
