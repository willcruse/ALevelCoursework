package dbOpertations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	var uID = 1
	db, err := sql.Open("mysql", "root:roottoor@/educationWebsite")
	checkError(err)
	defer db.Close()

	errCon := db.Ping()
	checkError(errCon)
	rows, err := db.Query("SELECT * FROM users WHERE uID=?", uID)
	checkError(err)
	for rows.Next() {
		var uID int
		var email string
		var uName string
		var pw string
		err = rows.Scan(&uID, &email, &uName, &pw)
		fmt.Println(uID)
		fmt.Println(email)
		fmt.Println(uName)
		fmt.Println(pw)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
