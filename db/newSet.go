package dbOpertations

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	var uID int
	setName := "TestSet"
	term1 := "Hello"
	term2 := "world"
	term3 := "print"
	term4 := "function"
	db, err := sql.Open("mysql", "will:somePass@/educationWebsite")
	checkError(err)
	defer db.Close()
	errCon := db.Ping()
	checkError(errCon)
	stmtC, err := db.Prepare("INSERT INTO cards (userOwn, setName) VALUES(3, ?)")
	checkError(err)
	res, err := stmtC.Exec(setName)
	checkError(err)
	//From here on adds new terms ----> own file
	id, err := res.LastInsertId()
	checkError(err)
	stmtT, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res2, err := stmtT.Exec(id, term1, term2)
	checkError(err)
	stmtT1, err := db.Prepare("INSERT INTO terms VALUES(?, ?, ?)")
	checkError(err)
	res3, err := stmtT1.Exec(id, term3, term4)
	checkError(err)
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
