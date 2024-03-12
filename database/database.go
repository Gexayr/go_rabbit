package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
    var err error
    dsn := "root:password@tcp(localhost:3306)/rabbit_game"

    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to the database")
}

func Query(statement string) (*sql.Rows, error) {
    rows, err := db.Query(statement)
    if err != nil {
        log.Fatal(err)
    }
    return rows, err
}

func Exec(statement string, args ...interface{}) (sql.Result, error) {
    res, err := db.Exec(statement, args...)
    if err != nil {
        log.Fatal(err)
    }
    return res, err
}
