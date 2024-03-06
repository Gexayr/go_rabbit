// database/database.go
package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// DB holds the connection pool to the MySQL database
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    // Set up the connection string
    connectionString := "root:password@tcp(localhost:3306)/rabbit_game"

    // Open a connection to the database
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return err
    }

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        return err
    }

    fmt.Println("Connected to the database")

    // Assign the database connection to the global variable
    DB = db

    return nil
}
