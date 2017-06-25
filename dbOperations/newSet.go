package dbOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewSet(setName string, uID int) int64 { //function to create a new set
	var db *sql.DB
	var noRows error
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	checkError(err)                                                                                  //Checks for errors in the dbconnection
	rows, err := db.Query("SELECT setID FROM cards WHERE setName = ? AND userOwn = ?", setName, uID) //Query for sets which share the setName and user
	if err == sql.ErrNoRows {                                                                        //if the user does not aready have a set sharing a name then it will create a new set
		stmtC, err := db.Prepare("INSERT INTO cards (userOwn, setName) VALUES(?, ?)") //Prepares statement that creates new set
		checkError(err)
		res, err := stmtC.Exec(uID, setName)
		checkError(err)
		fmt.Println(res)
		lastID, err := res.LastInsertId()
		checkError(err)
		return lastID //Returns id of the set
	}
	checkError(err)
	for rows.Next() { //returns the setID only Runs once as return statement exits the loop
		var setID int64
		noRows = rows.Scan(&setID)
		return setID
	}
	fmt.Println(noRows) //If error prints it
}
