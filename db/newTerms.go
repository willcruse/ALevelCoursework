package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var uID int

func main() {
	term1 := "Hello"
	term2 := "world"
	term3 := "print"
	term4 := "function"
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	id := 16
	checkError(err)
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(id, term1, term2)
	checkError(err)
	stmtT1, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res3, err := stmtT1.Exec(id, term3, term4)
	checkError(err)
	fmt.Println(res2)
	fmt.Println(res3)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
