package main

import(
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

var db *sql.DB

func main()  {
    email := "somemail@address.com"
    uName := "theREALguy"
    pw := "UNHACKABLE"
    db, err := sql.Open("mysql", "root:roottoor@/educationWebsite")
    checkError(err)
    defer db.Close()
    errCon := db.Ping()
    checkError(errCon)
    stmt, err := db.Prepare("INSERT INTO users (email, uName, pw) VALUES(?, ?, ?)")
    checkError(err)
    res, err := stmt.Exec(email, uName, pw)
    checkError(err)
    affect, err := res.RowsAffected()
    checkError(err)
    fmt.Println("Rows:", affect)
}

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}
